package main

import (
	"time"
)

// =====================设置查询字符串example==================
// gout使用SetQuery来设置查询字符串
// 其中SetQuery支持多种数据类型 map/struct/string/array

type testQuery struct {
	// struct里面的form tag是gin用来绑定数据用的，query才是gout需要的tag
	Q1 string    `query:"q1" form:"q1"`
	Q2 int       `query:"q2" form:"q2"`
	Q3 float32   `query:"q3" form:"q3"`
	Q4 float64   `query:"q4" form:"q4"`
	Q5 time.Time `query:"q5" form:"q5" time_format:"unix" time_location:"Asia/Shanghai"`
	Q6 time.Time `query:"q6" form:"q6" time_format:"unixNano" time_location:"Asia/Shanghai"`
	Q7 time.Time `query:"q7" form:"q7" time_format:"2006-01-02" time_location:"Asia/Shanghai"`
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
	// 使用结构体需要设置query tag
	return
}

func stringExample() {
	_ = "STUB: not implemented"
	// 4.使用string
	return
}

func bytesExample() {
	_ = "STUB: not implemented"
	// 4.使用string
	return
}

func main() {
	go server()

	time.Sleep(time.Millisecond)
	mapExample()
	structExample()
	arrayExample()
	stringExample()
	bytesExample()
}

func server() { _ = "STUB: not implemented"; return }
