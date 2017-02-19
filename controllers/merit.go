// 注意：在Go的标准库encoding/json包中，允许使用
// map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
//在线价值登记
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	"merit/models"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MeritController struct {
	beego.Controller
}

type MeritMark struct {
	Choose string `json:"choose"`
	Mark1  string `json:"mark1"` //打分1
}

type MeritList struct { //项目负责人——链接——大、中、小
	Id    int64  `json:"Id"`
	Pid   int64  `form:"-"`
	Title string `json:"text"`
	// Href    string `json:"href"`
	Tags     [2]int `json:"tags"`
	Mark     string `json:"mark2"` //打分2
	List     string //大型、中型……
	ListMark string //对应列表打分
	Level    string `json:"Level"` //4
}

type MeritSecoffice struct { //专业室：水工、施工……
	Id    int64  `json:"Id"`
	Pid   int64  `form:"-"`
	Title string `json:"text"`
	// Selectable bool    `json:"selectable"`
	MeritList []MeritList `json:"nodes"`
	Level     string      `json:"Level"` //2
}

type MeritDepartment struct { //分院：施工预算、水工分院……
	Id    int64  `json:"Id"`
	Pid   int64  `form:"-"`
	Title string `json:"text"` //这个后面json仅仅对于encode解析有用
	// Selectable bool    `json:"selectable"`
	Secoffice []MeritSecoffice `json:"nodes"`
	Level     string           `json:"Level"` //1
}

type Person struct {
	Id         int64  `json:"Id"`
	Name       string `json:"Name"`
	Department string `json:"Department"`
	Keshi      string `json:"Keshi"` //当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //记录个数
	Marks      int    //分值
}

//struct排序
type person1 []Person

func (list person1) Len() int {
	return len(list)
}

func (list person1) Less(i, j int) bool {
	if list[i].Marks > list[j].Marks {
		return true
	} else if list[i].Marks < list[j].Marks {
		return false
	} else {
		return list[i].Name > list[j].Name
	}
}

func (list person1) Swap(i, j int) {
	var temp Person = list[i]
	list[i] = list[j]
	list[j] = temp
}

