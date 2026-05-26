package encode

import (
	"net/http"
	"reflect"
)

var _ Adder = (*HeaderEncode)(nil)

// HeaderEncode http header encoder structure
type HeaderEncode struct {
	r         *http.Request
	rawHeader bool
}

// NewHeaderEncode create a new http header encoder
func NewHeaderEncode(req *http.Request, rawHeader bool) *HeaderEncode {
	_ = "STUB: not implemented"
	return nil
}

// Add Encoder core function, used to set each key / value into the http header
func (h *HeaderEncode) Add(key string, v reflect.Value, sf reflect.StructField) error {
	_ = "STUB: not implemented"
	return nil
}

// Name header Encoder name
func (h *HeaderEncode) Name() string { _ = "STUB: not implemented"; return "" }
