package encode

import (
	"io"
	"net/url"
	"reflect"

	"github.com/guonaihong/gout/setting"
)

var _ Adder = (*WWWFormEncode)(nil)

// WWWFormEncode x-www-form-urlencoded encoder structure
type WWWFormEncode struct {
	values url.Values
	setting.Setting
}

// NewWWWFormEncode create a new x-www-form-urlencoded encoder
func NewWWWFormEncode(s setting.Setting) *WWWFormEncode { _ = "STUB: not implemented"; return nil }

// Encode x-www-form-urlencoded encoder
func (we *WWWFormEncode) Encode(obj interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// Add Encoder core function, used to set each key / value into the http x-www-form-urlencoded
// 这里value的设置暴露 reflect.Value和 reflect.StructField原因如下
// reflect.Value把value转成字符串
// reflect.StructField主要是可以在Add函数里面获取tag相关信息
func (we *WWWFormEncode) Add(key string, v reflect.Value, sf reflect.StructField) error {
	_ = "STUB: not implemented"
	return nil
}

func (we *WWWFormEncode) End(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Name x-www-form-urlencoded Encoder name
func (we *WWWFormEncode) Name() string { _ = "STUB: not implemented"; return "" }
