package config

import (
    "github.com/daodao97/egin/config/routes"
    "github.com/daodao97/egin/pkg/route"
)

// 由此次向外导出所有路由及路由组
// 新增的路由必须在此注册
var Routes = []route.RoutesMap{
    routes.UserRoute,
}

var RoutesGroup = []route.RoutesGroup{
    routes.UserRouteGroup,
}
