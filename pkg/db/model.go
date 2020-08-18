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

type BaseModel struct {
    Driver     string
    Table      string
    Connection string
    db         *gorm.DB
    Entity     interface{}
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

func (m *BaseModel) Get(filter ...Filter) interface{} {
    if m.db == nil {
        m.init()
    }
    var count int
    m.db.Count(&count)

    // result := reflect.New(reflect.TypeOf(m.Entity)).Interface()

    sliceType := reflect.SliceOf(reflect.TypeOf(m.Entity))
    result := reflect.New(sliceType).Interface()

    m.db.Model(m.Entity).Where("id > ?", 0).Scan(result)

    // defer m.db.Close()
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
