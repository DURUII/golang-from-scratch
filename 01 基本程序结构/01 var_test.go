package ch01

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Go 的文件命名规则是用全小写字母形式的短小单词命名，
下划线在 Go 语言中特殊作用，如_test表示测试代码。
*/

// 测试
func TestFibonacciSeries(t *testing.T) {
	// 类型关键字在后面（一般用于变量声明，主要是全局/外部）
	// 静态类型（与 Python/JavaScript 不同）
	// var a int = 1
	// var b int = 1

	// 虽然是静态类型，但具有自动类型推断* （与 Java 不同）
	// var (
	//	a = 1
	//	b = 1
	// )

	// 短声明，支持多个变量赋值* （与 Python 相同）
	a, b := 1, 1

	fmt.Print(a, ",")
	for i := 0; i < 10; i++ {
		// C 语言的写法可能是：
		// tmp := a
		// a = b
		// b = tmp + a
		// Go 支持多个变量赋值* （与 Python 相同）
		a, b = b, a+b
		fmt.Print(a, ",")
	}
	fmt.Println(b)
}

func TestExchangeNum(t *testing.T) {
	a, b := 1, 3
	// a= 1 b= 3
	t.Log("a=", a, "b=", b)
	a, b = b, a
	// a= 3 b= 1
	t.Log("a=", a, "b=", b)
}

func TestKeywords(t *testing.T) {
	// 指针、接口、切片、channel、map 和函数的默认值为 nil
	var intFunc func(int) int
	fmt.Println(intFunc == nil)

	// 问：nil 是不是关键字？答：不是，在作用域内可被重定义。
	var a [3]int
	var b []int
	// 数组切片的初始值为元素零值
	fmt.Println(reflect.TypeOf(a), "a=", a)
	// 切片的初始值为 nil
	fmt.Println(reflect.TypeOf(b), "b=", b, b == nil)
	// nil 不是 25 个关键字之一，因为可以被覆盖作为变量名 (不要在实际代码中使用)
	// shadows declaration at builtin.go
	nil := []int{1, 2, 3}
	// append 要赋值给新的变量
	b = append(b, 1, 2, 3)
	fmt.Println(reflect.TypeOf(b), "b=", b, reflect.DeepEqual(b, nil))
}
