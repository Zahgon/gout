package encode

import (
	"errors"
	"io"

	"github.com/guonaihong/gout/encoder"
)

var ErrNotXML = errors.New("Not xml data")

// XMLEncode xml encoder structure
type XMLEncode struct {
	obj interface{}
}

// NewXMLEncode create a new xml encoder
func NewXMLEncode(obj interface{}) encoder.Encoder {
	_ = "STUB: not implemented"
	return *new(encoder.Encoder)
}

// Encode xml encoder
func (x *XMLEncode) Encode(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

// Name xml Encoder name
func (x *XMLEncode) Name() string { _ = "STUB: not implemented"; return "" }

func XMLValid(b []byte) bool { _ = "STUB: not implemented"; return false }
