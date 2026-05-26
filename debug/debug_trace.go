package debug

import (
	"io"
	"net/http"
	"time"

	"github.com/guonaihong/gout/middler"
)

type TraceInfo struct {
	DnsDuration          time.Duration
	ConnDuration         time.Duration
	TLSDuration          time.Duration
	RequestDuration      time.Duration
	WaitResponseDuration time.Duration
	ResponseDuration     time.Duration
	TotalDuration        time.Duration
	w                    io.Writer
}

func (t *TraceInfo) StartTrace(opt *Options, needTrace bool, req *http.Request, do middler.Do) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 格式化成json

func (t *TraceInfo) output(opt *Options) { _ = "STUB: not implemented"; return }
