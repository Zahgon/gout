package main

import (
	"fmt"
	"time"

	"github.com/guonaihong/gout"
)

const (
	benchNumber     = 30000
	benchConcurrent = 20
)

func server() { _ = "STUB: not implemented"; return }

func main() {
	go server()
	time.Sleep(time.Millisecond)

	err := gout.
		POST(":8080").
		SetJSON(gout.H{"hello": "world"}). //设置请求body内容
		Filter().                          //打开过滤器
		Bench().                           //选择bench功能
		Concurrent(benchConcurrent).       //并发数
		Number(benchNumber).               //压测次数
		Do()

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
