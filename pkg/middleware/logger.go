package middleware

import (
    "github.com/daodao97/egin/pkg/utils"
    "github.com/gin-gonic/gin"
    "time"
)

// 日志记录到文件
func HttpLog() gin.HandlerFunc {

    return func(c *gin.Context) {
        startTime := time.Now()
        // 处理请求
        c.Next()
        // 结束时间
        endTime := time.Now()
        // 执行时间
        latencyTime := endTime.Sub(startTime)
        // 请求方式
        reqMethod := c.Request.Method
        // 请求路由
        reqUri := c.Request.RequestURI
        // 状态码
        statusCode := c.Writer.Status()
        // 请求IP
        clientIP := c.ClientIP()
        // 日志格式
        utils.Logger.Info("http end", map[string]interface{}{
            "status_code":  statusCode,
            "latency_time": latencyTime,
            "client_ip":    clientIP,
            "req_method":   reqMethod,
            "req_uri":      reqUri,
        })
    }
}
