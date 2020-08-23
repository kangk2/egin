package middleware

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/daodao97/egin/pkg/utils"
)

func jwtAbort(c *gin.Context, msg string) {
    c.JSON(http.StatusUnauthorized, gin.H{
        "status":  "error",
        "message": msg,
    })
    c.Abort()
}

func JWTMiddleware(db interface{}) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
        fmt.Println(authHeader)
        if authHeader == "" {
            jwtAbort(c, "Authorization Failed.")
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            jwtAbort(c, "Authorization Failed.")
            return
        }

        claims, err := utils.ParseToken(parts[1])
        if err != nil {
            jwtAbort(c, "无效的Token")
            return
        }

        if time.Now().Unix() > claims.ExpiresAt {
            jwtAbort(c, "Token已过期")
            return
        }

        // 根据 claims.UserID 从库中查询
        user := struct {
            Id int
        }{}

        if user.Id != claims.UserID {
            jwtAbort(c, "无效的Token")
            return
        }

        c.Set("user", user)
        c.Next()
    }
}
