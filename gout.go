package gout

import (
	"net/http"
	"time"

	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/debug"
	_ "github.com/guonaihong/gout/export"
	_ "github.com/guonaihong/gout/filter"
)

// debug
type DebugOption = debug.Options //不推荐gout.DebugOption方式引用, 推荐debug.Options引用
type DebugOpt = debug.Apply      //不推荐gout.DebugOpt方式引用，推荐debug.Apply方式引用
type DebugFunc = debug.Func      //不推荐gout.DebugFunc方式引用，推荐debug.Func方式引用

func NoColor() DebugOpt { _ = "STUB: not implemented"; return *new(DebugOpt) }

func Trace() DebugOpt { _ = "STUB: not implemented"; return *new(DebugOpt) }

type Context = dataflow.Context

// New function is mainly used when passing custom http client
func New(c ...*http.Client) *dataflow.Gout { _ = "STUB: not implemented"; return nil }

// GET send HTTP GET method
// 第一种情况
// gout.GET("wwww.demo.xx/test-appkey")
//
// 第二种情况
//
//	type host struct {
//	 Host string
//	 AppKey string
//	}
//
// gout.GET("http://{{.Host}/{{.AppKey}}}", &host{Host:"www.demo.xx", AppKey:"test-appkey"})
func GET(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// POST send HTTP POST method
func POST(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// PUT send HTTP PUT method
func PUT(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// DELETE send HTTP DELETE method
func DELETE(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// PATCH send HTTP PATCH method
func PATCH(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// HEAD send HTTP HEAD method
func HEAD(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// OPTIONS send HTTP OPTIONS method
func OPTIONS(url string, urlStruct ...interface{}) *dataflow.DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// 设置不忽略空值
func NotIgnoreEmpty() { _ = "STUB: not implemented"; return }

// 设置忽略空值
func IgnoreEmpty() { _ = "STUB: not implemented"; return }

// 设置超时时间,
// d > 0, 设置timeout
// d == 0，取消全局变量
func SetTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

func SetDebug(b bool) { _ = "STUB: not implemented"; return }
