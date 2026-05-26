package dataflow

import (
	"net/http"
)

// Gout is the data structure at the beginning of everything
type Gout struct {
	*http.Client
	DataFlow // TODO 优化
}

var (
	// DefaultClient The default http client, which has a connection pool
	DefaultClient = http.Client{}
	// DefaultBenchClient is the default http client used by the benchmark,
	// which has a connection pool
	DefaultBenchClient = http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10000,
		},
	}
)

// New function is mainly used when passing custom http client
func New(c ...*http.Client) *Gout { _ = "STUB: not implemented"; return nil }

// TODO 这一层可以直接删除
// v0.3.3版本开始算起， v0.3.7版本将会删除
// GET send HTTP GET method
func GET(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// POST send HTTP POST method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func POST(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// PUT send HTTP PUT method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func PUT(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// DELETE send HTTP DELETE method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func DELETE(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// PATCH send HTTP PATCH method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func PATCH(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// HEAD send HTTP HEAD method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func HEAD(url string) *DataFlow { _ = "STUB: not implemented"; return nil }

// OPTIONS send HTTP OPTIONS method
// v0.3.3版本开始算起， v0.3.7版本将会删除
func OPTIONS(url string) *DataFlow { _ = "STUB: not implemented"; return nil }
