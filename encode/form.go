package encode

import (
	"bytes"
	"io"
	"mime/multipart"
	"reflect"
	"strings"
)

type formContent struct {
	fileName     string //filename
	contentType  string //Content-Type:Mime-Type
	data         []byte
	isFormFile   bool
	needOpenFile bool
}

var _ Adder = (*FormEncode)(nil)

// FormEncode form-data encoder structure
type FormEncode struct {
	*multipart.Writer
}

// NewFormEncode create a new form-data encoder
func NewFormEncode(b *bytes.Buffer) *FormEncode { _ = "STUB: not implemented"; return nil }

func genFormContext(key string, val reflect.Value, sf reflect.StructField, fc *formContent) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//说明是结构体

// `form-file:"mem"`  从内存中读取

// `form-file:"file"` 从文件中读取 `form-file:"true"` 也是从文件中读取 兼容老接口

func genFormContextCore(key string, val reflect.Value, sf reflect.StructField, fc *formContent) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//这里是type的类型转换，没有任何内存拷贝

//已经得到data的值，直接返回

// 下方为原函数附带的方法
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }

// CreateFormFile 重写原来net/http里面的CreateFormFile函数
func (f *FormEncode) CreateFormFile(fieldName, fileName, contentType string) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

// Add Encoder core function, used to set each key / value into the http form-data
func (f *FormEncode) Add(key string, v reflect.Value, sf reflect.StructField) (err error) {
	_ = "STUB: not implemented"
	// 1.提取数据
	return nil
}

// 2.生成formdata格式数据

func (f *FormEncode) createForm(key string, fc *formContent) error {
	_ = "STUB: not implemented"
	return nil
}

// End refresh data
func (f *FormEncode) End() error {
	_ = "STUB: not implemented"

	// Name form-data Encoder name
	return nil
}

func (f *FormEncode) Name() string { _ = "STUB: not implemented"; return "" }
