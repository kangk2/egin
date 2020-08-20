package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "strings"
)

type Filter map[string]interface{}

type Attr struct {
    Select  []string
    Limit   int
    Offset  int
    OrderBy string
    GroupBy string
    Having  string
}

type Record map[string]interface{}

type Records map[int]Record

func InsertRecodeToQuery(record Record) (string, []interface{}) {
    var fields []string
    var value []string
    var args []interface{}
    for k, v := range record {
        fields = append(fields, k)
        value = append(value, "?")
        args = append(args, v)
    }

    _sql := fmt.Sprintf(
        "insert into %s (%s) values (%s)",
        "%s",
        fmt.Sprintf("`%s`", strings.Join(fields, "`,`")),
        fmt.Sprintf("%s", strings.Join(value, ",")),
    )

    return _sql, args
}

func FilterToQuery(filter Filter) (string, []interface{}) {
    var sql string
    var args []interface{}
    if len(filter) == 0 {
        return sql, args
    }
    logic := filter["__logic"]
    if logic == nil {
        logic = "and"
    }
    var scopes []string

    for k, v := range filter {
        switch v.(type) {
        case []int, []string:
            length := len(v.([]interface{}))
            pl := make([]string, length)
            for _, val := range v.([]interface{}) {
                args = append(args, val)
            }
            scopes = append(scopes, fmt.Sprintf("`%s` in (%s)", k, strings.Join(pl, ",")))
        case map[string]string:
            var _scope []string
            var _logic string = "and"
            for op, val := range v.(map[string]string) {
                _, found := lib.Find([]string{">", ">=", "<", "<=", "=", "<>", "!="}, op)
                if found {
                    _scope = append(_scope, fmt.Sprintf("`%s` %s ?", k, op))
                }
                if op == "__logic" {
                    _logic = val
                }
                args = append(args, val)
            }
            scopes = append(scopes, fmt.Sprintf("(%s)", strings.Join(_scope, fmt.Sprintf(" %s ", _logic))))
        default:
            args = append(args, v)
            scopes = append(scopes, fmt.Sprintf("`%s` = ?", k))
        }
    }
    sql = fmt.Sprintf("where %s", strings.Join(scopes, fmt.Sprintf(" %s ", logic.(string))))
    return sql, args
}

func AttrToQuery(attr Attr) (string, []interface{}) {
    var sql string
    var scopes []string
    var args []interface{}
    if attr.Offset != 0 {
        args = append(args, attr.Offset)
        scopes = append(scopes, "offset ?")
    }
    if attr.Limit != 0 {
        args = append(args, attr.Limit)
        scopes = append(scopes, "limit ?")
    }
    if attr.OrderBy != "" {
        args = append(args, attr.OrderBy)

        scopes = append(scopes, "order by ?")
    }
    if attr.GroupBy != "" {
        args = append(args, attr.GroupBy)

        scopes = append(scopes, "group by ?")
    }
    sql = strings.Join(scopes, " ")
    return sql, args
}

func AttrToSelectQuery(attr Attr) string {
    if len(attr.Select) == 0 {
        return "*"
    }
    return fmt.Sprintf("`%s`", strings.Join(attr.Select, "`,`"))
}
