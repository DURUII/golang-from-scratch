package ch01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareOperator(t *testing.T) {
	var i int
	// Go 语言没有前置的 ++/-- （Python 连后置都没有）
	// 不支持 assert.Equal(t, 1, i++)，没有后缀自增表达式 i++ 的值语义
	i++
	assert.Equal(t, 1, i)

	// 在比较数组时，主流语言比较引用是否相同，但是 Go 语言中：
	// 1. 相同维数可以比较
	// 2. 每个元素都相同才相等 (与 Python 相同)

	// 只有指针、接口、切片、channel、map 和函数的默认值为 nil
	// 而数组、结构体是复合类型，是它们组成元素都为零值的结果
	// 切⽚是不可以⽐较的，运⾏后报错
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 4}
	t.Log(&a == &b)
	t.Log(a == b)
	b[2] = 3
	t.Log(a == b)
}

func TestBitwiseClear(t *testing.T) {
	const (
		Readable = 1 << iota
		Writable
		Executable
	)

	a := 7 // 111
	t.Log(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
		a&^Readable == Readable, // 按位清零
		a&^Writable == Writable,
		a&^Executable == Executable,
	)
}
