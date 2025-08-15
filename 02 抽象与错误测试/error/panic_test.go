package error

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestOSExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally")
	}()
	fmt.Println("Started")
	// defer 函数是不会执行的
	os.Exit(-1)
}

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("Finally")
	}()
	fmt.Println("Started")
	// defer 函数会执行
	panic(errors.New("fatal error"))
}

func TestRecover(t *testing.T) {
	defer func() {
		// 尽管如此，错误恢复机制 & “Let it Crash!”：避免僵尸进程
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	fmt.Println("Started")
	// defer 函数会执行
	panic(errors.New("fatal error"))
}
