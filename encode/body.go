package encode

import (
	"io"

	"github.com/guonaihong/gout/encoder"
)

// BodyEncode body encoder structure
type BodyEncode struct {
	obj interface{}
}

// NewBodyEncode create a new body encoder
func NewBodyEncode(obj interface{}) encoder.Encoder {
	_ = "STUB: not implemented"
	return *new(encoder.Encoder)
}

// Encode Add Encoder core function, used to set io.Writer into the http body
func (b *BodyEncode) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Name http body Encoder name
func (b *BodyEncode) Name() string { _ = "STUB: not implemented"; return "" }
