package ch00

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Go 语言是
// 静态【编译期间而非运行时进行类型检查】
// 强类型【不允许不合理的隐式类型混合转换】系统
type myInt int

func (i myInt) String() string {
	return fmt.Sprintf("myInt: %d", -i)
}

type myMap map[int]string

func TestTypeConversion(t *testing.T) {
	var i myInt
	var j = 1
	// 无法将 'i' (类型 myInt) 用作类型 int
	// j = i
	i = myInt(j)
	fmt.Println(i)

	// 无类型常量 (Untyped Constants)
	// 在 Go 里，当你直接写下 1, 3.14, "hello" 这样的字面量时，
	// 它们并不是你想象中的 int, float64, string。
	// 它们是一种特殊的“理想”状态，没有具体的类型就像一块可以捏成任何形状的橡皮泥
	assert.Equal(t, true, i == 1)
	assert.Equal(t, false, reflect.DeepEqual(i, 1))

	var p myMap
	var q = map[int]string{1: "one"}
	// 如果变量 x 的类型 V 与类型 T 具有 相同的底层类型
	// 并且 V 和 T 中至少有一个不是 defined type
	// 那么 x 就可以直接赋值给类型为 T 的变量
	// 否则需要显式转换
	// Go 内置类型中，所有数值类型（int, float64 等）、string 类型、bool 类型都是 defined type
	// map、slice、array、channel、func 这些复合类型本身不是 defined type
	p = q
	assert.Equal(t, false, reflect.DeepEqual(p, q))
}
