package docs

import (
    "encoding/json"
    "strings"

    "github.com/astaxie/beego"
    "Gin_API_Framework/api/docs/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/cms","description":"CMS API\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/cms":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/cms","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/staticblock/:key","description":"","operations":[{"httpMethod":"GET","nickname":"getStaticBlock","type":"","summary":"get all the staticblock by key","parameters":[{"paramType":"path","name":"key","description":"\"The email for login\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"Invalid email supplied","responseModel":""},{"code":404,"message":"User not found","responseModel":""}]}]},{"path":"/products","description":"","operations":[{"httpMethod":"GET","nickname":"Get Product list","type":"","summary":"Get Product list by some info","parameters":[{"paramType":"query","name":"category_id","description":"\"category id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"brand_id","description":"\"brand id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"query","description":"\"query of search\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"segment","description":"\"segment\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sort","description":"\"sort option\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"dir","description":"\"direction asc or desc\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"offset\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"count limit\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"price","description":"\"price\"","dataType":"float","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"special_price","description":"\"whether this is special price\"","dataType":"bool","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"size","description":"\"size filter\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"color","description":"\"color filter\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"format","description":"\"choose return format\"","dataType":"bool","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"{object}","responseModel":""},{"code":400,"message":"no enough input","responseModel":""},{"code":500,"message":"get products common error","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
    if beego.BConfig.WebConfig.EnableDocs {
        err := json.Unmarshal([]byte(Rootinfo), &rootapi)
        if err != nil {
            beego.Error(err)
        }
        err = json.Unmarshal([]byte(Subapi), &apilist)
        if err != nil {
            beego.Error(err)
        }
        beego.GlobalDocAPI["Root"] = rootapi
        for k, v := range apilist {
            for i, a := range v.APIs {
                a.Path = urlReplace(k + a.Path)
                v.APIs[i] = a
            }
            v.BasePath = BasePath
            beego.GlobalDocAPI[strings.Trim(k, "/")] = v
        }
    }
}


func urlReplace(src string) string {
    pt := strings.Split(src, "/")
    for i, p := range pt {
        if len(p) > 0 {
            if p[0] == ':' {
                pt[i] = "{" + p[1:] + "}"
            } else if p[0] == '?' && p[1] == ':' {
                pt[i] = "{" + p[2:] + "}"
            }
        }
    }
    return strings.Join(pt, "/")
}
