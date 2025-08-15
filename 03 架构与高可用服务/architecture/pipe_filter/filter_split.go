package pipe_filter

import (
	"errors"
	"strings"
)

var ErrSplitFilterWrongFormat = errors.New("SplitFilterWrongFormatError")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(r Request) (Response, error) {
	str, ok := r.(string)
	// Let it crash 原则
	if !ok {
		return nil, ErrSplitFilterWrongFormat
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
