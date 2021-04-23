package main

import "toolDB/mysql/utils"

func main()  {
	urlSQL :=  "trustkeys_dev:trustkeys.dev@tcp(192.168.2.8:4000)/trustkeys_dev"
	urlBS :=  "192.168.2.8:18733"
	err := utils.BSToMySQLWithPrefix(urlSQL, urlBS, "/dev/openstars/cryptocurrency/crypto-price/services/thrift", "USDT")
	if err != nil {
		return
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