//管理员进行人员价值排序查看
//排序第一排序为部门，第二排序为科室，第三排序为分值
func (c *MeritController) GetPerson() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	//3.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	//4.进行逻辑分析
	if role > 2 { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
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

//管理员登录显示侧栏结构，方便查看科室里员工总体情况，以及查看员工个人详细
func (c *MeritController) GetMerit() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	//2.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	//3.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	if role > 4 {               //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	achemployee := make([]AchEmployee, 0)
	achsecoffice := make([]AchSecoffice, 0)
	achdepart := make([]AchDepart, 0)
	//由uname取得user,获得user的分院名称
	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var depcount int
	switch role {
	case 1: //管理员登录显示的侧栏是全部的
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
					users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
					if err != nil {
						beego.Error(err)
					}
					for i3, _ := range users {
						cc := make([]AchEmployee, 1)
						cc[0].Id = users[i3].Id
						cc[0].Level = "3"
						cc[0].Pid = category2[i2].Id
						cc[0].Nickname = users[i3].Nickname //名称
						// beego.Info(users[i3].Nickname)
						// cc[0].Selectable = false
						achemployee = append(achemployee, cc...)
					}
					bb[0].Tags[0] = strconv.Itoa(count)
					bb[0].Employee = achemployee
					bb[0].Selectable = true
					achemployee = make([]AchEmployee, 0) //再把slice置0
					achsecoffice = append(achsecoffice, bb...)
					depcount = depcount + count //部门人员数=科室人员数相加
				}
				// aa[0].Secoffice = achsecoffice
				// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				// achdepart = append(achdepart, aa...)
			}
			//查出所有有这个部门但科室名为空的人员
			//根据分院查所有员工
			// beego.Info(category1[i1].Title)
			users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(users)
			for i3, _ := range users {
				dd := make([]AchSecoffice, 1)
				dd[0].Id = users[i3].Id
				dd[0].Level = "3"
				// dd[0].Href = users[i3].Ip + ":" + users[i3].Port
				dd[0].Pid = category1[i1].Id
				dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
				dd[0].Selectable = true
				achsecoffice = append(achsecoffice, dd...)
			}
			aa[0].Tags[0] = count + depcount
			// count = 0
			depcount = 0
			aa[0].Secoffice = achsecoffice
			aa[0].Selectable = true                //默认是false点击展开
			achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			achdepart = append(achdepart, aa...)
		}
	case 2: //分院领导登录显示的侧栏是本分院的所有科室
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		// aa[0].Selectable = false
		category2, err := models.GetAdminDepartTitle(user.Department) //得到多个科室
		if err != nil {
			beego.Error(err)
		}
		for i2, _ := range category2 {
			bb := make([]AchSecoffice, 1)
			bb[0].Id = category2[i2].Id
			bb[0].Level = "2"
			bb[0].Pid = category1.Id
			bb[0].Title = category2[i2].Title //科室名称
			//根据分院和科室查所有员工
			users, count, err := models.GetUsersbySec(category1.Title, category2[i2].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range users {
				cc := make([]AchEmployee, 1)
				cc[0].Id = users[i3].Id
				cc[0].Level = "3"
				cc[0].Pid = category2[i2].Id
				cc[0].Nickname = users[i3].Nickname //名称
				// cc[0].Selectable = false
				achemployee = append(achemployee, cc...)
			}
			bb[0].Tags[0] = strconv.Itoa(count)
			bb[0].Employee = achemployee
			achemployee = make([]AchEmployee, 0) //再把slice置0
			achsecoffice = append(achsecoffice, bb...)
		}
		aa[0].Secoffice = achsecoffice
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	case 3: //主任登录显示的侧栏是本科室的所有人
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		//由分院id和科室名称取得科室
		category2, err := models.GetAdminDepartbyidtitle(category1.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		bb := make([]AchSecoffice, 1)
		bb[0].Id = category2.Id
		bb[0].Level = "2"
		bb[0].Pid = category1.Id
		bb[0].Title = category2.Title //科室名称
		//根据分院和科室查所有员工
		users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			cc := make([]AchEmployee, 1)
			cc[0].Id = users[i3].Id
			cc[0].Level = "3"
			cc[0].Pid = category2.Id
			cc[0].Nickname = users[i3].Nickname //名称
			// cc[0].Selectable = false
			achemployee = append(achemployee, cc...)
		}
		bb[0].Tags[0] = strconv.Itoa(count)
		bb[0].Employee = achemployee
		achemployee = make([]AchEmployee, 0) //再把slice置0
		achsecoffice = append(achsecoffice, bb...)
		aa[0].Secoffice = achsecoffice
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	case 4: //个人登录显示自己
		//由uname取得分院名称和科室名称
		// user := models.GetUserByUsername(uname)
		//由分院名称取得分院属性
		category1, err := models.GetAdminDepartName(user.Department)
		if err != nil {
			beego.Error(err)
		}
		aa := make([]AchDepart, 1)
		aa[0].Id = category1.Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1.Title //分院名称
		//由分院id和科室名称取得科室
		category2, err := models.GetAdminDepartbyidtitle(category1.Id, user.Secoffice)
		if err != nil {
			beego.Error(err)
		}
		bb := make([]AchSecoffice, 1)
		bb[0].Id = category2.Id
		bb[0].Level = "2"
		bb[0].Pid = category1.Id
		bb[0].Title = category2.Title //科室名称

		cc := make([]AchEmployee, 1)
		cc[0].Id = user.Id
		cc[0].Level = "3"
		cc[0].Pid = category2.Id
		cc[0].Nickname = user.Nickname //名称

		achemployee = append(achemployee, cc...)

		bb[0].Tags[0] = "1"
		bb[0].Employee = achemployee
		achemployee = make([]AchEmployee, 0) //再把slice置0
		achsecoffice = append(achsecoffice, bb...)
		aa[0].Secoffice = achsecoffice
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	}
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = achdepart
	// beego.Info(achdepart)
	c.TplName = "merit.tpl"
}

