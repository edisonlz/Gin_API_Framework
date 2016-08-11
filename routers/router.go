// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers


import (
    _ "github.com/astaxie/beego"
    "github.com/gin-gonic/gin"
    "github.com/DeanThompson/ginpprof"
    "Gin_API_Framework/controllers"
    "net/http"
    "path"
    "runtime"
    "fmt"
)

func callerSourcePath() string {
    _, callerPath, _, _ := runtime.Caller(1)
    return path.Dir(callerPath)
}

func InitRouter() http.Handler {

    curpath := callerSourcePath()
    temp_path := path.Join(curpath, "..", "controllers/templates/")
    fmt.Println("[Register Template Path]",temp_path)

    static_path := path.Join(curpath, "..", "static","docs")
    fmt.Println("[Register Static Path]",static_path)

	router := gin.New()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    router.LoadHTMLGlob(temp_path + "/*")
    router.Static("/static", static_path)
    router.StaticFS(static_path, http.Dir("static"))

    /*   register router   */
    // router.GET("/users/:name", controllers.UserHandler)
    // router.GET("/users/:name/*action", controllers.UserActionHandler)

    //group
    v1 := router.Group("/user")
    {
        v1.GET("/login", controllers.UserLoginHandler)
        v1.GET("/logout", controllers.UserLogoutHandler)
        v1.GET("/create", controllers.CreateUserHandler)
        v1.GET("/query", controllers.UserQueryByIdHandler)
        v1.GET("/list", controllers.UserListHandler)
    }

    
    router.GET("/doc", controllers.TemplateDocHandler)

    router.GET("/test", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
    })

    fmt.Println("[Plugin Router Profile]...")
    ginpprof.Wrapper(router)

    return router

}