package main

import (
	"time"
)

type data struct {
	Id   int    `json:"id" xml:"id"`
	Data string `json:"data xml:"data""`
}

func useStruct() { _ = "STUB: not implemented"; return }

func useString() { _ = "STUB: not implemented"; return }

func useBytes() { _ = "STUB: not implemented"; return }

func main() {
	go server()
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好

	useStruct()
	useString()
	useBytes()
}

func server() { _ = "STUB: not implemented"; return }
