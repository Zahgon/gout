package autodecodebody

import (
	"net/http"
)

func AutoDecodeBody(rsp *http.Response) (err error) { _ = "STUB: not implemented"; return nil }

//case "gzip": // net/http包里面已经做了gzip自动解码
//rc, err = gzip.NewReader(rsp.Body)

// compress 是一种浏览器基本不使用的压缩格式，暂不考虑支持
