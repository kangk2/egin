package conf

import (
    "../apis"
    "../pkg"
)

func Routes() pkg.RouteMap {
    routes := pkg.RouteMap{
        "/user": pkg.Handle{Controller: &apis.User{}, Method: "Get"},
    }
    return routes
}
