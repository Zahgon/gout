package debug

import (
	"io"
)

// 默认，不需要调用
func DefaultDebug(o *Options) { _ = "STUB: not implemented"; return }

// NoColor Turn off color highlight debug mode
func NoColor() Apply { _ = "STUB: not implemented"; return *new(Apply) }

type file struct{ fileName string }

func (f *file) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// 第一个参数是要写入的文件名, 第二个参数是否颜色高亮
func ToFile(fileName string, color bool) Apply { _ = "STUB: not implemented"; return *new(Apply) }

// 第一个参数是要写入的io.Writer对象， 第二个参数是否颜色高亮
func ToWriter(w io.Writer, color bool) Apply { _ = "STUB: not implemented"; return *new(Apply) }

func OnlyTraceFlag() Apply { _ = "STUB: not implemented"; return *new(Apply) }

// trace信息格式化成json输出至标准输出
func TraceJSON() Apply { _ = "STUB: not implemented"; return *new(Apply) }

// trace信息格式化成json输出至w
func TraceJSONToWriter(w io.Writer) Apply { _ = "STUB: not implemented"; return *new(Apply) }

// 打开Trace()
func Trace() Apply { _ = "STUB: not implemented"; return *new(Apply) }
