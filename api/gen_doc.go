package main

import (
    "fmt"
    "log"
    "path"
    "runtime"
    "os"

    _ "github.com/gin-gonic/gin"
    "Gin_API_Framework/api/routers"
    _ "Gin_API_Framework/models"
    _ "Gin_API_Framework/api/docs"
    //"api_project/controllers"
    _ "github.com/astaxie/beego"
)


func callerSourcePath() string {
    _, callerPath, _, _ := runtime.Caller(1)
    return path.Dir(callerPath)
}

func main() {
    fmt.Println("[Doc Gen...]")
    curpath := callerSourcePath()
    fmt.Println("[curpath]",curpath)
    routers.GenerateDocs(curpath)
    log.Println("doc generated...")
    os.Exit(0)
}

