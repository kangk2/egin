package config

import (
    "github.com/daodao97/egin/pkg/middleware"
    "github.com/gin-gonic/gin"
)

type MiddlewaresSlice []func() gin.HandlerFunc

// 由上而下顺序执行
var HttpMiddlewares = MiddlewaresSlice{
    middleware.IPAuth,
    middleware.HttpLog,
}