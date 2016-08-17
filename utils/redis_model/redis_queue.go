package redis_model

import (
    "github.com/garyburd/redigo/redis"
    "encoding/json"
    "time"
    "log"
)

// Wraps the Redis client to meet the Cache interface.
type RedisQueue struct {
    queue_client              *RedisClient
    key string
}

type register_func func(string)


// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisQueue(key string) *RedisClient {
    rclient:= NewRedisQueue("127.0.0.1:6379","",time.Second * 60)
    return &RedisQueue{rclient,key}
}

/* Queue Function */
func (c *RedisQueue) ASync(value map[string]interface{}) error {

    v, _ := string(json.Marshal(value))
    return c.queue_client.LPush(c.key,v)
}

func (c *RedisClient) Do(f register_func) {
     for {

        v , err := c.queue_client.BRpop(c.key)

        var dat map[string]interface{}
        if err := json.Unmarshal(v, &dat); err != nil {
            panic(err)
        }

        f(dat)

        if err!=nil {
            log.Println("[redis queue err]")
            log.Println(err)
        }
    }
}

