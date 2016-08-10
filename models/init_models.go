package models

//doc --- http://beego.me/docs/mvc/model/orm.md

import  (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
    "api_project/models/user"
    "fmt"
)
    

/*
    run 

    go run orm_sync.go orm syncdb
*/

func init() {

    fmt.Println("[init database]......")

    orm.Debug = true
    //regiter driver
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // register model
    orm.RegisterModel(new(user.User))
    orm.RegisterModel(new(user.Post))

    mysql_config := "root:@/go_platform?charset=utf8"
    // set default database
    orm.RegisterDataBase("default", "mysql", mysql_config)
    //set db params

    orm.SetMaxIdleConns("default", 240)
    orm.SetMaxOpenConns("default", 240)

    // set go
    fmt.Println("[end init database]......")

}