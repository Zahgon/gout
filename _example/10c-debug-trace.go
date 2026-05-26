package main

import (
	"time"
)

func openDebugTrace() { _ = "STUB: not implemented"; return }

func main() {
	go server()

	time.Sleep(time.Millisecond * 200)

	openDebugTrace()
}

func server() { _ = "STUB: not implemented"; return }
