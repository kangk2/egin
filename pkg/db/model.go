package db

import (
    "database/sql"
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    "strings"
    "time"
)

type BaseModel struct {
    Driver     string
    Table      string
    Connection string
    db         *sql.DB
    Entity     interface{}
    LastSql    string
    ready      bool
}

func (m *BaseModel) init() bool {
    if m.ready {
        return true
    }
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
    m.db = db
    m.ready = true
    return true
}

func getDBInPool(key string) (*sql.DB, bool) {
    val, ok := pool.Load(key)
    return val.(*sql.DB), ok
}

func (m *BaseModel) Get(filter Filter, attr Attr) ([]map[string]string, error) {
    m.init()

    var args []interface{}

    sqlField := AttrToSelectQuery(attr)
    sqlWhere, args1 := FilterToQuery(filter)
    sqlAttr, args2 := AttrToQuery(attr)

    _sql := strings.Trim(fmt.Sprintf("select %s from %s %s %s", sqlField, m.Table, sqlWhere, sqlAttr), " ")

    logger().Info(_sql)

    args = append(args1, args2...)

    return Query(m.db, _sql, args...)
}

func (m BaseModel) Insert(record Record) (int64, error) {
    m.init()
    sqlInsert, args := InsertRecodeToQuery(record)

    _sql := fmt.Sprintf(sqlInsert, m.Table)

    stmt, err := Prepare(m.db, _sql)

    logger := logger().Channel("mysql")
    if err != nil {
        logger.Error(err)
        return 0, err
    }

    res, err := stmt.Exec(args...)
    if err != nil {
        logger.Error(err)
        return 0, err
    }

    lastId, err := res.LastInsertId()
    if err != nil {
        logger.Error(err)
        return 0, err
    }

    return lastId, nil
}

func (m BaseModel) Update(entity map[string]interface{}) {}

func (m BaseModel) Del() {}

func (m BaseModel) Exec(sql string, result interface{}) interface{} {
    m.init()
    defer timeCost()(m)
    m.LastSql = sql
    return result
}

func (m *BaseModel) DB() *sql.DB {
    m.init()
    return m.db
}

func logger() utils.LoggerInstance {
    return utils.Logger.Channel("mysql")
}

func timeCost() func(m BaseModel) {
    start := time.Now()
    return func(m BaseModel) {
        tc := time.Since(start)
        logger().Info(
            fmt.Sprintf("use time %v", tc),
            map[string]interface{}{
                "table":      m.Table,
                "connection": m.Connection,
                "sql":        m.LastSql,
                "ums ":       fmt.Sprintf("%v", tc),
            })
    }
}
