package main

import (
    "./conf"
    "./pkg"
)

func main() {
    boot := pkg.Bootstrap{
        RouteMap: conf.Routes(),
    }
    boot.Start()
}
