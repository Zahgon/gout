package main

import (
	"time"
)

// 可以使用BindBody解析返回的非结构化http body(结构化指json/xml/yaml等)
// 可以做到基础类型自动绑定, 下面是string/[]byte/int的example
func bindString() {
	_ = "STUB: not implemented"
	// 1.解析string
	return
}

func bindBytes() {
	_ = "STUB: not implemented"
	// 2.解析[]byte
	return
}

func bindInt() {
	_ = "STUB: not implemented"
	// 3.解析int
	return
}

//BindBody支持的更多基础类型有int, int8, int16, int32, int64
//uint, uint8, uint16, uint32, uint64
//float32, float64

func main() {
	go server()

	time.Sleep(time.Millisecond)

	bindString()
	bindBytes()
	bindInt()
}

func server() { _ = "STUB: not implemented"; return }
