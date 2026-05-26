package main

import (
	"time"
)

// ============== 解析http header
// 使用BindHeader接口解析http header, 基本数据类型可自动绑定
type rspHeader struct {
	Total int       `header:"total"`
	Sid   string    `header:"sid"`
	Time  time.Time `header:"time" time_format:"2006-01-02"`
}

func bindHeader() { _ = "STUB: not implemented"; return }

//解析请求header

func main() {
	go server()

	time.Sleep(time.Millisecond)
	bindHeader()
}

func server() { _ = "STUB: not implemented"; return }
