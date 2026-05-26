package main

import (
	"time"
)

var first bool

func useRetryFuncCode() { _ = "STUB: not implemented"; return }

func main() {
	first = true
	go server()
	time.Sleep(time.Millisecond * 200)
	useRetryFuncCode()
}

// mock 服务端函数
func server() { _ = "STUB: not implemented"; return }
