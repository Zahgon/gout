package main

import (
	"time"
)

// 使用SetWWWForm 接口设置x-www-form-urlencoded格式数据
// 下面的代码有 map/array/struct 使用example
type testWWWForm struct {
	Int     int     `form:"int" www-form:"int"`
	Float64 float64 `form:"float64" www-form:"float64"`
	String  string  `form:"string" www-form:"string"`
}

func mapExample() { _ = "STUB: not implemented"; return }

// 1.第一种方式，使用gout.H

func arrayExample() { _ = "STUB: not implemented"; return }

// 2.第一种方式，使用gout.A

func structExample() { _ = "STUB: not implemented"; return }

// 3.第一种方式，使用结构体

func main() {

	go server()

	time.Sleep(time.Millisecond * 500)

	mapExample()
	arrayExample()
	structExample()
}

func server() { _ = "STUB: not implemented"; return }
