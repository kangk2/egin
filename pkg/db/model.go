package db

import (
    "database/sql"
    "fmt"
    "strings"
    "time"

    "github.com/daodao97/egin/pkg/utils"
)

// 从池子中捞出一个db对象, 注意, 这里并非 mysql 连接池
func getDBInPool(key string) (*sql.DB, bool) {
    val, ok := pool.Load(key)
    return val.(*sql.DB), ok
}

// Model 的基础封装
type BaseModel struct {
    Driver     string
    Table      string
    Connection string
    db         *sql.DB
    Entity     interface{}
    LastSql    string
    FakeDelete bool
    FakeDelKey string
    ready      bool
    ActionType string
}

// Model 基础信息的初始化
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
    if m.FakeDelKey == "" {
        m.FakeDelKey = "is_deleted"
    }
    _, ok := utils.Config.Database[m.Connection]
    if !ok {
        logger().Error(fmt.Sprintf("database %s not found", m.Connection))
    }

    db, ok := getDBInPool(m.Connection)
    if !ok {
        panic("not found db conn")
    }
    m.db = db
    m.ready = true
    return true
}

// 数据库查询
// filter {sex:0, class:{in:[1,2]}}
// attr {Select:[id,name], OrderBy:"id desc"}
// sql: select `id`, `name` from ${table} where `sex` = ? and class in (?, ?)
// result: 查询结果, 错误信息
func (m *BaseModel) Get(filter Filter, attr Attr) ([]map[string]string, error) {
    m.init()
    defer timeCost()(m)

    var args []interface{}
    if m.FakeDelete {
        filter[m.FakeDelKey] = 0
    }
    sqlWhere, args1 := FilterToQuery(filter)
    sqlField := AttrToSelectQuery(attr)
    sqlAttr, args2 := AttrToQuery(attr)

    _sql := strings.Trim(fmt.Sprintf("select %s from %s %s %s", sqlField, m.Table, sqlWhere, sqlAttr), " ")
    args = append(args1, args2...)
    m.LastSql = _sql
    m.ActionType = "read"

    return Query(m.db, _sql, args...)
}

// record: {name:"Joke", sex:1}
// sql: insert into ${table} (`name`, `sex`) values (?, ?)
// return: 最后的记录id, 受影响行数, 错误信息
func (m *BaseModel) Insert(record Record) (int64, int64, error) {
    m.init()
    defer timeCost()(m)

    sqlInsert, args := InsertRecordToQuery(record)
    _sql := fmt.Sprintf("insert into %s %s", m.Table, sqlInsert)
    m.LastSql = _sql
    m.ActionType = "create"

    return Exec(m.db, _sql, args...)
}

// filter: {id: 1}
// record: {name:"Joke", sex:1}
// sql: update ${table} set `name` = ?, `sex` = ? where `id` = ?
// return: 最后的记录id, 受影响行数, 错误信息
func (m *BaseModel) Update(filter Filter, record Record) (int64, int64, error) {
    m.init()
    defer timeCost()(m)

    sqlUpdate, argsUpdate := UpdateRecordToQuery(record)
    sqlWhere, argsWhere := FilterToQuery(filter)
    _sql := fmt.Sprintf("update `%s` set %s %s", m.Table, sqlUpdate, sqlWhere)
    args := append(argsUpdate, argsWhere...)
    m.LastSql = _sql
    m.ActionType = "update"

    return Exec(m.db, _sql, args...)
}

// filter: {id: 1}
// record: {name:"Joke", sex:1}
// 物理删除sql: delete from ${table} where `id` = ?
// 伪删除sql: update ${table} set ${FakeDelKey} = 1 where `id` = ?
// return: 最后的记录id, 受影响行数, 错误信息
func (m *BaseModel) Delete(filter Filter) (int64, int64, error) {
    m.init()
    defer timeCost()(m)

    if m.FakeDelete {
        return m.Update(filter, Record{m.FakeDelKey: 1})
    }

    sqlWhere, args := FilterToQuery(filter)

    _sql := fmt.Sprintf("delete from %s %s", m.Table, sqlWhere)
    m.LastSql = _sql
    m.ActionType = "delete"

    return Exec(m.db, _sql, args...)
}

// 获取 db原生对象, 可以执行原生sql语句等更多操作
func (m *BaseModel) DB() *sql.DB {
    m.init()
    return m.db
}

func logger() utils.LoggerInstance {
    return utils.Logger.Channel("mysql")
}

// 记录 sql, 耗时 等信息
func timeCost() func(m *BaseModel) {
    start := time.Now()
    return func(m *BaseModel) {
        tc := time.Since(start)
        logger().Info(
            fmt.Sprintf("use time %v", tc),
            map[string]interface{}{
                "table":      m.Table,
                "connection": m.Connection,
                "sql":        m.LastSql,
                "type":       m.ActionType,
                "ums":        fmt.Sprintf("%v", tc),
            })
    }
}
