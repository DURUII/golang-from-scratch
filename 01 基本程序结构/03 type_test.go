package ch01

import (
	"math"
	"testing"
)

// 布局规范建议顺序：常量、类型、变量、函数
var (
	a int32   = 1
	b         = int64(a) // Go 对于类型转换很严苛，只支持显式类型转换
	c         = MyInt(a)
	d MyFloat = 1.0
)

type MyInt int64
type MyFloat float64

func TestImplicitCasting(t *testing.T) {
	// Go强调显式：无括号、无三元运算符、无继承、无指针运算
	t.Log(math.MaxUint16)
	t.Log(a, b, c, d)
}
