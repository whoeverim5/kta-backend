// Package service /*
/*
  @Author:      吴贤权
  @Date:        2022/10/3
  @Description: 业务层初始化
*/
package service

// 状态返回码
var (
	Error   = -1  // 通用错误
	Success = 200 // 通用成功

	PasswordWrong   = 10033 // 密码错误
	AccountNotFound = 10034 // 账号不存在
	CodeNotExist    = 10036 // 验证码错误

	TokenExpired = 10031 // Token过期
	TokenValid   = 10032 // Token无效

	UserExisted = 10035 // 用户已存在
)

// H 数据格式
type H = map[string]any

// Resp 统一响应返回的数据格式
type Resp struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 返回信息
	Data    H      `json:"data"`    // 返回数据
	Error   bool   `json:"error"`   // 只有当系统异常时返回True，否则返回False
}
