package decode

import (
	//"net/http"

	"reflect"
)

type setter interface {
	Set(value reflect.Value,

		sf reflect.StructField,

		tagValue string) error
}

var emptyField = reflect.StructField{}

func setForm(m map[string][]string,
	value reflect.Value,
	sf reflect.StructField,
	tagValue string,
) error {
	_ = "STUB: not implemented"
	return nil
}

//fmt.Printf("tagName = %s:%v\n", tagValue, m)

func decode(d setter, obj interface{}, tagName string) error { _ = "STUB: not implemented"; return nil }

// todo delete
func parseTag(tag string) (string, []string) { _ = "STUB: not implemented"; return "", nil }

func parseTagAndSet(val reflect.Value, sf reflect.StructField, setter setter, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeCore(val reflect.Value, sf reflect.StructField, setter setter, tagName string) (err error) {
	_ = "STUB: not implemented"
	return nil

	// elem pointer
}

// 每个类型都会先尝试set
// 如果不是结构体才设置。那time.Time类型该如何呢?
// (time.Time是标准库里面用于表示时间的类型, 用结构体实现)？

//todo 是否已经设置过

type convert struct {
	bitSize int
	cb      func(val string, bitSize int, sf reflect.StructField, field reflect.Value) error
}

var convertFunc = map[reflect.Kind]convert{
	reflect.Uint:    {bitSize: 0, cb: setUintField},
	reflect.Uint8:   {bitSize: 8, cb: setUintField},
	reflect.Uint16:  {bitSize: 16, cb: setUintField},
	reflect.Uint32:  {bitSize: 32, cb: setUintField},
	reflect.Uint64:  {bitSize: 64, cb: setUintField},
	reflect.Int:     {bitSize: 0, cb: setIntField},
	reflect.Int8:    {bitSize: 8, cb: setIntField},
	reflect.Int16:   {bitSize: 16, cb: setIntField},
	reflect.Int32:   {bitSize: 32, cb: setIntField},
	reflect.Int64:   {bitSize: 64, cb: setIntDurationField},
	reflect.Bool:    {bitSize: 0, cb: setBoolField},
	reflect.Float32: {bitSize: 32, cb: setFloatField},
	reflect.Float64: {bitSize: 64, cb: setFloatField},
	reflect.Struct:  {bitSize: 0, cb: setStructField},
	reflect.Map:     {bitSize: 0, cb: setMapField},
}

func setIntDurationField(val string, bitSize int, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setIntField(val string, bitSize int, sf reflect.StructField, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setUintField(val string, bitSize int, sf reflect.StructField, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setBoolField(val string, bitSize int, sf reflect.StructField, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setFloatField(val string, bitSize int, sf reflect.StructField, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimeField(val string, bitSize int, structField reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setStructField(val string, bitSize int, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setArray(vals []string, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setSlice(vals []string, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setMapField(val string, bitSize int, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimeDuration(val string, bitSize int, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setBase(val string, sf reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}
