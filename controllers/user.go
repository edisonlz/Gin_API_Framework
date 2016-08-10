package controllers


import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "api_project/controllers/msg_struct"
    "api_project/models/user"
    "strconv"
    _ "github.com/astaxie/beego"
    _ "io/ioutil"
    _ "path"
    _ "runtime"
)



// @Title User Login
// @API_GROUP User
// @Description 用户登录接口
// @Success 200 {object} 
// @Param   name     query   string true  "user name"  "username"
// @Param   pwd      query   string  true "password"  "password"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/login [get]
func UserLoginHandler(c *gin.Context) {
        name := c.Query("name")
        pwd := c.Query("pwd")
        message := name + " is " + pwd
        var msg = new(msg_struct.Msg)
        msg.Message = message
        fmt.Println(msg.Message)

        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": message,
        })
}


// @Title User Logout
// @API_GROUP User
// @Description 用户登出接口
// @Success 200 {object} 
// @Param   name     query   string false       "user name"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/logout [get]
func UserLogoutHandler(c *gin.Context) {
    name := c.Query("name")
    message := name + " is logout"
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": message,
    })
}

// @Title User Create
// @API_GROUP User
// @Description 创建用户接口
// @Success 200 {object} 
// @Param   name     query   string false       "user name"
// @Param   gender      query   string  false       "gender"
// @Param   phone      query   string  false       "phone"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/create [get]
func CreateUserHandler(c *gin.Context) {

    name := c.Query("name")
    gender := c.Query("gender")
    phone := c.Query("phone")

    user := new(user.User)

    success := user.CreateUser(name ,gender ,phone)
    
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "is_created": success,
    })

}


// @Title User Query By ID
// @API_GROUP User
// @Description 查询用户接口通过用户ID
// @Success 200 {object} 
// @Param   uid     query   string false       "user id"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/query [get]
func UserQueryByIdHandler(c *gin.Context) {

    suid := c.Query("uid")
    uid , error := strconv.Atoi(suid)
    if error != nil {
        c.JSON(400, gin.H{
            "status":  "fail",
            "msg": "字符串转换成整数失败",
        })
        return
    }

    u := user.UserQueryById(uid)

    c.JSON(http.StatusOK, gin.H {
        "status":  "success",
        "user": u,
    })

}

// @Title User List
// @API_GROUP User
// @Description 查询用户列表接口
// @Success 200 {object} 
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/create [get]
func UserListHandler(c *gin.Context) {
    users := user.UserList()

    c.JSON(http.StatusOK, gin.H {
        "status":  "success",
        "users": users,
    })

}


