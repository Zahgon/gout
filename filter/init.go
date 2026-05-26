package filter

import (
	"math/rand"
	"time"

	"github.com/guonaihong/gout/dataflow"
)

var (
	defaultBench = Bench{}
	defaultRetry = Retry{}
)

func init() {
	dataflow.Register("bench", &defaultBench)
	dataflow.Register("retry", &defaultRetry)

	rand.Seed(time.Now().UnixNano())
}
