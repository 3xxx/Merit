package routers

import (
	"github.com/astaxie/beego"
	"merit/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{}, "get:Test")

	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	beego.Router("/json", &controllers.JsonController{})
	beego.Router("/importjson", &controllers.JsonController{}, "post:ImportJson")
	beego.Router("/add", &controllers.JsonController{}, "get:Add")

	beego.Router("/AddMeritTopic", &controllers.MeritTopicController{}, "post:AddMeritTopic")
}
