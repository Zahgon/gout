package gout

import (
	"net/http"
	"time"

	"github.com/guonaihong/gout/setting"
)

type options struct {
	hc *http.Client
	setting.Setting
	err error
}

type Option interface {
	apply(*options)
}

// 1.start
type insecureSkipVerifyOption bool

func (i insecureSkipVerifyOption) apply(opts *options) { _ = "STUB: not implemented"; return }

// 1.忽略ssl验证
func WithInsecureSkipVerify() Option { _ = "STUB: not implemented"; return *new(Option) }

// 2. start
type client http.Client

func (c *client) apply(opts *options) { _ = "STUB: not implemented"; return }

// 2.自定义http.Client
func WithClient(c *http.Client) Option {
	_ = "STUB: not implemented"
	return *

	// 3.start
	new(Option)
}

type close3xx struct{}

func (c close3xx) apply(opts *options) { _ = "STUB: not implemented"; return }

// 3.关闭3xx自动跳转
func WithClose3xxJump() Option {
	_ = "STUB: not implemented"

	// 4.timeout
	return *new(Option)
}

type timeout time.Duration

func WithTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func (t *timeout) apply(opts *options) { _ = "STUB: not implemented"; return }

// 5. 设置代理
type proxy string

func WithProxy(p string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (p *proxy) apply(opts *options) { _ = "STUB: not implemented"; return }

// 6. 设置socks5代理
type socks5 string

func WithSocks5(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (s *socks5) apply(opts *options) { _ = "STUB: not implemented"; return }

// 7. 设置unix socket
type unixSocket string

func WithUnixSocket(u string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (u *unixSocket) apply(opts *options) { _ = "STUB: not implemented"; return }
