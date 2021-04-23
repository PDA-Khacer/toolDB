package utils

import (
	"context"
	"fmt"
	"log"
	"sync"
	"toolDB/serverDB/thrift/gen-go/openstars/core/bigset/generic"
)

var(
	UPDATE_EXIST_SQL = "UPDATE %s set Val = ? where BsKey = ? and BsItemKey = ?;"
	DELETE_SQL = "DELETE FROM %s where BsKey = ? and BsItemKey = ?;"
	INSERT_SQL = "Insert into %s values(?,?,?);"
	INSERT_BULK_SQL = "Insert into %s values"
	SELECT_ONE_SQL = "SELECT BsItemKey, Val FROM %s WHERE BsKey = ? and BsItemKey = ?;"
	SELECT_SQL = "SELECT BsItemKey, Val FROM %s WHERE BsKey = ?"
	CREATE_TABLE = "CREATE TABLE %s ( BsKey nvarchar(100) not null , BsItemKey nvarchar(100) not null , Val nvarchar(100), PRIMARY KEY(BsKey,BsItemKey))"
)

type DataBs struct {
	BsKey string
	BsItemKey string
	BsItemVal string
}


func (d *DataBase) GetItem(BsKey, BsItemKey string) (*generic.TItem,error){
	key := ""
	value := ""
	row := d.Mysql.QueryRow(fmt.Sprintf(SELECT_ONE_SQL, d.Sid), BsKey, BsItemKey)
	if row.Err() != nil{
		log.Println("[ERROR SELECT]",fmt.Sprintf(SELECT_ONE_SQL, d.Sid), BsKey, BsItemKey)
		return nil, row.Err()
	}
	err := row.Scan(&key, &value)
	if err != nil {
		log.Println("[ERROR SCAN] Mysql/utils/handler.go:31")
	}
	return &generic.TItem{
		Key:   []byte(key),
		Value: []byte(value),
	}, err
}

func (d *DataBase) GetItems(BsKey string) ([]*generic.TItem, error) {
	var result []*generic.TItem
	row, err := d.Mysql.Query(fmt.Sprintf(SELECT_SQL, d.Sid), BsKey)
	if err != nil{
		log.Println("[ERROR SELECT]",fmt.Sprintf(SELECT_SQL, d.Sid), BsKey)
		return nil, err
	}
	for row.Next(){
		key := ""
		value := ""
		err = row.Scan(&key, &value)
		if err != nil {
			log.Println("[ERROR SCAN] Mysql/utils/handler.go:31")
		}
		item := &generic.TItem{
			Key:   []byte(key),
			Value: []byte(value),
		}
		result = append(result, item)
	}
	return result, nil
}

func (d *DataBase) PutItem(BsKey, BsItemKey, BsItemVal string) error {
	// check Exists
	if item, err := d.GetItem(BsKey, BsItemKey); err != nil {
		return err
	} else {
		// exist
		if item != nil{
			if _, err := d.Mysql.Exec(fmt.Sprintf(UPDATE_EXIST_SQL, d.Sid), BsItemVal, BsKey, BsItemKey); err != nil{
				log.Println("[ERROR UPDATE] ", fmt.Sprintf(UPDATE_EXIST_SQL, d.Sid), BsItemVal, BsKey, BsItemKey)
				return err
			}
		} else {
			if _, err := d.Mysql.Exec(fmt.Sprintf(INSERT_SQL, d.Sid), BsKey, BsItemKey, BsItemVal); err != nil{
				log.Println("[ERROR INSERT] ", fmt.Sprintf(INSERT_SQL, d.Sid), BsKey, BsItemKey, BsItemVal)
				return err
			}
		}
	}
	return nil
}

func (d *DataBase) PutItems(items []DataBs) error {
	for _, item := range items{
		if err := d.PutItem(item.BsKey, item.BsItemKey, item.BsItemVal); err != nil{
			return err
		}
	}
	return nil
}