//上面那个是显示侧栏
//这个是显示右侧iframe框架内容——科室内人员情况统计
func (c *MeritController) Secofficeshow() {
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		return
	}
	//2.取得客户端用户名
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
	//3.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	if role > 4 {               //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	if secid == "" { //如果为空，则用登录的
		secid = strconv.FormatInt(user.Id, 10)
	}
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	level := c.Input().Get("level")
	key := c.Input().Get("key")
	daterange := c.Input().Get("datefilter")
	// beego.Info(daterange)
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
	if len(daterange) > 19 {
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
		c.TplName = "merit_institute.tpl"
	case "1": //如果是分院，则显示全部科室
		categoryname, err := models.GetAdminDepartbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		//权限判断，并且属于这个分院
		if role == 1 || role == 2 && user.Department == categoryname.Title { //
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["Deptitle"] = categoryname.Title
			c.TplName = "merit_depoffice.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	case "2": //如果是科室，则显示全部人员情况
		//取得科室名称
		categoryname, err := models.GetAdminDepartbyId(secid1)
		if err != nil {
			beego.Error(err)
		}
		// 取得分院名称
		categoryname1, err := models.GetAdminDepartbyId(categoryname.ParentId)
		if err != nil {
			beego.Error(err)
		}
		//1.进行权限读取,属于这个科室，或者属于这个分院
		if role == 1 || role == 3 && user.Secoffice == categoryname.Title || role == 2 && user.Department == categoryname1.Title {
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			c.Data["Secid"] = secid
			c.Data["Sectitle"] = categoryname.Title
			c.Data["Level"] = level

			c.TplName = "merit_secoffice.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	default:
		// case "3": //如果是个人，则显示个人详细情况
		//分2部分，一部分是已经完成状态的，state是4，另一部分是状态分别是3待审查通过,2，1的
		usernickname := models.GetUserByUserId(secid1)
		//1.进行权限读取，室主任以上并且属于这个科室，或者或本人
		if role == 1 || role == 3 && user.Secoffice == usernickname.Secoffice || role == 2 && user.Department == usernickname.Department || user.Nickname == usernickname.Nickname { //
			c.Data["Starttime"] = t1
			c.Data["Endtime"] = t2
			//下面这个catalogs用于employee_show.tpl
			c.Data["Secid"] = secid
			c.Data["Level"] = level
			c.Data["UserNickname"] = usernickname.Nickname

			if key == "modify" { //新窗口显示处理页面
				//如果是本人，则显示
				c.TplName = "merit_employeework.tpl"
			} else { //直接查看页面
				//如果是本人，则显示带处理按钮的
				if usernickname.Nickname == user.Nickname {
					c.Data["IsMe"] = true
				} else { //别人查看，不显示处理按钮
					c.Data["IsMe"] = false
				}
				c.TplName = "merit_employee.tpl"
			}
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			return
		}
	}
}

//上面那个是显示右侧页面
//这个是填充数据——科室内人员成果情况统计
func (c *MeritController) SecofficeData() {
	//分院——科室——人员甲（乙、丙……）——绘制——设计——校核——审查——合计——排序
	secid := c.Input().Get("secid")
	secid1, err := strconv.ParseInt(secid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	level := c.Input().Get("level")
	daterange := c.Input().Get("datefilter")
	// beego.Info(daterange)
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
	if len(daterange) > 19 {
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

	//取得科室名称
	categoryname, err := models.GetAdminDepartbyId(secid1)
	if err != nil {
		beego.Error(err)
	}

	employeevalue := make([]models.Employeeachievement, 0)
	//根据科室id查所有员工
	users, _, err := models.GetUsersbySecId(secid) //得到员工姓名
	// beego.Info(users)
	if err != nil {
		beego.Error(err)
	}
	for _, v := range users {
		//由username查出所有编制成果总分、设计总分……合计
		employee, _, err := models.Getemployeevalue(v.Nickname, t1, t2)
		if err != nil {
			beego.Error(err)
		}
		employeevalue = append(employeevalue, employee...)
	}
	//排序
	pList := graphictopics(employeevalue)
	sort.Sort(pList)
	c.Data["Starttime"] = t1
	c.Data["Endtime"] = t2
	c.Data["Secid"] = secid
	c.Data["Sectitle"] = categoryname.Title
	c.Data["Level"] = level
	c.Data["json"] = pList
	c.ServeJSON()
}

//上面那个是显示侧栏
//这个是显示右侧iframe框架内容——自己的价值内容列表
//要修改*************
func (c *MeritController) Myself() {
	//如果是主任以上权限人查看，则id代表用户名id，个人查看，id则代表价值id
	//1.首先判断是否注册
	if !checkAccount(c.Ctx) {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
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
	//4.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	if role > 4 {               //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}

	user, err := models.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	var numbers1, marks1 int
	slice1 := make([]Person, 0)
	id := c.Input().Get("id")
	secid1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	level := c.Input().Get("level")
	if role < 4 { //主任以上，则id代表userid
		switch level {
		case "0": //如果是总院，则显示全部分院情况
			c.TplName = "merit_myself.tpl"
		case "1":
			c.TplName = "merit_myself.tpl"
		case "2": //id代表的是科室id，取得所有人的价值进行排序
			//取得科室名称
			categoryname, err := models.GetCategory(secid1)
			if err != nil {
				beego.Error(err)
			}
			// 取得分院名称
			// categoryname1, err := models.GetCategory(categoryname.ParentId)
			// if err != nil {
			// 	beego.Error(err)
			// }
			//1.进行权限读取,属于这个科室，或者属于这个分院
			// if role == 1 || role == 3 && user.Secoffice == categoryname.Title || role == 2 && user.Department == categoryname1.Title { //
			// employeevalue := make([]models.Employeeachievement, 0)
			// depid := c.Input().Get("depid")
			//根据分院和科室查所有员工
			// users, count, err := models.GetUsersbySec(category1.Title, category2.Title) //得到员工姓名
			//根据科室id查所有员工
			users, _, err := models.GetUsersbySecId(id) //得到员工姓名
			// beego.Info(users)
			if err != nil {
				beego.Error(err)
			}
			// for _, v := range users {
			//由username查出所有价值总分，排序
			// 	employee, err := models.Getemployeevalue(v.Nickname, t1, t2)
			// 	if err != nil {
			// 		beego.Error(err)
			// 	}
			// 	employeevalue = append(employeevalue, employee...)
			// }
			for _, v1 := range users {
				//根据价值id和用户id，得到成果，统计数量和分值
				//取得用户的价值数量和分值
				_, numbers, marks, err := models.GetMeritTopic(0, v1.Id)
				if err != nil {
					beego.Error(err)
				}
				marks1 = marks1 + marks
				numbers1 = numbers1 + numbers
				aa := make([]Person, 1)
				aa[0].Id = v1.Id //这里用for i1,v1,然后用v1.Id一样的意思
				aa[0].Name = v1.Nickname
				aa[0].Department = v1.Department
				aa[0].Keshi = v1.Secoffice
				aa[0].Numbers = numbers1
				aa[0].Marks = marks1
				slice1 = append(slice1, aa...)
				marks1 = 0
				numbers1 = 0
			}
			//排序
			pList := person1(slice1)
			sort.Sort(pList)
			c.Data["Sectitle"] = categoryname.Title
			c.Data["Level"] = level
			c.Data["Employee"] = pList //employeevalue

			c.TplName = "merit_secoffice.tpl"
		case "3": //查看用户详情
			userid, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			topics, err := models.GetAllMeritTopic(userid)
			c.Data["topics"] = topics
			usernickname := models.GetUserByUserId(userid)
			c.Data["UserNickname"] = usernickname.Nickname
			//查出所有用户的价值资料
			//前提是价值资料里要带用户id
			category5, _ := models.GetAllCategory()
			if err != nil {
				beego.Error(err)
			}
			c.Data["category"] = category5
			c.TplName = "merit_myself.tpl"
		default: //显示自己的价值情况
			topics, err := models.GetAllMeritTopic(user.Id)
			c.Data["topics"] = topics
			c.Data["UserNickname"] = user.Nickname
			//查出所有用户的价值资料
			//前提是价值资料里要带用户id
			category5, _ := models.GetAllCategory()
			if err != nil {
				beego.Error(err)
			}
			c.Data["category"] = category5
			c.TplName = "merit_myself.tpl"
		}

	} else if role == 4 {
		switch level {
		case "0": //如果是总院，则显示全部分院情况
			c.TplName = "merit_myself.tpl"
		case "1":
			c.TplName = "merit_myself.tpl"
		case "2":
			c.TplName = "merit_myself.tpl"
		case "3":
			c.TplName = "merit_myself.tpl"
		case "4":
			//添加页面
			c.TplName = "merit_add.tpl"
			// idNum, err := strconv.ParseInt(id, 10, 64)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// category11, err := models.GetCategory(idNum) //得到价值
			// if err != nil {
			// 	beego.Error(err)
			// }
			// //进行选择列表拆分
			// array1 := strings.Split(category11.List, ",")
			// slice1 := make([]List1, 0)
			// for _, v := range array1 {
			// 	ee := make([]List1, 1)
			// 	ee[0].Choose = v
			// 	slice1 = append(slice1, ee...)
			// }
			// //2.取得客户端用户名
			// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
			// defer sess.SessionRelease(c.Ctx.ResponseWriter)
			// v := sess.Get("uname")
			// var uname string
			// if v != nil {
			// 	uname = v.(string)
			// 	c.Data["Uname"] = v.(string)
			// }
			// //先由uname取得uid
			// user11, err := models.GetUserByUsername(uname)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// //取得父级id和用户id下的价值
			// topics11, _, _, err := models.GetMeritTopic(idNum, user11.Id)
			// c.Data["topics"] = topics11
			// // c.ServeJSON()
			// c.Data["category"] = category11
			// c.Data["list"] = slice1
		default: //显示自己的价值情况
			topics, err := models.GetAllMeritTopic(user.Id)
			c.Data["topics"] = topics
			c.Data["UserNickname"] = user.Nickname
			//查出所有用户的价值资料
			//前提是价值资料里要带用户id
			category5, _ := models.GetAllCategory()
			if err != nil {
				beego.Error(err)
			}
			c.Data["category"] = category5
			c.TplName = "merit_myself.tpl"
		}
	}
}

//添加价值结构中的项目
func (c *MeritController) AddMerit() {
	id := c.Input().Get("pid")
	text1 := c.Input().Get("title")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//存入数据库——分院
	_, err = models.AddCategory(idNum, text1, "", "", "")
	if err != nil {
		beego.Error(err)
	} else {
		data := text1
		c.Ctx.WriteString(data)
	}
}

//显示——修改价值结构中的项目
// func (c *MeritController) UpdateMerit() {
// 	//1.首先判断是否注册
// 	if !checkAccount(c.Ctx) {
// 		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/login?url="+route, 302)
// 		// c.Redirect("/login", 302)
// 		return
// 	}
// 	//2.取得文章的作者
// 	//3.由用户id取得用户名
// 	//4.取得客户端用户名
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		c.Data["Uname"] = v.(string)
// 	}
// 	// uname := v.(string) //ck.Value
// 	//4.取出用户的权限等级
// 	role, _ := checkRole(c.Ctx) //login里的
// 	// beego.Info(role)
// 	//5.进行逻辑分析：
// 	// rolename, _ := strconv.ParseInt(role, 10, 64)
// 	if role > 2 { //
// 		// port := strconv.Itoa(c.Ctx.Input.Port()) //c.Ctx.Input.Site() + ":" + port +
// 		route := c.Ctx.Request.URL.String()
// 		c.Data["Url"] = route
// 		c.Redirect("/roleerr?url="+route, 302)
// 		// c.Redirect("/roleerr", 302)
// 		return
// 	}

// 	//4.取得价值列表choose和mark
// 	id := c.Input().Get("id")
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	category, err := models.GetCategory(idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//如果mark为空，则寻找选择列表的分值
// 	//进行选择列表拆分
// 	array1 := strings.Split(category.List, ",")
// 	array2 := strings.Split(category.ListMark, ",")
// 	//进行选择列表拆分
// 	// array1 := strings.Split(category.List, ",")
// 	slice1 := make([]List1, 0)
// 	for i, v := range array1 {
// 		ee := make([]List1, 1)
// 		ee[0].Choose = v
// 		ee[0].Mark1 = array2[i]
// 		slice1 = append(slice1, ee...)
// 	}

// 	slice2 := make([]List1, 0)
// 	for _, w := range array2 {
// 		ff := make([]List1, 1)
// 		ff[0].Mark1 = w
// 		slice2 = append(slice2, ff...)
// 	}
// 	c.TplName = "admin_json_modify.tpl"
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	c.Data["category"] = category
// 	c.Data["list"] = slice1
// 	c.Data["mark"] = slice2
// 	// beego.Info(slice2)
// 	c.Data["Cid"] = id
// 	c.Data["IsCategory"] = true
// }

//提交修改—修改价值结构中的项目
func (c *MeritController) UpdateMerit() {
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
	c.Redirect("/json", 301)
}

//删除价值结构中的项目
func (c *MeritController) DeleteMerit() {
	id := c.Input().Get("id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	err = models.Deletejson(idNum)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/json", 301)
}

//传递merit
//退回
//已提交——待审核
//待提交
//待审核
//导入

//导入json数据
// func (c *MeritController) ImportJson() {
// 	//获取上传的文件
// 	_, h, err := c.GetFile("json")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	var path string
// 	if h != nil {
// 		//保存附件
// 		path = ".\\attachment\\" + h.Filename
// 		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
// 		err = c.SaveToFile("json", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	contents, _ := ioutil.ReadFile(path)
// 	var r List6
// 	//var r JsonStruct//空结构对于系统unmarshal不行。
// 	//	var r map[string]interface{}//空接口可行
// 	//	var r []interface{}//这个对于系统unmarshal不行
// 	err = json.Unmarshal([]byte(contents), &r)
// 	if err != nil {
// 		fmt.Printf("err was %v", err)
// 	}
// 	// fmt.Println(r)
// 	// beego.Info(r)

// 	js, err := simplejson.NewJson([]byte(contents))
// 	if err != nil {
// 		panic("json format error")
// 	}
// 	//1.获取省水利院
// 	text, err := js.Get("text").String()
// 	//存入数据库——单位
// 	Id, err := models.AddCategory(0, text, "", "", "")
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	arr, err := js.Get("nodes").Array()
// 	if err != nil {
// 		fmt.Println("decode error: get array failed!")
// 		// return
// 	}
// 	for i, _ := range arr {
// 		// beego.Info(v)是map[string]interface{}
// 		text1, _ := js.Get("nodes").GetIndex(i).Get("text").String()
// 		//存入数据库——分院
// 		Id1, err := models.AddCategory(Id, text1, "", "", "")
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		arr1, err := js.Get("nodes").GetIndex(i).Get("nodes").Array()
// 		for i1, _ := range arr1 {
// 			text2, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("text").String()
// 			//存入数据库——科室
// 			Id2, err := models.AddCategory(Id1, text2, "", "", "")
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			arr2, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").Array()
// 			for i2, _ := range arr2 {
// 				text3, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("text").String()
// 				//存入数据库——管理类
// 				Id3, err := models.AddCategory(Id2, text3, "", "", "")
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				arr3, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").Array()
// 				for i3, _ := range arr3 {
// 					text4, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("text").String()
// 					text5, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("mark").String()
// 					//循环取出选择项，拼接字符串
// 					//循环取出每个选择项的打分，拼接字符串
// 					arr4, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").Array()
// 					var text8, text9 string
// 					for i4, _ := range arr4 {
// 						text6, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("text").String()
// 						text7, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("mark").String()
// 						if i == 0 {
// 							text8 = text6
// 							text9 = text7
// 						} else {
// 							text8 = text8 + "," + text6
// 							text9 = text9 + "," + text7
// 						}
// 					}
// 					//存入数据库——项目负责人
// 					// url:="/"+"add?id="+
// 					_, err = models.AddCategory(Id3, text4, text5, text8, text9)
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 				}

// 			}
// 		}
// 	}
// }

//根据conf目录下的json.json文件初始化价值结构
// func (c *MeritController) InitJson() {
// 	contents, _ := ioutil.ReadFile("./conf/json.json")
// 	var r List6
// 	err := json.Unmarshal([]byte(contents), &r)
// 	if err != nil {
// 		fmt.Printf("err was %v", err)
// 	}
// 	// fmt.Println(r)
// 	// beego.Info(r)

// 	js, err := simplejson.NewJson([]byte(contents))
// 	if err != nil {
// 		panic("json format error")
// 	}
// 	//1.获取省水利院
// 	text, err := js.Get("text").String()
// 	//存入数据库——单位
// 	Id, err := models.AddCategory(0, text, "", "", "")
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	arr, err := js.Get("nodes").Array()
// 	if err != nil {
// 		fmt.Println("decode error: get array failed!")
// 		// return
// 	}
// 	for i, _ := range arr {
// 		// beego.Info(v)是map[string]interface{}
// 		text1, _ := js.Get("nodes").GetIndex(i).Get("text").String()
// 		//存入数据库——分院
// 		Id1, err := models.AddCategory(Id, text1, "", "", "")
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		arr1, err := js.Get("nodes").GetIndex(i).Get("nodes").Array()
// 		for i1, _ := range arr1 {
// 			text2, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("text").String()
// 			//存入数据库——科室
// 			Id2, err := models.AddCategory(Id1, text2, "", "", "")
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			arr2, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").Array()
// 			for i2, _ := range arr2 {
// 				text3, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("text").String()
// 				//存入数据库——管理类
// 				Id3, err := models.AddCategory(Id2, text3, "", "", "")
// 				if err != nil {
// 					beego.Error(err)
// 				}
// 				arr3, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").Array()
// 				for i3, _ := range arr3 {
// 					text4, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("text").String()
// 					text5, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("mark").String()
// 					//循环取出选择项，拼接字符串
// 					//循环取出每个选择项的打分，拼接字符串
// 					arr4, err := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").Array()
// 					var text8, text9 string
// 					for i4, _ := range arr4 {
// 						text6, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("text").String()
// 						text7, _ := js.Get("nodes").GetIndex(i).Get("nodes").GetIndex(i1).Get("nodes").GetIndex(i2).Get("nodes").GetIndex(i3).Get("nodes").GetIndex(i4).Get("mark").String()
// 						if i == 0 {
// 							text8 = text6
// 							text9 = text7
// 						} else {
// 							text8 = text8 + "," + text6
// 							text9 = text9 + "," + text7
// 						}
// 					}
// 					//存入数据库——项目负责人
// 					// url:="/"+"add?id="+
// 					_, err = models.AddCategory(Id3, text4, text5, text8, text9)
// 					if err != nil {
// 						beego.Error(err)
// 					}
// 				}

// 			}
// 		}
// 	}
// }

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
