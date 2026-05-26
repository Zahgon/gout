package gout

import (
	"github.com/guonaihong/gout/dataflow"
)

type Client struct {
	options
}

// NewWithOpt 设计哲学
// 1.一些不经常变化的配置放到NewWithOpt里面实现
// 2.一些和http.Client深度绑定的放到NewWithOpt里面实现
// 3.一些可以提升使用体验的放到NewWithOpt里面实现
func NewWithOpt(opts ...Option) *Client { _ = "STUB: not implemented"; return nil }

// GET send HTTP GET method
func (c *Client) GET(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// POST send HTTP POST method
func (c *Client) POST(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// PUT send HTTP PUT method
func (c *Client) PUT(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// DELETE send HTTP DELETE method
func (c *Client) DELETE(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// PATCH send HTTP PATCH method
func (c *Client) PATCH(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// HEAD send HTTP HEAD method
func (c *Client) HEAD(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }

// OPTIONS send HTTP OPTIONS method
func (c *Client) OPTIONS(url string) *dataflow.DataFlow { _ = "STUB: not implemented"; return nil }
