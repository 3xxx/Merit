package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "merit/models"
	_ "merit/routers"
)

//自定义模板函数，序号加1
func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}

func main() {
	beego.AddFuncMap("indexaddone", Indexaddone) //模板中使用{{indexaddone $index}}或{{$index|indexaddone}}

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
