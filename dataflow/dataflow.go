package dataflow

import (
	"context"
	"net/http"
	"time"

	"github.com/guonaihong/gout/decode"
	"github.com/guonaihong/gout/middler"
	"github.com/guonaihong/gout/setting"
)

const (
	get     = "GET"
	post    = "POST"
	put     = "PUT"
	delete2 = "DELETE"
	patch   = "PATCH"
	head    = "HEAD"
	options = "OPTIONS"
)

// DataFlow is the core data structure,
// including the encoder and decoder of http data
type DataFlow struct {
	Req
	out *Gout
}

// GET send HTTP GET method
func (df *DataFlow) GET(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// POST send HTTP POST method
func (df *DataFlow) POST(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// PUT send HTTP PUT method
func (df *DataFlow) PUT(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// DELETE send HTTP DELETE method
func (df *DataFlow) DELETE(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// PATCH send HTTP PATCH method
func (df *DataFlow) PATCH(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// HEAD send HTTP HEAD method
func (df *DataFlow) HEAD(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// OPTIONS send HTTP OPTIONS method
func (df *DataFlow) OPTIONS(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// SetSetting
func (df *DataFlow) SetSetting(s setting.Setting) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetHost set host
func (df *DataFlow) SetHost(host string) *DataFlow { _ = "STUB: not implemented"; return nil }

// GetHost return value->host or host:port
func (df *DataFlow) GetHost() (string, error) { _ = "STUB: not implemented"; return "", nil }

// SetMethod set method
func (df *DataFlow) SetMethod(method string) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetURL set url
func (df *DataFlow) SetURL(url string, urlStruct ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

func (df *DataFlow) SetRequest(req *http.Request) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetBody set the data to the http body, Support string/bytes/io.Reader
func (df *DataFlow) SetBody(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetForm send form data to the http body, Support struct/map/array/slice
func (df *DataFlow) SetForm(obj ...interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetWWWForm send x-www-form-urlencoded to the http body, Support struct/map/array/slice types
func (df *DataFlow) SetWWWForm(obj ...interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetQuery send URL query string, Support string/[]byte/struct/map/slice types
func (df *DataFlow) SetQuery(obj ...interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetHeader send http header, Support struct/map/slice types
func (df *DataFlow) SetHeader(obj ...interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetHeader send http header, Support struct/map/slice types
func (df *DataFlow) SetHeaderRaw(obj ...interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// SetJSON send json to the http body, Support raw json(string, []byte)/struct/map types
func (df *DataFlow) SetJSON(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetJSON send json to the http body, Support raw json(string, []byte)/struct/map types
// 与SetJSON的区一区别就是不转义HTML里面的标签
func (df *DataFlow) SetJSONNotEscape(obj interface{}) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// SetXML send xml to the http body
func (df *DataFlow) SetXML(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetYAML send yaml to the http body, Support struct,map types
func (df *DataFlow) SetYAML(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetProtoBuf send yaml to the http body, Support struct types
// obj必须是结构体指针或者[]byte类型
func (df *DataFlow) SetProtoBuf(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// UnixSocket 函数会修改Transport, 请像对待全局变量一样对待UnixSocket
// 对于全局变量的解释可看下面的链接
// https://github.com/guonaihong/gout/issues/373
func (df *DataFlow) UnixSocket(path string) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetProxy 函数会修改Transport，请像对待全局变量一样对待SetProxy
// 对于全局变量的解释可看下面的链接
// https://github.com/guonaihong/gout/issues/373
func (df *DataFlow) SetProxy(proxyURL string) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetSOCKS5 函数会修改Transport,请像对待全局变量一样对待SetSOCKS5
// 对于全局变量的解释可看下面的链接
// https://github.com/guonaihong/gout/issues/373
func (df *DataFlow) SetSOCKS5(addr string) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetCookies set cookies
func (df *DataFlow) SetCookies(c ...*http.Cookie) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindHeader parse http header to obj variable.
// obj must be a pointer variable
// Support string/int/float/slice ... types
func (df *DataFlow) BindHeader(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindBody parse the variables in http body to obj.
// obj must be a pointer variable
func (df *DataFlow) BindBody(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindJSON parse the json string in http body to obj.
// obj must be a pointer variable
func (df *DataFlow) BindJSON(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindYAML parse the yaml string in http body to obj.
// obj must be a pointer variable
func (df *DataFlow) BindYAML(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindXML parse the xml string in http body to obj.
// obj must be a pointer variable
func (df *DataFlow) BindXML(obj interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// BindDecoder allow user parse data by their own decoder
func (df *DataFlow) BindDecoder(decode decode.Decoder) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// Code parse the http code into the variable httpCode
func (df *DataFlow) Code(httpCode *int) *DataFlow { _ = "STUB: not implemented"; return nil }

// Callback parse the http body into obj according to the condition (json or string)
func (df *DataFlow) Callback(cb func(*Context) error) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// Chunked
func (df *DataFlow) Chunked() *DataFlow { _ = "STUB: not implemented"; return nil }

// SetTimeout set timeout, and WithContext are mutually exclusive functions
func (df *DataFlow) SetTimeout(d time.Duration) *DataFlow { _ = "STUB: not implemented"; return nil }

// WithContext set context, and SetTimeout are mutually exclusive functions
func (df *DataFlow) WithContext(c context.Context) *DataFlow { _ = "STUB: not implemented"; return nil }

// SetBasicAuth
func (df *DataFlow) SetBasicAuth(username, password string) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// Request middleware
func (df *DataFlow) RequestUse(reqModify ...middler.RequestMiddler) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// Response middleware
func (df *DataFlow) ResponseUse(responseModify ...middler.ResponseMiddler) *DataFlow {
	_ = "STUB: not implemented"
	return nil
}

// Debug start debug mode
func (df *DataFlow) Debug(d ...interface{}) *DataFlow { _ = "STUB: not implemented"; return nil }

// https://github.com/guonaihong/gout/issues/264
// When calling SetWWWForm(), the Content-Type header will be added automatically,
// and calling NoAutoContentType() will not add an HTTP header
//
// SetWWWForm "Content-Type", "application/x-www-form-urlencoded"
// SetJSON "Content-Type", "application/json"
func (df *DataFlow) NoAutoContentType() *DataFlow { _ = "STUB: not implemented"; return nil }

// https://github.com/guonaihong/gout/issues/343
// content-encoding会指定response body的压缩方法，支持常用的压缩，gzip, deflate, br等
func (df *DataFlow) AutoDecodeBody() *DataFlow { _ = "STUB: not implemented"; return nil }

func (df *DataFlow) IsDebug() bool { _ = "STUB: not implemented"; return false }

// Do send function
func (df *DataFlow) Do() (err error) {
	_ = "STUB: not implemented"

	// Filter filter function, use this function to turn on the filter function
	return nil
}

func (df *DataFlow) Filter() *filter { _ = "STUB: not implemented"; return nil }

// F filter function, use this function to turn on the filter function
func (df *DataFlow) F() *filter {
	_ = "STUB: not implemented"

	// Export filter function, use this function to turn on the filter function
	return nil
}

func (df *DataFlow) Export() *export { _ = "STUB: not implemented"; return nil }

// E filter function, use this function to turn on the filter function
func (df *DataFlow) E() *export { _ = "STUB: not implemented"; return nil }

func (df *DataFlow) SetGout(out *Gout) { _ = "STUB: not implemented"; return }
