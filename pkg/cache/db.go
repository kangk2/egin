package cache

import (
    "fmt"
    "sync"

    "github.com/go-redis/redis/v8"

    "github.com/daodao97/egin/pkg/utils"
)

var pool sync.Map

func init() {
    InitDb()
}

// 系统启动时, 初始化所有连接的 db
// 注意此时, database/sql 此时并不会产生任何mysql连接, 连接将会在使用时创建
func InitDb() {
    dbConf := utils.Config.Redis
    for key, conf := range dbConf {
        db := makeDb(conf)
        pool.Store(key, db)
    }
}

func makeDb(conf utils.Redis) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
        Password: conf.Password,
        DB:       conf.DB,
    })

    rdb.AddHook(&logger{})

    return rdb
}

func getDBInPool(key string) (*redis.Client, bool) {
    val, ok := pool.Load(key)
    return val.(*redis.Client), ok
}
