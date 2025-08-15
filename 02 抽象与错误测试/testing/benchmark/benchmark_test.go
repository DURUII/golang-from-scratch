package benchmark

import (
	"strings"
	"testing"
)

// 基准测试：go test -bench <to-test-func> -benchmem
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	// https://go.dev/blog/testing-b-loop
	// Go 1.24 中，b.Loop() 自动调用 b.ResetTimer() 和 b.StopTimer()
	for b.Loop() {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
		_ = ret
	}
}

func BenchmarkConcatStringByBuilder(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	for b.Loop() {
		// 也可以写 bytes.Buffer 类型
		var s strings.Builder
		for _, elem := range elems {
			s.WriteString(elem)
		}
		_ = s.String()
	}
}
