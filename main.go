package main

import (
    "github.com/daodao97/egin/config"
    "github.com/daodao97/egin/pkg"
)

// config/*** 存放启动时配置, 如路由, 中间件
// app.json 存放运行时配置, 如 db/redis 配置, 由 pkg/config 持有
// pkg 不能依赖 config/**.go 否则很容易出现循环依赖
func main() {
    boot := pkg.Bootstrap{
        HttpMiddlewares: config.HttpMiddlewares,
        RouteMap:        config.Routes(),
    }
    boot.Start()
}
