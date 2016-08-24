package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"strconv"
)

type MainController struct {
	beego.Controller
}

//（2）建立一个全局session mananger对象
var globalSessions *session.Manager

//（3）在初始化“全局session mananger对象”
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	// go globalSessions.GC()
}
func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Admin() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
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
	role, err := checkRole(c.Ctx) //login里的
	if err != nil {
		beego.Error(err)
	} else {
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
	}
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin.tpl"
}

func (c *MainController) Test() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user_show.tpl"
}

func (c *MainController) Test1() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.tpl"
}
func (c *MainController) Jsoneditor() {
	c.TplName = "jsoneditor.tpl"
}
