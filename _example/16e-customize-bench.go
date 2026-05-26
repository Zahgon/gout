package main

import (
	"time"
)

const (
	benchNumber     = 30000
	benchConcurrent = 30
)

func server() { _ = "STUB: not implemented"; return }

func customize() { _ = "STUB: not implemented"; return }

// 下面的代码，每次生成不一样的http body 用于压测

func main() {
	go server()
	time.Sleep(300 * time.Millisecond)

	customize()
}
