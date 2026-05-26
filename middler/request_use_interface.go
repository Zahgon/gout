package middler

import "net/http"

type RequestMiddlerFunc func(req *http.Request) error

type RequestMiddler interface {
	ModifyRequest(req *http.Request) error
}

func (f RequestMiddlerFunc) ModifyRequest(req *http.Request) error {
	_ = "STUB: not implemented"

	// WithRequestMiddlerFunc 是创建一个 RequestMiddler 的helper
	// 如果我们只需要简单的逻辑，只关注闭包本身，则可以使用这个helper快速创建一个 RequestMiddler
	return nil
}

func WithRequestMiddlerFunc(f RequestMiddlerFunc) RequestMiddler {
	_ = "STUB: not implemented"
	return *new(RequestMiddler)
}
