package controller

import (
    "encoding/json"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"

    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/cache"
    "github.com/daodao97/egin/pkg/consts"
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/utils"
)

type BaseApi struct {
    name string
}

type User struct {
    BaseApi
}

// FIXME time_format not working
type ParamsValidate struct {
    CheckIn  time.Time `form:"check_in" json:"check_in" binding:"required,bookabledate" time_format:"2006-01-02" label:"输入时间"`
    CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02" label:"输出时间"`
}

var Bookabledate = utils.CustomValidateFunc{
    Handle: func(fl validator.FieldLevel) bool {
        date, ok := fl.Field().Interface().(time.Time)
        if ok {
            today := time.Now()
            if today.After(date) {
                return false
            }
        }
        return true
    },
    TagName: "bookabledate",
    Message: "{0}不能早于当前时间或{1}格式错误!",
}

func (u User) Get(c *gin.Context) (interface{}, consts.ErrCode, error) {

    var p *ParamsValidate
    if params, exists := c.Get("params"); exists {
        p = params.(*ParamsValidate)
    }

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
    _cache, err := redis.Get("egin:test")

    return []interface{}{result, p, _cache}, 0, err
}

func (u User) Post(c *gin.Context) (interface{}, consts.ErrCode, error) {
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

func (u User) Put(c *gin.Context) (interface{}, consts.ErrCode, error) {
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

func (u User) Delete(c *gin.Context) (interface{}, consts.ErrCode, error) {
    user := model.User
    _, affected, err := user.Delete(db.Filter{
        "id": 22,
    })
    return affected, 0, err
}
