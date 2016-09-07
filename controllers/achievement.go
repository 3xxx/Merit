// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
//本控制器用于侧栏的显示和修改等

//侧栏是直接用user表生成：总院——分院——科室——姓名，这里还是要构造三层结构
//还是用json表来生成？这里利用已有的三层结构，但是要专门的总院-分院-科室的数据表，即json表
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	m "merit/models"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"github.com/astaxie/beego/logs"
	"merit/models"
	"strconv"
	"strings"
	"time"
)

// type JsonStruct struct { //空结构体？？
// }

type Achievement struct {
	beego.Controller
}

type Sidebar1 struct {
	Choose string `json:"choose"`
	Mark1  string `json:"mark1"` //打分1
}

// type Sidebar2 struct { //项目负责人——链接——大、中、小
// 	Project string `json:"text"`
// 	Href    string
// 	Mark2   string  //打分2
// 	Xuanze  []Sidebar1 `json:"nodes"` //大型、中型……
// }
// type Sidebar2 struct { //项目负责人——链接——大、中、小
// 	Id      int64  `form:"-"`
// 	Pid     int64  `form:"-"`
// 	Project string `json:"text"`
// 	Href    string `json:"href"`
// 	Tags    [2]int `json:"tags"`
// 	Mark2   string `json:"mark2"` //打分2
// 	Xuanze  string //大型、中型……
// 	Mark1   string //对应列表打分
// }
type Sidebar3 struct { //员工姓名
	Id       int64  `json:"Id"` //`form:"-"`
	Pid      int64  `form:"-"`
	Nickname string `json:"text"` //这个是侧栏显示的内容
	Level    string `json:"Level"`
	// Href     string `json:"href"`这个不能要，虽然没赋值，但是也导致返回根路径
	// Selectable bool   `json:"selectable"`
	// Tags       [2]int `json:"tags"`
	// Mark2      string `json:"mark2"` //打分2
	// Mark1      string //对应列表打分
	// Xuanze     string //大型、中型……
}

type Sidebar4 struct { //专业室：水工、施工……
	Id       int64      `json:"Id"` //`form:"-"`
	Pid      int64      `form:"-"`
	Keshi    string     `json:"text"`
	Tags     [1]int     `json:"tags"` //显示员工数量
	Yuangong []Sidebar3 `json:"nodes"`
	Level    string     `json:"Level"`
	// Href       string     `json:"href"` //点击科室，显示总体情况
	// Selectable bool       `json:"selectable"`这个不能要，虽然没赋值。否则点击node，没反应，即默认false？？
}

type Sidebar5 struct { //分院：施工预算、水工分院……
	Id         int64      `json:"Id"` //`form:"-"`
	Pid        int64      `form:"-"`
	Department string     `json:"text"` //这个后面json仅仅对于encode解析有用
	Bumen      []Sidebar4 `json:"nodes"`
	Level      string     `json:"Level"`
	// Selectable bool       `json:"selectable"`
}

type Sidebar6 struct { //总院：水利设计院……
	Id      int64      `json:"Id"` //`form:"-"`
	Pid     int64      `form:"-"`
	Danwei  string     `json:"text"` //这个后面json仅仅对于encode解析有用
	Fenyuan []Sidebar5 `json:"nodes"`
	Level   string     `json:"Level"` //部门级别：总院0级，分院1级，科室2级，用户3级
	// Selectable bool       `json:"selectable"`

}

type Employee struct { //职员的分院和科室属性
	Id         int64  `form:"-"`
	Name       string `json:"Name"`
	Department string `json:"Department"` //分院
	Keshi      string `json:"Keshi"`      //科室。当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //分值
	Marks      int    //记录个数
}

