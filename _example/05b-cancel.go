package main

import (
	"time"
)

func cancelExample() {
	_ = "STUB: not implemented"
	// 给http请求 设置超时
	return
}

//取消

func main() {
	go server()
	time.Sleep(time.Millisecond)
	cancelExample()
}

func server() { _ = "STUB: not implemented"; return }
