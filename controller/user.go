package controller

import (
    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/consts"
    "github.com/daodao97/egin/pkg/db"
    "github.com/gin-gonic/gin"
)

type BaseApi struct {
    name string
}

type User struct {
    BaseApi
}

func (u User) Get(c *gin.Context) (interface{}, error, int) {
    user := model.UserModel
    result, err := user.Get(db.Filter{
        "id": 1,
    }, db.Attr{
        Select: []string{"realname", "id"},
    })
    // config := model.ConfigModel
    // result["config"] = config.Get()
    return result, err, 0
}

func (u User) Post(c *gin.Context) (interface{}, error, int) {
    user := model.UserModel
    result, err := user.Insert(db.Record{
        "username": "test",
        "realname": "你好",
        "password": "cool",
    })
    var code int
    if err != nil {
        code = consts.ErrorSystem
    }
    return result, err, code
}
