package decode

import (
	"io"
)

// XMLDecode xml decoder core data structure
type XMLDecode struct {
	obj interface{}
}

// NewXMLDecode create a new xml decoder
func NewXMLDecode(obj interface{}) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

// Decode xml decoder
func (x *XMLDecode) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Decode object
func (x *XMLDecode) Value() interface{} {
	_ = "STUB: not implemented"

	// XML xml decoder
	return nil
}

func XML(r io.Reader, obj interface{}) error { _ = "STUB: not implemented"; return nil }
