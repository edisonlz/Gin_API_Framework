# 欢迎使用GO Gin API Framework

@(示例笔记本)[马克飞象|帮助|Markdown]

**GIN API Framework**是一款专为Go Gin 框架打造的API Framework，通过精心的设计与技术实现，集成了大部分稳定开发组件，memcache consistance Hash，redis，nsq，api doc ，mysql 等。特点概述：
 
- **功能丰富** ：支持大部分服务器组件，支持API Doc；
- **得心应手** ：简单的实例，非常容易上手；
- **深度整合** ：深度整合memcache，redis，mysql，beego ，gin 框架。


-------------------

[TOC]

## 开始使用

执行服务
>  go run main.go
生成文档
>  go run gen_doc.go
同步数据库模型
>  go run orm_sync.go


### 代码块
``` go
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

```

Go 版本：
Go 1.6.3

环境配置：
export GOPATH=/Users/用户x/go/go_path
export GOBIN=/usr/local/go/bin
export GOROOT=/usr/local/go
export GOBIN=/Users/用户x/go/go_path/bin


readme 在线：https://maxiang.io/