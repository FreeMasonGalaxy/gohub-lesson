// Package config
// descr 负责配置信息
// author fm
// date 2022/11/14 16:49
package config

import (
	"os"

	"github.com/spf13/cast"
	viper13 "github.com/spf13/viper"
	"gohub-lesson/pkg/helpers"
)

var (
	// viper 实例
	viper *viper13.Viper

	// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
	ConfigFuncs map[string]ConfigFunc
)

// ConfigFunc 动态加载配置
type ConfigFunc func() map[string]any

func init() {

	// 1. 初始化 viper 库
	viper = viper13.New()
	// 2. 配置类型
	// 支持类型： "json", "toml", "yaml", "yml", "properties",
	// "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")
	// 3. 环境变量查找路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用于区分 go 的系统环境变量
	viper.SetEnvPrefix("appEnv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
	loadConfig()
}

func internalGet(path string, defaultValue ...any) any {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

func loadEnv(envSuffix string) {

	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"

	if len(envSuffix) > 0 {
		filePath := ".env." + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filePath
		}
	}

	// 加载 env
	viper.SetConfigName(envPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...any) string {
	return GetString(path, defaultValue...)
}

// GetDefaultAddr 获取默认 addr
func GetDefaultAddr() string {
	return ":" + Get("app.port")
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...any) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...any) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...any) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...any) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...any) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...any) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
