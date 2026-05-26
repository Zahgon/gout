package main

import (
	"time"
)

func debugExample() { _ = "STUB: not implemented"; return }

func noColorExample() { _ = "STUB: not implemented"; return }

func main() {
	go server()                        // 起测试服务
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好

	debugExample()
	noColorExample()
}

func server() { _ = "STUB: not implemented"; return }
