package db

import (
    "fmt"
    "github.com/daodao97/egin/pkg/lib"
    "strconv"
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

// INSERT INTO products (name, code) VALUES ("name", "code")
func InsertSql(record Record) {

}

func InsertMultiSql(records Records) {

}

func updateSql(filter Filter, up Record) {

}

func deleteSql(filter Filter) {

}

func SelectSql(filter Filter, attr Attr) string {
    _fields := "*"
    if len(attr.Select) != 0 {
        _fields = fmt.Sprintf("`%s`", strings.Join(attr.Select, "`,`"))
    }
    _where := ""
    if len(filter) != 0 {
        _where = fmt.Sprintf("where %s", FilterToQuery(filter))
    }
    _attr := AttrToQuery(attr)

    sql := fmt.Sprintf("select %s from %s %s %s;", _fields, "%s", _where, _attr)

    return sql
}

func FilterToQuery(filter Filter) string {
    var sql string
    logic := filter["__logic"]
    if logic == nil {
        logic = "and"
    }
    var scopes []string

    for k, v := range filter {
        switch v.(type) {
        case []int:
            op := make([]string, 0)
            for _, val := range v.([]int) {
                op = append(op, strconv.Itoa(int(val)))
            }
            scopes = append(scopes, fmt.Sprintf("`%s` in (%s)", k, strings.Join(op, ",")))
        case []string:
            op := make([]string, 0)
            op = v.([]string)
            scopes = append(scopes, fmt.Sprintf("`%s` in ('%s')", k, EscapeSingleQuote(strings.Join(op, "','"))))
        case int32, int64, string, bool:
            scopes = append(scopes, fmt.Sprintf("`%s` = '%s'", k, EscapeSingleQuote(v.(string))))
        case map[string]string:
            var _scope []string
            var _logic string = "and"
            for op, val := range v.(map[string]string) {
                _, found := lib.Find([]string{">", ">=", "<", "<=", "=", "<>", "!="}, op)
                if found {
                    _scope = append(_scope, fmt.Sprintf("`%s` %s '%s'", k, op, EscapeSingleQuote(val)))
                }
                if op == "__logic" {
                    _logic = val
                }
            }
            fmt.Println(_scope)
            scopes = append(scopes, fmt.Sprintf("(%s)", strings.Join(_scope, fmt.Sprintf(" %s ", _logic))))
        }
    }
    sql = strings.Join(scopes, fmt.Sprintf(" %s ", logic.(string)))
    return sql
}

func AttrToQuery(attr Attr) string {
    var sql string
    var scopes []string
    if attr.Offset != 0 {
        scopes = append(scopes, fmt.Sprintf("offset %d", attr.Offset))
    }
    if attr.Limit != 0 {
        scopes = append(scopes, fmt.Sprintf("limit %d", attr.Limit))
    }
    if attr.OrderBy != "" {
        scopes = append(scopes, fmt.Sprintf("order by %s", attr.OrderBy))
    }
    if attr.GroupBy != "" {
        scopes = append(scopes, fmt.Sprintf("group by %s", attr.GroupBy))
    }
    sql = strings.Join(scopes, " ")
    return sql
}

func EscapeSingleQuote(str string) string {
    return strings.Replace(str, "'", "\\'", -1)
}
