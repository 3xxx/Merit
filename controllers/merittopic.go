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

func (c *MeritTopicController) GetMeritTopic() {

}

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

	_, err = models.AddMeritTopic(idNum, name, choose, content, ff)

	topics, err := models.GetMeritTopic(idNum)
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
