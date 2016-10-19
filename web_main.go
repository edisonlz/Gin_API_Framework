package main

import (
	_ "Gin_API_Framework/models"
	_ "Gin_API_Framework/web-routers"
	"github.com/astaxie/beego"
	"fmt"
    "path"
    "runtime"
)


func callerSourcePath() string {
    _, callerPath, _, _ := runtime.Caller(1)
    return path.Dir(callerPath)
}


func main() {

    curpath := callerSourcePath()
    static_path := path.Join(curpath, "/", "static")
    template_path := path.Join(curpath, "/web-controllers/templates")

    beego.LoadAppConfig("ini", path.Join(curpath, "/conf/app.conf"))
	beego.SetStaticPath("/static",static_path)
	beego.SetViewsPath(template_path)

    fmt.Println(beego.AppConfig.Int("HttpPort"))
	fmt.Println("[static path]" , static_path)
	fmt.Println("[template path]" , template_path)

	beego.Run()
}