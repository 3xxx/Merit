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
	"merit/models"
	"strconv"
	// "strings"
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
	rolename, _ := strconv.ParseInt(role, 10, 64)
	if rolename > 2 { //
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

	// b, err := json.Marshal(Sidebar7) //不需要转成json格式
	c.Data["Input"] = Sidebar7
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
	//水工室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	level := c.Input().Get("level")
	switch level {
	case "0": //如果是总院，则显示全部分院情况
	case "1": //如果是分院，则显示全部科室
	case "2": //如果是科室，则显示全部人员情况
		employeevalue := make([]models.Employeeachievement, 0)
		// depid := c.Input().Get("depid")
		//根据科室id查所有员工
		users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for _, v := range users {
			//由username查出所有编制成果总分、设计总分……合计
			// beego.Info(v.Username)
			employee, err := models.Getemployeevalue(v.Nickname)
			if err != nil {
				beego.Error(err)
			}
			employeevalue = append(employeevalue, employee...)
		}
		user1 := models.GetUserByUsername("蔡碧红")
		beego.Info(user1.Id)
		c.Data["Employee"] = employeevalue
		c.TplName = "secoffice_show.tpl"
	case "3": //如果是个人，则显示个人详细情况
		// employeecatalog := make([]models.Catalog, 0)
		//根据员工id和成果类型查出所有成果，设计成果，校核成果，审查成果
		//1、查图纸
		catalogtuzhi, err := models.Getcatalogbyuserid(secid, "图纸")
		if err != nil {
			beego.Error(err)
		}
		catalogbaogao, err := models.Getcatalogbyuserid(secid, "报告")

		if err != nil {
			beego.Error(err)
		}
		catalogjisuanshu, err := models.Getcatalogbyuserid(secid, "计算书")
		if err != nil {
			beego.Error(err)
		}
		catalogxiugaidan, err := models.Getcatalogbyuserid(secid, "修改单")
		if err != nil {
			beego.Error(err)
		}
		catalogdagang, err := models.Getcatalogbyuserid(secid, "大纲")
		if err != nil {
			beego.Error(err)
		}
		catalogbiaoshu, err := models.Getcatalogbyuserid(secid, "标书")
		if err != nil {
			beego.Error(err)
		}
		c.Data["Catalogtuzhi"] = catalogtuzhi
		c.Data["Catalogbaogao"] = catalogbaogao
		c.Data["Catalogjisuanshu"] = catalogjisuanshu
		c.Data["Catalogxiugaidan"] = catalogxiugaidan
		c.Data["Catalogdagang"] = catalogdagang
		c.Data["Catalogbiaoshu"] = catalogbiaoshu
		c.TplName = "employee_show.tpl"
	default: //默认显示全部人员情况
		employeevalue := make([]models.Employeeachievement, 0)
		// depid := c.Input().Get("depid")
		//根据科室id查所有员工
		users, _, err := models.GetUsersbySecId("3") //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for _, v := range users {
			//由username查出所有编制成果总分、设计总分……合计
			// beego.Info(v.Username)
			employee, err := models.Getemployeevalue(v.Nickname)
			if err != nil {
				beego.Error(err)
			}
			employeevalue = append(employeevalue, employee...)
		}
		c.Data["Employee"] = employeevalue
		c.TplName = "secoffice_show.tpl"
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
		rolename, err := strconv.ParseInt(role, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		if rolename > 5 { //
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
		user = models.GetUserByUsername(uname) //得到用户的id、分院和科室等
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
				if len(row.Cells) >= 1 { //总列数，从1开始
					catalog.ProjectNumber, err = row.Cells[j+1].String() //第一列从0开始,忽略第一列序号
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 2 {
					catalog.ProjectName, err = row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 3 {
					catalog.DesignStage, err = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 4 {
					catalog.Section, err = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 5 {
					catalog.Tnumber, err = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 6 {
					catalog.Name, err = row.Cells[j+6].String()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 7 {
					catalog.Category, err = row.Cells[j+7].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 8 {
					catalog.Page, err = row.Cells[j+8].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 9 {
					catalog.Count, err = row.Cells[j+9].Float()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 10 {
					catalog.Drawn, err = row.Cells[j+10].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 11 {
					catalog.Designd, err = row.Cells[j+11].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 12 {
					catalog.Checked, err = row.Cells[j+12].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 13 {
					catalog.Examined, err = row.Cells[j+13].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 14 {
					catalog.Verified, err = row.Cells[j+14].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 15 {
					catalog.Approved, err = row.Cells[j+15].String()
					if err != nil {
						beego.Error(err)
					}
				}

				if len(row.Cells) >= 16 {
					catalog.Complex, err = row.Cells[j+16].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 17 {
					catalog.Drawnratio, err = row.Cells[j+17].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 18 {
					catalog.Designdratio, err = row.Cells[j+18].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 19 {
					catalog.Checkedratio, err = row.Cells[j+19].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 20 {
					catalog.Examinedratio, err = row.Cells[j+20].Float()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 21 {
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

				catalog.Author = uname
				_, err := m.SaveCatalog(catalog)
				if err != nil {
					beego.Error(err)
				}
			}
		}
	}

	c.TplName = "catalog.tpl"
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
