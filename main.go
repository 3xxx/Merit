package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "merit/models"
	_ "merit/routers"
)

func main() {
	//开启orm调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

//错误描述：当controllers中的func中没有使用models中的func时，提示没有定义default数据库。
//也就是controllers中不使用models时，models中的init()不起作用

// <orm.RegisterModel> table name `category` repeat register, must be unique
//因为"merit/models"没有改到，原来是quick/models
