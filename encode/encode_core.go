package encode

import (
	"errors"
	"reflect"
	"time"
)

// ErrUnsupported Unsupported type error returned
var ErrUnsupported = errors.New("Encode:Unsupported type")

var emptyField = reflect.StructField{}

// Adder interface
type Adder interface {
	Add(key string, v reflect.Value, sf reflect.StructField) error
	Name() string
}

// Encode core entry function
// in 的类型可以是
// struct
// map
// []string
func Encode(in interface{}, a Adder) error { _ = "STUB: not implemented"; return nil }

func parseTag(tag string) (string, tagOptions) {
	_ = "STUB: not implemented"
	return "", *new(tagOptions)
}

func timeToStr(v reflect.Value, sf reflect.StructField) string {
	_ = "STUB: not implemented"
	return ""
}

func valToStr(v reflect.Value, sf reflect.StructField) string { _ = "STUB: not implemented"; return "" }

// 看:
// https://github.com/guonaihong/gout/issues/322
/*
	if v.IsZero() {
		return ""
	}
*/

func setMoreTypes(val reflect.Value, sf reflect.StructField, a Adder, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTagAndSet(val reflect.Value, sf reflect.StructField, a Adder) error {
	_ = "STUB: not implemented"
	return nil
}

func encode(val reflect.Value, sf reflect.StructField, a Adder) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO使用接口解耦具体类型

type tagOptions []string

func (t tagOptions) Contains(tag string) bool { _ = "STUB: not implemented"; return false }

var timeType = reflect.TypeOf(time.Time{})

func valueIsEmpty(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
