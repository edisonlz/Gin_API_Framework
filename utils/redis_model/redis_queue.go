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

type register_func func(map[string]interface{})

type dic_type map[string]interface{}

// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisQueue(key string) *RedisQueue {
    rclient:= NewRedisCache("127.0.0.1:6379","",time.Second * 60)
    return &RedisQueue{rclient,key}
}

/* Queue Function */
func (c *RedisQueue) ASync(value dic_type) {

    v, _ := json.Marshal(value)
    vv := string(v)
    log.Println("[Marshal]",vv)
    c.queue_client.LPush(c.key,vv)
}

func (c *RedisQueue) Do(f register_func) {
     for {

        v := c.queue_client.BRpop(c.key)

        if v!=nil {
            value := v[1]
            var dat dic_type
            if err := json.Unmarshal([]byte(value), &dat); err != nil {
                log.Println("[Do error]",err)
                continue
            }
            f(dat)
        }
    }
}

