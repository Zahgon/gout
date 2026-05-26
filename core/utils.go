package core

// ReadCloseFail 内部测试使用
type ReadCloseFail struct{}

// Read 供测试使用
func (r *ReadCloseFail) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close 内部测试使用
func (r *ReadCloseFail) Close() error { _ = "STUB: not implemented"; return nil }
