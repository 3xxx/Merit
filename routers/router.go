package routers

import (
	"github.com/astaxie/beego"
	"merit/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	beego.Router("/test1", &controllers.MainController{}, "get:Test1")

	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")

	beego.Router("/json", &controllers.JsonController{})
	beego.Router("/importjson", &controllers.JsonController{}, "post:ImportJson")
	beego.Router("/user", &controllers.JsonController{}, "get:GetMeritUser")
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

}
