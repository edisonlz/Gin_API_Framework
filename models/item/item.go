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

func UpdateById(id int,filed map[string] interface{})bool{
	o := orm.NewOrm()
	_, err := o.QueryTable(
		"item").Filter(
		"Id", id).Update(
		filed)
	if err == nil{
		return true
	}
	return false
}

func QueryById(id int) (Item) {
	o := orm.NewOrm()

	itemObj := Item{Id: id}
	err := o.Read(&itemObj)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(itemObj.Id, itemObj.Name)
	}

	return itemObj
}


func ItemList() (item []Item) {

	o := orm.NewOrm()
	qs := o.QueryTable("item")

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