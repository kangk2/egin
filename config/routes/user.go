package routes

import (
    "github.com/daodao97/egin/controller"
    "github.com/daodao97/egin/pkg/middleware"
    "github.com/daodao97/egin/pkg/route"
    "github.com/gin-gonic/gin"
)

var UserRoute route.RoutesMap

var UserRouteGroup route.RoutesGroup

func init() {
    UserRoute = route.RoutesMap{
        "/user": route.SingleRoute{
            Handel: controller.User{}.Get,
            Middlewares: []gin.HandlerFunc{
                middleware.HttpLog(),
            },
        },
        "POST::/user": route.SingleRoute{
            Handel: controller.User{}.Post,
        },
        "PUT::/user": route.SingleRoute{
            Handel: controller.User{}.Put,
        },
        "DELETE::/user": route.SingleRoute{
            Handel: controller.User{}.Delete,
        },
    }

    UserRouteGroup = route.RoutesGroup{
        "/v1": {
            RoutesMap: route.RoutesMap{
                "/user": route.SingleRoute{
                    Handel: controller.User{}.Get,
                },
                "POST::/user": route.SingleRoute{
                    Handel: controller.User{}.Post,
                },
                "PUT::/user": route.SingleRoute{
                    Handel: controller.User{}.Put,
                },
                "DELETE::/user": route.SingleRoute{
                    Handel: controller.User{}.Delete,
                },
            },
        },
    }
}
