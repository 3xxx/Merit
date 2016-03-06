// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
package controllers

import (
	json "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"merit/models"
	"strconv"
	"strings"
)

// type JsonStruct struct { //空结构体？？
// }

type JsonController struct {
	beego.Controller
}

type List1 struct {
	Choose string `json:"text"`
	Mark1  string //打分1
}

// type List2 struct { //项目负责人——链接——大、中、小
// 	Project string `json:"text"`
// 	Href    string
// 	Mark2   string  //打分2
// 	Xuanze  []List1 `json:"nodes"` //大型、中型……
// }
type List2 struct { //项目负责人——链接——大、中、小
	Project string `json:"text"`
	Href    string `json:"href"`
	Mark2   string //打分2
	Xuanze  string //大型、中型……
	Mark1   string //对应列表打分
}
type List3 struct { //项目管理类：项目负责人、课题……
	Category string  `json:"text"`
	Fenlei   []List2 `json:"nodes"`
	Parent2  string
}

type List4 struct { //专业室：水工、施工……
	Keshi string  `json:"text"`
	Kaohe []List3 `json:"nodes"`
}

type List5 struct { //分院：施工预算、水工分院……
	Department string  `json:"text"` //这个后面json仅仅对于encode解析有用
	Bumen      []List4 `json:"nodes"`
}

type List6 struct { //分院：施工预算、水工分院……
	Danwei  string  `json:"text"` //这个后面json仅仅对于encode解析有用
	Fenyuan []List5 `json:"nodes"`
}

func (c *JsonController) Get() {
	// contents, _ := ioutil.ReadFile("./conf/json.json")
	// var r List6
	//var r JsonStruct//空结构对于系统unmarshal不行。
	//	var r map[string]interface{}//空接口可行
	//	var r []interface{}//这个对于系统unmarshal不行
	// err := json.Unmarshal([]byte(contents), &r)
	// if err != nil {
	// fmt.Printf("err was %v", err)
	// }

	// js, err := simplejson.NewJson([]byte(contents))
	// if err != nil {
	// panic("json format error")
	// }

	// arr, err := js.Get("nodes").Array()
	// if err != nil {
	// fmt.Println("decode error: get array failed!")
	// return
	// }

	//从数据库取得parentid为0的单位名称和ID
	//然后查询所有parentid为ID的名称——得到分院名称和分院id
	//查询所有parentid为分院id的名称——得到科室名称和科室id
	//查询所有pid为科室id的名称和id——得到价值分类名称和id
	//查询所有pid为价值分类id——得到价值名称和id，分值
	//查询所有pid为价值id——得到选择项和分值——进行字符串分割
	//构造struct——转json数据b, err := json.Marshal(group) fmt.Println(string(b))
	// slice1 := make([]List1, 0)
	slice2 := make([]List2, 0)
	slice3 := make([]List3, 0)
	slice4 := make([]List4, 0)
	slice5 := make([]List5, 0)
	// slice6 := make([]List6, 0)
	category, err := models.GetPids(0) //得到单位
	if err != nil {
		beego.Error(err)
	}
	var List7 List6
	List7.Danwei = category[0].Title                 //单位名称
	category1, err := models.GetPids(category[0].Id) //得到多个分院
	// beego.Info(category[0].Id)
	if err != nil {
		beego.Error(err)
	}
	for i1, _ := range category1 {
		aa := make([]List5, 1)
		aa[0].Department = category1[i1].Title             //分院名称
		category2, err := models.GetPids(category1[i1].Id) //得到多个科室
		// beego.Info(category1[i1].Id)
		if err != nil {
			beego.Error(err)
		}
		for i2, _ := range category2 {
			bb := make([]List4, 1)
			bb[0].Keshi = category2[i2].Title                  //科室名称
			category3, err := models.GetPids(category2[i2].Id) //得到多个价值分类
			// beego.Info(category2[i2].Id)
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range category3 {
				cc := make([]List3, 1)
				cc[0].Category = category3[i3].Title               //价值分类名称
				category4, err := models.GetPids(category3[i3].Id) //得到多个价值
				// beego.Info(category3[i3].Id)
				if err != nil {
					beego.Error(err)
				}
				for i4, _ := range category4 {
					dd := make([]List2, 1)
					dd[0].Project = category4[i4].Title                               //得到价值名称
					dd[0].Mark2 = category4[i4].Mark                                  //得到价值得分
					dd[0].Xuanze = category4[i4].List                                 //得到选择列表
					dd[0].Mark1 = category4[i4].ListMark                              //得到选择列表得分
					dd[0].Href = "/add?id=" + strconv.FormatInt(category4[i4].Id, 10) //得到Id用于添加成果
					//进行选择列表拆分
					// array1 := strings.Split(category4[i4].List, ",")
					// for _, v := range array1 {
					// 	ee := make([]List1, 1)
					// 	ee[0].Choose = v
					// 	slice1 = append(slice1, ee...)
					// }
					// dd[0].Xuanze = slice1
					// array2 := strings.Split(category4[i4].ListMark, ",")
					// for _, v := range array2 {
					// 	ff := make([]List1, 1)
					// 	ff[0].Mark1 = v
					// 	slice0 = append(slice0, ff...)
					// }
					// dd[0].Xuanze = slice1
					slice2 = append(slice2, dd...)
				}
				// var cc1 List3
				cc[0].Fenlei = slice2
				slice2 = make([]List2, 0) //再把slice置0
				slice3 = append(slice3, cc...)
			}
			bb[0].Kaohe = slice3
			slice3 = make([]List3, 0) //再把slice置0
			slice4 = append(slice4, bb...)
		}
		aa[0].Bumen = slice4
		slice4 = make([]List4, 0) //再把slice置0
		slice5 = append(slice5, aa...)
	}
	List7.Fenyuan = slice5
	slice5 = make([]List5, 0) //再把slice置0
	// beego.Info(List7)
	// beego.Info(contents)二进制的东西
	// c.Data["Input"] = r
	// b, err := json.Marshal(List7)//不需要转成json格式
	// beego.Info(string(b))
	// fmt.Println(string(b))
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = List7
	// c.ServeJSON()
	c.TplName = "json_show.tpl"
}

