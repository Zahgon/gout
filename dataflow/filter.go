package dataflow

type filter struct {
	df *DataFlow
}

func (f *filter) Bench() Bencher { _ = "STUB: not implemented"; return *new(Bencher) }

func (f *filter) Retry() Retry { _ = "STUB: not implemented"; return *new(Retry) }
