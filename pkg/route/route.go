package route

import (
    "fmt"
    "github.com/daodao97/egin/pkg/consts"
    "github.com/daodao97/egin/pkg/lib"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "strings"
)

type SingleRoute struct {
    Handler     func(c *gin.Context, param interface{}) (interface{}, consts.ErrCode, error)
    Middlewares []gin.HandlerFunc
    Param       interface{}
}

type routePath string
type RoutesMap map[routePath]SingleRoute

type groupName string
type RoutesGroup map[groupName]struct {
    RoutesMap   RoutesMap
    Middlewares []gin.HandlerFunc
}

func wrap(singleRoute SingleRoute) func(c *gin.Context) {
    return func(c *gin.Context) {
        param := singleRoute.Param
        fmt.Println(1111111, param, param == nil, singleRoute)
        if param != nil {
            err := c.ShouldBind(param)
            if err != nil {
                c.JSON(200, gin.H{
                    "err": err.Error(),
                })
                return
            }
            c.Set("params", param)
        }

        result, code, err := singleRoute.Handler(c, param)

        response := gin.H{
            "code":    code,
            "payload": result,
        }

        code = consts.ErrCode(code)
        if err != nil {
            if utils.Config.Mode != "release" {
                response["message"] = err.Error()
            } else {
                response["message"] = code.String()
            }
        }

        c.JSON(200, response)
    }
}

var httpMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}

func RegRoutes(r *gin.Engine, routesMap []RoutesMap) {
    for _, item := range routesMap {
        for path, singleRoute := range item {
            tokens := strings.Split(string(path), "::")
            method := "GET"
            _path := string(path)
            if len(tokens) > 1 {
                method = tokens[0]
                _path = tokens[1]
            }
            method = strings.ToUpper(method)
            if _, hasIt := lib.Find(httpMethods, method); !hasIt {
                continue
            }
            var handlers []gin.HandlerFunc
            handlers = append(handlers, singleRoute.Middlewares...)
            handlers = append(handlers, wrap(singleRoute))
            r.Handle(method, _path, handlers...)
        }
    }
}

func RegRouteGroup(r *gin.Engine, routesGroup []RoutesGroup) {
    for _, item := range routesGroup {
        for groupName, groupInfo := range item {
            g := r.Group(string(groupName), groupInfo.Middlewares...)
            for path, singleRoute := range groupInfo.RoutesMap {
                tokens := strings.Split(string(path), "::")
                method := "GET"
                _path := string(path)
                if len(tokens) > 1 {
                    method = tokens[0]
                    _path = tokens[1]
                }
                method = strings.ToUpper(method)
                if _, hasIt := lib.Find(httpMethods, method); !hasIt {
                    continue
                }
                var handlers []gin.HandlerFunc
                handlers = append(handlers, singleRoute.Middlewares...)
                handlers = append(handlers, wrap(singleRoute))
                g.Handle(method, _path, handlers...)
            }
        }
    }
}
