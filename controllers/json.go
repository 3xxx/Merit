// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
package controllers

import (
	json "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"strconv"
)

type JsonStruct struct { //空结构体？？
}

type JsonController struct {
	beego.Controller
}

type List1 struct {
	Choose string `json:"text"`
}
type List2 struct { //项目负责人——链接——大、中、小
	Project string `json:"text"`
	Href    string
	Xuanze  []List1 `json:"nodes"` //大型、中型……
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
	Bumeng     []List4 `json:"nodes"`
}

func (c *JsonController) Get() {
	contents, _ := ioutil.ReadFile("./conf/json.json")
	var r List5
	//var r JsonStruct//空结构对于系统unmarshal不行。
	//	var r map[string]interface{}//空接口可行
	//	var r []interface{}//这个对于系统unmarshal不行
	err := json.Unmarshal([]byte(contents), &r)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	// fmt.Println(r)
	beego.Info(r)
	js, err := simplejson.NewJson([]byte(contents))
	if err != nil {
		panic("json format error")
	}
	arr, err := js.Get("nodes").Array()
	if err != nil {
		fmt.Println("decode error: get array failed!")
		// return
	}
	// fmt.Println(arr)
	beego.Info(arr)
	if err != nil {
		panic("json format error")
	}

	// beego.Info(input)
	c.Data["Input"] = r
	// aws1 := js.Get("nodes").GetIndex(0).Get("nodes").GetIndex(0).Get("nodes")
	// awsval1, _ := aws1.GetIndex(1).Get("text").String() //Int()
	// beego.Info(awsval1)                                 //施工室

	for _, v := range arr {
		var iv int
		switch v.(type) {
		case float64:
			iv = int(v.(float64))
			fmt.Println(iv)
		case string:
			iv, _ = strconv.Atoi(v.(string)) //string to int
			beego.Info(iv)
			fmt.Println(iv)
		}
		// beego.Info(v)
		// for _, w := range v {
		// }
	}
	// beego.Info(contents)二进制的东西
	// beego.Info(js)

	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
	// book, ok := inter.(map[string]interface{})
	//然后通过for循环一一访问解码后的目标数据
	// if ok {
	// for k, v := range book {
	// 	switch vt := v.(type) {
	// 	case float64:
	// 		fmt.Println(k, " is float64 ", vt)
	// 	case string:
	// 		fmt.Println(k, " is string ", vt)
	// 	case []interface{}:
	// 		fmt.Println(k, " is an array:")
	// 		for i, iv := range vt {
	// 			fmt.Println(i, iv)
	// 		}
	// 	default:
	// 		fmt.Println("illegle type")
	// 	}
	// }
	// }

	// tt := []byte(&v)
	// var r interface{}
	// json.Unmarshal(tt, &r) //这个byte要解码

	c.Data["json"] = js //r
	// c.ServeJSON()
	c.TplName = "user_show.tpl"
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (self *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil {
		return
	}
}

type ValueTestAtmp struct {
	StringValue    string
	NumericalValue int
	BoolValue      bool
}

type testdata struct {
	ValueTestA ValueTestAtmp
}

// package main
// import (
//     "encoding/json"
//     "fmt"
// )
func test() {
	b := []byte(`{  
    "Title":"go programming language",  
    "Author":["john","ada","alice"],  
    "Publisher":"qinghua",  
    "IsPublished":true,  
    "Price":99  
  }`)
	//先创建一个目标类型的实例对象，用于存放解码后的值
	var inter interface{}
	err := json.Unmarshal(b, &inter)
	if err != nil {
		fmt.Println("error in translating,", err.Error())
		return
	}
	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
	book, ok := inter.(map[string]interface{})
	//然后通过for循环一一访问解码后的目标数据
	if ok {
		for k, v := range book {
			switch vt := v.(type) {
			case float64:
				fmt.Println(k, " is float64 ", vt)
			case string:
				fmt.Println(k, " is string ", vt)
			case []interface{}:
				fmt.Println(k, " is an array:")
				for i, iv := range vt {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println("illegle type")
			}
		}
	}
}

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

func test2() {
	//拼凑json   body为map数组
	var rbody []map[string]interface{}
	t := make(map[string]interface{})
	t["device_id"] = "dddddd"
	t["device_hid"] = "ddddddd"

	rbody = append(rbody, t)
	t1 := make(map[string]interface{})
	t1["device_id"] = "aaaaa"
	t1["device_hid"] = "aaaaa"

	rbody = append(rbody, t1)

	cnnJson := make(map[string]interface{})
	cnnJson["code"] = 0
	cnnJson["request_id"] = 123
	cnnJson["code_msg"] = ""
	cnnJson["body"] = rbody
	cnnJson["page"] = 0
	cnnJson["page_size"] = 0

	b, _ := json.Marshal(cnnJson)
	cnnn := string(b)
	fmt.Println("cnnn:%s", cnnn)
	cn_json, _ := simplejson.NewJson([]byte(cnnn))
	cn_body, _ := cn_json.Get("body").Array()

	for _, di := range cn_body {
		//就在这里对di进行类型判断
		newdi, _ := di.(map[string]interface{})
		device_id := newdi["device_id"]
		device_hid := newdi["device_hid"]
		fmt.Println(device_hid, device_id)
	}

}

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
