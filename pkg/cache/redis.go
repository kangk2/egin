package cache

import (
    "context"
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/go-redis/redis/v8"
    "time"
)

var ctx = context.Background()

type Redis struct {
    rdb        *redis.Client
    Connection string
}

func (r *Redis) init() {
    if r.rdb != nil {
        return
    }

    logger := utils.Logger.Channel("redis")
    _, ok := utils.Config.Database[r.Connection]
    if !ok {
        logger.Error(fmt.Sprintf("redis connection %s not found", r.Connection))
    }

    rdb, ok := getDBInPool(r.Connection)
    if ok {
        r.rdb = rdb
    } else {
        logger.Error(fmt.Sprintf("get %s rdb failed", r.Connection))
    }
}

func (r *Redis) Get(key string) (string, error) {
    r.init()
    return r.rdb.Get(ctx, key).Result()
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
    r.init()
    return r.rdb.Set(ctx, key, value, expiration).Err()
}