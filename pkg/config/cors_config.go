// Package config /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午3:10
 * @Description: cors配置
**/
package config

import (
	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	var (
		allOrigins    = Yaml.GetBool("cors.all-origins")
		origins       = Yaml.GetStringSlice("cors.origins")
		methods       = Yaml.GetStringSlice("cors.methods")
		headers       = Yaml.GetStringSlice("cors.headers")
		exposeHeaders = Yaml.GetStringSlice("cors.expose-headers")
		credential    = Yaml.GetBool("cors.credential")
		maxAge        = Yaml.GetDuration("cors.max-age")
	)

	return cors.Config{
		AllowAllOrigins:  allOrigins,
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowHeaders:     headers,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: credential,
		MaxAge:           maxAge,
	}
}
