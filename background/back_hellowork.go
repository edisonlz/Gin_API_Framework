package main

import ( "Gin_API_Framework/utils/redis_model"
        "log")


func sync_hello(value map[string]interface{}) {

    log.Println("[sync_hello]...")
    log.Println(value)
}

func main(){

    queue := NewRedisQueue("channel.test")
    value := map[string]int{"hello": 1, "world": 2}
    queue.ASync(value)
    queue.Do(sync_hello)

}
