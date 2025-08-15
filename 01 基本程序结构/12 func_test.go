package ch01

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 自定义类型，让程序简单可读
type IntConv func(int) int

func cubic(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return x * x * x
}

func timeit(inner IntConv) IntConv {
	// 装饰器，函数作为返回值一般在闭包和构建功能中间件时使用得比较多
	return func(x int) int {
		tic := time.Now()
		ret := inner(x)
		fmt.Println("time spent =>", time.Since(tic).Seconds())
		return ret
	}
}

func TestFunctionAsValue(t *testing.T) {
	// 值传递，函数可以作为变量值、参数、返回值，可以有多个返回值
	t.Log(timeit(cubic)(9))
}

func sum(ops ...int) int {
	// 可变长参数会被转换为一个数组
	if len(ops) == 0 {
		return 0
	}
	// 语法将切片解包传递给递归调用（Python 用 * 表示）
	return sum(ops[1:]...) + ops[0]
}

func sumIter(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVariadicArguments(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5))
	t.Log(sumIter(1, 2, 3, 4, 5))
}

func TestDeferWithPanic(t *testing.T) {
	// 函数（异常）返回前，类似于 Java 的 finally
	// 通常用于（安全地）释放资源，LIFO 执行顺序
	defer func() {
		fmt.Println("Clear resources")
	}()
	fmt.Println("Started")
	// defer 仍然会执行
	panic("Fatal error")
}

func TestDeferExecutionOrder(t *testing.T) {
	var f = func() {
		defer fmt.Println("D")
		fmt.Println("F")
	}

	f()
	fmt.Println("M")
	// output: F D M
}

func TestDeferModifiesReturn(t *testing.T) {
	var f = func(i int) (r int) {
		defer func() {
			r += i
		}()

		/*
			流程：
			先将返回值result设为2
			执行defer语句，将result更新
			真正返回给调用方
		*/
		return 2
	}

	fmt.Println(f(10))
}

func GetFn() func() {
	fmt.Print("[outside]")
	return func() {
		fmt.Print("[inside]")
	}
}

func TestDeferWithFunctionCall(t *testing.T) {
	// 先执行 GetFn()，打印 [outside]
	// 返回打印 [inside] 的函数，延迟执行
	// 打印 [here]，打印 [inside]
	defer GetFn()()
	fmt.Print("[here]")
}

func TestBlockSpace(t *testing.T) {
	x := 11
	fmt.Println(x) // 11
	{
		fmt.Println(x) // 11
		x := 12
		fmt.Println(x) // 12
	}
	fmt.Println(x) // 11
}
