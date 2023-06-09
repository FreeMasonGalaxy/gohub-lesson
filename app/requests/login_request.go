// Package requests
// descr
// author fm
// date 2022/11/21 11:22
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

// ValidateLoginByPhone 验证表单，返回长度等于零即通过
func ValidateLoginByPhone(ctx *gin.Context) (data LoginByPhoneRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs = validate(&data, rules, messages)

	// 手机验证码
	ValidateVerifyCode(data.Phone, data.VerifyCode, errs)

	return
}

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	LoginID  string `valid:"login_id" json:"login_id"`
	Password string `valid:"password" json:"password,omitempty"`
}

// ValidateLoginByPassword 验证表单，返回长度等于零即通过
func ValidateLoginByPassword(ctx *gin.Context) (data LoginByPasswordRequest, errs MapErrs) {

	ShouldBindJSON(&data, ctx)

	rules := govalidator.MapData{
		"login_id":       []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"login_id": []string{
			"required:登录 ID 为必填项，支持手机号、邮箱和用户名",
			"min:登录 ID 长度需大于 3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs = validate(&data, rules, messages)

	// 图片验证码
	errs = ValidateCaptcha(data.CaptchaID, data.CaptchaAnswer, errs)

	return
}
