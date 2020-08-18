package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "reflect"
)

type Model struct {
    gorm.Model
}

var pool map[string]*gorm.DB

type BaseModel struct {
    Driver   string
    Database string
    Table    string
    Connect  string
    db       *gorm.DB
    Entity   interface{}
}

type Filter map[string]interface{}

func init() {
    pool = make(map[string]*gorm.DB)
}

func (m *BaseModel) init() {
    if m.Driver == "" {
        m.Driver = "mysql"
    }
    if m.Connect == "" {
        m.Connect = "default"
    }

    dbConf, ok := utils.Config.Database[m.Connect]
    if !ok {
        panic(fmt.Sprintf("database %s not found", m.Connect))
    }

    m.db = getDBInPool(dbConf, m)
}

func getDBInPool(conf utils.Database, m *BaseModel) *gorm.DB {
    server := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
    dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Passwd, server, m.Database)
    //if _db := pool[dsn]; _db != nil {
    //    return _db
    //}
    logger().Info(dsn)
    logger().Info(m.Table)
    db, err := gorm.Open(m.Driver, dsn)
    if err != nil {
        panic("failed to connect database")
    }
    // TODO 连接池
    db.DB().SetMaxOpenConns(10) //设置数据库连接池最大连接数
    db.DB().SetMaxIdleConns(2)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
    db = db.Table(m.Table).Debug()
    pool[dsn] = db
    return pool[dsn]
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
