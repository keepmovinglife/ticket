package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ticket/model"
)

const mysql="root:12345678@tcp(192.168.101.255)/market?charset=utf8"
var EngineXorm *xorm.Engine
func init()  {
	var err error
	EngineXorm, err = xorm.NewEngine("mysql", mysql)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}

func MarketIsExit(pair string,chanelname string ) bool  {
	market:=&model.Market{Pair:string(pair),ChanelName:chanelname}
	has, err := EngineXorm.Exist(market)
	if err!=nil{
		log.Fatal("数据库查询失败:", err)
	}
	return has
}
func TickerIsExit(pair string ,chanelname string ) bool{
	var ticket =model.Ticket{Pair:pair,ChanelName:chanelname}
	has, err := EngineXorm.Exist(&ticket)
	if err!=nil{
		log.Fatal("数据库查询失败:", err)
	}
	return has
}

