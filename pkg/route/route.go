package route

import (
    "fmt"
    "github.com/daodao97/egin/pkg/consts"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "strings"
)

func handle(api interface{}, methodName string) func(c *gin.Context) {
    return func(c *gin.Context) {
        vals, err := lib.Invoke(api, methodName, c)
        if err != nil {
            fmt.Println(err)
            return
        }

        code := vals[2].Interface().(consts.ErrCode)
        response := gin.H{
            "code":    code,
            "payload": vals[0].Interface(),
        }

        message := vals[1].Interface()
        if message != nil {
            if utils.Config.Mode != "release" {
                response["message"] = vals[1].Interface()
            } else {
                response["message"] = code.String()
            }
        }

        c.JSON(200, response)
    }
}

type Handle struct {
    Controller interface{}
    Method     string
}

type RouteMap map[string]Handle

func RegRoutes(r *gin.Engine, routeMap RouteMap) {
    for path, control := range routeMap {
        tokens := strings.Split(path, "::")
        if len(tokens) == 1 {
            r.GET(path, handle(control.Controller, control.Method))
        } else {
            switch strings.ToUpper(tokens[0]) {
            case "POST":
                r.POST(tokens[1], handle(control.Controller, control.Method))
            }
        }
    }
}
