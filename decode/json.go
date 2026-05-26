package decode

import (
	"io"
)

// JSONDecode json decoder core data structure
type JSONDecode struct {
	obj interface{}
}

// NewJSONDecode create a new json decoder
func NewJSONDecode(obj interface{}) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

// Decode json decoder
func (j *JSONDecode) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Decode obj
func (j *JSONDecode) Value() interface{} {
	_ = "STUB: not implemented"

	// JSON json decoder
	return nil
}

func JSON(r io.Reader, obj interface{}) error { _ = "STUB: not implemented"; return nil }
