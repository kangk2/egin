package main

import (
    "database/sql"
    "fmt"
    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/davecgh/go-spew/spew"
    _ "github.com/go-sql-driver/mysql"
    "html/template"
    "log"
    "sync"
)

var wg sync.WaitGroup

func main() {
    // testBaseSelectSql()
    // testInsert()
    api()
    fmt.Println("over")
}

func testBaseSelectSql() {
    result, err := baseSelectSql()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
    spew.Dump(result)
}
func baseSelectSql() ([]map[string]string, error) {
    mydb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hyperf_admin")
    var result []map[string]string
    if err != nil {
        log.Fatal(err)
        return result, err
    }
    _sql := fmt.Sprintf("select `id`, `username`,`realname`,`password` from `user` where id > ?")

    stmt, err := mydb.Prepare(_sql)
    if err != nil {
        log.Fatal(err)
        return result, err
    }

    var args = []interface{}{0}
    rows, err := stmt.Query(args...)
    if err != nil {
        log.Fatal(err)
        return result, err
    }

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
    //
    return
}

func testInsert() {
    lastId, err := baseInertSql()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(lastId)
}
func baseInertSql() (int64, error) {
    mydb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hyperf_admin")
    if err != nil {
        log.Fatal(err)
        return 0, err
    }

    _sql := fmt.Sprintf("insert into user (`username`,`realname`,`password`) values (?, ?, ?)")

    stmt, err := mydb.Prepare(_sql)
    if err != nil {
        log.Fatal(err)
        return 0, err
    }

    var args = []interface{}{"test%d", "你好", "cool"}
    res, err := stmt.Exec(args...)
    if err != nil {
        log.Fatal(err)
        return 0, err
    }

    lastId, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
        return 0, err
    }

    return lastId, nil
}

func cust() {
    // filter := db.Filter{
    //     // "name": "nike",
    //     // "id":   2,
    //     // "type": []string{"on", "off"},
    //     // "sex":  []int{1, 2},
    //     // "age": map[string]string{
    //     //     ">":       "1",
    //     //     "<":       "2",
    //     //     "__logic": "or",
    //     // },
    // }
    // fmt.Println(db.FilterToQuery(filter))
    // attr := db.Attr{
    //     // Select: []string{"name", "age"},
    //     // OrderBy: "age desc",
    // }
    // fmt.Println(db.AttrToQuery(attr))
    //
    // fmt.Println(db.SelectSql(filter, attr))

    // fmt.Println(db.FilteredSQLInject("SELECT * FROM user WHERE username='myuser' or 'foo' = 'foo' --'' AND password='xxx'"))

    filter := db.Filter{
        "username": "daodao",
        // "password":      "fo'ot",
    }

    attr := db.Attr{
        Select: []string{"username"},
    }

    user := model.UserModel
    spew.Dump(user.Get(filter, attr))

    fmt.Println(template.HTMLEscapeString("s234@#%@$%'"))
}

func api() {
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(index int) {
            res, err := lib.Get("http://127.0.0.1:8080/user", map[string]string{}, map[string]string{})
            if err != nil {
                log.Println("ERROR", err)
            } else {
                log.Println("RESULT", res)
                fmt.Println(index, res.StatusCode == 200)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
