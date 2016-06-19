package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"merit/models"
	"strconv"
	"strings"
)

type MeritTopicController struct {
	beego.Controller
}

type List11 struct {
	Choose string `json:"text"`
	Mark1  string //打分1
}

//取得所有价值
func (c *MeritTopicController) GetMeritTopic() {
	// topics, err := models.GetMeritTopic(idNum)
}

//显示添加界面
func (c *MeritTopicController) Add() {
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	category, err := models.GetCategory(idNum) //得到价值
	if err != nil {
		beego.Error(err)
	}
	//进行选择列表拆分
	array1 := strings.Split(category.List, ",")
	slice1 := make([]List1, 0)
	for _, v := range array1 {
		ee := make([]List1, 1)
		ee[0].Choose = v
		slice1 = append(slice1, ee...)
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	var uname string
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	//先由uname取得uid
	user := models.GetUserByUsername(uname)
	//取得父级id和用户id下的价值
	topics, _, _, err := models.GetMeritTopic(idNum, user.Id)
	c.Data["topics"] = topics
	// c.ServeJSON()
	c.Data["category"] = category
	c.Data["list"] = slice1
	c.TplName = "add.tpl"
}

//用户进行价值添加
func (c *MeritTopicController) AddMeritTopic() {
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	name := c.Input().Get("name")
	choose := c.Input().Get("choose")
	content := c.Input().Get("editorValue")

	category, err := models.GetCategory(idNum) //得到价值
	if err != nil {
		beego.Error(err)
	}

	var ff string
	//如果mark为空，则寻找选择列表的分值

	//进行选择列表拆分
	array1 := strings.Split(category.List, ",")
	array2 := strings.Split(category.ListMark, ",")
	if category.Mark == "" {
		for i, v := range array1 {
			if v == choose {
				ff = array2[i]
			}
		}
	} else {
		ff = category.Mark
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	var uname string
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}

	_, err = models.AddMeritTopic(idNum, uname, name, choose, content, ff)
	//先由uname取得uid
	user := models.GetUserByUsername(uname)
	topics, _, _, err := models.GetMeritTopic(idNum, user.Id)
	// c.Data["category"] = category
	// c.Data["list"] = slice1
	// // c.ServeJSON()
	// 	id := c.Input().Get("id")
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// category, err := models.GetCategory(idNum) //得到价值
	// if err != nil {
	// 	beego.Error(err)
	// }
	//进行选择列表拆分
	// array1 := strings.Split(category.List, ",")
	slice1 := make([]List1, 0)
	for _, v := range array1 {
		ee := make([]List1, 1)
		ee[0].Choose = v
		slice1 = append(slice1, ee...)
	}
	c.Data["category"] = category
	c.Data["list"] = slice1
	c.Data["topics"] = topics
	c.TplName = "add.tpl"
}

//查看详情
func (c *MeritTopicController) ViewMeritTopic() {
	id := c.Input().Get("id")
	// beego.Info(id)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}

	topic, err := models.GetMeritTopicbyId(id)
	// beego.Info(topic)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "merittopic_view.tpl"
	c.Data["Topic"] = topic
	c.Data["Tid"] = id
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
}

//修改页面
func (c *MeritTopicController) ModifyMeritTopic() { //这个也要登陆验证
	tid := c.Input().Get("id")
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得文章的作者
	merittopic, err := models.GetMeritTopicbyId(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.由用户id取得用户名
	username := models.GetUserByUserId(merittopic.UserId)
	//4.由文章parentid取得价值列表choose和mark
	category, err := models.GetCategory(merittopic.ParentId) //得到价值
	if err != nil {
		beego.Error(err)
	}
	// var ff string
	//如果mark为空，则寻找选择列表的分值
	//进行选择列表拆分
	array1 := strings.Split(category.List, ",")
	// array2 := strings.Split(category.ListMark, ",")
	// if category.Mark == "" {
	// 	for i, v := range array1 {
	// 		if v == choose {
	// 			ff = array2[i]
	// 		}
	// 	}
	// } else {
	// 	ff = category.Mark
	// }
	//进行选择列表拆分
	// array1 := strings.Split(category.List, ",")
	slice1 := make([]List1, 0)
	for _, v := range array1 {
		ee := make([]List1, 1)
		ee[0].Choose = v
		slice1 = append(slice1, ee...)
	}

	//4.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	uname := v.(string) //ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 && uname != username.Username { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	c.TplName = "merittopic_modify.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["category"] = category
	c.Data["list"] = slice1
	// beego.Info(slice1)
	c.Data["Topic"] = merittopic
	c.Data["Tid"] = tid
	c.Data["IsTopic"] = true
}

//提交修改
func (c *MeritTopicController) ModifyPost() { //这个post属于topic_modify.html提交修改。
	//选项修改后需要重新计算分值
	//应该由选项去查找分值，而不应该存储分值到topic中
	//解析表单
	tid := c.Input().Get("id")
	title := c.Input().Get("name")
	choose := c.Input().Get("choose")
	content := c.Input().Get("editorValue")
	//2.取得文章的作者
	merittopic, err := models.GetMeritTopicbyId(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.由用户id取得用户名
	// username := models.GetUserByUserId(merittopic.UserId)
	//4.由文章parentid取得价值列表choose和mark
	category, err := models.GetCategory(merittopic.ParentId) //得到价值
	if err != nil {
		beego.Error(err)
	}
	var ff string
	//如果mark为空，则寻找选择列表的分值

	//进行选择列表拆分
	array1 := strings.Split(category.List, ",")
	array2 := strings.Split(category.ListMark, ",")
	if category.Mark == "" {
		for i, v := range array1 {
			if v == choose {
				ff = array2[i]
			}
		}
	} else {
		ff = category.Mark
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	// var uname string
	if v != nil {
		// uname = v.(string)
		c.Data["Uname"] = v.(string)
	}

	err = models.ModifyMeritTopic(tid, title, choose, content, ff)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/view?id="+tid, 302) //回到修改后的文章
}

//删除价值
func (c *MeritTopicController) DeleteMeritTopic() { //应该显示警告
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	//2.取得文章的作者
	merittopic, err := models.GetMeritTopicbyId(c.Input().Get("id"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	//3.由用户id取得用户名
	username := models.GetUserByUserId(merittopic.UserId)
	//4.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	uname := v.(string)
	//5.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	beego.Info(role)
	//5.进行逻辑分析：
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 && uname != username.Username { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	err = models.DeletMeritTopic(c.Input().Get("id")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/user", 302) //这里增加topic
}
