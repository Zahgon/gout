package encode

import (
	"errors"
	"io"

	"github.com/guonaihong/gout/encoder"
)

var ErrNotYAML = errors.New("Not yaml data")

// YAMLEncode yaml encoder structure
type YAMLEncode struct {
	obj interface{}
}

// NewYAMLEncode create a new yaml encoder
func NewYAMLEncode(obj interface{}) encoder.Encoder {
	_ = "STUB: not implemented"
	return *new(encoder.Encoder)
}

// Encode yaml encoder
func (y *YAMLEncode) Encode(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// Name yaml Encoder name
func (y *YAMLEncode) Name() string { _ = "STUB: not implemented"; return "" }

func YAMLValid(b []byte) bool { _ = "STUB: not implemented"; return false }
