package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(50 * time.Millisecond)
	return "done"
}

// 类似于 Java 中的 FutureTask
func AsyncService() chan string {
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retChan <- ret
		fmt.Println("service exited")
	}()
	return retChan
}

func TestAsync(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	// 超时机制：slow response 是比 quick failure 还可怕的错误
	case <-time.After(100 * time.Millisecond):
		t.Error("timeout")
		// 添加 default 会使 select 非阻塞
	}
}
