package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "github.com/daodao97/egin/pkg/lib"
    "github.com/daodao97/egin/pkg/utils"
)

// cors
func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        conf := utils.Config.Auth.Cors
        if !conf.Enable {
            return
        }
        allowOrigin := conf.AllowOrigins
        if len(allowOrigin) == 0 {
            allowOrigin = []string{"*"}
        }
        allCredentials := "true"
        if !conf.AllowCredentials {
            allCredentials = "false"
        }
        allowHeaders := []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
        allowMethods := []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}
        c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(allowOrigin, ","))
        c.Writer.Header().Set("Access-Control-Allow-Credentials", allCredentials)
        c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ","))
        c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(allowMethods, ","))
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

// ip白名单
func IPAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        conf := utils.Config.Auth.IpAuth
        if !conf.Enable {
            return
        }
        clientIp := c.ClientIP()
        _, hasIt := lib.Find(conf.AllowedIpList, clientIp)
        if !hasIt {
            c.String(http.StatusUnauthorized, "%s, not in ipList", clientIp)
            c.Abort()
        }
        c.Next()
    }
}
