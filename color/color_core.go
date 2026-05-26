package color

import (
	"bytes"
	"io"
)

// BodyType 区分body的类型
type BodyType int

const (
	// JSONType http body是json类型的
	JSONType BodyType = iota + 1
	// XMLType http body是xml类型的
	XMLType
	// YAMLType http body是yaml类型的
	YAMLType
	// TxtType http body是txt类型的
	TxtType
)

// 本文件来自github.com/TylerBrock/colorjson, 感谢TylerBrock
// TODO
// * 修复原来代码bug
// * 支持更多数据类型
// * 扩展原有功能以支持xml和yaml

const initialDepth = 0
const valueSep = ","
const null = "null"
const startMap = "{"
const endMap = "}"
const startArray = "["
const endArray = "]"

const emptyMap = startMap + endMap
const emptyArray = startArray + endArray

// Formatter 是颜色高亮核心结构体
type Formatter struct {
	escapeHTML      bool
	KeyColor        *Color // 设置key的颜色
	StringColor     *Color // 设置string的颜色
	BoolColor       *Color // 设置bool的颜色
	NumberColor     *Color // 设置数字的颜色
	NullColor       *Color // 设置null的颜色
	StringMaxLength int
	Indent          int
	DisabledColor   bool
	RawStrings      bool

	r io.Reader
}

func strToObject(all []byte) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// NewFormatEncoder 着色json/yaml/xml构造函数
func NewFormatEncoder(r io.Reader, openColor bool, bodyType BodyType, escapeHTML bool) *Formatter {
	_ = "STUB: not implemented"
	// 如果颜色没打开，或者bodyType为txt
	return nil
}

//todo xmlType and yamlType

func (f *Formatter) sprintColor(c *Color, s string) string { _ = "STUB: not implemented"; return "" }

func (f *Formatter) writeIndent(buf *bytes.Buffer, depth int) { _ = "STUB: not implemented"; return }

func (f *Formatter) writeObjSep(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

// Marshal 给原始的结构化数据着色
func (f *Formatter) Marshal(jsonObj interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Formatter) marshalMap(m map[string]interface{}, buf *bytes.Buffer, depth int) {
	_ = "STUB: not implemented"
	return
}

func (f *Formatter) marshalArray(a []interface{}, buf *bytes.Buffer, depth int) {
	_ = "STUB: not implemented"
	return
}

func (f *Formatter) marshalValue(val interface{}, buf *bytes.Buffer, depth int) {
	_ = "STUB: not implemented"
	return
}

func (f *Formatter) marshalString(str string, buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (f *Formatter) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
