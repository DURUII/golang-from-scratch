package error

import (
	"errors"
	"fmt"
	"testing"
)

// 一般预置错误，为了后续区分错误类型
var TooSmallError = errors.New("n must be no less than than 2")

// 我们改造 GetFibonacciList 函数，支持错误校验 (error 是个接口)
func GetFibonacciList(n int) ([]int, error) {
	if n < 2 {
		return nil, TooSmallError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList[:], nil
}

func TestError(t *testing.T) {
	// Go 没有异常机制，没有 try ... catch ... finally
	if val, err := GetFibonacciList(1); err != nil {
		// 结构体的所有成员变量都是可比较的，那么结构体就可比较
		if err == TooSmallError {
			fmt.Println("It is less than 2.")
		}
		t.Log(err)
	} else {
		t.Log(val)
	}
}
