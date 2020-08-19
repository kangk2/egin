package main

import (
    "fmt"
    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/davecgh/go-spew/spew"
    "log"
    "sync"
    "html/template"
)

var wg sync.WaitGroup

func main() {
    sql()
    fmt.Println("over")
}

func sql() {
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
    for i := 0; i < 1000; i++ {
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
