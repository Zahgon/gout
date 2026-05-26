package core

import (
	"errors"
	"net/http"
	"reflect"
)

// FormFile 用于formdata类型数据编码
// 从文件读取数据流
type FormFile string

// FormMem 用于formdata类型数据编码
// 从[]byte里面读取数据流
type FormMem []byte

// FormType 自定义formdata文件名和流的类型
type FormType struct {
	FileName    string      //filename
	ContentType string      //Content-Type:Mime-Type
	File        interface{} //FromFile | FromMem (这里就是您的从文件地址中读取和从内存中读取)
}

// H 是map[string]interface{} 简写
type H map[string]interface{}

// A是[]interface{} 简写
type A []interface{}

// ErrUnknownType 未知错误类型
var ErrUnknownType = errors.New("unknown type")

// LoopElem 不停地对指针解引用
func LoopElem(v reflect.Value) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

// BytesToString 没有内存开销的转换
func BytesToString(b []byte) string { _ = "STUB: not implemented"; return "" }

// StringToBytes 没有内存开销的转换
func StringToBytes(s string) (b []byte) { _ = "STUB: not implemented"; return nil }

// NewPtrVal 新建这个类型的指针变量并赋值
func NewPtrVal(defValue interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func CloneRequest(r *http.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetBytes(v interface{}) (b []byte, ok bool) { _ = "STUB: not implemented"; return nil, false }

func GetString(v interface{}) (s string, ok bool) { _ = "STUB: not implemented"; return "", false }
