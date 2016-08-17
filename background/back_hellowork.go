package main

import ( 
        "Gin_API_Framework/utils/redis_model"
        "log"
)


func sync_hello(dic []interface {}) {

    log.Println("[sync_hello]...")
    log.Println(dic)

    for key,value:=range dic {
        log.Println(key,value)
    }

}



func main(){

    queue := redis_model.NewRedisQueue("channel.test")
    value := map[string]interface{}{}
    value["hello"] = 1
    value["world"] = 2

    queue.ASync(value)

    queue.Do(sync_hello)

}
