package main

import (
	"time"
)

func rawhttp() { _ = "STUB: not implemented"; return }

func main() {
	go server()
	time.Sleep(time.Millisecond * 200)
	rawhttp()
}

func server() { _ = "STUB: not implemented"; return }
