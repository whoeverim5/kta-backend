// Package route /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 接口路径
*/
package route

import (
	"github.com/gin-gonic/gin"
	"ktaexam/pkg/controller"
)

func Router(r *gin.Engine) {
	/*--- Test ---*/

	/*--- 用户认证 ---*/
	restful := r.Group("/api")
	login := restful.Group("/login")
	register := restful.Group("/register")

	/*--- 登录 ---*/
	login.POST("", controller.LoginHandler) // 登录

	/*--- 注册 ---*/
	register.POST("/:code", controller.RegisterHandler) // 注册
	register.GET("/code", controller.SendCodeHandler)   // 发送邮箱验证码
}
