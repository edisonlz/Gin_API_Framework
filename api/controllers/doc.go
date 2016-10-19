package controllers


import (
    "github.com/gin-gonic/gin"
    "net/http"
    //"fmt"
    _ "github.com/astaxie/beego"
    _ "io/ioutil"
    _ "path"
    _ "runtime"
    "encoding/json"
    "Gin_API_Framework/utils"
    "io/ioutil"
    "path"
)



// @Title Doc
// @API_GROUP doc
// @Description 文档
// @Success 200 {object} 
// @router /doc [get]
func TemplateDocHandler(c *gin.Context) {

    //dat := `{"Api":{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":null,"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}},"Subapi":{"User":[{"id":"8c230a505af1f73ea03e4fe84a679dd0","path":"/user/login","description":"","operations":[{"httpMethod":"GET","nickname":"User Login","type":"","summary":"用户登录接口","parameters":[{"paramType":"query","name":"name","description":"\"user name\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"pwd","description":"\"password\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get  common error","responseModel":""}]}]},{"id":"b54784ebcd9b6c0e1052ec48bde5df5f","path":"/user/logout","description":"","operations":[{"httpMethod":"GET","nickname":"User Logout","type":"","summary":"用户登出接口","parameters":[{"paramType":"query","name":"name","description":"\"user name\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get  common error","responseModel":""}]}]},{"id":"1b6c901f51f52722451a564ae0d8af5d","path":"/user/create","description":"","operations":[{"httpMethod":"GET","nickname":"User Create","type":"","summary":"创建用户接口","parameters":[{"paramType":"query","name":"name","description":"\"user name\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"gender","description":"\"gender\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"phone","description":"\"phone\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get  common error","responseModel":""}]}]},{"id":"53519e6725588e2528fad1b2b08c83ff","path":"/user/query","description":"","operations":[{"httpMethod":"GET","nickname":"User Query By ID","type":"","summary":"查询用户接口通过用户ID","parameters":[{"paramType":"query","name":"uid","description":"\"user id\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get  common error","responseModel":""}]}]},{"id":"1b6c901f51f52722451a564ae0d8af5d","path":"/user/create","description":"","operations":[{"httpMethod":"GET","nickname":"User List","type":"","summary":"查询用户列表接口","responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get  common error","responseModel":""}]}]}],"doc":[{"id":"ffa365519ecac3ed03b9aa7125239e5b","path":"/api_docs","description":"","operations":[{"httpMethod":"GET","nickname":"User List","type":"","summary":"查询用户列表接口","responseMessages":[{"code":200,"message":"{object}","responseModel":""}]}]}]}}`
    curpath := utils.CallerSourcePath()
    //fmt.Println(curpath)
    tpath := path.Join(curpath, "api_doc.json")
    //fmt.Println(tpath)

    data, _ := ioutil.ReadFile(tpath)
    dat := string(data)

    var s map[string]interface{}
    json.Unmarshal([]byte(dat),&s)

    // fmt.Println(" * * * * * * * * * * * *")
    // fmt.Println(s)
    
    c.HTML(http.StatusOK, "api_docs.tmpl", gin.H{
            "api" : s["Subapi"],
            "api_base" : "",
    })

}



