// Package middleware /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:45
 * @Description: 跨域配置中间件
**/
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ktaexam/pkg/config"
)

func Cors() gin.HandlerFunc {
	return cors.New(config.CorsConfig())
}
