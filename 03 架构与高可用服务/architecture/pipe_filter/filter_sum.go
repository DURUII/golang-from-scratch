package pipe_filter

import (
	"errors"
)

var ErrSumFilterWrongFormat = errors.New("WrongFormatError")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, ErrSumFilterWrongFormat
	}
	ret := 0
	for _, elm := range elems {
		ret += elm
	}
	return ret, nil
}
