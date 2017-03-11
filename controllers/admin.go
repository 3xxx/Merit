package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
	"merit/models"
	"net"
	"strconv"
	"strings"
)

type AdminController struct {
	beego.Controller
}

//（2）建立一个全局session mananger对象
var globalSessions *session.Manager

//（3）在初始化“全局session mananger对象”
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 36000, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 36000, "providerConfig": ""}`)
	go globalSessions.GC()
	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	// go globalSessions.GC()
}

// func (c *AdminController) Get() {
// 	// c.Data["Website"] = "beego.me"
// 	// c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (c *AdminController) Get() {
	role := Getiprole(c.Ctx.Input.IP())
	if role == 1 {
		c.TplName = "admin.tpl"
	} else {
		c.Data["json"] = "权限不够！"
		c.ServeJSON()
	}
	c.Data["Ip"] = c.Ctx.Input.IP()
}

func (c *AdminController) Admin() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["Id"] = id
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// //1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	return
	// }
	// //4.取得客户端用户名
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// if v != nil {
	// 	c.Data["Uname"] = v.(string)
	// }
	// //4.取出用户的权限等级
	// role, err := checkRole(c.Ctx) //login里的
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	//5.进行逻辑分析：
	// 	if role > 2 { //
	// 		route := c.Ctx.Request.URL.String()
	// 		c.Data["Url"] = route
	// 		c.Redirect("/roleerr?url="+route, 302)
	// 		return
	// 	}
	// }
	switch id {
	case "010": //日历实物
		c.TplName = "admin_calendar.tpl"
	case "011": //基本设置
		c.TplName = "admin_base.tpl"
	case "012": //组织
		c.TplName = "admin_department.tpl"
	case "013": //定义价值
		c.TplName = "admin_merit.tpl"
	case "014": //科室价值
		achsecoffice := make([]AchSecoffice, 0)
		achdepart := make([]AchDepart, 0)
		category1, err := models.GetAdminDepart(0) //得到多个分院
		if err != nil {
			beego.Error(err)
		}
		for i1, _ := range category1 {
			aa := make([]AchDepart, 1)
			aa[0].Id = category1[i1].Id
			aa[0].Level = "1"
			// aa[0].Pid = category[0].Id
			aa[0].Title = category1[i1].Title //分院名称
			// beego.Info(category1[i1].Title)
			category2, err := models.GetAdminDepart(category1[i1].Id) //得到多个科室
			if err != nil {
				beego.Error(err)
			}
			//如果返回科室为空，则直接取得员工
			//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
			//应该是反过来找出所有没有科室字段的人员，把他放在部门下
			if len(category2) > 0 {
				for i2, _ := range category2 {
					bb := make([]AchSecoffice, 1)
					bb[0].Id = category2[i2].Id
					bb[0].Level = "2"
					bb[0].Pid = category1[i1].Id
					bb[0].Title = category2[i2].Title //科室名称
					// beego.Info(category2[i2].Title)
					//根据分院和科室查所有员工
					// users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
					// if err != nil {
					// 	beego.Error(err)
					// }
					// for i3, _ := range users {
					// 	cc := make([]AchEmployee, 1)
					// 	cc[0].Id = users[i3].Id
					// 	cc[0].Level = "3"
					// 	cc[0].Pid = category2[i2].Id
					// 	cc[0].Nickname = users[i3].Nickname //名称
					// 	// beego.Info(users[i3].Nickname)
					// 	// cc[0].Selectable = false
					// 	achemployee = append(achemployee, cc...)
					// }
					// bb[0].Tags[0] = strconv.Itoa(count)
					// bb[0].Employee = achemployee
					bb[0].Selectable = true
					// achemployee = make([]AchEmployee, 0) //再把slice置0
					achsecoffice = append(achsecoffice, bb...)
					// depcount = depcount + count //部门人员数=科室人员数相加
				}
			}
			//查出所有有这个部门但科室名为空的人员
			//根据分院查所有员工
			// beego.Info(category1[i1].Title)
			// users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
			// if err != nil {
			// 	beego.Error(err)
			// }
			// beego.Info(users)
			// for i3, _ := range users {
			// 	dd := make([]AchSecoffice, 1)
			// 	dd[0].Id = users[i3].Id
			// 	dd[0].Level = "3"
			// 	// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
			// 	dd[0].Pid = category1[i1].Id
			// 	dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
			// 	dd[0].Selectable = true
			// 	achsecoffice = append(achsecoffice, dd...)
			// }
			// aa[0].Tags[0] = count + depcount
			// count = 0
			// depcount = 0
			aa[0].Secoffice = achsecoffice
			aa[0].Selectable = true                //默认是false点击展开
			achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			achdepart = append(achdepart, aa...)
		}
		c.Data["json"] = achdepart
		c.TplName = "admin_secofficemerit.tpl"
	case "015": //成果类型
		c.TplName = "admin_achievcategory.tpl"
	case "016": //科室成果类型
		c.TplName = "admin_departachievcate.tpl"
	case "021": //系统权限
		c.TplName = "admin_systemrole.tpl"
	case "022": //项目权限
		c.TplName = "admin_projectrole.tpl"
	case "031": //用户
		c.TplName = "admin_users.tpl"
	case "032": //IP地址段
		c.TplName = "admin_ipsegment.tpl"
	case "033": //用户组
		c.TplName = "admin_usergroup.tpl"
	case "041": //本周成果编辑
		c.TplName = "admin_achievementseditor.tpl"
	case "042": //本月成果编辑
		c.TplName = "admin_projectsynch.tpl"
	case "043": //上月成果编辑
		c.TplName = "admin_projectcaterole.tpl"
	case "044": //当年成果编辑
		c.TplName = "admin_projectcaterole.tpl"
	default:
		c.TplName = "admin_calendar.tpl"
	}
}

//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Department() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminDepartment
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	categories, err := models.GetAdminDepart(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据名称title查询分级表
func (c *AdminController) DepartmentTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminDepartTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//添加
func (c *AdminController) AddDepartment() {
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}

	_, err = models.AddAdminDepart(pidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdateDepartment() {
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateAdminDepart(cidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteDepartment() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminDepart(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//**********价值***********
//取得所有价值分类，或没有下级的价值
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Merit() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	merits, err := models.GetAdminMerit(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = merits
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据科室id得到价值分类，填充table
func (c *AdminController) SecofficeMerit() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	merits, err := models.GetSecofficeMerit(idNum)
	if err != nil {
		beego.Error(err)
	}
	// meritcate := make([]*models.AdminMerit, 0)
	meritcate := make([]MeritCategory, 0)

	allmerits, err := models.GetAdminMerit(0)
	// beego.Info(allmerits)
	if err != nil {
		beego.Error(err)
	}
	var level string

	level = "2"
	for _, v1 := range allmerits {
		for _, v2 := range merits {
			if v2.MeritId == v1.Id {
				beego.Info(v2.MeritId)
				level = "1"
				// merittitle, err := models.GetAdminMeritbyId(v2.MeritId) //因为这个数据库只是科室和分类的对应表
				// if err != nil {
				// 	beego.Error(err)
				// }
				// aa := make([]MeritCategory, 1)
				// aa[0].Id = merittitle.Id
				// aa[0].Title = merittitle.Title
				// aa[0].Level = "1"
				// meritcate = append(meritcate, aa...)
			}
		}
		aa := make([]MeritCategory, 1)
		aa[0].Id = v1.Id
		aa[0].Title = v1.Title
		aa[0].Level = level
		meritcate = append(meritcate, aa...)
		aa = make([]MeritCategory, 0)
		level = "2"
	}
	c.Data["json"] = meritcate
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//向科室id里添加价值分类
func (c *AdminController) AddSecofficeMerit() {
	sid := c.GetString("sid") //secofficeid
	//id转成64位
	sidNum, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取出所有sidnum的merit
	merits, err := models.GetSecofficeMerit(sidNum)
	if err != nil {
		beego.Error(err)
	}

	ids := c.GetString("ids") //meritid
	array := strings.Split(ids, ",")
	bool := false
	for _, v1 := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v1, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		for _, v2 := range merits {
			//没有找到则插入
			if v2.MeritId == idNum {
				bool = true
			}
		}
		if bool == false {
			//存入数据库
			err = models.AddSecofficeMerit(sidNum, idNum)
			if err != nil {
				beego.Error(err)
			}
		}
		bool = false
	}

	for _, v3 := range merits {
		for _, v4 := range array {
			//id转成64位
			idNum, err := strconv.ParseInt(v4, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//没有找到则删除
			if v3.MeritId == idNum {
				bool = true
			}
		}

		if bool == false {
			//存入数据库
			err = models.DeleteSecofficeMerit(sidNum, v3.MeritId)
			if err != nil {
				beego.Error(err)
			}
		}
		bool = false
	}
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

// func (c *AdminController) MeritCategory() {
// 	c.Data["Ip"] = c.Ctx.Input.IP()
// 	merits, err := models.GetAdminMeritCategory()
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["json"] = merits
// 	c.ServeJSON()
// }

// //根据价值分类，取得价值
// func (c *AdminController) MeritList() {
// 	c.Data["Ip"] = c.Ctx.Input.IP()
// 	merits, err := models.GetAdminMeritList()
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["json"] = merits
// 	c.ServeJSON()
// }

//添加价值结构
func (c *AdminController) AddMerit() {
	pid := c.Input().Get("pid")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}
	title := c.Input().Get("title")
	mark := c.Input().Get("mark")
	list := c.Input().Get("list")
	listmark := c.Input().Get("listmark")

	//存入数据库
	_, err = models.AddAdminMerit(pidNum, title, mark, list, listmark)
	if err != nil {
		beego.Error(err)
	} else {
		data := title
		c.Ctx.WriteString(data)
	}
}

//修改
func (c *AdminController) UpdateMerit() {
	title := c.Input().Get("title")
	mark := c.Input().Get("mark")
	list := c.Input().Get("list")
	listmark := c.Input().Get("listmark")
	cid := c.Input().Get("cid")
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminMerit(cidNum, title, mark, list, listmark)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
		logs := logs.NewLogger(1000)
		logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.Info(c.Ctx.Input.IP() + " " + "修改ratio" + cid)
		logs.Close()
	}
}

//删除
func (c *AdminController) DeleteMerit() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminMerit(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminMerit(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminMerit(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
			logs := logs.NewLogger(1000)
			logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
			logs.EnableFuncCallDepth(true)
			logs.Info(c.Ctx.Input.IP() + " " + "删除价值" + ids)
			logs.Close()
		}
	}
}

// func (c *AdminController) Admin() {
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	//1.首先判断是否注册
// 	if !checkAccount(c.Ctx) {
// 		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/login?url="+route, 302)
// 		// c.Redirect("/login", 302)
// 		return
// 	}
// 	//4.取得客户端用户名
// 	// var uname string
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		// uname = v.(string)
// 		c.Data["Uname"] = v.(string)
// 	}
// 	// uname := v.(string) //ck.Value
// 	//4.取出用户的权限等级
// 	role, err := checkRole(c.Ctx) //login里的
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		// beego.Info(role)
// 		//5.进行逻辑分析：
// 		// rolename, _ := strconv.ParseInt(role, 10, 64)
// 		if role > 2 { //
// 			// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 			route := c.Ctx.Request.URL.String()
// 			c.Data["Url"] = route
// 			c.Redirect("/roleerr?url="+route, 302)
// 			// c.Redirect("/roleerr", 302)
// 			return
// 		}
// 	}
// 	// c.Data["Website"] = "beego.me"
// 	// c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "admin.tpl"
// }

func (c *AdminController) Test() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user_show.tpl"
	c.Data["json"] = map[string]interface{}{
		"id":    2,
		"name":  "111",
		"price": "demo.jpg",
	}
	// c.Data["json"] = catalogs
	c.ServeJSON()
}

func (c *AdminController) Test1() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.tpl"
}
func (c *AdminController) Jsoneditor() {
	c.TplName = "jsoneditor.tpl"
}

//添加ip地址段
func (c *AdminController) AddIpsegment() {
	// pid := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminIpsegment(title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//修改ip地址段
func (c *AdminController) UpdateIpsegment() {
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminIpsegment(cidNum, title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//删除
func (c *AdminController) DeleteIpsegment() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteAdminIpsegment(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//查询IP地址段
func (c *AdminController) Ipsegment() {
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = ipsegments
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
// 2 每个包可以拥有多个init函数
// 3 包的每个源文件也可以拥有多个init函数
// 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
// 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
// 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
//读取iprole.txt文件，作为全局变量Iprolemaps，供调用访问者ip的权限用
var (
	Iprolemaps map[string]int
)

func init() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

func Createip() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

//取得访问者的权限
func Getiprole(ip string) (role int) {
	role = Iprolemaps[ip]
	return role
}

//获取下一个IP
func nextIp(ip string) string {
	ips := strings.Split(ip, ".")
	var i int
	for i = len(ips) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(ips[i])
		if n >= 255 {
			//进位
			ips[i] = "1"
		} else {
			//+1
			n++
			ips[i] = strconv.Itoa(n)
			break
		}
	}
	if i == -1 {
		//全部IP段都进行了进位,说明此IP本身已超出范围
		return ""
	}
	ip = ""
	leng := len(ips)
	for i := 0; i < leng; i++ {
		if i == leng-1 {
			ip += ips[i]
		} else {
			ip += ips[i] + "."
		}
	}
	return ip
}

//生成IP地址列表
func processIp(startIp, endIp string) []string {
	var ips = make([]string, 0)
	for ; startIp != endIp; startIp = nextIp(startIp) {
		if startIp != "" {
			ips = append(ips, startIp)
		}
	}
	ips = append(ips, startIp)
	return ips
}

//port代替权限role
func processFlag(arg []string) (maps map[string]int) {
	//开始IP,结束IP
	var startIp, endIp string
	//端口
	var ports []int = make([]int, 0)
	index := 0
	startIp = arg[index]
	si := net.ParseIP(startIp)
	if si == nil {
		//开始IP不合法
		// fmt.Println("'startIp' Setting error")
		beego.Error("开始IP不合法")
		return nil
	}
	index++
	endIp = arg[index]
	ei := net.ParseIP(endIp)
	if ei == nil {
		//未指定结束IP,即只扫描一个IP
		endIp = startIp
	} else {
		index++
	}

	tmpPort := arg[index]
	if strings.Index(tmpPort, "-") != -1 {
		//连续端口
		tmpPorts := strings.Split(tmpPort, "-")
		var startPort, endPort int
		var err error
		startPort, err = strconv.Atoi(tmpPorts[0])
		if err != nil || startPort < 1 || startPort > 65535 {
			//开始端口不合法
			return nil
		}
		if len(tmpPorts) >= 2 {
			//指定结束端口
			endPort, err = strconv.Atoi(tmpPorts[1])
			if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
				//结束端口不合法
				// fmt.Println("'endPort' Setting error")
				beego.Error("'endPort' Setting error")
				return nil
			}
		} else {
			//未指定结束端口
			endPort = 65535
		}
		for i := 0; startPort+i <= endPort; i++ {
			ports = append(ports, startPort+i)
		}
	} else {
		//一个或多个端口
		ps := strings.Split(tmpPort, ",")
		for i := 0; i < len(ps); i++ {
			p, err := strconv.Atoi(ps[i])
			if err != nil {
				//端口不合法
				// fmt.Println("'port' Setting error")
				beego.Error("'port' Setting error")
				return nil
			}
			ports = append(ports, p)
		}
	}

	//生成扫描地址列表
	ips := processIp(startIp, endIp)
	il := len(ips)
	m1 := make(map[string]int)
	for i := 0; i < il; i++ {
		pl := len(ports)
		for j := 0; j < pl; j++ {
			//			ipAddrs <- ips[i] + ":" + strconv.Itoa(ports[j])
			//			ipAddrs := ips[i] + ":" + strconv.Itoa(ports[j])
			m1[ips[i]] = ports[j]
		}
	}
	//	fmt.Print(slice1)
	return m1
	//	close(ipAddrs)
}
