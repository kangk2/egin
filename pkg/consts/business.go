package consts

const (
    ErrorSystem = 500
    ErrorParam  = 400
)

var MessageMap = map[int]string{
    ErrorSystem: "服务器内部错误",
    ErrorParam:  "参数错误",
}
