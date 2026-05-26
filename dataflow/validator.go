package dataflow

import (
	"reflect"
	"sync"

	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

var valid *defaultValidator = &defaultValidator{}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *defaultValidator) Engine() interface{} { _ = "STUB: not implemented"; return nil }

func (v *defaultValidator) lazyinit() { _ = "STUB: not implemented"; return }

// see universal-translator for details

// add any custom validations etc. here

func kindOfData(data interface{}) reflect.Kind {
	_ = "STUB: not implemented"
	return *new(reflect.Kind)
}
