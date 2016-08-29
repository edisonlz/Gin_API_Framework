package web_controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"Gin_API_Framework/models/item"
	"strings"
	//"github.com/astaxie/beego/orm"
	"strconv"
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
	this.Data["Items"] = item.ItemList()
	this.Render()
}

type ItemPositionCtroller struct {
	beego.Controller
}

func (this *ItemPositionCtroller) Post()  {
	ids := this.GetString("gift_ids")

	idList := strings.Split(ids,",")
	for i, v := range idList{
		intId ,_:= strconv.Atoi(v)
		fields := make(map[string] interface{})
		fields["Seq"] = i
		item.UpdateById(intId, fields)
	}

	resultMap := make(map[string]string)
	resultMap["status"] = "success"
	resultMap["data"] = ids
	this.Data["json"]= resultMap

	this.ServeJSON()
}


type ItemEditCtroller struct {
	beego.Controller
}

func (this *ItemEditCtroller) Get()  {
	itemId,err := this.GetInt("itemId")
	name := this.GetString("name")
	fmt.Println(itemId,name)
	fmt.Println(err)
	//this.Redirect()
	if err == nil{
		itemObj := item.QueryById(itemId)
		this.Data["Item"] = itemObj
	}else {
		this.Data["Item"] = nil
	}

	this.TplName = "itemedit/itemedit.html"
	this.Layout = "layout/layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["JsFileExtend"] = "itemedit/js.tpl"
	this.Render()
}

func (this *ItemEditCtroller) Post()  {
	itemId,err := this.GetInt("itemId")
	name := this.GetString("name")

	if err==nil{
		fields := make(map[string] interface{})
		fields["Name"] = name
		item.UpdateById(itemId,fields)
	}else {
		itemObj := new(item.Item)
		itemObj.CreateItem(name)
	}
	this.Redirect("/item/list",302)
}
