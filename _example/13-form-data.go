package main

import (
	"time"
)

type testForm struct {
	Mode string `form:"mode"`
	Text string `form:"text"`
	//Voice []byte `form:"voice" form-mem:"true"` //todo open
}

type testForm2 struct {
	Mode   string `form:"mode"`
	Text   string `form:"text"`
	Voice  string `form:"voice" form-file:"file"` //从文件中读取
	Voice2 []byte `form:"voice2" form-file:"mem"` //从内存中构造
}

// 使用map装载数据
func mapExample() {
	_ = "STUB: not implemented"

	// 1.使用gout.H
	return
}

// 使用结构体装载数据
func structExample() {
	_ = "STUB: not implemented"

	// 2.使用结构体里面的数据
	return
}

// 自定义filename
func mapExample2() {
	_ = "STUB: not implemented"

	// 2.使用结构体里面的数据
	return
}

func main() {
	go server()
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好

	mapExample()
	structExample()
	mapExample2()
}

func server() { _ = "STUB: not implemented"; return }
