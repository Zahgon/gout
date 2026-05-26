package bench

import (
	"sync"
	"time"
)

// SubTasker 是task模块的核心接口
type SubTasker interface {
	Init()
	Process(chan struct{})
	Cancel()
	WaitAll()
}

// Task Task模块的核心数据结构
type Task struct {
	Duration   time.Duration //压测时间
	Number     int           //压测次数
	Concurrent int           //并发数
	Rate       int           //压测频率

	work chan struct{}

	ok bool

	wg sync.WaitGroup
}

func (t *Task) init() {
	t.work = make(chan struct{})
	if t.Concurrent == 0 {
		t.Concurrent = 1
	}
	t.ok = true
}

func (t *Task) producer() { _ = "STUB: not implemented"; return }

// 控制压测时间

// t.Number < 0

func (t *Task) run(sub SubTasker) { _ = "STUB: not implemented"; return }

//time.Sleep(next.Sub(time.Now()))

//default:

//select里面包含default:会产生一个bug，试问t.Rate如果是很大的值, time.Sleep这句相当于没有
//决定消费者可以消费多少条是消费者自己决定，消费有多块，就可以产生多少令牌给消费者使用
//这和一开始的设计初衷相悖，消费者消费多少条需由t.Number或 t.Duration决定
//注释可以让t.Number 或 t.Duration更准确

// Run Task模块的入口函数
func (t *Task) Run(sub SubTasker) { _ = "STUB: not implemented"; return }
