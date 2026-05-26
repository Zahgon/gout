package main

import (
	"time"

	"github.com/guonaihong/gout"
)

// 自定义debug example，下面使用环境变量输出日志输出
// 日志输出功能，使用环境变量打开
func IOSDebug() gout.DebugOpt { _ = "STUB: not implemented"; return *new(gout.DebugOpt) }

//打开颜色高亮

func customExample() { _ = "STUB: not implemented"; return }

// 运行 example(其中env IOS_DEBUG=on 用于设置环境变量)
// env IOS_DEBUG=on go run 10b-debug-custom.go
func main() {
	go server()                        // 起测试服务
	time.Sleep(time.Millisecond * 500) //sleep下等服务端真正起好

	customExample()
}

func server() { _ = "STUB: not implemented"; return }
