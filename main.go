package main

import (
	"flag"
	"toolDB/mysql/utils"
)

var (
	urlSQL = flag.String("urlSQL", "trustkeys_dev:trustkeys.dev@tcp(192.168.2.8:4000)/trustkeys_dev", "Ket noi den SQL")
	urlBS = flag.String("urlBS", "192.168.2.8:18733", "Ket noi den BigSet")
	sid = flag.String("sid", "/dev/openstars/cryptocurrency/crypto-price/services/thrift", "sid cua bs")
	option = flag.Int("option", 0, "chon kieu tao ra")
	key = flag.String("key", "key", "chon kieu tao ra")
	fullTransfer = 1
	subPrefix = 2
	include = 3
	exactly = 4
)

func main()  {
	urlSQL := urlSQL
	urlBS :=  urlBS
	sid := sid
	option := option
	key := key
	switch *option {
		case fullTransfer:
			err := utils.BSToMySQL(*urlSQL, *urlBS, *sid)
			if err != nil {
				return
			}
			break
		case exactly:
			err := utils.BSToMySQLWithBSKey(*urlSQL, *urlBS, *sid, *key)
			if err != nil {
				return
			}
			break
		case include:
			err := utils.BSToMySQLWithInclude(*urlSQL, *urlBS, *sid, *key)
			if err != nil {
				return
			}
			break
		case subPrefix:
			err := utils.BSToMySQLWithPrefix(*urlSQL, *urlBS, *sid, *key)
			if err != nil {
				return
			}
			break
	}
	//var	db utils.DataBase
	//err := db.ConnectSql("trustkeys_dev:trustkeys.dev@tcp(192.168.2.8:4000)/trustkeys_dev")
	//if err != nil{
	//	print("[ERR CONNECT] ", err.Error()," mysql/main.go:9 ")
	//}
	//err = db.ConnectBSWith("192.168.2.8:18733")
	//if err != nil {
	//	print("[ERR CONNECT] ", err.Error()," mysql/main.go:17 ")
	//}
	//
	//count, err := db.Bs.TotalStringKeyCount(context.Background())
	//if err != nil {
	//	return
	//}
	//slices, err := db.Bs.GetListKeyFrom(context.Background(), generic.TStringKey(0), int32(count-1))
	//if err != nil {
	//	return
	//}
	//
	//for _, item := range slices{
	//	println(string(item))
	//}

	//db.Sid = "testTool"
	//
	//var d []utils.DataBs
	//item := utils.DataBs{BsKey: "testTool", BsItemKey: "1", BsItemVal: "zxc"}
	//d = append(d, item)
	//item = utils.DataBs{BsKey: "testTool", BsItemKey: "2", BsItemVal: "qwe"}
	//d = append(d, item)
	//item = utils.DataBs{BsKey: "testTool", BsItemKey: "3", BsItemVal: "rty"}
	//d = append(d, item)
	//item = utils.DataBs{BsKey: "testTool", BsItemKey: "4", BsItemVal: "sdf"}
	//d = append(d, item)


	//err = db.AddItems(d)
	//if err != nil {
	//	print(err.Error())
	//}

	//err = db.PutItem("testTool","1", "1236")
	//if err != nil {
	//	print(err.Error())
	//}

	//itemBS := generic.NewTItem()
	//itemBS.Key = []byte("test")
	//itemBS.Value = []byte("tesst")
	//_, err = db.Bs.BsPutItem(context.Background(), "testTool", itemBS)
	//if err != nil{
	//	print(err," main.go:43")
	//}
	//err = db.ReadAndMultiPutToMySQL("testTool")
	//if err != nil {
	//	print(err.Error())
	//}
}