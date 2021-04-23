package utils

import (
	"context"
	"strings"
	"toolDB/serverDB/thrift/gen-go/openstars/core/bigset/generic"
)

func BSToMySQL(urlMySql, urlBS, sid string) error {
	// ================Connection============
	db := &DataBase{}
	err := db.ConnectSql(urlMySql)
	if err != nil {
		return err
	}
	err = db.ConnectBSWith(urlBS)
	if err != nil {
		return err
	}
	// =======================================
	db.Sid = strings.ReplaceAll(sid, "/", "_")
	db.Sid = strings.ReplaceAll(db.Sid, "-", "_")

	// ======== Create table =================
	_ = db.CreateTable()

	// ================Get All BsKey
	count, err := db.Bs.TotalStringKeyCount(context.Background())
	if err != nil {
		return err
	}
	slices, err := db.Bs.GetListKeyFrom(context.Background(), generic.TStringKey(0), int32(count-1))
	if err != nil {
		return err
	}
	for _, item := range slices{
		err := db.ReadAndMultiPutToMySQL(string(item))
		if err != nil {
			return err
		}
	}
	return nil
}

func BSToMySQLWithBSKey(urlMySql, urlBS, sid, bsKey string) error {
	// ================Connection============
	db := &DataBase{}
	err := db.ConnectSql(urlMySql)
	if err != nil {
		return err
	}
	err = db.ConnectBSWith(urlBS)
	if err != nil {
		return err
	}
	// =======================================
	db.Sid = strings.ReplaceAll(sid, "/", "_")
	db.Sid = strings.ReplaceAll(db.Sid, "-", "_")

	// ======== Create table =================
	_ = db.CreateTable()

	// ================Get All BsKey
	err = db.ReadAndMultiPutToMySQL(bsKey)
	if err != nil {
		return err
	}

	return nil
}

func BSToMySQLWithPrefix(urlMySql, urlBS, sid, prefix string) error {
	// ================Connection============
	db := &DataBase{}
	err := db.ConnectSql(urlMySql)
	if err != nil {
		return err
	}
	err = db.ConnectBSWith(urlBS)
	if err != nil {
		return err
	}
	// =======================================
	db.Sid = strings.ReplaceAll(sid, "/", "_")
	db.Sid = strings.ReplaceAll(db.Sid, "-", "_")

	// ======== Create table =================
	_ = db.CreateTable()

	// ================Get All BsKey
	count, err := db.Bs.TotalStringKeyCount(context.Background())
	if err != nil {
		return err
	}
	slices, err := db.Bs.GetListKeyFrom(context.Background(), generic.TStringKey(0), int32(count-1))
	if err != nil {
		return err
	}
	for _, item := range slices{
		//println("[0. CHECK CONTAIN] ", string(item), )
		if strings.Contains(string(item), prefix) {
			err := db.ReadAndMultiPutToMySQL(string(item))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func BSToMySQLWithInclude(urlMySql, urlBS, sid, sub string) error {
	// ================Connection============
	db := &DataBase{}
	err := db.ConnectSql(urlMySql)
	if err != nil {
		return err
	}
	err = db.ConnectBSWith(urlBS)
	if err != nil {
		return err
	}
	// =======================================
	db.Sid = strings.ReplaceAll(sid, "/", "_")
	db.Sid = strings.ReplaceAll(db.Sid, "-", "_")

	// ======== Create table =================
	_ = db.CreateTable()

	// ================Get All BsKey
	count, err := db.Bs.TotalStringKeyCount(context.Background())
	if err != nil {
		return err
	}
	slices, err := db.Bs.GetListKeyFrom(context.Background(), generic.TStringKey(0), int32(count-1))
	if err != nil {
		return err
	}
	for _, item := range slices{
		//println("[0. CHECK CONTAIN] ", string(item), )
		if strings.ContainsAny(string(item), sub) {
			err := db.ReadAndMultiPutToMySQL(string(item))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
