// Package requests
// descr
// author fm
// date 2022/11/18 18:19
package requests

import (
	"github.com/thedevsaddam/govalidator"
)

// init 初始化
func init() {
	registerRules()
}

// registerRules 注册自定义规则
func registerRules() {
	govalidator.AddCustomRule("not_exists", RuleNotExists)
	govalidator.AddCustomRule("max_cn", RuleMaxCn)
	govalidator.AddCustomRule("min_cn", RuleMinCn)
	govalidator.AddCustomRule("exists", RuleExists)
}
