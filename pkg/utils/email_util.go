// Package utils /*
/*
  @Author:      吴贤权
  @Date:        2022/10/4 上午11:41
  @Description: 发送邮箱验证码工具
*/
package utils

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"ktaexam/pkg/config"
)

var (
	port     = config.Yaml.GetInt("email.port")
	from     = config.Yaml.GetString("email.from")
	host     = config.Yaml.GetString("email.host")
	secret   = config.Yaml.GetString("email.secret")
	nickname = config.Yaml.GetString("email.nickname")
	subject  = config.Yaml.GetString("email.subject")
)

func SendEmail(to, code string) error {
	html := fmt.Sprintf("<h1>您的验证码为：\n\n<h2 style=\"font-weight: 700\">%v</h2>\n\n<p>即使提交此请求的不是您，没有此验证码也无法访问您的帐户。<p></h1>", code)
	auth := smtp.PlainAuth("", from, secret, host)

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", nickname, from)
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(html)

	hostAddr := fmt.Sprintf("%s:%d", host, port)
	return e.Send(hostAddr, auth)
}
