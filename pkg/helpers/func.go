// Package helpers
// descr 助手函数
// author fm
// date 2022/11/14 17:04
package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"reflect"
	"time"
)

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {

	mathrand.Seed(time.Now().UnixNano())

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)

	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}

	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// IsError 是否 err
func IsError(err error) bool {
	return err != nil
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {

	var (
		table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		b     = make([]byte, length)
	)

	n, err := io.ReadAtLeast(rand.Reader, b, length)

	if n != length {
		panic(err)
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}


// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// Empty 是否空
// 类似 PHP 的 empty 函数
func Empty(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Slice, reflect.Map:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
}