//管理员进行人员价值排序查看
//排序第一排序为部门，第二排序为科室，第三排序为分值
func (c *Achievement) GetEmployee() {
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
	slice1 := make([]Employee, 0)
	users, _ := models.GetAllusers(1, 2000, "Id")
	for i1, _ := range users {
		//根据价值id和用户id，得到成果，统计数量和分值
		//取得用户的价值数量和分值
		_, numbers, marks, err := models.GetAchievementTopic(0, users[i1].Id)
		if err != nil {
			beego.Error(err)
		}
		marks1 = marks1 + marks
		numbers1 = numbers1 + numbers
		aa := make([]Employee, 1)
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

//管理员登录显示侧栏结构，方便查看科室里员工总体情况，以及查看员工个人详细
func (c *Achievement) GetAchievement() {
	//从数据库取得parentid为0的单位名称和ID
	//然后查询所有parentid为ID的名称——得到分院名称和分院id
	//查询所有parentid为分院id的名称——得到科室名称和科室id
	//查询所有pid为科室id的名称和id——得到价值分类名称和id
	//查询所有pid为价值分类id——得到价值名称和id，分值
	//查询所有pid为价值id——得到选择项和分值——进行字符串分割
	//构造struct——转json数据b, err := json.Marshal(group) fmt.Println(string(b))
	// slice1 := make([]Sidebar1, 0)
	// slice2 := make([]Sidebar2, 0)
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
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
	if role > 4 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// user = models.GetUserByUsername(uname) //得到用户的id、分院和科室等
	//基于Ip权限的控制：根据ip，查出姓名，得到科室、分院，得到权限等级，根据等级，显示不同侧栏
	//然后交给iframe里结合权限进行内容显示，点击了无权访问的node，显示权限不够。

	slice3 := make([]Sidebar3, 0)
	slice4 := make([]Sidebar4, 0)
	slice5 := make([]Sidebar5, 0)
	// slice6 := make([]Sidebar6, 0)
	category, err := models.GetPids(0) //得到单位
	if err != nil {
		beego.Error(err)
	}
	var Sidebar7 Sidebar6
	Sidebar7.Id = category[0].Id
	Sidebar7.Level = "0"
	Sidebar7.Danwei = category[0].Title //单位名称
	// Sidebar7.Selectable = false
	//由uname取得user,获得user的分院名称
	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	switch role {
	case 1: //管理员登录显示的侧栏是全部的
		category1, err := models.GetPids(category[0].Id) //得到多个分院
		// beego.Info(category[0].Id)
		if err != nil {
			beego.Error(err)
		}
		for i1, _ := range category1 {
			aa := make([]Sidebar5, 1)
			aa[0].Id = category1[i1].Id
			aa[0].Level = "1"
			aa[0].Pid = category[0].Id
			aa[0].Department = category1[i1].Title //分院名称
			// aa[0].Selectable = false
			category2, err := models.GetPids(category1[i1].Id) //得到多个科室
			// beego.Info(category1[i1].Id)
			if err != nil {
				beego.Error(err)
			}
			for i2, _ := range category2 {
				bb := make([]Sidebar4, 1)
				bb[0].Id = category2[i2].Id
				bb[0].Level = "2"
				bb[0].Pid = category1[i1].Id
				bb[0].Keshi = category2[i2].Title //科室名称
				// bb[0].Selectable = false
				//根据分院和科室查所有员工
				users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
				if err != nil {
					beego.Error(err)
				}
				for i3, _ := range users {
					cc := make([]Sidebar3, 1)
					cc[0].Id = users[i3].Id
					cc[0].Level = "3"
					cc[0].Pid = category2[i2].Id
					cc[0].Nickname = users[i3].Nickname //名称
					// cc[0].Selectable = false
					slice3 = append(slice3, cc...)
				}
				//64位转string
				// bb[0].Href = "/secofficeshow?secid=" + strconv.FormatInt(category2[i2].Id, 10) + "&depid=" + strconv.FormatInt(category1[i1].Id, 10) //得到Id用于添加成果 + " target='_blank'"
				bb[0].Tags[0] = count
				bb[0].Yuangong = slice3
				slice3 = make([]Sidebar3, 0) //再把slice置0
				slice4 = append(slice4, bb...)
			}
			aa[0].Bumen = slice4
			slice4 = make([]Sidebar4, 0) //再把slice置0
			slice5 = append(slice5, aa...)
		}
		Sidebar7.Fenyuan = slice5
		slice5 = make([]Sidebar5, 0) //再把slice置0
		// beego.Info(Sidebar7)
		// beego.Info(contents)二进制的东西
	case 2: //分院领导登录显示的侧栏是本分院的所有科室
		//由分院名称取得分院属性
		category1, err := models.GetCategorybyname(user.Department)
		// category1, err := models.GetPids(category[0].Id) //得到多个分院
		// beego.Info(category[0].Id)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]Sidebar5, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		aa[0].Pid = category[0].Id
		aa[0].Department = category1.Title //分院名称
		// aa[0].Selectable = false
		category2, err := models.GetPids(category1.Id) //得到多个科室
		// beego.Info(category1[i1].Id)
		if err != nil {
			beego.Error(err)
		}
		for i2, _ := range category2 {
			bb := make([]Sidebar4, 1)
			bb[0].Id = category2[i2].Id
			bb[0].Level = "2"
			bb[0].Pid = category1.Id
			bb[0].Keshi = category2[i2].Title //科室名称
			// bb[0].Selectable = false
			//根据分院和科室查所有员工
			users, count, err := models.GetUsersbySec(category1.Title, category2[i2].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range users {
				cc := make([]Sidebar3, 1)
				cc[0].Id = users[i3].Id
				cc[0].Level = "3"
				cc[0].Pid = category2[i2].Id
				cc[0].Nickname = users[i3].Nickname //名称
				// cc[0].Selectable = false
				slice3 = append(slice3, cc...)
			}
			//64位转string
			// bb[0].Href = "/secofficeshow?secid=" + strconv.FormatInt(category2[i2].Id, 10) + "&depid=" + strconv.FormatInt(category1[i1].Id, 10) //得到Id用于添加成果 + " target='_blank'"
			bb[0].Tags[0] = count
			bb[0].Yuangong = slice3
			slice3 = make([]Sidebar3, 0) //再把slice置0
			slice4 = append(slice4, bb...)
		}
		aa[0].Bumen = slice4
		slice4 = make([]Sidebar4, 0) //再把slice置0
		slice5 = append(slice5, aa...)
		Sidebar7.Fenyuan = slice5
		slice5 = make([]Sidebar5, 0) //再把slice置0
	case 3: //主任登录显示的侧栏是本科室的所有人
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetCategorybyname(user.Department)
		if err != nil {
			beego.Error(err)
		}
		//由分院id和科室名称取得科室
		category2, err := models.GetCategorybyidtitle(category1.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]Sidebar5, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		aa[0].Pid = category[0].Id
		aa[0].Department = category1.Title //分院名称

		bb := make([]Sidebar4, 1)
		bb[0].Id = category2.Id
		bb[0].Level = "2"
		bb[0].Pid = category1.Id
		bb[0].Keshi = category2.Title //科室名称
		// bb[0].Selectable = false
		//根据分院和科室查所有员工
		users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			cc := make([]Sidebar3, 1)
			cc[0].Id = users[i3].Id
			cc[0].Level = "3"
			cc[0].Pid = category2.Id
			cc[0].Nickname = users[i3].Nickname //名称
			// cc[0].Selectable = false
			slice3 = append(slice3, cc...)
		}
		//64位转string
		// bb[0].Href = "/secofficeshow?secid=" + strconv.FormatInt(category2[i2].Id, 10) + "&depid=" + strconv.FormatInt(category1[i1].Id, 10) //得到Id用于添加成果 + " target='_blank'"
		bb[0].Tags[0] = count
		bb[0].Yuangong = slice3
		slice3 = make([]Sidebar3, 0) //再把slice置0
		slice4 = append(slice4, bb...)
		aa[0].Bumen = slice4
		slice4 = make([]Sidebar4, 0) //再把slice置0
		slice5 = append(slice5, aa...)
		Sidebar7.Fenyuan = slice5
		slice5 = make([]Sidebar5, 0) //再把slice置0
	case 4: //个人登录显示自己
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetCategorybyname(user.Department)
		if err != nil {
			beego.Error(err)
		}
		//由分院id和科室名称取得科室
		category2, err := models.GetCategorybyidtitle(category1.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]Sidebar5, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		aa[0].Pid = category[0].Id
		aa[0].Department = category1.Title //分院名称

		bb := make([]Sidebar4, 1)
		bb[0].Id = category2.Id
		bb[0].Level = "2"
		bb[0].Pid = category1.Id
		bb[0].Keshi = category2.Title //科室名称
		// bb[0].Selectable = false
		//根据分院和科室查所有员工
		// users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
		// if err != nil {
		// 	beego.Error(err)
		// }
		// for i3, _ := range users {
		cc := make([]Sidebar3, 1)
		cc[0].Id = user.Id
		cc[0].Level = "3"
		cc[0].Pid = category2.Id
		cc[0].Nickname = user.Nickname //名称
		// cc[0].Selectable = false
		slice3 = append(slice3, cc...)
		// }
		//64位转string
		// bb[0].Href = "/secofficeshow?secid=" + strconv.FormatInt(category2[i2].Id, 10) + "&depid=" + strconv.FormatInt(category1[i1].Id, 10) //得到Id用于添加成果 + " target='_blank'"
		bb[0].Tags[0] = 1
		bb[0].Yuangong = slice3
		slice3 = make([]Sidebar3, 0) //再把slice置0
		slice4 = append(slice4, bb...)
		aa[0].Bumen = slice4
		slice4 = make([]Sidebar4, 0) //再把slice置0
		slice5 = append(slice5, aa...)
		Sidebar7.Fenyuan = slice5
		slice5 = make([]Sidebar5, 0) //再把slice置0
	}
	// b, err := json.Marshal(Sidebar7) //不需要转成json格式
	c.Data["Input"] = Sidebar7 //这个没用吧
	// beego.Info(string(b))
	// fmt.Println(string(b))
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = Sidebar7
	// c.ServeJSON()
	c.TplName = "admin_achievement_show.tpl"
}

//上面那个是显示侧栏
//这个是显示右侧iframe框架内容——科室内人员成果情况统计
func (c *Achievement) Secofficeshow() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	//由uname取得user
	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
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
	if role > 4 { //
		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	if secid == "" {
		secid = strconv.FormatInt(user.Id, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	level := c.Input().Get("level")
	key := c.Input().Get("key")
	daterange := c.Input().Get("datefilter")
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	hours := 0
	var t1, t2 time.Time
	if daterange != "" {
		array := strings.Split(daterange, " - ")
		starttime1 := array[0]
		endtime1 := array[1]
		const lll = "2006-01-02"
		starttime, _ := time.Parse(lll, starttime1)
		endtime, _ := time.Parse(lll, endtime1)
		t1 = starttime.Add(-time.Duration(hours) * time.Hour)
		// beego.Info(t1)：2016-08-19 00:00:00 +0000 UTC
		t2 = endtime.Add(-time.Duration(hours) * time.Hour)
		beego.Info(t2)
	} else {
		t2 = time.Now()
		// beego.Info(t1):2016-08-19 23:27:29.7463081 +0800 CST
		// starttime, _ := time.Parse("2006-01-02", starttime1)
		t1 = t2.Add(-time.Duration(720) * time.Hour) //往前一个月时间
		// beego.Info(t2)
	}

	switch level {
	case "0": //如果是总院，则显示全部分院情况
		// c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["Starttime"] = t1
		c.Data["Endtime"] = t2
		c.TplName = "institute_show.tpl"
	case "1": //如果是分院，则显示全部科室
		categoryname, err := models.GetCategory(secid1)
		if err != nil {
			beego.Error(err)
		}
		//权限判断，并且属于这个分院
		if role == 1 || role == 2 && user.Department == categoryname.Title { //
			//根据分院id得到科室id
			//循环构造分院数据，view中进行循环显示各个科室情况
			Secofficevalue := make([]models.Secofficeachievement, 0)
			category, err := models.GetPids(secid1) //得到多个科室
			if err != nil {
				beego.Error(err)
			}
			employeevalue := make([]models.Employeeachievement, 0)
			for _, v1 := range category {
				aa := make([]models.Secofficeachievement, 1)
				//根据科室id查所有员工
				secid2 := strconv.FormatInt(v1.Id, 10)
				users, _, err := models.GetUsersbySecId(secid2) //得到员工姓名
				if err != nil {
					beego.Error(err)
				}
				for _, v := range users {
					//由username查出所有编制成果总分、设计总分……合计
					employee, err := models.Getemployeevalue(v.Nickname, t1, t2)
					if err != nil {
						beego.Error(err)
					}
					employeevalue = append(employeevalue, employee...)
				}
				aa[0].Id = v1.Id      //科室Id
				aa[0].Name = v1.Title //科室名称
				aa[0].Employee = employeevalue
				Secofficevalue = append(Secofficevalue, aa...)
				aa = make([]models.Secofficeachievement, 0) //再把slice置0
				employeevalue = make([]models.Employeeachievement, 0)
			}
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["Secoffice"] = Secofficevalue
			c.Data["Deptitle"] = categoryname.Title
			c.TplName = "depoffice_show.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	case "2": //如果是科室，则显示全部人员情况
		//取得科室名称
		categoryname, err := models.GetCategory(secid1)
		if err != nil {
			beego.Error(err)
		}
		// 取得分院名称
		categoryname1, err := models.GetCategory(categoryname.ParentId)
		if err != nil {
			beego.Error(err)
		}
		//1.进行权限读取,属于这个科室，或者属于这个分院
		if role == 1 || role == 3 && user.Secoffice == categoryname.Title || role == 2 && user.Department == categoryname1.Title { //
			employeevalue := make([]models.Employeeachievement, 0)
			// depid := c.Input().Get("depid")
			//根据分院和科室查所有员工
			// users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
			//根据科室id查所有员工
			users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
			beego.Info(users)
			if err != nil {
				beego.Error(err)
			}
			for _, v := range users {
				//由username查出所有编制成果总分、设计总分……合计
				employee, err := models.Getemployeevalue(v.Nickname, t1, t2)
				if err != nil {
					beego.Error(err)
				}
				employeevalue = append(employeevalue, employee...)
			}
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Sectitle"] = categoryname.Title
			c.Data["Level"] = level
			c.Data["Employee"] = employeevalue
			c.TplName = "secoffice_show.tpl"
		} else {
			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	case "3": //如果是个人，则显示个人详细情况
		//分2部分，一部分是已经完成状态的，state是4，另一部分是状态分别是3待审查通过,2，1的
		usernickname := models.GetUserByUserId(secid1)
		//1.进行权限读取，室主任以上并且属于这个科室，或者或本人
		if role == 1 || role == 3 && user.Secoffice == usernickname.Secoffice || role == 2 && user.Department == usernickname.Department || user.Nickname == usernickname.Nickname { //
			// employeecatalog := make([]models.Catalog, 0)
			//根据员工id和成果类型查出所有成果，设计成果，校核成果，审查成果
			//1、查图纸、报告……补充时间段secid即为userid
			//这里根据成果类型表循环查找
			//取得成果类型
			ratios, err := models.GetRatios()
			//下面这些没有用
			// for _, v := range category {
			// 	catalog, err := models.Getcatalogbyuserid(secid, v.Category, t1, t2)
			// }

			// catalogtuzhi, err := models.Getcatalogbyuserid(secid, "图纸", t1, t2)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// catalogbaogao, err := models.Getcatalogbyuserid(secid, "报告", t1, t2)

			// if err != nil {
			// 	beego.Error(err)
			// }
			// catalogjisuanshu, err := models.Getcatalogbyuserid(secid, "计算书", t1, t2)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// catalogxiugaidan, err := models.Getcatalogbyuserid(secid, "修改单", t1, t2)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// catalogdagang, err := models.Getcatalogbyuserid(secid, "大纲", t1, t2)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// catalogbiaoshu, err := models.Getcatalogbyuserid(secid, "标书", t1, t2)
			// if err != nil {
			// 	beego.Error(err)
			// }
			//根据userid得到user
			// Id, err := strconv.ParseInt(secid, 10, 64)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// user := models.GetUserByUserId(Id)
			//根据userid得到所有成果,时间段，在模板里，根据catalogs的类型与category匹配进行显示即可
			catalogs, err := models.Getcatalog2byuserid(secid, t1, t2)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Ratio"] = ratios //定义的成果类型
			//下面这个catalogs用于employee_show.tpl
			c.Data["Catalogs"] = catalogs
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["UserNickname"] = usernickname.Nickname
			//下面这个分类显示用于employee_show.tpl
			// c.Data["Catalogtuzhi"] = catalogtuzhi
			// c.Data["Catalogbaogao"] = catalogbaogao
			// c.Data["Catalogjisuanshu"] = catalogjisuanshu
			// c.Data["Catalogxiugaidan"] = catalogxiugaidan
			// c.Data["Catalogdagang"] = catalogdagang
			// c.Data["Catalogbiaoshu"] = catalogbiaoshu
			//如果在线编辑则显示self_show
			if key == "editor" { //新窗口显示添加页面
				c.TplName = "employeeself_show.tpl"
			} else if key == "modify" { //新窗口显示处理页面
				c.TplName = "employeeselfmodify_show.tpl"
			} else { //直接查看页面
				c.TplName = "employee_show.tpl"
			}
		} else {
			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
		//如果是本人，则显示可编辑状态
		// if uname == user.Username {
		// 	c.TplName = "employeeself_show.tpl"
		// } else {
		// 	c.TplName = "employee_show.tpl"
		// }
	default: //默认显示个人
		usernickname := models.GetUserByUserId(secid1)
		ratios, err := models.GetRatios()
		//根据userid得到所有成果,时间段，在模板里，根据catalogs的类型与category匹配进行显示即可
		catalogs, err := models.Getcatalog2byuserid(secid, t1, t2)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Starttime"] = t1
		c.Data["Endtime"] = t2
		c.Data["Ratio"] = ratios //定义的成果类型
		//下面这个catalogs用于employee_show.tpl
		c.Data["Catalogs"] = catalogs
		c.Data["Secid"] = secid
		c.Data["Level"] = "3"
		c.Data["UserNickname"] = usernickname.Nickname
		c.TplName = "employee_show.tpl"
	}

}

//用户登录后获得自己所在的分院和科室，然后显示对应的菜单
//同时显示所有的成果记录
func (c *Achievement) GetAchievementUser() {
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
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/login?url="+route, 302)
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
		//4.取出用户的权限等级
		role, _ := checkRole(c.Ctx) //login里的
		//5.进行逻辑分析：
		// rolename, err := strconv.ParseInt(role, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		if role > 5 { //
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
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
	c.Data["category"] = user
}

//上传excel文件，导入成果到数据库
func (c *Achievement) Import_Xls_Catalog() {
	// type Duration int64
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	// hours := 8
	// const lll = "2006/01/02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24

	//解析表单
	//获取上传的文件
	_, h, err := c.GetFile("catalog")
	if err != nil {
		beego.Error(err)
	}
	var path string
	if h != nil {
		//保存附件
		path = ".\\attachment\\" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("catalog", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	var uname string
	if v != nil {
		uname = v.(string)
	} else {
		beego.Error(err)
	}

	var catalog m.Catalog
	// id1 := c.Input().Get("id")
	// cid, _ := strconv.ParseInt(id1, 10, 64)
	// catalog.ParentId = cid
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	j := 0
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows { //行数,第一行从0开始
			if i != 0 { //忽略第一行标题
				// 1ProjectNumber string    //项目编号
				// 2ProjectName   string    //项目名称
				// 3DesignStage   string    //阶段
				// 4Section       string    //专业
				// 5Tnumber       string    //成果编号
				// 6Name          string    //成果名称
				// 7Drawn         string    //编制、绘制
				// 8Designd       string    //设计
				// 9Checked       string    //校核
				// 10Examined      string    //审查
				// 11Verified      string    //核定
				// 12Approved      string    //批准
				// 13Data          time.Time `orm:"null;auto_now_add;type(datetime)"`
				// 14Created       time.Time `orm:"index;auto_now_add;type(datetime)"`
				// 15Updated       time.Time `orm:"index;auto_now_add;type(datetime)"`
				// 16Author        string    //上传者
				if len(row.Cells) >= 2 { //总列数，从1开始
					catalog.ProjectNumber, err = row.Cells[j+1].String() //第一列从0开始,忽略第一列序号
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 3 {
					catalog.ProjectName, err = row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 4 {
					catalog.DesignStage, err = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 5 {
					catalog.Section, err = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 6 {
					catalog.Tnumber, err = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 7 {
					catalog.Name, err = row.Cells[j+6].String()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 8 {
					catalog.Category, err = row.Cells[j+7].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 9 {
					catalog.Page, err = row.Cells[j+8].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 10 {
					catalog.Count, err = row.Cells[j+9].Float()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 11 {
					catalog.Drawn, err = row.Cells[j+10].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 12 {
					catalog.Designd, err = row.Cells[j+11].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 13 {
					catalog.Checked, err = row.Cells[j+12].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 14 {
					catalog.Examined, err = row.Cells[j+13].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 15 {
					catalog.Verified, err = row.Cells[j+14].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 16 {
					catalog.Approved, err = row.Cells[j+15].String()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 17 {
					catalog.Complex, err = row.Cells[j+16].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 18 {
					catalog.Drawnratio, err = row.Cells[j+17].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 19 {
					catalog.Designdratio, err = row.Cells[j+18].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 20 {
					catalog.Checkedratio, err = row.Cells[j+19].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 21 {
					catalog.Examinedratio, err = row.Cells[j+20].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 22 {
					// endtime1, err := row.Cells[j+13].String()
					// beego.Info(endtime1)
					endtime2, err := row.Cells[j+21].Float()
					date := xlsx.TimeFromExcelTime(endtime2, false)
					// beego.Info(date)
					if err != nil {
						beego.Error(err)
					}
					// endtime, _ := time.Parse(lll, date)
					// t2 := endtime.Add(-time.Duration(hours) * time.Hour)
					catalog.Data = date //t2

				}

				catalog.Created = time.Now()
				catalog.Updated = time.Now()
				catalog.State = "4"
				catalog.Author = uname
				_, err := m.SaveCatalog(catalog)
				if err != nil {
					beego.Error(err)
				}
			}
		}
	}
	c.TplName = "catalog.tpl"
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "上传成果文件")
	logs.Close()
	c.Redirect("/admin", 302)
}

//在线添加目录，即插入一条目录,保存state=1；提交state=2
//任何人只能填写自己是设计和绘图或编制的成果，无权填写自己是校核或审查的成果
func (c *Achievement) AddCatalog() {
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// c.Data["Uname"] = uname

	// rolename, _ = strconv.Atoi(role)
	// 	if rolename > 2 {
	// 		port := strconv.Itoa(c.Ctx.Input.Port())
	// 			route = c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
	// 			c.Data["Url"] = route
	// 			c.Redirect("/roleerr?url="+route, 302)
	// 			return
	// 	}
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var catalog m.Catalog
	catalog.ProjectNumber = c.Input().Get("Pnumber")
	catalog.ProjectName = c.Input().Get("Pname")
	catalog.DesignStage = c.Input().Get("Stage")
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Category = c.Input().Get("Category")
	catalog.Page = c.Input().Get("Page")

	count1 := c.Input().Get("Count")
	if count1 != "" {
		catalog.Count, err = strconv.ParseFloat(count1, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Examined = c.Input().Get("Examined")
	// catalog.Verified = c.Input().Get("Verified")
	// catalog.Approved = c.Input().Get("Approved")

	// complex, err := strconv.ParseFloat(c.Input().Get("Complex"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Complex = complex
	drawnratio1 := c.Input().Get("Drawnratio")
	if drawnratio1 != "" {
		catalog.Drawnratio, err = strconv.ParseFloat(drawnratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	// designdratio, err := strconv.ParseFloat(c.Input().Get("Designdratio"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Designdratio = designdratio
	// checkedratio, err := strconv.ParseFloat(c.Input().Get("Checkedratio"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Checkedratio = checkedratio
	// examinedratio, err := strconv.ParseFloat(c.Input().Get("Examinedratio"), 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// catalog.Examinedratio = examinedratio

	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	printtime, err := time.Parse(lll, c.Input().Get("Data"))
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(printtime)
	catalog.Data = printtime
	catalog.Created = time.Now()
	catalog.Updated = time.Now()
	catalog.State = "1"
	catalog.Author = uname
	//只能添加自己是设计者或绘图者的成果
	if user.Nickname == catalog.Drawn || user.Nickname == catalog.Designd {
		cid, err := m.SaveCatalog(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "添加记录" + strconv.FormatInt(cid, 10))
		logs.Close()
	}
	// err := models.ModifyCatalog(tid, title, tnumber)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Redirect("/catalog", 302)
}

//直接添加提交，状态加1——这个也要改成添加，然后保存，然后再提交。不要一步提交
func (c *Achievement) AddSendCatalog() {
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// c.Data["Uname"] = uname

	// rolename, _ = strconv.Atoi(role)
	// 	if rolename > 2 {
	// 		port := strconv.Itoa(c.Ctx.Input.Port())
	// 			route = c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
	// 			c.Data["Url"] = route
	// 			c.Redirect("/roleerr?url="+route, 302)
	// 			return
	// 	}
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}

	var catalog m.Catalog
	catalog.ProjectNumber = c.Input().Get("Pnumber")
	catalog.ProjectName = c.Input().Get("Pname")
	catalog.DesignStage = c.Input().Get("Stage")
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Category = c.Input().Get("Category")
	catalog.Page = c.Input().Get("Page")
	count, err := strconv.ParseFloat(c.Input().Get("Count"), 64)
	if err != nil {
		beego.Error(err)
	}
	catalog.Count = count

	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Examined = c.Input().Get("Examined")

	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	printtime, _ := time.Parse(lll, c.Input().Get("Data"))
	catalog.Data = printtime
	catalog.Created = time.Now()
	catalog.Updated = time.Now()
	catalog.State = "2"
	catalog.Author = uname
	//只能添加自己是设计者或绘图者的成果
	if uname == catalog.Drawn || uname == catalog.Designd {
		cid, err := m.SaveCatalog(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "添加记录" + strconv.FormatInt(cid, 10))
		logs.Close()
	}
	// err := models.ModifyCatalog(tid, title, tnumber)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Redirect("/catalog", 302)
}

//修改保存一条目录，状态不改变
//自己只能修改状态为1的，即未提交的
//其他人只能修改自己是校核或审查的
func (c *Achievement) ModifyCatalog() {
	//4.取得客户端用户名——是简写，不是昵称
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(user.Nickname)
	var catalog m.Catalog
	catalog.ProjectNumber = c.Input().Get("Pnumber")
	catalog.ProjectName = c.Input().Get("Pname")
	catalog.DesignStage = c.Input().Get("Stage")
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Category = c.Input().Get("Category")
	catalog.Page = c.Input().Get("Page")
	count1 := c.Input().Get("Count")
	if count1 != "" {
		catalog.Count, err = strconv.ParseFloat(count1, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Examined = c.Input().Get("Examined")
	complex1 := c.Input().Get("Complex")
	if complex1 != "" {
		catalog.Complex, err = strconv.ParseFloat(complex1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	drawnratio1 := c.Input().Get("Drawnratio")
	if drawnratio1 != "" {
		catalog.Drawnratio, err = strconv.ParseFloat(drawnratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	designdratio1 := c.Input().Get("Designdratio")
	if designdratio1 != "" {
		catalog.Designdratio, err = strconv.ParseFloat(designdratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	checkratio1 := c.Input().Get("Checkedratio")
	if checkratio1 != "" {
		catalog.Checkedratio, err = strconv.ParseFloat(checkratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	examinedratio1 := c.Input().Get("Examinedratio")
	if examinedratio1 != "" {
		catalog.Examinedratio, err = strconv.ParseFloat(examinedratio1, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	printtime, err := time.Parse(lll, c.Input().Get("Data"))
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(printtime)
	catalog.Data = printtime
	catalog.Updated = time.Now()
	catalog.Author = uname
	cid := c.Input().Get("CatalogId")
	// beego.Info(cid)
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		// beego.Info(id)
	}
	cidNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	if user.Nickname == catalog.Drawn || user.Nickname == catalog.Designd {
		err = m.ModifyCatalog(cidNum, catalog, "1") //只能修改状态为1的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + id)
		logs.Close()
	} else if user.Nickname == catalog.Checked {
		err = m.ModifyCatalog(cidNum, catalog, "2") //只能修改状态为2的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改保存校核记录" + id)
		logs.Close()
	} else if user.Nickname == catalog.Examined {
		beego.Info(catalog.Examined)
		err = m.ModifyCatalog(cidNum, catalog, "3") //只能修改状态为3的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改保存审查记录" + id)
		logs.Close()
	}
}

//提交一条目录，即状态加1而已——这个没用，修改不要直接提交，要先保存，再提交
//自己只能提交状态为1的，即未提交的
//其他人只能提交自己是校核或审查的
func (c *Achievement) ModifySendCatalog() {
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var catalog m.Catalog
	catalog.ProjectNumber = c.Input().Get("Pnumber")
	catalog.ProjectName = c.Input().Get("Pname")
	catalog.DesignStage = c.Input().Get("Stage")
	catalog.Tnumber = c.Input().Get("Tnumber")
	catalog.Name = c.Input().Get("Name")
	catalog.Category = c.Input().Get("Category")
	catalog.Page = c.Input().Get("Page")
	count, err := strconv.ParseFloat(c.Input().Get("Count"), 64)
	if err != nil {
		beego.Error(err)
	}
	catalog.Count = count
	catalog.Drawn = c.Input().Get("Drawn")
	catalog.Designd = c.Input().Get("Designd")
	catalog.Checked = c.Input().Get("Checked")
	catalog.Examined = c.Input().Get("Examined")
	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
	printtime, err := time.Parse(lll, c.Input().Get("Data"))
	if err != nil {
		beego.Error(err)
	}
	beego.Info(printtime)
	catalog.Data = printtime
	catalog.Updated = time.Now()
	catalog.Author = uname
	cid := c.Input().Get("CatalogId")
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		// beego.Info(id)
	}
	cidNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	if user.Nickname == catalog.Drawn || user.Nickname == catalog.Designd {
		err = m.ModifyCatalog(cidNum, catalog, "2") //只能提交状态为1的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改记录" + id)
		logs.Close()
	} else if user.Nickname == catalog.Checked {
		err = m.ModifyCatalog(cidNum, catalog, "3") //只能提交状态为2的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改提交记录" + id)
		logs.Close()
	} else if user.Nickname == catalog.Examined {
		err = m.ModifyCatalog(cidNum, catalog, "4") //只能提交状态为3的
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}

		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改记录" + id)
		logs.Close()
	}
}

//直接提交一条目录，即状态加1
//如果提交的内容没有审查人员，则直接由校核人员提交，直接改为4
//比如会务等，交组长确认，或项目负责人、专业负责人，只有一个人确认即可
func (c *Achievement) SendCatalog() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//取得用户名

	// if rolename > 2 && uname != username {
	cid := c.Input().Get("CatalogId")
	var id string
	var cidNum int64
	var err error
	if cid != "" {
		id = string(cid[3:len(cid)])
		// beego.Info(id)
		cidNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	//查出cidnum这个有没有审查人员，没有就直接state+2
	catalog, err := m.GetCatalog(id)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(catalog.State)
	// beego.Info(catalog.Examined)
	if catalog.Examined == "" && catalog.State == "2" { //catalog.Examined == "" && catalog.State == "1" ||
		err = m.ModifyCatalogState(cidNum, "2") //记得填上难度系数为1
		if err != nil {
			beego.Error(err)
		} else {
			data := "提交汇总ok!"
			c.Ctx.WriteString(data)
		}
	} else {
		err = m.ModifyCatalogState(cidNum, "1")
		if err != nil {
			beego.Error(err)
		} else {
			data := "提交下一级ok!"
			c.Ctx.WriteString(data)
		}
	}
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "提交记录" + cid)
	logs.Close()
}

//退回一条目录，即状态减1
func (c *Achievement) DownSendCatalog() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//取得用户名

	// if rolename > 2 && uname != username {
	cid := c.Input().Get("CatalogId")
	var id string
	var cidNum int64
	var err error
	if cid != "" {
		id = string(cid[3:len(cid)])
		cidNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(id)
	}
	err = m.ModifyCatalogState(cidNum, "-1")
	if err != nil {
		beego.Error(err)
	} else {
		data := "退回前一级ok!"
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "退回记录" + cid)
	logs.Close()
}

//删除一条目录
func (c *Achievement) DeleteCatalog() {
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, _, _ := checkRolewrite(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//取得用户名

	// if rolename > 2 && uname != username {
	cid := c.Input().Get("CatalogId")
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		beego.Info(id)
	}
	err := m.DeletCatalog(id)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "添加记录" + cid)
	logs.Close()
}

//进入编辑ratio页面
func (c *Achievement) Ratio() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
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
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	//5.进行逻辑分析：
	// rolename, err := strconv.ParseInt(role, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	if role > 2 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	ratio, _ := m.GetRatios()
	c.Data["Ratio"] = ratio
	c.TplName = "ratio.tpl"
}
func (c *Achievement) AddRatio() {
	var ratio m.Ratio
	ratio.Category = c.Input().Get("Category")
	ratio.Unit = c.Input().Get("Unit")
	ratio1, err := strconv.ParseFloat(c.Input().Get("Ratio"), 64)
	if err != nil {
		beego.Error(err)
	}
	ratio.Rationum = ratio1
	ratio.Created = time.Now()
	ratio.Updated = time.Now()
	cid, err := m.SaveRatio(ratio)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok"
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "添加ratio" + strconv.FormatInt(cid, 10))
	logs.Close()
}
func (c *Achievement) ModifyRatio() {
	var ratio m.Ratio
	ratio.Category = c.Input().Get("Category")
	ratio.Unit = c.Input().Get("Unit")
	ratio1, err := strconv.ParseFloat(c.Input().Get("Ratio"), 64)
	if err != nil {
		beego.Error(err)
	}
	ratio.Rationum = ratio1
	// ratio.Created = time.Now()
	ratio.Updated = time.Now()

	cid := c.Input().Get("RatioId")
	var id string
	if cid != "" {
		id = string(cid[3:len(cid)])
		// beego.Info(id)
	}
	err = m.ModifyRatio(id, ratio)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改ratio" + id)
	logs.Close()
}

//参考时间的转换：由string转成time data
// func (c *TaskController) AddTask() {
// 	//解析表单     表示时间的变量和字段，应为time.Time类型
// 	type Duration int64
// 	const (
// 		Nanosecond  Duration = 1
// 		Microsecond          = 1000 * Nanosecond
// 		Millisecond          = 1000 * Microsecond
// 		Second               = 1000 * Millisecond
// 		Minute               = 60 * Second
// 		Hour                 = 60 * Minute
// 	)
// 	//seconds := 10
// 	//fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
// 	hours := 8
// 	//	time.Duration(hours) * time.Hour
// 	// t1 := t.Add(time.Duration(hours) * time.Hour)
// 	// datestring = t1.Format(layout) //{{dateformat .Created "2006-01-02 15:04:05"}}
// 	// return

// 	var err error
// 	tid := c.Input().Get("tid")
// 	title := c.Input().Get("title")
// 	content := c.Input().Get("content")
// 	daterange := c.Input().Get("datefilter")
// 	array := strings.Split(daterange, " - ")
// 	// for _, v := range array {
// 	starttime1 := array[0]
// 	beego.Info(array[0])
// 	endtime1 := array[1]
// 	beego.Info(array[1])
// 	// }
// 	// starttime1 := c.Input().Get("starttime")
// 	// endtime1 := c.Input().Get("endtime")
// 	const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
// 	starttime, _ := time.Parse(lll, starttime1)
// 	endtime, _ := time.Parse(lll, endtime1)
// 	t1 := starttime.Add(-time.Duration(hours) * time.Hour)
// 	t2 := endtime.Add(-time.Duration(hours) * time.Hour)
// 	//12-19-2015 22:40:24
// 	// ck, err := c.Ctx.Request.Cookie("uname")
// 	// if err != nil {
// 	// beego.Error(err)
// 	// }
// 	// uname := ck.Value
// 	if len(tid) == 0 {
// 		err = models.AddTask(title, content, t1, t2)
// 		// beego.Info(attachment)
// 		// } else {
// 		// 	err = models.UpdateTask(tid, title, content)
// 	}
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Redirect("/todo", 302)
// }
