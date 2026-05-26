package setting

import (
	"time"

	"github.com/guonaihong/gout/debug"
)

// 设置
type Setting struct {
	// debug相关字段
	debug.Options
	// 控制是否使用空值
	NotIgnoreEmpty bool

	//是否自动加ContentType
	NoAutoContentType bool
	//超时时间
	Timeout time.Duration

	UseChunked bool
}

// 使用chunked数据
func (s *Setting) Chunked() { _ = "STUB: not implemented"; return }

func (s *Setting) SetTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

func (s *Setting) SetDebug(b bool) { _ = "STUB: not implemented"; return }

func (s *Setting) Reset() { _ = "STUB: not implemented"; return }

//s.TimeoutIndex = 0