func (d *DataBase) AddItems(items []DataBs) error {
	bath_items := ""
	for _, item := range items{
		bath_items += "('" + item.BsKey + "','" + item.BsItemKey + "','" + item.BsItemKey + "') ,"
	}
	if _, err := d.Mysql.Exec(fmt.Sprintf(INSERT_BULK_SQL, d.Sid) + bath_items[:len(bath_items)-1]); err != nil{
		//log.Println("[ERROR INSERT] ", fmt.Sprintf(INSERT_BULK_SQL, d.Sid), bath_items[:len(bath_items)-1])
		return err
	}
	println("DONE ADD")
	return nil
}

func (d *DataBase) CreateTable() error{
	if _, err := d.Mysql.Exec(fmt.Sprintf(CREATE_TABLE, d.Sid)); err != nil{
		log.Println("[ERROR CREATE] ",err.Error(), fmt.Sprintf(CREATE_TABLE, d.Sid))
		return err
	}
	return nil
}

func (d *DataBase) RemoveItem(BsKey, BsItemKey string) error{
	if item, err := d.GetItem(BsKey, BsItemKey); err != nil {
		return err
	} else {
		if item != nil{
			if _, err := d.Mysql.Exec(fmt.Sprintf(DELETE_SQL, d.Sid), BsKey, BsItemKey); err != nil{
				log.Println("[ERROR UPDATE] ", fmt.Sprintf(UPDATE_EXIST_SQL, d.Sid), BsKey, BsItemKey)
				return err
			}
		}
		return nil
	}
}

// =======================================

func (d *DataBase)ReadAndPutToBs(BsKey, BsItemKey string) error {
	if item, err := d.GetItem(BsKey, BsItemKey); err != nil{
		return err
	} else {
		if _, err = d.Bs.BsPutItem(context.Background(), generic.TStringKey(BsKey), item); err != nil{
			return err
		}
	}
	return nil
}

func (d *DataBase)ReadAndMultiPutToBs(BsKey string) error {
	if items, err := d.GetItems(BsKey); err != nil{
		return err
	} else {
		print("INFO", len(items))
		itemsAdd := generic.NewTItemSet()
		itemsAdd.Items = items
		if _, err = d.Bs.BsMultiPut(context.Background(), generic.TStringKey(BsKey), itemsAdd, true, false); err != nil{
			return err
		}
	}
	return nil
}


func (d *DataBase)ReadAndMultiPutToMySQL(BsKey string) error {
	count, err := d.Bs.GetTotalCount(context.Background(), generic.TStringKey(BsKey))
	if err != nil {
		return err
	}
	log.Println("COUNT ", BsKey , count)
	if count > 0{
		numberThread := int(count / 10000)
		log.Println(numberThread)
		var wg sync.WaitGroup
		for i := 0; i <= numberThread; i++ {
			println("INDEX",i, "Number ", numberThread)
			wg.Add(1)
			if i == numberThread{
				go d.subFunction1(&wg, i*10000, int(count) - i*10000,BsKey)
			} else {
				go d.subFunction1(&wg, i*10000, 10000-1,BsKey)
			}
			wg.Wait()
		}
		//wg.Wait()
	}
	return nil
}

func (d *DataBase)subFunction1(wg *sync.WaitGroup, start,count int, BsKey string)  {
	defer wg.Done()
	log.Println("BACKUP FOR ", start, "To", count)
	slice, err := d.Bs.BsGetSlice(context.Background(), generic.TStringKey(BsKey), int32(start), int32(count))
	if err != nil {
		log.Panic("[TRANS] ", err.Error(), " mysql/utils/handler.go:189 ", start, count, BsKey)
	} else {
		log.Println("LEN SLICCE", len(slice.Items.Items))
		var itemsbs = make([]DataBs, 0)
		for _, item := range slice.Items.Items{
			temp := DataBs{BsKey: BsKey, BsItemKey: string(item.Key), BsItemVal: string(item.Value)}
			itemsbs = append(itemsbs, temp)
		}
		log.Println("ADD ITEM ", len(itemsbs))
		err = d.AddItems(itemsbs)
		if err != nil {
			log.Panic("[ADD ITEM] ", err.Error(), " mysql/utils/handler.go:200")
		}
	}
}

