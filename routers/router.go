package routers

import (
	"github.com/astaxie/beego"
	"merit/controllers"
)

func init() {
	//1.首页index
	beego.Router("/", &controllers.MainController{})
	//管理员
	beego.Router("/admin", &controllers.MainController{}, "get:Admin")
	beego.Router("/jsoneditor", &controllers.MainController{}, "get:Jsoneditor")

	//2.1首页进入价值——直接进入自己的价值页面
	beego.Router("/merit", &controllers.JsonController{}, "get:GetMeritUser")
	//2.2首页进入成果登记
	beego.Router("/achievement", &controllers.Achievement{}, "get:GetAchievement")
	//这个同上面一样
	beego.Router("/getachievement", &controllers.Achievement{}, "get:GetAchievement")
	//个人在线添加成果
	beego.Router("/achievement/addcatalog", &controllers.Achievement{}, "post:AddCatalog")
	//个人在线直接提交成果
	beego.Router("/achievement/sendcatalog", &controllers.Achievement{}, "post:SendCatalog")
	//在线退回成果
	beego.Router("/achievement/downsendcatalog", &controllers.Achievement{}, "post:DownSendCatalog")

	//个人在线修改保存成果
	beego.Router("/achievement/modifycatalog", &controllers.Achievement{}, "post:ModifyCatalog")

	//个人在线删除一条成功
	beego.Router("/achievement/delete", &controllers.Achievement{}, "post:DeleteCatalog")
	//编辑成果类型和折标系数表
	beego.Router("/achievement/ratio", &controllers.Achievement{}, "get:Ratio")
	beego.Router("/achievement/addratio", &controllers.Achievement{}, "post:AddRatio")
	beego.Router("/achievement/modifyratio", &controllers.Achievement{}, "post:ModifyRatio")

	beego.Router("/test", &controllers.MainController{}, "get:Test")
	beego.Router("/test1", &controllers.MainController{}, "get:Test1")

	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")

	beego.Router("/getperson", &controllers.JsonController{}, "get:GetPerson")

	beego.Router("/get", &controllers.JsonController{})
	beego.Router("/json", &controllers.JsonController{}) //这个和上面等价
	beego.Router("/importjson", &controllers.JsonController{}, "post:ImportJson")

	beego.Router("/addjson", &controllers.JsonController{}, "post:Addjson")
	beego.Router("/modifyjson", &controllers.JsonController{}, "get:Modifyjson")          //显示修改页面
	beego.Router("/modifyjsonpost", &controllers.JsonController{}, "post:ModifyjsonPost") //提交修改
	beego.Router("/deletejson", &controllers.JsonController{}, "get:Deletejson")

	beego.Router("/add", &controllers.MeritTopicController{}, "get:Add")
	beego.Router("/AddMeritTopic", &controllers.MeritTopicController{}, "post:AddMeritTopic")
	beego.Router("/view", &controllers.MeritTopicController{}, "get:ViewMeritTopic")
	beego.Router("/modify", &controllers.MeritTopicController{}, "get:ModifyMeritTopic")
	beego.Router("/ModifyPost", &controllers.MeritTopicController{}, "post:ModifyPost")
	beego.Router("/delete", &controllers.MeritTopicController{}, "get:DeleteMeritTopic")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")

	beego.Router("/regist", &controllers.RegistController{})
	// beego.Router("/registerr", &controllers.RegistController{}, "get:RegistErr")
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "post:GetUname")

	//成果登记系统
	//管理员登录查看分院整体情况
	beego.Router("/getachievement", &controllers.Achievement{}, "get:GetAchievement")
	//成果登记表导入数据库
	beego.Router("/import_xls_catalog", &controllers.Achievement{}, "post:Import_Xls_Catalog")
	// 主页里显示iframe——科室总体情况
	beego.Router("/secofficeshow", &controllers.Achievement{}, "get:Secofficeshow")

	//人员管理
	beego.Router("/user/AddUser", &controllers.UserController{}, "*:AddUser")
	beego.Router("/user/UpdateUser", &controllers.UserController{}, "*:UpdateUser")
	beego.Router("/user/deluser", &controllers.UserController{}, "*:DelUser")
	beego.Router("/user/index", &controllers.UserController{}, "*:Index")
	//管理员修改用户资料
	beego.Router("/user/view", &controllers.UserController{}, "get:View")
	beego.Router("/user/view/*", &controllers.UserController{}, "get:View")
	beego.Router("/user/importexcel", &controllers.UserController{}, "post:ImportExcel")

	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")

	beego.Router("/role/AddAndEdit", &controllers.RoleController{}, "*:AddAndEdit")
	beego.Router("/role/DelRole", &controllers.RoleController{}, "*:DelRole")
	// beego.Router("/role/AccessToNode", &controllers.RoleController{}, "*:AccessToNode")
	// beego.Router("/role/AddAccess", &controllers.RoleController{}, "*:AddAccess")
	beego.Router("/role/RoleToUserList", &controllers.RoleController{}, "*:RoleToUserList")
	beego.Router("/role/AddRoleToUser", &controllers.RoleController{}, "*:AddRoleToUser")
	beego.Router("/role/Getlist", &controllers.RoleController{}, "*:Getlist")
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/roleerr", &controllers.RoleController{}, "*:Roleerr") //显示权限不够

}
