package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/daodao97/egin/pkg/cache"
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

func ApiLimiter(limiter *utils.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter.Incr()
		if limiter.CheckOverLimit() {
			c.String(http.StatusBadGateway, "reject")
			c.Abort()
		}
		c.Next()
	}
}

func IpLimiter() gin.HandlerFunc {
	red := cache.Redis{}
	return func(c *gin.Context) {
		conf := utils.Config.Auth.IpLimiter
		if !conf.Enable {
			return
		}
		ip := c.ClientIP()
		mu := sync.Mutex{}
		mu.Lock()
		limitCount, ok := conf.IPLimit[ip]
		mu.Unlock()
		if !ok {
			return
		}
		key := fmt.Sprintf("%s:%s", "egin_ip_limiter", c.ClientIP())
		currentCount, _ := red.Get(key)
		_currentCount, _ := strconv.Atoi(currentCount)
		// FIXME 由于 incr 在后, 所以会比实际limit多一次
		// incr放在前又会每次请求都透传到redis, 所有选择后置
		if _currentCount > limitCount {
			c.String(http.StatusBadGateway, "reject")
			c.Abort()
			return
		}
		_ = red.Incr(key)
		_ = red.PExpire(key, time.Second)
		c.Next()
	}
}
