package unit

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert" // 引入断言
)

func malSquare(i int) int {
	return i*i + rand.Intn(3)
}

// 单元测试：对最小可测试单元进行检查和验证
// 原则：自动化（至少不能通过打印输出来检查）、独立性（不依赖其他测试用例）、可重复（Mock 框架模拟第三方资源）
func TestSquare(t *testing.T) {
	// Arrange
	// 表格测试/表驱动测试
	cases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 4},
		{3, 9},
	}
	for _, c := range cases {
		// Act
		got := malSquare(c.input)
		want := c.expected
		// Assert
		assert.Equal(t, want, got) // 也可以使用 t.Fail/Error/Fatal
	}
}
