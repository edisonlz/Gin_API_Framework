package item

import  (
	"github.com/astaxie/beego/orm"
	"fmt"
)


type Item struct {
	Id     int    `orm:"auto"`
	Name   string `orm:"size(100)"`
	Seq    int `orm:"default(0)"`
}


func (item *Item) CreateItem(
	name string,
	)  bool {

	o := orm.NewOrm()
	o.Using("default")
	item.Name = name
	o.Insert(item)
	return true
}


func ItemList() (item []Item) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var items []Item
	cnt, err :=  qs.Filter(
		"id__gt", 0).OrderBy(
		"seq").Limit(
		10, 0).All(
		&items)

	if err == nil {
		fmt.Printf("count", cnt)
	}
	return items
}