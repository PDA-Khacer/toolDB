package utils

import (
	"database/sql"
	"github.com/apache/thrift/lib/go/thrift"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"toolDB/serverDB/thrift/gen-go/openstars/core/bigset/generic"
)

type DataBase struct {
	Mysql *sql.DB
	Bs *generic.TStringBigSetKVServiceClient
	Sid string
}



func (d *DataBase) ConnectSql(url string) error {
	if db, err := sql.Open("mysql", url); err != nil{
		return  err
	} else {
		d.Mysql = db
		return nil
	} 
}

func (d *DataBase) CloseSql() error {
	return d.Mysql.Close()
}

func (d *DataBase) ConnectBS() error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:18990")
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportBuff := thrift.NewTBufferedTransportFactory(8192)
	transportFactory1 := thrift.NewTFramedTransportFactory(transportBuff)
	transport, err = transportFactory1.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	d.Bs = generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot))
	return nil
}

func (d *DataBase) ConnectBSWith(url string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(url)
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportBuff := thrift.NewTBufferedTransportFactory(8192)
	transportFactory1 := thrift.NewTFramedTransportFactory(transportBuff)
	transport, err = transportFactory1.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	d.Bs = generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot))
	return nil
}