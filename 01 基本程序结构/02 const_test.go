package ch01

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Go 没有原生枚举类型，但可以通过常量 + iota模拟
// tips: 容易弄混地，itoa 是 C语言传统命名的整数转字符串的函数，即 integer to ASCII
type Weekday int

const (
	_ Weekday = iota // 常量自动递增声明
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func TestIotaConstSpec(t *testing.T) {
	t.Log(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
	require.Equal(t, Sun == Mon, false)
	// 1. Wed = 3 是不允许的 2. Wed+3 会转换为 Weekday
	require.Equal(t, Wed+3, Weekday(6))
}
