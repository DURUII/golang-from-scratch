package ch01

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {
	// switch 不限制常量整数/布尔值，可以是字符串
	// 支持局部变量初始化
	switch os := runtime.GOOS; os {
	// 支持多结果选项
	// switch 句末不需要额外加 break （与 C 不同）
	case "darwin", "linux":
		fmt.Println("Linux-Based System")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func TestFallthrough(t *testing.T) {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			// 穿透 keyword：继续执行下一个case中的代码，而不管下一个case条件是否满足
			fallthrough
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}
