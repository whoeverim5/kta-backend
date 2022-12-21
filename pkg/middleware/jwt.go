// Package middleware /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:45
 * @Description: jwt用户鉴权中间件
**/
package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ktaexam/pkg/service"
	"ktaexam/pkg/utils"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, service.Resp{
				Code:    service.TokenExpired,
				Message: "Login failed",
				Data:    nil,
				Error:   false,
			})
			c.Abort()
			return
		}
		log.Print("\nget token: ", token)

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				c.JSON(http.StatusOK, service.Resp{
					Code:    service.TokenExpired,
					Message: "Token had expired",
					Data:    nil,
					Error:   false,
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, service.Resp{
				Code:    service.TokenValid,
				Message: "Token valid",
				Data:    nil,
				Error:   false,
			})
			c.Abort()
			return
		}
		// 将token中解析出来的信息传递下去
		c.Set("claims", claims)
	}
}
