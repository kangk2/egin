package controller

import (
    "github.com/daodao97/egin/model"
    "github.com/davecgh/go-spew/spew"
    "github.com/gin-gonic/gin"
)

type BaseApi struct {
    name string
}

type User struct {
    BaseApi
}

func (u User) Get(c *gin.Context) interface{} {
    spew.Dump(c.GetQuery("id"))
    result := make(map[string]interface{})
    user := model.UserModel
    result["user"] = user.Get()
    config := model.ConfigModel
    result["config"] = config.Get()
    return result
}