func (c *JsonController) Add() {
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
	c.Data["category"] = category
	c.Data["list"] = slice1

	topics, err := models.GetMeritTopic(idNum)
	c.Data["topics"] = topics
	// c.ServeJSON()
	c.TplName = "add.tpl"
}

func (c *JsonController) ImportJson() {
	contents, _ := ioutil.ReadFile("./conf/json.json")
	var r List6
	//var r JsonStruct//空结构对于系统unmarshal不行。
	//	var r map[string]interface{}//空接口可行
	//	var r []interface{}//这个对于系统unmarshal不行
	err := json.Unmarshal([]byte(contents), &r)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	// fmt.Println(r)
	// beego.Info(r)

	js, err := simplejson.NewJson([]byte(contents))
	if err != nil {
		panic("json format error")
	}
	//1.获取省水利院
	text, err := js.Get("text").String()
	//存入数据库——单位
	Id, err := models.AddCategory(0, text, "", "", "", "")
	if err != nil {
		beego.Error(err)
	}

	arr, err := js.Get("nodes").Array()
	if err != nil {
		fmt.Println("decode error: get array failed!")
		// return
	}
	for i, _ := range arr {
		// beego.Info(v)是map[string]interface{}
		text1, _ := js.Get("nodes").GetIndex(i).Get("text").String()
		//存入数据库——分院
		Id1, err := models.AddCategory(Id, text1, "", "", "", "")
		if err != nil {
			beego.Error(err)
		}
		arr1, err := js.Get("nodes").GetIndex(i).Get("nodes").Array()
		for i1, _ := range arr1 {
			text2, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("text").String()
			//存入数据库——科室
			Id2, err := models.AddCategory(Id1, text2, "", "", "", "")
			if err != nil {
				beego.Error(err)
			}
			arr2, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").Array()
			for i2, _ := range arr2 {
				text3, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("text").String()
				//存入数据库——管理类
				Id3, err := models.AddCategory(Id2, text3, "", "", "", "")
				if err != nil {
					beego.Error(err)
				}
				arr3, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").Array()
				for i3, _ := range arr3 {
					text4, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("text").String()
					text5, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("mark").String()
					//循环取出选择项，拼接字符串
					//循环取出每个选择项的打分，拼接字符串
					arr4, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").Array()
					var text8, text9 string
					for i4, _ := range arr4 {
						text6, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("text").String()
						text7, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("mark").String()
						text8 = text8 + "," + text6
						text9 = text9 + "," + text7
					}
					//存入数据库——项目负责人
					// url:="/"+"add?id="+
					_, err = models.AddCategory(Id3, text4, text5, "", text8, text9)
					if err != nil {
						beego.Error(err)
					}
				}

			}
		}
	}
}

