package decode

import (
	"io"
)

// YAMLDecode yaml decoder core data structure
type YAMLDecode struct {
	obj interface{}
}

// NewYAMLDecode create a new yaml decoder
func NewYAMLDecode(obj interface{}) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

// Decode yaml decoder
func (y *YAMLDecode) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Decode obj
func (y *YAMLDecode) Value() interface{} {
	_ = "STUB: not implemented"

	// YAML yaml decoder
	return nil
}

func YAML(r io.Reader, obj interface{}) error { _ = "STUB: not implemented"; return nil }
