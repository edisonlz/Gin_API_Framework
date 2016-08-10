package main

import (
    "github.com/gin-gonic/gin"
    "github.com/fvbock/endless"
    "net/http"
    "controllers"
    "time"
    "fmt"
    "log"
    "os"
    _ "models"
)


func main() {
    router := gin.New()

    fmt.Println("starting......")
    // Global middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.Static("/assets", "./assets")

    //register router
    // router.GET("/user/:name", controllers.UserHandler)
    // router.GET("/user/:name/*action", controllers.UserActionHandler)

    v1 := router.Group("/user")
    {
        v1.GET("/login", controllers.UserLoginHandler)
        v1.GET("/logout", controllers.UserLogoutHandler)
    }

    router.GET("/test", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
    })

    server := endless.NewServer("127.0.0.1:8080", router)
    server.ReadTimeout = 3 * time.Second
    server.WriteTimeout = 3 * time.Second

    server.BeforeBegin = func(add string) {
        log.Printf(add)
    }

    err := server.ListenAndServe()

    if err != nil {
        log.Println(err)
    }

    log.Println("Server stopped")
    os.Exit(0)
}

