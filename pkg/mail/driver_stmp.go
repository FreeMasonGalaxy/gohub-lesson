// Package mail
// descr
// author fm
// date 2022/11/18 14:18
package mail

import (
	"fmt"
	"net/smtp"

	jordanEmail "github.com/jordan-wright/email"
	"gohub-lesson/pkg/logger"
)

type SMTP struct {
}

// Send 发送邮件
func (class *SMTP) Send(email Email, config map[string]string) bool {
	e := jordanEmail.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发送详情", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth("", config["username"], config["password"], config["host"]),
	)

	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发送成功", "")

	return true
}
