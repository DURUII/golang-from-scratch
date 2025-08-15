package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	splitter := NewSplitFilter(",")
	converter := NewToIntFilter()
	adder := NewSumFilter()
	sp := NewStraightPipeline("p1", splitter, converter, adder)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatal(ret, "is not expected")
	}
}
