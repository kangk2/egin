package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/daodao97/egin/pkg/utils"
)

var ctx = context.Background()

type Redis struct {
	rdb        *redis.Client
	Connection string
}

var logger = utils.Logger.Channel("redis")

func (r *Redis) init() {
	if r.rdb != nil {
		return
	}

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
