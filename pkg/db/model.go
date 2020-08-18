package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "reflect"
    "sync"
)

type Model struct {
    gorm.Model
}

var pool sync.Map

type BaseModel struct {
    Driver  string
    Table   string
    Connection string
    db      *gorm.DB
    Entity  interface{}
}

type Filter map[string]interface{}

func init() {
    InitDb()
}

func (m *BaseModel) init() {
    if m.Driver == "" {
        m.Driver = "mysql"
    }
    if m.Connection == "" {
        m.Connection = "default"
    }
    conf, ok := utils.Config.Database[m.Connection]
    if !ok {
        panic(fmt.Sprintf("database %s not found", m.Connection))
    }

    key := fmt.Sprintf("%s:%d:%s", conf.Host, conf.Port, conf.Database)

    db, ok := getDBInPool(key)
    if !ok {
        panic("not found db conn")
    }
    m.db = db.Table(m.Table)
}

func getDBInPool(key string) (*gorm.DB, bool) {
    val, ok := pool.Load(key)
    return val.(*gorm.DB), ok
}

func InitDb() {
    dbConf := utils.Config.Database
    for _, conf := range dbConf {
        go func(conf utils.Database) {
            key := fmt.Sprintf("%s:%d:%s", conf.Host, conf.Port, conf.Database)
            pool.Store(key, makeDb(conf))
        }(conf)
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

    //设置数据库连接池最大连接数
    var MaxOpenConns int
    if conf.Pool.MaxOpenConns == 0 {
        MaxOpenConns = 100
    } else {
        MaxOpenConns = conf.Pool.MaxOpenConns
    }
    db.DB().SetMaxOpenConns(MaxOpenConns)

    //连接池最大允许的空闲连接数
    //如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
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

func (m *BaseModel) Get(filter ...Filter) interface{} {
    if m.db == nil {
        m.init()
    }
    var count int
    m.db.Count(&count)

    //result := reflect.New(reflect.TypeOf(m.Entity)).Interface()

    sliceType := reflect.SliceOf(reflect.TypeOf(m.Entity))
    result := reflect.New(sliceType).Interface()

    m.db.Model(m.Entity).Where("id > ?", 0).Scan(result)

    //defer m.db.Close()
    return result
}
func (m *BaseModel) Insert() {

}
func (m *BaseModel) Update() {

}
func (m *BaseModel) Del() {

}

func logger() utils.LoggerInstance {
    return utils.Logger.Channel("mysql")
}
