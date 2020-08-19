package config

import (
    "github.com/daodao97/egin/controller"
    "github.com/daodao97/egin/pkg/utils"
)

func Routes() utils.RouteMap {
    routes := utils.RouteMap{
        "/user":       utils.Handle{Controller: &controller.User{}, Method: "Get"},
        "POST::/user": utils.Handle{Controller: &controller.User{}, Method: "Post"},
    }
    return routes
}
