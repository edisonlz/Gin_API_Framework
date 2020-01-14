package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "Gin_API_Framework/models"
)


func main() {
    fmt.Println("starting......")
    orm.RunCommand()
    //orm.RunSyncdb("default", false, true)
    fmt.Println("ended......")
}

