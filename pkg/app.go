// Package pkg /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 应用启动项
*/
package pkg

import (
	"github.com/gin-gonic/gin"
	"ktaexam/pkg/config"
	"ktaexam/pkg/middleware"
	"ktaexam/pkg/route"
)

var (
	port    = config.Yaml.GetString("server.port")
	proxies = config.Yaml.GetStringSlice("server.trusted-proxies")
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	app := gin.Default()
	// 设置受信任的代理IP
	HandleError(app.SetTrustedProxies(proxies))
	// 全局中间件需要在路由设置前
	app.Use(middleware.Cors())
	// 路由
	route.Router(app)
	HandleError(app.Run(port))
}
