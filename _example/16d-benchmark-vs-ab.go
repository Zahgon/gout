package main

import (
	"fmt"
	"time"
)

const (
	benchCount      = 100000
	benchConcurrent = 30
)

func server() { _ = "STUB: not implemented"; return }

func runGout() { _ = "STUB: not implemented"; return }

//设置请求body内容
//打开过滤器
//选择bench功能
//并发数
//压测次数

// sudo apt install apache2-utils
var abCmd = fmt.Sprintf(`ab -c %d -n %d -p ../testdata/voice.pcm http://127.0.0.1:8080/`, benchConcurrent, benchCount)

func runAb() { _ = "STUB: not implemented"; return }

func main() {
	go server()
	time.Sleep(300 * time.Millisecond)

	// 设为false，可看ab性能
	startGout := true
	if startGout {
		runGout()
	} else {
		runAb()
	}
}
