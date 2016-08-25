package web_routers

import (
	"github.com/astaxie/beego"
	"Gin_API_Framework/web-controllers"
)


func init() {
	beego.Router("/", &web_controllers.MainController{})
}
