// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
//本控制器用于侧栏的显示和修改等
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
	Choose string `json:"choose"`
	Mark1  string `json:"mark1"` //打分1
}

// type List2 struct { //项目负责人——链接——大、中、小
// 	Project string `json:"text"`
// 	Href    string
// 	Mark2   string  //打分2
// 	Xuanze  []List1 `json:"nodes"` //大型、中型……
// }
type List2 struct { //项目负责人——链接——大、中、小
	Id      int64  `form:"-"`
	Pid     int64  `form:"-"`
	Project string `json:"text"`
	Href    string `json:"href"`
	Tags    [2]int `json:"tags"`
	Mark2   string `json:"mark2"` //打分2
	Xuanze  string //大型、中型……
	Mark1   string //对应列表打分
}
type List3 struct { //项目管理类：项目负责人、课题……
	Id         int64   `form:"-"`
	Pid        int64   `form:"-"`
	Category   string  `json:"text"`
	Selectable bool    `json:"selectable"`
	Tags       [2]int  `json:"tags"`
	Fenlei     []List2 `json:"nodes"`
	Parent2    string
}

type List4 struct { //专业室：水工、施工……
	Id         int64   `form:"-"`
	Pid        int64   `form:"-"`
	Keshi      string  `json:"text"`
	Selectable bool    `json:"selectable"`
	Kaohe      []List3 `json:"nodes"`
}

type List5 struct { //分院：施工预算、水工分院……
	Id         int64   `form:"-"`
	Pid        int64   `form:"-"`
	Department string  `json:"text"` //这个后面json仅仅对于encode解析有用
	Selectable bool    `json:"selectable"`
	Bumen      []List4 `json:"nodes"`
}

type List6 struct { //总院：水利设计院……
	Id         int64   `form:"-"`
	Pid        int64   `form:"-"`
	Danwei     string  `json:"text"` //这个后面json仅仅对于encode解析有用
	Selectable bool    `json:"selectable"`
	Fenyuan    []List5 `json:"nodes"`
}

type Person struct { //总院：水利设计院……
	Id         int64  `form:"-"`
	Name       string `json:"Name"`
	Department string `json:"Department"`
	Keshi      string `json:"Keshi"` //当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //分值
	Marks      int    //记录个数
}

//管理员进行人员价值排序查看
//排序第一排序为部门，第二排序为科室，第三排序为分值
func (c *JsonController) GetPerson() {
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
	//3.由用户id取得用户名
	//4.取得客户端用户名
	// var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		// uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	// uname := v.(string) //ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	// rolename, _ := strconv.ParseInt(role, 10, 64)
	if role > 2 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	var numbers1, marks1 int
	slice1 := make([]Person, 0)
	users, _ := models.GetAllusers(1, 2000, "Id")
	for i1, _ := range users {
		//根据价值id和用户id，得到成果，统计数量和分值
		//取得用户的价值数量和分值
		_, numbers, marks, err := models.GetMeritTopic(0, users[i1].Id)
		if err != nil {
			beego.Error(err)
		}
		marks1 = marks1 + marks
		numbers1 = numbers1 + numbers
		aa := make([]Person, 1)
		aa[0].Id = users[i1].Id //这里用for i1,v1,然后用v1.Id一样的意思
		aa[0].Name = users[i1].Nickname
		aa[0].Department = users[i1].Department
		aa[0].Keshi = users[i1].Secoffice
		aa[0].Numbers = numbers1
		aa[0].Marks = marks1
		slice1 = append(slice1, aa...)
		marks1 = 0
		numbers1 = 0
	}
	c.Data["person"] = slice1
	c.TplName = "admin_person.tpl"
}

