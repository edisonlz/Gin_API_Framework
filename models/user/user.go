package user

import  (
    "github.com/astaxie/beego/orm"
    "fmt"
)

type User struct {
    Id     int    `orm:"auto"`
    Name   string `orm:"size(100)"`
    Gender string `orm:"size(10)"`
    Phone  string `orm:"size(32)"`
    Email  string `orm:"size(32)"`
    Address string `orm:"size(64)"`
    Uuid string `orm:"size(128)"`
}


func (user *User) CreateUser(name string ,gender string ,phone string) bool { 

    o := orm.NewOrm()
    o.Using("default")

    user.Name = name
    user.Gender = gender
    user.Phone = phone

    fmt.Println(o.Insert(user))
    return true
}


func UserQueryById(uid int) (user User) {

    o := orm.NewOrm()
    u := User{Id: uid}

    err := o.Read(&u)

    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
    } else if err == orm.ErrMissPK {
        fmt.Println("找不到主键")
    } else {
        fmt.Println(u.Id, u.Name)
    }

    return u
}

func UserList() (users []User) {

    o := orm.NewOrm()
    qs := o.QueryTable("user")

    var us []User
    cnt, err :=  qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
    if err == nil {
        fmt.Printf("count", cnt)
        // for _, u := range us {
        //     fmt.Println(u)
        // }
    }
    return us
}

type Post struct {
    Id    int    `orm:"auto"`
    Title string `orm:"size(100)"`
    User  *User  `orm:"rel(fk)"`
}
