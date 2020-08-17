package apis

import (
    "github.com/daodao97/egin/module"
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
    user := module.UserModel
    result["user"] = user.Get()
    config := module.ConfigModel
    result["config"] = config.Get()
    return result
}
