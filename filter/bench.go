package filter

import (
	"net/http"
	"time"

	"github.com/guonaihong/gout/bench"
	"github.com/guonaihong/gout/dataflow"
)

// Bench provide benchmark features
type Bench struct {
	bench.Task

	df *dataflow.DataFlow

	r          *bench.Report
	getRequest func() (*http.Request, error)
}

func NewBench() *Bench {
	_ = "STUB: not implemented"

	// New
	return nil
}

func (b *Bench) New(df *dataflow.DataFlow) interface{} { _ = "STUB: not implemented"; return nil }

// Concurrent set the number of benchmarks for concurrency
func (b *Bench) Concurrent(c int) dataflow.Bencher {
	_ = "STUB: not implemented"
	return *new(dataflow.Bencher)
}

// Number set the number of benchmarks
func (b *Bench) Number(n int) dataflow.Bencher {
	_ = "STUB: not implemented"
	return *new(dataflow.Bencher)
}

// Rate set the frequency of the benchmark
func (b *Bench) Rate(rate int) dataflow.Bencher {
	_ = "STUB: not implemented"
	return *new(dataflow.Bencher)
}

// Durations set the benchmark time
func (b *Bench) Durations(d time.Duration) dataflow.Bencher {
	_ = "STUB: not implemented"
	return *new(dataflow.Bencher)
}

func (b *Bench) Loop(cb func(c *dataflow.Context) error) dataflow.Bencher {
	_ = "STUB: not implemented"
	return *new(dataflow.Bencher)
}

// TODO 优化，这里创建了两个dataflow对象
// c.SetGout和 c.getDataFlow 都依赖gout对象
// 后面的版本要先梳理下gout对象的定位

func (b *Bench) GetReport(r *bench.Report) dataflow.Bencher {
	_ = "STUB: not implemented"

	// Do benchmark startup function
	return *new(dataflow.Bencher)
}

func (b *Bench) Do() error {
	_ = "STUB: not implemented"
	// 报表插件
	return nil
}

// task是并发控制模块