// func NewJsonStruct() *JsonStruct {
// 	return &JsonStruct{}
// }

// func (self *JsonStruct) Load(filename string, v interface{}) {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return
// 	}
// 	datajson := []byte(data)
// 	err = json.Unmarshal(datajson, v)
// 	if err != nil {
// 		return
// 	}
// }

// type ValueTestAtmp struct {
// 	StringValue    string
// 	NumericalValue int
// 	BoolValue      bool
// }

// type testdata struct {
// 	ValueTestA ValueTestAtmp
// }

// package main
// import (
//     "encoding/json"
//     "fmt"
// )
// func test() {
// 	b := []byte(`{
//     "Title":"go programming language",
//     "Author":["john","ada","alice"],
//     "Publisher":"qinghua",
//     "IsPublished":true,
//     "Price":99
//   }`)
// 	//先创建一个目标类型的实例对象，用于存放解码后的值
// 	var inter interface{}
// 	err := json.Unmarshal(b, &inter)
// 	if err != nil {
// 		fmt.Println("error in translating,", err.Error())
// 		return
// 	}
// 	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
// 	book, ok := inter.(map[string]interface{})
// 	//然后通过for循环一一访问解码后的目标数据
// 	if ok {
// 		for k, v := range book {
// 			switch vt := v.(type) {
// 			case float64:
// 				fmt.Println(k, " is float64 ", vt)
// 			case string:
// 				fmt.Println(k, " is string ", vt)
// 			case []interface{}:
// 				fmt.Println(k, " is an array:")
// 				for i, iv := range vt {
// 					fmt.Println(i, iv)
// 				}
// 			default:
// 				fmt.Println("illegle type")
// 			}
// 		}
// 	}
// }

// 今天遇到个接口需要处理一个json的map类型的数组，开始想法是用simple—json里的Array读取数组，然后遍历数组取出每个map，然后读取对应的值，在进行后续操作，貌似很简单的工作，却遇到了一个陷阱。
// Json格式类似下边：
// {"code":0
// ,"request_id": xxxx
// ,"code_msg":""
// ,"body":[{
//         "device_id": "xxxx"
//         ,"device_hid": "xxxx"
// }]
// , "count":0}
//     很快按上述想法写好了带码，但是以外发生了，编译不过，看一看代码逻辑没有问题，问题出在哪里呢？
//     原来是interface{} Array方法返回的是一个interface{}类型的，我们都在golang里interface是一个万能的接受者可以保存任意类型的参数，但是却忽略了一点，它是不可以想当然的当任意类型来用，在使用之前一定要对interface类型进行判断。我开始就忽略了这点，想当然的使用interface变量造成了错误。
//     下面写了个小例子

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/bitly/go-simplejson"
// )

// func test2() {
// 	//拼凑json   body为map数组
// 	var rbody []map[string]interface{}
// 	t := make(map[string]interface{})
// 	t["device_id"] = "dddddd"
// 	t["device_hid"] = "ddddddd"

// 	rbody = append(rbody, t)
// 	t1 := make(map[string]interface{})
// 	t1["device_id"] = "aaaaa"
// 	t1["device_hid"] = "aaaaa"

// 	rbody = append(rbody, t1)

// 	cnnJson := make(map[string]interface{})
// 	cnnJson["code"] = 0
// 	cnnJson["request_id"] = 123
// 	cnnJson["code_msg"] = ""
// 	cnnJson["body"] = rbody
// 	cnnJson["page"] = 0
// 	cnnJson["page_size"] = 0

// 	b, _ := json.Marshal(cnnJson)
// 	cnnn := string(b)
// 	fmt.Println("cnnn:%s", cnnn)
// 	cn_json, _ := simplejson.NewJson([]byte(cnnn))
// 	cn_body, _ := cn_json.Get("body").Array()

// 	for _, di := range cn_body {
// 		//就在这里对di进行类型判断
// 		newdi, _ := di.(map[string]interface{})
// 		device_id := newdi["device_id"]
// 		device_hid := newdi["device_hid"]
// 		fmt.Println(device_hid, device_id)
// 	}

// }

// 第一步，得到json的内容
// contents, _ := ioutil.ReadAll(res.Body)
// js, js_err := simplejson.NewJson(contents)

// 第二部，根据json的格式，选择使用array或者map储存数据
// var nodes = make(map[string]interface{})
// nodes, _ = js.Map()

// 第三步，将nodes当作map处理即可，如果map的value仍是一个json机构，回到第二步。
// for key,_ := range nodes {
// ...
// }
