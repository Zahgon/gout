package decode

import (
	"io"
	"reflect"
)

// BodyDecode body decoder structure
type BodyDecode struct {
	obj interface{}
}

// NewBodyDecode create a new body decoder
func NewBodyDecode(obj interface{}) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

var convertBodyFunc = map[reflect.Kind]convert{
	reflect.Uint:    {bitSize: 0, cb: setUintField},
	reflect.Uint8:   {bitSize: 8, cb: setUintField},
	reflect.Uint16:  {bitSize: 16, cb: setUintField},
	reflect.Uint32:  {bitSize: 32, cb: setUintField},
	reflect.Uint64:  {bitSize: 64, cb: setUintField},
	reflect.Int:     {bitSize: 0, cb: setIntField},
	reflect.Int8:    {bitSize: 8, cb: setIntField},
	reflect.Int16:   {bitSize: 16, cb: setIntField},
	reflect.Int32:   {bitSize: 32, cb: setIntField},
	reflect.Int64:   {bitSize: 64, cb: setIntDurationField},
	reflect.Float32: {bitSize: 32, cb: setFloatField},
	reflect.Float64: {bitSize: 64, cb: setFloatField},
}

// Decode body decoder
func (b *BodyDecode) Decode(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Decode obj
func (b *BodyDecode) Value() interface{} {
	_ = "STUB: not implemented"

	// Body body decoder
	return nil
}

func Body(r io.Reader, obj interface{}) error { _ = "STUB: not implemented"; return nil }
