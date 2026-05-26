package main

import (
	"time"
)

func useRetryFunc() {
	_ = "STUB: not implemented"
	// 获取一个没有服务绑定的端口
	return
}

//必须是存在的端口

func main() {
	go server()
	time.Sleep(time.Millisecond * 200)
	useRetryFunc()
}

func server() { _ = "STUB: not implemented"; return }
