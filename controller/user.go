package controller

import (
    "encoding/json"
    "fmt"
    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/cache"
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

type ParamsValidate struct {
    A string `form:"a" binding:"required"`
    B int    `form:"b" binding:"required"`
}

func (u User) Get(c *gin.Context, params interface{}) (interface{}, consts.ErrCode, error) {

    var p *ParamsValidate
    p = params.(*ParamsValidate)

    user := model.User
    result, err := user.Get(db.Filter{
        "id": map[string]int{
            ">": 20,
        },
    }, db.Attr{
        Select:  []string{"realname", "id", "username", "password"},
        OrderBy: "id desc",
    })

    redis := cache.Redis{Connection: "default"}
    setV, _ := json.Marshal([]int{1, 2, 4})
    err = redis.Set("egin:test", setV, 0)
    fmt.Println(222222222, err)
    _cache, err := redis.Get("egin:test")
    fmt.Println(3333333333, err)

    return []interface{}{result, p, _cache}, 0, err
}

func (u User) Post(c *gin.Context, params interface{}) (interface{}, consts.ErrCode, error) {
    user := model.User
    result, _, err := user.Insert(db.Record{
        "username": "test33333",
        "realname": "你好",
        "password": "cool",
    })
    var code consts.ErrCode
    if err != nil {
        code = consts.ErrorSystem
    }
    return []interface{}{result}, code, err
}

func (u User) Put(c *gin.Context, params interface{}) (interface{}, consts.ErrCode, error) {
    user := model.User
    _, affected, err := user.Update(
        db.Filter{
            "id": 13,
        },
        db.Record{
            "username": "test12",
        })
    var code consts.ErrCode
    if err != nil {
        code = consts.ErrorSystem
    }
    return affected, code, err
}

func (u User) Delete(c *gin.Context, params interface{}) (interface{}, consts.ErrCode, error) {
    user := model.User
    _, affected, err := user.Delete(db.Filter{
        "id": 22,
    })
    return affected, 0, err
}
