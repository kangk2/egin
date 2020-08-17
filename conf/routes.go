package conf

import (
    "github.com/daodao97/egin/apis"
    "github.com/daodao97/egin/pkg"
)

func Routes() pkg.RouteMap {
    routes := pkg.RouteMap{
        "/user": pkg.Handle{Controller: &apis.User{}, Method: "Get"},
    }
    return routes
}
