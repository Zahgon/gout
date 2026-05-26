package main

import (
	"time"
)

// 使用SetCookies接口设置两个cookie
func twoCookieExample() {
	_ = "STUB: not implemented"
	// 发送两个cookie
	return
}

// 使用SetCookies接口设置一个cookie
func oneCookieExample() {
	_ = "STUB: not implemented"

	// 发送一个cookie
	return
}

func main() {
	go server()                        // 起测试服务
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好
	twoCookieExample()
	oneCookieExample()
}

func server() { _ = "STUB: not implemented"; return }
