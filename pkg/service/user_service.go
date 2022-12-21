// Package service /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 用户认证业务处理
*/
package service

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"ktaexam/pkg/dao"
	"ktaexam/pkg/model"
	"ktaexam/pkg/utils"
)

func Login(account, password string) Resp {
	user, err := dao.GetUser(account)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Message: "Get user failed",
			Error:   true,
		}
	}
	// 账号不存在
	if user.Password == "" {
		return Resp{
			Code:    AccountNotFound,
			Message: "Login failed",
		}
	}
	md5Password := utils.Md5Crypto(password, user.Salt) // 密码加盐加密
	// 密码错误
	if md5Password != user.Password {
		return Resp{
			Code:    PasswordWrong,
			Message: "Login failed",
		}
	}
	// 密码正确
	claims := utils.CustomClaims{
		Account: account,
		UserId:  user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(utils.TokenExpireTime)).Unix(),
			Issuer:    "wu-" + time.Now().Format("06-01-02")},
	}
	token, err := utils.NewJWT().CreateToken(claims)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Message: "Create token error",
			Error:   true,
		}
	}
	return Resp{
		Code:    Success,
		Message: "Login success",
		Data: H{
			"token":         token,
			"user_identity": user.UserIdentity,
			"email":         user.Email,
		},
	}
}

func SendCode(email string) Resp {
	// 生成随机验证码
	code, err := utils.GetCode(6)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Message: "Generate code failed",
			Error:   true,
		}
	}
	// 设置验证码缓存
	err = dao.SetEmailCodeCache(code, email)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Error:   true,
			Message: "Redis operation error",
		}
	}
	// 向邮箱发送验证码
	err = utils.SendEmail(email, code)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Error:   true,
			Message: "Send email error",
		}
	}
	return Resp{
		Code:    Success,
		Message: "Send code success",
	}
}

// Register 用户注册
func Register(user *model.User, code string) Resp {
	// 从数据库中查找是否该用户已存在
	u, err := dao.JudgeUserExisted(user.Account)
	if err != nil {
		log.Println(err)
		return Resp{
			Code:    Error,
			Message: "Get user error",
			Error:   true,
		}
	}
	if u.Account == user.Account {
		return Resp{
			Code:    UserExisted,
			Message: "Register failed",
		}
	}
	// 在缓存中查找该用户的特定验证码是否存在
	isExist := dao.QueryCodeFromRedis(code, user.Email)
	if !isExist {
		return Resp{
			Code:    CodeNotExist,
			Message: "Code wrong",
		}
	}
	// 向数据库中新增一个用户
	salt := utils.GetRandomSalt()
	user.Salt = salt
	user.UserId = utils.NewUserId()
	user.Password = utils.Md5Crypto(user.Password, salt)
	err = dao.CreateNewUser(user)
	if err != nil {
		return Resp{
			Code:    Error,
			Message: "Create new user error",
			Error:   true,
		}
	}
	return Resp{
		Code:    Success,
		Message: "Register success",
	}
}
