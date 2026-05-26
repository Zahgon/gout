package main

import (
	"time"
)

type Result struct {
	Errmsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

// Callback接口用于处理服务段会返回多种数据结构，比如404返回出错html, 200返回json
// 客户端example
func callbackExample() { _ = "STUB: not implemented"; return }

//http code为200时，服务端返回的是json 结构

//http code为404时，服务端返回是html 字符串

func main() {
	go server()                        //等会起测试服务
	time.Sleep(time.Millisecond * 500) //用时间做个等待同步

	callbackExample()
}

// 模拟 API网关
func server() { _ = "STUB: not implemented"; return }

//使用随机函数模拟某个服务有一定概率出现404

// 模拟 404 找不到资源

// 正确业务返回结果
