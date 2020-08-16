package pkg

import (
    "github.com/gin-gonic/gin"
)

type Bootstrap struct {
    RouteMap RouteMap
}

func (boot *Bootstrap) Start() {
    r := gin.Default()
    RegRoutes(r, boot.RouteMap)
    err := r.Run()
    if err != nil {
        return
    }
}
