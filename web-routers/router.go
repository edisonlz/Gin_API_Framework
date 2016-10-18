package web_routers

import (
	"github.com/astaxie/beego"
	"Gin_API_Framework/web-controllers"
)


func init() {
	beego.Router("/", &web_controllers.MainController{})
	beego.Router("/item/list", &web_controllers.ItemListCtroller{})
	beego.Router("/item/pos", &web_controllers.ItemPositionCtroller{})
	beego.Router("/item/edit", &web_controllers.ItemEditCtroller{})
}
