package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "api_project/models"
)


func main() {
    fmt.Println("starting......")
    orm.RunCommand()
    fmt.Println("ended......")
}

