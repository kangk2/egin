package pkg

import (
    "./lib"
    "fmt"
    "github.com/gin-gonic/gin"
    "strings"
)

func handle(api interface{}, methodName string) func(c *gin.Context) {
    return func(c *gin.Context) {
        val, err := lib.Invoke(api, methodName, c)
        if err != nil {
            fmt.Println(err)
            return
        }

        c.JSON(200, gin.H{
            "payload": val.Interface(),
        })
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
