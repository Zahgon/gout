package main

import (
	"fmt"
	"io"
	"os"

	"github.com/guonaihong/gout"
)

func SaveFile(w io.Writer) gout.DebugOpt { _ = "STUB: not implemented"; return *new(gout.DebugOpt) }

func main() {

	fd, err := os.Create("debug.log")
	if err != nil {
		return
	}

	defer fd.Close()

	s := ""
	err = gout.
		// 发起GET请求
		GET("www.baidu.com").
		//打开debug模式
		Debug(SaveFile(fd)).
		//解析响应body至s变量中
		BindBody(&s).
		//结束函数
		Do()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("html size:%d", len(s))
}
