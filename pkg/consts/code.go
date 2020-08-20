package consts

import (
    "github.com/daodao97/egin/pkg/utils"
)

type ErrCode int

const (
    ErrorSystem ErrCode = 500
    ErrorParam  ErrCode = 400
)

var msgMap = map[string]map[ErrCode]string{
    "zh-CN": {
        ErrorSystem: "服务器内部错误",
        ErrorParam:  "参数错误",
    },
    "en": {
        ErrorSystem: "system error",
        ErrorParam:  "param error",
    },
}

func (e ErrCode) String() string {
    lan := "zh-CN"
    if c := utils.Config.Lan; c != "" {
        lan = c
    }
    msg := msgMap[lan][e]
    if msg == "" {
        return "未知错误"
    }
    return msg
}
