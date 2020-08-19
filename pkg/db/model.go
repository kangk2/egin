package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/davecgh/go-spew/spew"
    "github.com/jinzhu/gorm"
    "reflect"
)

type Model struct {
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

func (m *BaseModel) Get(filter Filter, attr Attr) interface{} {
    if m.db == nil {
        m.init()
    }

    // result := reflect.New(reflect.TypeOf(m.Entity)).Interface()

    sliceType := reflect.SliceOf(reflect.TypeOf(m.Entity))
    result := reflect.New(sliceType).Interface()

    sql := fmt.Sprintf(SelectSql(filter, attr), m.Table)
    spew.Dump(sql)
    db := m.db.Raw(sql)
    db.Scan(result)

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
