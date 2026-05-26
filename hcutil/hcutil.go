package hcutil

import (
	"net/http"
)

func ModifyURL(url string) string { _ = "STUB: not implemented"; return "" }

// socks5://username:password@localhost:7890
// socks5://localhost:7890
// localhost:7890
func SetSOCKS5(c *http.Client, socks5URL string) error { _ = "STUB: not implemented"; return nil }

func SetProxy(c *http.Client, proxyURL string) error { _ = "STUB: not implemented"; return nil }

func UnixSocket(c *http.Client, path string) error { _ = "STUB: not implemented"; return nil }
