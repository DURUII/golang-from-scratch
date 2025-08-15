package series

import "fmt"

// 一个包中，可以包含多个 init 函数
// Go 运行时自动调用，不允许手动调用
func init() {
	fmt.Print("A ")
}

func init() {
	fmt.Print("B ")
}

func GetFibonacciList(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList[:]
}

func init() {
	fmt.Print("C ")
}
