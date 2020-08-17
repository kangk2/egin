package pkg

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type Bootstrap struct {
    RouteMap RouteMap
}

func (boot *Bootstrap) Start() {
    gin.SetMode(Config.Mode)
    r := gin.Default()
    RegRoutes(r, boot.RouteMap)
    fmt.Println(Config)
    err := r.Run(Config.Address)
    if err != nil {
        return
    }
}
