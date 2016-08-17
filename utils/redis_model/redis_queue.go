package redis_model

import (
    "encoding/json"
    "time"
    "log"
)

// Wraps the Redis client to meet the Cache interface.
type RedisQueue struct {
    queue_client              *RedisClient
    key string
}

type register_func func([]interface {})

type dic_type map[string]interface{}

// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisQueue(key string) *RedisQueue {
    rclient:= NewRedisCache("127.0.0.1:6379","",time.Second * 60)
    return &RedisQueue{rclient,key}
}

/* Queue Function */
func (c *RedisQueue) ASync(value dic_type) error {

    v, _ := json.Marshal(value)
    log.Println("[Marshal]",v)
    return c.queue_client.LPush(c.key,v)
}

func (c *RedisQueue) Do(f register_func) {
     for {

        v , err := c.queue_client.BRpop(c.key)

        //log.Println("[brpop]",v,err)

        // var dat dic_type
        // if err := json.Unmarshal(v, &dat); err != nil {
        //     log.Println("[Do error]",err)
        //     continue
        // }

        f(v)

        if err!=nil {
            log.Println("[redis queue err]")
            log.Println(err)
        }
    }
}

