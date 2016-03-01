package routers

import (
	"github.com/astaxie/beego"
	"merit/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/json", &controllers.JsonController{})
}
