package main

import (
	"time"
)

func uploadExample() { _ = "STUB: not implemented"; return }

// upload file

func main() {
	go server()
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好

	uploadExample()
}

func server() { _ = "STUB: not implemented"; return }
