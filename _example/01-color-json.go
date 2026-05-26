package main

import (
	"time"
)

// ================SetJSON=接口用法===============
// gout使用SetJSON函数发送json请求至服务端
// 亮点有SetJSON函数支持多种数据类型，map/struct/array/string/bytes
// 下面的xxxExample对应该类型的用法

func mapExample() { _ = "STUB: not implemented"; return }

func structExample() { _ = "STUB: not implemented"; return }

var query = `
{
  "query": {
    "bool": {
      "must": [
        {
          "exists": {
            "field": "voice"
          }
        },
        {
          "match": {
              "errcode": 3
          }
        },
        {
          "range": {
            "time": {
              "lt": "2020-01-13T23:16:04+08:00",
              "gt": "2020-01-13T00:00:00+08:00"
            }
          }
        }
      ]
    }
  }
}
`

func stringExample() { _ = "STUB: not implemented"; return }

func bytesExample() { _ = "STUB: not implemented"; return }

func main() {
	go server()

	time.Sleep(time.Millisecond * 200)

	mapExample()
	structExample()
	stringExample()
	bytesExample()
}

func server() { _ = "STUB: not implemented"; return }
