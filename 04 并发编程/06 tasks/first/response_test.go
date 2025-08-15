package first

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("Result from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// 存在协程泄漏，可能导致服务器资源耗尽
	// ch := make(chan string)
	
	// 采用 buffered channel 将 发送/接受 解耦，防止协程泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 当第一个人放消息后，接受者就会被唤醒，代码会直接返回
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(1 * time.Second)
	t.Log("After: ", runtime.NumGoroutine())
}
