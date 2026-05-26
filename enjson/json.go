package enjson

import (
	"errors"
	"io"

	"github.com/guonaihong/gout/encoder"
)

var ErrNotJSON = errors.New("Not json data")

// JSONEncode json encoder structure
type JSONEncode struct {
	obj        interface{}
	escapeHTML bool
}

// NewJSONEncode create a new json encoder
func NewJSONEncode(obj interface{}, escapeHTML bool) encoder.Encoder {
	_ = "STUB: not implemented"
	return *new(encoder.Encoder)
}

func Marshal(obj interface{}, escapeHTML bool) (all []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encode结束之后会自作聪明的加'\n'
// 为了保持和json.Marshal一样的形为，手动删除最后一个'\n'

// Encode json encoder
func (j *JSONEncode) Encode(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// Name json Encoder name
func (j *JSONEncode) Name() string { _ = "STUB: not implemented"; return "" }
