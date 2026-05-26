package decode

import (
	"errors"
	"net/http"
	"reflect"
)

var ErrWrongParam = errors.New("Wrong parameter")

type headerDecode struct{}

func (h *headerDecode) Decode(rsp *http.Response, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// 如果是
// http.Header
// *http.Header
// 等类型, 直接把rsp.Header的字段都拷贝下

type headerSet map[string][]string

var _ setter = headerSet(nil)

func (h headerSet) Set(value reflect.Value, sf reflect.StructField, tagValue string) error {
	_ = "STUB: not implemented"
	return nil
}
