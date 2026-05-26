package filter

import (
	"errors"
	"net/http"
	"time"

	"github.com/guonaihong/gout/dataflow"
)

var (
	// RetryWaitTime retry basic wait time
	RetryWaitTime = 200 * time.Millisecond
	// RetryMaxWaitTime Maximum retry wait time
	RetryMaxWaitTime = 10 * time.Second
	// RetryAttempt number of retries
	RetryAttempt = 1
)

var (
	ErrRetryFail = errors.New("retry fail")
	ErrRetry     = errors.New("need to retry")
)

// Retry is the core data structure of the retry function
// https://amazonaws-china.com/cn/blogs/architecture/exponential-backoff-and-jitter/
type Retry struct {
	df          *dataflow.DataFlow
	attempt     int // Maximum number of attempts
	currAttempt int
	maxWaitTime time.Duration
	waitTime    time.Duration
	cb          func(c *dataflow.Context) error
}

func (r *Retry) New(df *dataflow.DataFlow) interface{} { _ = "STUB: not implemented"; return nil }

// Attempt set the number of retries
func (r *Retry) Attempt(attempt int) dataflow.Retry {
	_ = "STUB: not implemented"
	return *new(dataflow.Retry)
}

// WaitTime sets the basic wait time
func (r *Retry) WaitTime(waitTime time.Duration) dataflow.Retry {
	_ = "STUB: not implemented"
	return *new(dataflow.Retry)
}

// MaxWaitTime Sets the maximum wait time
func (r *Retry) MaxWaitTime(maxWaitTime time.Duration) dataflow.Retry {
	_ = "STUB: not implemented"
	return *new(dataflow.Retry)
}

func (r *Retry) Func(cb func(c *dataflow.Context) error) dataflow.Retry {
	_ = "STUB: not implemented"
	return *new(dataflow.Retry)
}

func (r *Retry) reset() { _ = "STUB: not implemented"; return }

func (r *Retry) init() {
	if r.attempt == 0 {
		r.attempt = RetryAttempt
	}

	if r.waitTime == 0 {
		r.waitTime = RetryWaitTime
	}

	if r.maxWaitTime == 0 {
		r.maxWaitTime = RetryMaxWaitTime
	}
}

// Does not pollute the namespace
func (r *Retry) min(a, b uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (r *Retry) getSleep() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

//对int64边界处理, 后面使用rand.Int63n所以,最大值只能是int64的最大值防止溢出

func (r *Retry) genContext(resp *http.Response, err error) *dataflow.Context {
	_ = "STUB: not implemented"
	return nil
}

// Do send function
func (r *Retry) Do() (err error) { _ = "STUB: not implemented"; return nil }

// 这里只要调用Func方法，且回调函数返回ErrRetry 会生成新的*http.Request对象
// 不使用DataFlow.Do()方法原因基于两方面考虑
// 1.为了效率只需经过一次编码器得到*http.Request,如果需要重试几次后面是多次使用解码器.Bind()函数
// 2.为了更灵活的控制

//为的是输出debug信息

// 外部可以使用context直接取消
