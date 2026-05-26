package encode

import (
	"errors"
	"io"

	"github.com/guonaihong/gout/encoder"
)

var ErrNotImplMessage = errors.New("The proto.Message interface is not implemented")

type ProtoBufEncode struct {
	obj interface{}
}

func NewProtoBufEncode(obj interface{}) encoder.Encoder {
	_ = "STUB: not implemented"
	return *new(encoder.Encoder)
}

func (p *ProtoBufEncode) Encode(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

//TODO找一个检测protobuf数据格式的函数

// 这里如果能把普通结构体转成指针类型结构体就

func (p *ProtoBufEncode) Name() string { _ = "STUB: not implemented"; return "" }
