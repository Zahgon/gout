package dataflow

type export struct {
	df *DataFlow
}

func (e *export) Curl() Curl { _ = "STUB: not implemented"; return *new(Curl) }
