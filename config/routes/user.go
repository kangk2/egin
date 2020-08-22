package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/daodao97/egin/controller"
    "github.com/daodao97/egin/pkg/middleware"
    "github.com/daodao97/egin/pkg/route"
    "github.com/daodao97/egin/pkg/utils"
)

var UserRoute route.RoutesMap

var UserRouteGroup route.RoutesGroup

func init() {
    UserRoute = route.RoutesMap{
        "/user": route.SingleRoute{
            Handler: controller.User{}.Get,
            Middlewares: []gin.HandlerFunc{
                middleware.HttpLog(),
            },
            Param: new(controller.ParamsValidate),
            CustomValidateFuncs: []utils.CustomValidateFunc{
                controller.Bookabledate,
            },
        },
        "POST::/user": route.SingleRoute{
            Handler: controller.User{}.Post,
        },
        "PUT::/user": route.SingleRoute{
            Handler: controller.User{}.Put,
        },
        "DELETE::/user": route.SingleRoute{
            Handler: controller.User{}.Delete,
        },
    }

    UserRouteGroup = route.RoutesGroup{
        "/v1": {
            RoutesMap: route.RoutesMap{
                "/user": route.SingleRoute{
                    Handler: controller.User{}.Get,
                },
                "POST::/user": route.SingleRoute{
                    Handler: controller.User{}.Post,
                },
                "PUT::/user": route.SingleRoute{
                    Handler: controller.User{}.Put,
                },
                "DELETE::/user": route.SingleRoute{
                    Handler: controller.User{}.Delete,
                },
            },
        },
    }
}
