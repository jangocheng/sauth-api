package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego/logs"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:yeqc@tcp(45.76.192.20:3306)/tjauth?charset=utf8")
	if nil != err {
		logs.Error("mysql引擎初始化失败", err)
	}
	Engine.ShowSQL(true)
}
