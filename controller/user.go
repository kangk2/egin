package controller

import (
    "github.com/daodao97/egin/model"
    "github.com/daodao97/egin/pkg/db"
    "github.com/gin-gonic/gin"
)

type BaseApi struct {
    name string
}

type User struct {
    BaseApi
}

func (u User) Get(c *gin.Context) interface{} {
    result := make(map[string]interface{})
    user := model.UserModel
    result["user"] = user.Get(db.Filter{}, db.Attr{})
    // config := model.ConfigModel
    // result["config"] = config.Get()
    return result
}
