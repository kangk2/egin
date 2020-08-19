package db

import (
    "database/sql"
    "fmt"
    "github.com/daodao97/egin/pkg/utils"
    _ "github.com/go-sql-driver/mysql"
    "log"
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
        db := makeDb(conf)
        pool.Store(key, db)
    }
}

func makeDb(conf utils.Database) *sql.DB {
    server := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
    dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Passwd, server, conf.Database)
    driver := conf.Driver
    if driver == "" {
        driver = "mysql"
    }

    db, err := sql.Open(driver, dsn)
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
    db.SetMaxOpenConns(MaxOpenConns)

    // 连接池最大允许的空闲连接数
    // 如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
    var MaxIdleConns int
    if conf.Pool.MaxIdleConns == 0 {
        MaxIdleConns = 20
    } else {
        MaxIdleConns = conf.Pool.MaxIdleConns
    }
    db.SetMaxIdleConns(MaxIdleConns)
    return db
}

func Prepare(db *sql.DB, _sql string) (*sql.Stmt, error) {
    return db.Prepare(_sql)
}

func Query(db *sql.DB, _sql string, args ...interface{}) ([]map[string]string, error) {
    var result []map[string]string

    stmt, err := db.Prepare(_sql)

    if err != nil {
        log.Fatal(err)
        return result, err
    }

    rows, err := stmt.Query(args...)
    if err != nil {
        log.Fatal(err)
        return result, err
    }

    // TODO 字段类型转换
    return Rows2SliceMap(rows), nil
}

func Rows2SliceMap(rows *sql.Rows) (list []map[string]string) {
    // 字段名称
    columns, _ := rows.Columns()
    // 多少个字段
    length := len(columns)
    // 每一行字段的值
    values := make([]sql.RawBytes, length)
    // 保存的是values的内存地址
    pointer := make([]interface{}, length)
    //
    for i := 0; i < length; i++ {
        pointer[i] = &values[i]
    }
    //
    for rows.Next() {
        // 把参数展开，把每一行的值存到指定的内存地址去，循环覆盖，values也就跟着被赋值了
        rows.Scan(pointer...)
        // 每一行
        row := make(map[string]string)
        for i := 0; i < length; i++ {
            row[columns[i]] = string(values[i])
        }
        list = append(list, row)
    }
    return list
}
