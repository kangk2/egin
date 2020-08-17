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
    Driver   string
    Database string
    Table    string
    Connect  string
    db       *gorm.DB
    Entity   interface{}
}

type Filter map[string]interface{}

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

    dsn := fmt.Sprintf(
        "%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        dbConf.User,
        dbConf.Passwd,
        fmt.Sprintf("%s:%d", dbConf.Host, dbConf.Port),
        m.Database,
    )
    utils.Logger.Info("222222222")
    fmt.Println(dsn)
    db, err := gorm.Open(m.Driver, dsn)
    if err != nil {
        panic("failed to connect database")
    }
    //defer db.Close()
    m.db = db.Table(m.Table).Debug()
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

    return result
}
func (m *BaseModel) Insert() {

}
func (m *BaseModel) Update() {

}
func (m *BaseModel) Del() {

}
