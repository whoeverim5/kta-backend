// Package controller /*
/**
 * @Author:      吴贤权
 * @Date:        2022/10/3 下午4:44
 * @Description: 用户控制器
**/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ktaexam/pkg/model"
	"ktaexam/pkg/service"
)

func LoginHandler(c *gin.Context) {
	// 从请求体中获取账号和密码
	//account := c.PostForm("account")
	//password := c.PostForm("password")
	var user = new(model.User)
	c.ShouldBindJSON(user)
	// 登录业务
	resp := service.Login(user.Account, user.Password)
	c.JSON(http.StatusOK, resp)
}

func RegisterHandler(c *gin.Context) {
	var user = new(model.User)
	code := c.Param("code")
	
	c.ShouldBindJSON(user)

	resp := service.Register(user, code)
	c.JSON(http.StatusOK, resp)
}

func SendCodeHandler(c *gin.Context) {
	email := c.Query("email")
	resp := service.SendCode(email)
	c.JSON(http.StatusOK, resp)
}
