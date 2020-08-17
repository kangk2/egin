package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "reflect"
)

type Model struct {
    gorm.Model
}

type BaseModel struct {
    Driver   string `default:"mysql"`
    Database string
    Table    string
    db       *gorm.DB
    Entity   interface{}
}

type Filter map[string]interface{}

func (m *BaseModel) init() {
    lib.CompareDefaultTag(m)
    user := "root"
    pwd := "root"
    dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, m.Database)
    db, err := gorm.Open(m.Driver, dsn)
    if err != nil {
        panic("failed to connect database")
    }
    //defer db.Close()
    m.db = db.Table(m.Table)
}

func (m *BaseModel) Get() interface{} {
    if m.db == nil {
        m.init()
    }
    var count int
    m.db.Count(&count)

    //result := reflect.New(reflect.TypeOf(m.Entity)).Interface()

    sliceType := reflect.SliceOf(reflect.TypeOf(m.Entity))
    result := reflect.New(sliceType).Interface()

    m.db.Where("id > ?", 0).Scan(result)

    return result
}
func (m *BaseModel) Insert() {

}
func (m *BaseModel) Update() {

}
func (m *BaseModel) Del() {

}
