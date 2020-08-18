package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/jinzhu/gorm"
    "sync"
)

var pool sync.Map

func init() {
    InitDb()
}

func InitDb() {
    dbConf := utils.Config.Database
    for _, conf := range dbConf {
        key := fmt.Sprintf("%s:%d:%s", conf.Host, conf.Port, conf.Database)
        pool.Store(key, makeDb(conf))
    }
}

func makeDb(conf utils.Database) *gorm.DB {
    server := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
    dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Passwd, server, conf.Database)
    driver := conf.Driver
    if driver == "" {
        driver = "mysql"
    }

    db, err := gorm.Open(driver, dsn)
    if err != nil {
        panic(fmt.Sprintf("failed Connection database: %s", err))
    }

    // 设置数据库连接池最大连接数
    var MaxOpenConns int
    if conf.Pool.MaxOpenConns == 0 {
        MaxOpenConns = 100
    } else {
        MaxOpenConns = conf.Pool.MaxOpenConns
    }
    db.DB().SetMaxOpenConns(MaxOpenConns)

    // 连接池最大允许的空闲连接数
    // 如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
    var MaxIdleConns int
    if conf.Pool.MaxIdleConns == 0 {
        MaxIdleConns = 20
    } else {
        MaxIdleConns = conf.Pool.MaxIdleConns
    }
    db.DB().SetMaxIdleConns(MaxIdleConns)
    db.Debug()
    return db
}
