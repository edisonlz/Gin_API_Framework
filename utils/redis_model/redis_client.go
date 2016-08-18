package redis_model

import (
    "gopkg.in/redis.v4"
    "time"
    _"encoding/json"
    "log"
)

// Wraps the Redis client to meet the Cache interface.
type RedisClient struct {
    pool     *redis.Client
}

// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(host string, password string, defaultExpiration time.Duration) *RedisClient {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "", 
        MaxRetries: 3,
        DialTimeout: 5 * time.Second,
        WriteTimeout:3 * time.Second,
        PoolSize:20,
        PoolTimeout: 0,
        IdleTimeout:0,
        DB:       0,  
    })

    pong, err := client.Ping().Result()
    log.Println(pong, err)

    c := &RedisClient{client}
    return c

}

/* Queue Function */
func (c *RedisClient) LPush(key string,  value string) {
    cmd := c.pool.LPush(key,value)
    log.Println(cmd)
}


func (c *RedisClient) BRpop(key string) []string {
    data := c.pool.BRPop(0,key)
    return data.Val() 
}

