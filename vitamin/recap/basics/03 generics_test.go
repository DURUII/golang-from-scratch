package ch00

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// Go 1.18 引入泛型
// T 是类型参数（通常用单个大写字母表示）
// [T any] 中的 any接口 是类型约束，还有例如，Comparable/Ordered、Stringer、TypeSet如Number
func Filter[T any](s []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

type Case[T constraints.Ordered] struct {
	input     []T
	predicate func(T) bool
	want      []T
}

func TestGenericFilter(t *testing.T) {
	cases := []Case[int]{
		{[]int{-2, 0, 3, -5, 2, -1}, func(x int) bool { return x < 0 }, []int{-2, -5, -1}},
	}

	for _, c := range cases {
		got := Filter(c.input, c.predicate)
		assert.Equal(t, c.want, got)
	}
}
