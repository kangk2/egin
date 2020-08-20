package middleware

import (
    "github.com/daodao97/egin/pkg/lib"
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

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
