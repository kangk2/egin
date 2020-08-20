package config

import (
    "github.com/daodao97/egin/controller"
    "github.com/daodao97/egin/pkg/route"
)

func Routes() route.RouteMap {
    routes := route.RouteMap{
        "/user":       route.Handle{Controller: &controller.User{}, Method: "Get"},
        "POST::/user": route.Handle{Controller: &controller.User{}, Method: "Post"},
    }
    return routes
}
