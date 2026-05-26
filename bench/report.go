package bench

import (
	"context"
	"net/http"
	"sync"
	"time"
)

var _ SubTasker = (*Report)(nil)

type result struct {
	time       time.Duration
	statusCode int
}

// 数据字段，每个字段都用于显示
type report struct {
	Concurrency     int    //并发数
	Failed          uint64 //出错的连接数
	CompleteRequest uint64 //正常的请求数
	TotalRead       uint64 //统计所有read的流量body+(request line)+(http header)
	TotalBody       uint64 //统计所有body的流量
	TotalWriteBody  uint64 //统计所有写入的body流量
	Tps             float64
	Duration        time.Duration // 连接总时间
	Kbs             float64
	Mean            float64
	AllMean         float64
	Percentage55    time.Duration
	Percentage66    time.Duration
	Percentage75    time.Duration
	Percentage80    time.Duration
	Percentage90    time.Duration
	Percentage95    time.Duration
	Percentage98    time.Duration
	Percentage99    time.Duration
	Percentage100   time.Duration
	StatusCodes     map[int]int
	ErrMsg          map[string]int
}

type ReportData struct {
	SendNum int // 已经发送的http 请求
	report
	Number     int // 发送总次数
	step       int // 动态报表输出间隔
	allResult  chan result
	waitQuit   chan struct{} //等待startReport函数结束
	allTimes   []time.Duration
	ctx        context.Context
	cancel     func()
	getRequest func() (*http.Request, error)

	startTime time.Time
	*http.Client
}

// Report 是报表核心数据结构
type Report struct {
	ReportData
	lerr  sync.Mutex
	lcode sync.Mutex
}

// NewReport is a report initialization function
func NewReport(ctx context.Context,
	c, n int,

	duration time.Duration,

	getRequest func() (*http.Request, error),

	client *http.Client) *Report {
	_ = "STUB: not implemented"
	return nil
}

// Cancel report logic
func (r *Report) Cancel() {
	_ = "STUB: not implemented"

	// Init 初始化报表模块, 后台会起一个统计go程
	return
}

func (r *Report) Init() { _ = "STUB: not implemented"; return }

func (r *Report) addComplete() { _ = "STUB: not implemented"; return }

// 统计错误消息
func (r *Report) addErrAndFailed(err error) { _ = "STUB: not implemented"; return }

// 统计http code数量
func (r *Report) addCode(code int) { _ = "STUB: not implemented"; return }

// Process 负责构造压测http 链接和统计压测元数据
func (r *Report) Process(work chan struct{}) { _ = "STUB: not implemented"; return }

// 统计http code数量

// WaitAll 等待结束
func (r *Report) WaitAll() {
	_ = "STUB: not implemented"

	// TODO 处理错误
	return
}

//输出最终报表

func (r *Report) calBody(resp *http.Response, bodySize uint64) { _ = "STUB: not implemented"; return }

//space
//\r\n

//:space
//\r\n

func genTimeStr(now time.Time) string { _ = "STUB: not implemented"; return "" }

func (r *Report) startReport() { _ = "STUB: not implemented"; return }

//if newInterval := next.Sub(time.Now()); newInterval > 0 {

func (r *Report) outputReport() error { _ = "STUB: not implemented"; return nil }