//管理员登录显示所有价值结构，方便后面操作
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
	List7.Id = category[0].Id
	List7.Danwei = category[0].Title //单位名称
	List7.Selectable = false
	category1, err := models.GetPids(category[0].Id) //得到多个分院
	// beego.Info(category[0].Id)
	if err != nil {
		beego.Error(err)
	}
	for i1, _ := range category1 {
		aa := make([]List5, 1)
		aa[0].Id = category1[i1].Id
		aa[0].Pid = category[0].Id
		aa[0].Department = category1[i1].Title //分院名称
		aa[0].Selectable = false
		category2, err := models.GetPids(category1[i1].Id) //得到多个科室
		// beego.Info(category1[i1].Id)
		if err != nil {
			beego.Error(err)
		}
		for i2, _ := range category2 {
			bb := make([]List4, 1)
			bb[0].Id = category2[i2].Id
			bb[0].Pid = category1[i1].Id
			bb[0].Keshi = category2[i2].Title //科室名称
			bb[0].Selectable = false
			category3, err := models.GetPids(category2[i2].Id) //得到多个价值分类
			// beego.Info(category2[i2].Id)
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range category3 {
				cc := make([]List3, 1)
				cc[0].Id = category3[i3].Id
				cc[0].Pid = category2[i2].Id
				cc[0].Category = category3[i3].Title //价值分类名称
				cc[0].Selectable = false
				cc[0].Tags[0] = 0
				cc[0].Tags[1] = 0
				category4, err := models.GetPids(category3[i3].Id) //得到多个价值
				// beego.Info(category3[i3].Id)
				if err != nil {
					beego.Error(err)
				}
				for i4, _ := range category4 {
					dd := make([]List2, 1)
					dd[0].Id = category4[i4].Id
					dd[0].Pid = category3[i3].Id
					dd[0].Project = category4[i4].Title //得到价值名称
					dd[0].Tags[0] = 0
					dd[0].Tags[1] = 0
					dd[0].Mark2 = category4[i4].Mark                                  //得到价值得分
					dd[0].Xuanze = category4[i4].List                                 //得到选择列表
					dd[0].Mark1 = category4[i4].ListMark                              //得到选择列表得分
					dd[0].Href = "/add?id=" + strconv.FormatInt(category4[i4].Id, 10) //得到Id用于添加成果 + " target='_blank'"
					// beego.Info(dd[0].Href)
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

	// b, err := json.Marshal(List7) //不需要转成json格式
	c.Data["Input"] = List7
	// beego.Info(string(b))
	// fmt.Println(string(b))
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = List7
	// c.ServeJSON()
	c.TplName = "admin_json_show.tpl"
}

//用户登录后获得自己所在的分院和科室，然后显示对应的菜单
//同时显示所有的价值内容
func (c *JsonController) GetMeritUser() {
	//读取用户id
	//查询用户分院名称和科室名称
	//查出分院id和科室id
	//从数据库取得parentid为0的单位名称和ID
	//然后
	//过滤科室下——得到价值分类名称和id
	//查询所有pid为价值分类id——得到价值名称和id，分值
	//查询所有pid为价值id——得到选择项和分值——进行字符串分割
	//构造struct——
	//这个不用：转json数据b, err := json.Marshal(group) fmt.Println(string(b))
	var user models.User
	var err error
	//管理员可以查看
	Uid := c.Input().Get("uid")
	if Uid == "" { //如果是技术人员自己进行查看，则Uid为空
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
		//3.由用户id取得用户名
		//4.取得客户端用户名
		var uname string
		sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		defer sess.SessionRelease(c.Ctx.ResponseWriter)
		v := sess.Get("uname")
		if v != nil {
			uname = v.(string)
			c.Data["Uname"] = v.(string)
		}
		// uname := v.(string) //ck.Value
		//4.取出用户的权限等级
		role, _ := checkRole(c.Ctx) //login里的
		// beego.Info(role)
		//5.进行逻辑分析：
		// rolename, err := strconv.ParseInt(role, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		if role > 5 { //
			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
		user, err = models.GetUserByUsername(uname) //得到用户的id、分院和科室等
		if err != nil {
			beego.Error(err)
		}
	} else { //如果是管理员进行查看，则uid是用户名
		userid, err := strconv.ParseInt(Uid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		user = models.GetUserByUserId(userid)
	}

	var numbers1, marks1 int
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
	List7.Danwei = category[0].Title //单位名称
	List7.Selectable = false
	// category1, err := models.GetPids(category[0].Id) //得到多个分院
	//查出分院id——再由分院id作为parentid和科室名称，得到科室id
	department, err := models.GetCategorybyname(user.Department)
	if err != nil {
		beego.Error(err)
	}
	secoffice, err := models.GetCategorybyidtitle(department.Id, user.Secoffice)
	if err != nil {
		beego.Error(err)
	}
	// for i1, _ := range category1 {
	// if category1[i1].Title == user.Department { //如果分院名=用户分院 名
	aa := make([]List5, 1)
	aa[0].Department = department.Title //user.Department //分院名称
	aa[0].Selectable = false
	// category2, err := models.GetPids(category1[i1].Id) //得到多个科室
	// beego.Info(category1[i1].Id)
	// if err != nil {
	// beego.Error(err)
	// }
	// for i2, _ := range category2 {
	// if category2[i2].Title == user.Secoffice { //如果科室名称=用户科室名称

	bb := make([]List4, 1)
	bb[0].Keshi = secoffice.Title //user.Secoffice //科室名称
	bb[0].Selectable = false
	//由分院名和科室名得到科室id——得到多个价值分类
	category3, err := models.GetPids(secoffice.Id) //得到多个价值分类
	// beego.Info(category2[i2].Id)
	if err != nil {
		beego.Error(err)
	}

	for i3, _ := range category3 {
		cc := make([]List3, 1)
		cc[0].Category = category3[i3].Title //价值分类名称
		cc[0].Selectable = false
		category4, err := models.GetPids(category3[i3].Id) //得到多个价值
		// beego.Info(category3[i3].Id)
		if err != nil {
			beego.Error(err)
		}
		for i4, _ := range category4 {
			dd := make([]List2, 1)
			dd[0].Project = category4[i4].Title //得到价值名称
			//根据价值id和用户id，得到成果，统计数量和分值
			//取得用户的价值数量和分值
			_, numbers, marks, err := models.GetMeritTopic(category4[i4].Id, user.Id)
			if err != nil {
				beego.Error(err)
			}
			dd[0].Tags[0] = marks
			dd[0].Tags[1] = numbers
			dd[0].Mark2 = category4[i4].Mark                                  //得到价值得分
			dd[0].Xuanze = category4[i4].List                                 //得到选择列表
			dd[0].Mark1 = category4[i4].ListMark                              //得到选择列表得分
			dd[0].Href = "/add?id=" + strconv.FormatInt(category4[i4].Id, 10) //得到Id用于添加成果
			marks1 = marks1 + marks
			numbers1 = numbers1 + numbers
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
		cc[0].Tags[0] = marks1
		marks1 = 0
		cc[0].Tags[1] = numbers1
		numbers1 = 0
		cc[0].Fenlei = slice2
		slice2 = make([]List2, 0) //再把slice置0
		slice3 = append(slice3, cc...)
	}
	bb[0].Kaohe = slice3
	slice3 = make([]List3, 0) //再把slice置0
	slice4 = append(slice4, bb...)
	// }
	// }
	aa[0].Bumen = slice4
	slice4 = make([]List4, 0) //再把slice置0
	slice5 = append(slice5, aa...)
	// }
	// }
	List7.Fenyuan = slice5
	slice5 = make([]List5, 0) //再把slice置0
	// beego.Info(List7)
	// beego.Info(contents)二进制的东西
	// c.Data["Input"] = r
	// b, err := json.Marshal(List7) //不需要转成json格式
	// beego.Info(string(b))
	// fmt.Println(string(b))
	// if err != nil {
	// 	beego.Error(err)
	// }
	c.Data["json"] = List7
	// c.ServeJSON()
	c.TplName = "json_show.tpl"
	//查出所有用户的价值资料
	//前提是价值资料里要带用户id
	topics, err := models.GetAllMeritTopic(user.Id)
	c.Data["topics"] = topics
	// for i5, _ := range topics {
	// 	ee := make([]List2, 1)
	// 	cagegory5, err := models.GetCategory(topics[i5].ParentId)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	ee[0].Project = cagegory5.Title
	// 	ee[0].Id = cagegory5.Id
	// 	slice2 = append(slice2, ee...)
	// }
	category5, _ := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	c.Data["category"] = category5
	// beego.Info(slice2)
}

//添加价值结构中的项目
func (c *JsonController) Addjson() {
	id := c.Input().Get("pid")
	text1 := c.Input().Get("title")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//存入数据库——分院
	_, err = models.AddCategory(idNum, text1, "", "", "", "")
	if err != nil {
		beego.Error(err)
	} else {
		// c.Data["json"] = map[string]interface{}{
		// 	"state":    "SUCCESS",
		// 	"data":     "111",
		// 	"original": "demo.jpg",
		// }
		// c.ServeJSON()
		//返回值给ajax的data
		data := text1
		c.Ctx.WriteString(data)
	}
}

//显示——修改价值结构中的项目
func (c *JsonController) Modifyjson() {
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
	//3.由用户id取得用户名
	//4.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// uname := v.(string) //ck.Value
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//5.进行逻辑分析：
	// rolename, _ := strconv.ParseInt(role, 10, 64)
	if role > 2 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	//4.取得价值列表choose和mark
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	category, err := models.GetCategory(idNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(category)
	// var ff string
	//如果mark为空，则寻找选择列表的分值
	//进行选择列表拆分
	array1 := strings.Split(category.List, ",")
	array2 := strings.Split(category.ListMark, ",")
	// beego.Info(category.ListMark)
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
	for i, v := range array1 {
		ee := make([]List1, 1)
		ee[0].Choose = v
		ee[0].Mark1 = array2[i]
		slice1 = append(slice1, ee...)
	}

	slice2 := make([]List1, 0)
	for _, w := range array2 {
		ff := make([]List1, 1)
		ff[0].Mark1 = w
		slice2 = append(slice2, ff...)
	}
	c.TplName = "admin_json_modify.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["category"] = category
	c.Data["list"] = slice1
	c.Data["mark"] = slice2
	// beego.Info(slice2)
	c.Data["Cid"] = id
	c.Data["IsCategory"] = true
}

//提交修改—修改价值结构中的项目
func (c *JsonController) ModifyjsonPost() {
	title := c.Input().Get("title")
	mark := c.Input().Get("mark")
	list := c.Input().Get("list")
	listmark := c.Input().Get("listmark")
	//4.取得价值列表choose和mark
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.Modifyjson(idNum, title, mark, "", list, listmark)
	if err != nil {
		beego.Error(err)
	}
	// c.Redirect("/view?id="+tid, 302) //回到修改后的文章
	c.Redirect("/json", 301)
}

//删除价值结构中的项目
func (c *JsonController) Deletejson() {
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.Deletejson(idNum)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/json", 301)
}

//导入json数据
func (c *JsonController) ImportJson() {
	//获取上传的文件
	_, h, err := c.GetFile("json")
	if err != nil {
		beego.Error(err)
	}
	var path string
	if h != nil {
		//保存附件
		path = ".\\attachment\\" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("json", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	contents, _ := ioutil.ReadFile(path)
	var r List6
	//var r JsonStruct//空结构对于系统unmarshal不行。
	//	var r map[string]interface{}//空接口可行
	//	var r []interface{}//这个对于系统unmarshal不行
	err = json.Unmarshal([]byte(contents), &r)
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
						if i == 0 {
							text8 = text6
							text9 = text7
						} else {
							text8 = text8 + "," + text6
							text9 = text9 + "," + text7
						}
						// for i, label2 := range label {
						// 		if i == 0 {
						// 			label1 = label2.Title
						// 		} else {
						// 			label1 = label1 + "," + label2.Title
						// 		}
						// 	}
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

//根据conf目录下的json.json文件初始化价值结构
func (c *JsonController) InitJson() {
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
						if i == 0 {
							text8 = text6
							text9 = text7
						} else {
							text8 = text8 + "," + text6
							text9 = text9 + "," + text7
						}
						// for i, label2 := range label {
						// 		if i == 0 {
						// 			label1 = label2.Title
						// 		} else {
						// 			label1 = label1 + "," + label2.Title
						// 		}
						// 	}
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

// 今天遇到个接口需要处理一个json的map类型的数组，开始想法是用simple—json
// 里的Array读取数组，然后遍历数组取出每个map，然后读取对应的值，在进行后续操作，
// 貌似很简单的工作，却遇到了一个陷阱。
// Json格式类似下边：
// {"code":0
// ,"request_id": xxxx
// ,"code_msg":""
// ,"body":[{
//         "device_id": "xxxx"
//         ,"device_hid": "xxxx"
// }]
// , "count":0}
//     很快按上述想法写好了带码，但是以外发生了，编译不过，看一看代码逻辑没有
// 问题，问题出在哪里呢？
//     原来是interface{} Array方法返回的是一个interface{}类型的，我们都在golang
// 里interface是一个万能的接受者可以保存任意类型的参数，但是却忽略了一点，它是
// 不可以想当然的当任意类型来用，在使用之前一定要对interface类型进行判断。我开始
// 就忽略了这点，想当然的使用interface变量造成了错误。
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

// 第三步，将nodes当作map处理即可，如果map的value仍是一个json结构，回到第二步。
// for key,_ := range nodes {
// ...
// }
