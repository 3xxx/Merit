package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	"merit/models"
	"time"
)

// type Userselect struct { //
// 	Id   int64  `json:"id"`
// 	Name string `json:"text"`
// }

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Get() {
	this.TplName = "regist.tpl"
}

func (this *RegistController) RegistErr() {
	this.TplName = "registerr.tpl"
}

func (this *RegistController) CheckUname() {
	var user models.User //这里修改
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	err := models.CheckUname(user) //这里修改
	if err == nil {
		this.Ctx.WriteString("false")
		// return false
	} else {
		this.Ctx.WriteString("true")
		// return true
	}
	// return
}

func (this *RegistController) Post() {
	var user models.User //这里修改
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	Pwd1 := inputs.Get("pwd")

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	user.Lastlogintime = time.Now()
	uid, err := models.SaveUser(user) //这里修改

	_, err = models.AddRoleUser(4, uid)
	if err == nil {
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "registerr.tpl"
	}
}

//post方法
func (this *RegistController) GetUname() {
	var user models.User //这里修改[]*models.User(uname string)
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err == nil {
		// this.Ctx.WriteString(string(b))
		this.Data["json"] = uname1 //string(b)
		this.ServeJSON()
	}
	// 	this.Ctx.WriteString(uname1[1].Username)
	// 	// return uname1[0].Username
	// }
	// return uname1[0].Username
}

//get方法，用于x-editable的select2方法
func (this *RegistController) GetUname1() {
	var user models.User //这里修改[]*models.User(uname string)
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err != nil {
		beego.Error(err)
	}

	slice1 := make([]Userselect, 0)

	for _, v := range uname1 {
		aa := make([]Userselect, 1)
		aa[0].Id = v.Id //这里用for i1,v1,然后用v1.Id一样的意思
		// aa[0].Ad = v.Id
		aa[0].Name = v.Username
		slice1 = append(slice1, aa...)
	}
	// b, err := json.Marshal(slice1) //不需要转成json格式
	// beego.Info(string(b))
	// fmt.Println(string(b))
	if err != nil {
		beego.Error(err)
	}
	// this.Data["Userselect"] = slice1
	this.Data["json"] = slice1 //string(b)
	this.ServeJSON()
	// this.TplName = "loginerr.html"
	// 	this.Ctx.WriteString(uname1[1].Username)
	// 	// return uname1[0].Username
	// }
	// return uname1[0].Username
}
