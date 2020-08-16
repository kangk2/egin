package apis

import (
    "github.com/gin-gonic/gin"
)

type BaseApi struct {
    name string
}

type User struct {
    BaseApi
}

func (this User) Get(c *gin.Context) interface{} {
    return map[string]string{
        "field": "ccc",
        "field2": "ccc",
        "field3": "ccc",
    }
}
