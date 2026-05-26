package main

import (
	"time"
)

// 写入非结构化数据至http.body里，使用SetBody接口
// 当然结构化数据主要有json/xml/yaml
func stringExample() {
	_ = "STUB: not implemented"
	// 1.发送string
	return
}

// string

func bytesExample() {
	_ = "STUB: not implemented"
	// 2.发送[]byte
	return
}

// []byte

func ReaderExample() {
	_ = "STUB: not implemented"
	// 3.发送实现io.Reader接口的变量
	return
}

// io.Reader

func baseTypeExample() {
	_ = "STUB: not implemented"
	// 4.发送基础类型的变量
	return
}

//float64

//SetBody支持的更多基础类型有int, int8, int16, int32, int64
//uint, uint8, uint16, uint32, uint64
//float32, float64

func main() {
	go server()

	time.Sleep(time.Millisecond)

	stringExample()
	bytesExample()
	ReaderExample()
	baseTypeExample()
}

func server() { _ = "STUB: not implemented"; return }
