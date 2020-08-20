package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "strings"
)

// sql 片段的构造

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

//record {"field":"val", "field2":"val2"}
//return arg1: (`field`, `field2`) values (?, ?)
//return arg2: ["val", "val2"]
func InsertRecordToQuery(record Record) (string, []interface{}) {
    var fields []string
    var value []string
    var args []interface{}
    for k, v := range record {
        fields = append(fields, k)
        value = append(value, "?")
        args = append(args, v)
    }

    _sql := fmt.Sprintf(
        "(%s) values (%s)",
        fmt.Sprintf("`%s`", strings.Join(fields, "`,`")),
        fmt.Sprintf("%s", strings.Join(value, ",")),
    )

    return _sql, args
}

//record {"field":"val", "field2":"val2"}
//return arg1: set `field` = ?, `field2` = ?
//return arg2: ["val", "val2"]
func UpdateRecordToQuery(record Record) (string, []interface{}) {
    var scopes []string
    var args []interface{}
    for k, v := range record {
        scopes = append(scopes, fmt.Sprintf("`%s` = ?", k))
        args = append(args, v)
    }
    return strings.Join(scopes, ", "), args
}

//field {"field":"val", "field2":{in:[1,3]}}
//return arg1: where `field` = ? and `field2` in (?, ?)
//return arg2: ["val", "val2"]
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

// attr: {OrderBy:"id desc",Limit: 1}
// return arg1: order by ? limit ?
// return arg2: ["id desc", 1]
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

// attr: {Select:[id,name]}
// return arg: `id`, `name`
func AttrToSelectQuery(attr Attr) string {
    if len(attr.Select) == 0 {
        return "*"
    }
    return fmt.Sprintf("`%s`", strings.Join(attr.Select, "`,`"))
}
