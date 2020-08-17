package pkg

import (
    "fmt"
    "github.com/daodao97/egin/pkg/middleware"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "io"
    "os"
)

type Bootstrap struct {
    RouteMap RouteMap
}

func (boot *Bootstrap) Start() {
    gin.SetMode(utils.Config.Mode)
    gin.DefaultWriter = ginLogger()
    r := gin.Default()
    r.Use(middleware.LoggerToFile())
    RegRoutes(r, boot.RouteMap)
    fmt.Println(utils.Config)
    err := r.Run(utils.Config.Address)
    if err != nil {
        return
    }
}

func ginLogger() io.Writer {
    f, _ := os.Create("gin.log")
    return io.MultiWriter(f, os.Stdout)
}
