package pkg

import (
    "github.com/daodao97/egin/pkg/db"
    "github.com/daodao97/egin/pkg/route"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "io"
    "os"
)

type Bootstrap struct {
    HttpMiddlewares []func() gin.HandlerFunc
    RouteMap        route.RouteMap
}

func (boot *Bootstrap) Start() {
    db.InitDb()
    gin.SetMode(utils.Config.Mode)
    gin.DefaultWriter = ginLogger()
    r := gin.Default()
    for _, midFunc := range boot.HttpMiddlewares {
        r.Use(midFunc())
    }
    route.RegRoutes(r, boot.RouteMap)
    err := r.Run(utils.Config.Address)
    if err != nil {
        return
    }
}

func ginLogger() io.Writer {
    f, _ := os.Create("gin.log")
    return io.MultiWriter(f, os.Stdout)
}
