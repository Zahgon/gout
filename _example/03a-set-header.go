package main

import (
	"time"
)

// ============== gout 设置http header example============
// 使用SetHeader接口 设置http header
// SetHeader支持的数据类型有map/array/struct
type testHeader struct {
	H1 string    `header:"h1"`
	H2 int       `header:"h2"`
	H3 float32   `header:"h3"`
	H4 float64   `header:"h4"`
	H5 time.Time `header:"h5" time_format:"unix"`
	H6 time.Time `header:"h6" time_format:"unixNano"`
	H7 time.Time `header:"h7" time_format:"2006-01-02"`
}

func mapExample() {
	_ = "STUB: not implemented"
	// 1.使用gout.H
	return
}

func arrayExample() {
	_ = "STUB: not implemented"
	// 2.使用数组变量
	return
}

func structExample() {
	_ = "STUB: not implemented"
	// 3.使用结构体
	// 使用结构体需要设置"header" tag
	return
}

func main() {
	go server()

	time.Sleep(time.Millisecond)
	mapExample()
	arrayExample()
	structExample()
}

func server() { _ = "STUB: not implemented"; return }
