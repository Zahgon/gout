package export

import (
	"io"

	"github.com/guonaihong/gout/dataflow"
)

var _ dataflow.Curl = (*Curl)(nil)

type Curl struct {
	w               io.Writer
	df              *dataflow.DataFlow
	longOption      bool
	generateAndSend bool
}

func (c *Curl) New(df *dataflow.DataFlow) interface{} { _ = "STUB: not implemented"; return nil }

func (c *Curl) LongOption() dataflow.Curl { _ = "STUB: not implemented"; return *new(dataflow.Curl) }

func (c *Curl) GenAndSend() dataflow.Curl { _ = "STUB: not implemented"; return *new(dataflow.Curl) }

func (c *Curl) SetOutput(w io.Writer) dataflow.Curl {
	_ = "STUB: not implemented"
	return *new(dataflow.Curl)
}

func (c *Curl) Do() (err error) { _ = "STUB: not implemented"; return nil }

// 清空状态，Setxxx函数拆开使用就不会有问题
