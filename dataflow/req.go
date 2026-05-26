package dataflow

import (
	"bytes"
	"context"
	"net/http"

	"github.com/guonaihong/gout/debug"
	"github.com/guonaihong/gout/decode"
	"github.com/guonaihong/gout/encode"
	"github.com/guonaihong/gout/encoder"
	"github.com/guonaihong/gout/middler"
	"github.com/guonaihong/gout/setting"
)

// Req controls core data structure of http request
type Req struct {
	method   string
	url      string
	host     string
	userName *string
	password *string

	form    []interface{}
	wwwForm []interface{}

	// http body
	bodyEncoder encoder.Encoder
	bodyDecoder []decode.Decoder

	// http header
	headerEncode []interface{}
	// raw header
	rawHeader bool

	headerDecode interface{}

	// query
	queryEncode []interface{}

	httpCode *int
	g        *Gout

	callback func(*Context) error

	// cookie
	cookies []*http.Cookie

	ctxIndex int

	c   context.Context
	Err error

	reqModify []middler.RequestMiddler

	responseModify []middler.ResponseMiddler

	req *http.Request

	// 内嵌字段
	setting.Setting

	cancel context.CancelFunc
}

// Reset 重置 Req结构体
// req 结构布局说明，以decode为例
// body 可以支持text, json, yaml, xml，所以定义成接口形式
// headerDecode只有一个可能，就定义为具体类型。这里他们的decode实现也不一样
// 有没有必要，归一化成一种??? TODO:
func (r *Req) Reset() { _ = "STUB: not implemented"; return }

func isAndGetString(x interface{}) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (r *Req) addDefDebug() { _ = "STUB: not implemented"; return }

func (r *Req) addContextType(req *http.Request) { _ = "STUB: not implemented"; return }

func (r *Req) selectRequest(body *bytes.Buffer) (req *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Req) encodeQuery() error { _ = "STUB: not implemented"; return nil }

func (r *Req) encodeForm(body *bytes.Buffer, f *encode.FormEncode) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Req) encodeWWWForm(body *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }

// Request Get the http.Request object
func (r *Req) Request() (req *http.Request, err error) { _ = "STUB: not implemented"; return nil, nil }

// 如果同时传递调用SetWWWForm和SetJSON函数，默认json优先级别比较高

// set http body

// set query header

// TODO
// 可以考虑和 bodyEncoder合并,
// 头疼的是f.FormDataContentType如何合并，每个encoder都实现这个方法???

// 放这个位置不会误删除SetForm的http header

// set http header

// 运行请求中间件

func (r *Req) encodeHeader(req *http.Request) (err error) { _ = "STUB: not implemented"; return nil }

func clearHeader(header http.Header) { _ = "STUB: not implemented"; return }

// retry模块需要context.Context，所以这里也返回context.Context
func (r *Req) GetContext() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

// TODO 优化代码，每个decode都有自己的指针偏移直接指向流，减少大body的内存使用
func (r *Req) decodeBody(req *http.Request, resp *http.Response) (err error) {
	_ = "STUB: not implemented"
	return nil

	// 当只有一个解码器时，直接在流上操作，避免读取整个响应体
}

// 确保在读取完成后关闭body

// 当有多个解码器需要处理响应体时，才读取整个响应体到内存中

// 已经取走数据，直接关闭body

func (r *Req) decode(req *http.Request, resp *http.Response, openDebug bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// This is code(output debug info) be placed here
// all, err := ioutil.ReadAll(resp.Body)
// respBody  = bytes.NewReader(all)

// 运行响应中间件。放到debug打印后面，避免混淆请求返回内容

func (r *Req) getDataFlow() *DataFlow { _ = "STUB: not implemented"; return nil }

const maxBodySlurpSize = 4 * (2 << 10) // 4KB
func clearBody(resp *http.Response) error {
	_ = "STUB: not implemented"
	// 这里限制下io.Copy的大小
	return nil
}

func (r *Req) Bind(req *http.Request, resp *http.Response) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// 注意这里的r.callback使用了r.DataFlow的地址, r.callback和r.decode操作的是同一个的DataFlow
// 执行r.callback只是装载解码器, 后面的r.decode才是真正的解码

// 如果没有设置解码器

func (r *Req) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (r *Req) getDebugOpt() *debug.Options { _ = "STUB: not implemented"; return nil }

func (r *Req) canTrace() bool { _ = "STUB: not implemented"; return false }

// 使用chunked方式
func (r *Req) maybeUseChunked(req *http.Request) { _ = "STUB: not implemented"; return }

// getReqAndRsp 内部函数获取req和resp
func (r *Req) getReqAndRsp() (req *http.Request, rsp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// 如果调用Chunked()接口, 就使用chunked的数据包

// resp, err := r.Client().Do(req)
// TODO r.Client() 返回Do接口

// Response 获取原始http.Response数据结构
func (r *Req) Response() (rsp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do Send function
func (r *Req) Do() (err error) { _ = "STUB: not implemented"; return nil }

// reset  Req

func modifyURL(url string) string { _ = "STUB: not implemented"; return "" }

func reqDef(method string, url string, g *Gout, urlStruct ...interface{}) (Req, error) {
	_ = "STUB: not implemented"
	return *new(Req), nil
}

// ReadAll returns the whole response body as bytes.
// This is an optimized version of `io.ReadAll`.
func ReadAll(resp *http.Response) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// if we know the body length we can allocate the buffer only once

// using `bytes.NewBuffer` + `io.Copy` is much faster than `io.ReadAll`
// see https://github.com/elastic/beats/issues/36151#issuecomment-1931696767
