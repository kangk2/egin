package conf

import (
    "github.com/daodao97/egin/controller"
    "github.com/daodao97/egin/pkg"
)

func Routes() pkg.RouteMap {
    routes := pkg.RouteMap{
        "/user": pkg.Handle{Controller: &controller.User{}, Method: "Get"},
    }
    return routes
}
