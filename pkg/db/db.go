package db

import (
    "database/sql"
    "fmt"
    "sync"

    _ "github.com/go-sql-driver/mysql"

    "github.com/daodao97/egin/pkg/utils"
)

var pool sync.Map

func init() {
    InitDb()
}

// 系统启动时, 初始化所有连接的 db
// 注意此时, database/sql 此时并不会产生任何mysql连接, 连接将会在使用时创建
func InitDb() {
    dbConf := utils.Config.Database
    for key, conf := range dbConf {
        db := makeDb(conf)
        pool.Store(key, db)
    }
}

// 生成 原生 DB 对象
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

// 一般用Prepared Statements和Exec()完成INSERT, UPDATE, DELETE操作
func Exec(db *sql.DB, _sql string, args ...interface{}) (int64, int64, error) {

    tx, err := db.Begin()
    if err != nil {
        logger().Error(err)
        return 0, 0, err
    }
    var flag bool
    addr := &flag
    defer func(flag *bool, errMsg *error) {
        if *flag {
            return
        }
        err := tx.Rollback()
        if err != nil {
            logger().Error(fmt.Sprintf("db.exec.rollback fail: %s", err), map[string]interface{}{
                "sql":  _sql,
                "args": args,
                "msg":  errMsg,
            })
        } else {
            logger().Info("db.exec.rollback", map[string]interface{}{
                "sql":  _sql,
                "args": args,
                "msg":  errMsg,
            })
        }
    }(addr, &err)

    stmt, err := tx.Prepare(_sql)
    if err != nil {
        logger().Error("db.exec.prepare fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return 0, 0, err
    }

    res, err := stmt.Exec(args...)
    if err != nil {
        logger().Error("db.exec.exec fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return 0, 0, err
    }

    err = tx.Commit()
    if err != nil {
        logger().Error("db.exec.commit fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return 0, 0, err
    }
    flag = true

    lastId, err := res.LastInsertId()
    if err != nil {
        logger().Error("db.exec.lastId fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return 0, 0, err
    }

    affected, err := res.RowsAffected()
    if err != nil {
        logger().Error("db.exec.affected fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return 0, 0, err
    }

    return lastId, affected, nil
}

func Query(db *sql.DB, _sql string, args ...interface{}) ([]map[string]string, error) {
    var result []map[string]string

    stmt, err := db.Prepare(_sql)

    if err != nil {
        logger().Error("db.query.prepare fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return result, err
    }

    rows, err := stmt.Query(args...)
    if err != nil {
        logger().Error("db.query.query fail", map[string]interface{}{
            "sql":  _sql,
            "args": args,
            "msg":  err,
        })
        return result, err
    }

    return Rows2SliceMap(rows)
}

// 将 sql.Rows 的结果转换为 map
// 注意这里所有的value 均为string
// TODO 是否可以根据 rows.ColumnTypes() databasesType 做类型转换?
func Rows2SliceMap(rows *sql.Rows) (list []map[string]string, err error) {
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
    if err := rows.Err(); err != nil {
        return list, err
    }
    return list, nil
}
