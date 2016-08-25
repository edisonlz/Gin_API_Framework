package main

import (
	_ "Gin_API_Framework/models"
	_ "Gin_API_Framework/web-routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.SetStaticPath("/static","/Users/yinxing/godir/go/src/Gin_API_Framework/static")
	beego.SetViewsPath("/Users/yinxing/godir/go/src/Gin_API_Framework/web-controllers/templates")
	fmt.Println(beego.AppConfig.String("StaticDir"))
	fmt.Println(beego.AppConfig.String("ViewsPath"))

	beego.Run()
}