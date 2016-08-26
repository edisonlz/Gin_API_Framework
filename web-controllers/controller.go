package web_controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.html"
	this.Layout = "layout/layout.html"
	this.Render()
}

type ItemListCtroller struct {
	beego.Controller
}

func (this *ItemListCtroller) Get()  {
	this.TplName = "itemlist/items.html"
	this.Layout = "layout/layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["JsFileExtend"] = "itemlist/js.tpl"
	this.Render()
}