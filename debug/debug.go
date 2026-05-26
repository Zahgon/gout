package debug

import (
	"io"
	"net/http"

	"github.com/guonaihong/gout/color"
)

// ToBodyType Returns the http body type,
// which mainly affects color highlighting
func ToBodyType(s string) color.BodyType { _ = "STUB: not implemented"; return *new(color.BodyType) }

// return color.XmlType //TODO open

// return color.YamlType //TODO open

// Options Debug mode core data structure
type Options struct {
	EscapeHTML      bool
	Write           io.Writer
	Debug           bool
	Color           bool
	Trace           bool
	FormatTraceJSON bool
	ReqBodyType     string
	RspBodyType     string
	TraceInfo
}

// Apply is an interface for operating Options
type Apply interface {
	Apply(*Options)
}

// Func Apply is a function that manipulates core data structures
type Func func(*Options)

// Apply is an interface for operating Options
func (f Func) Apply(o *Options) { _ = "STUB: not implemented"; return }

func (do *Options) ResetBodyAndPrint(req *http.Request, resp *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (do *Options) debugPrint(req *http.Request, rsp *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

// write request header

// write req body

// write rsp body
