package _close

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func isCanceledByChannel(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true // 收到取消消息

	default:
		return false // 未收到取消消息
	}
}

// 任意发一个结束信号，运行发现只有 1 个 Runner 被取消
// 和数量耦合，需要知道有多少个 Runner
func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}

}

func cancel2(cancelChan chan struct{}) {
	// 不需要知道有多少个 runner
	close(cancelChan)
}

func TestCancel1(t *testing.T) {
	cancelChan := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func(i int, cancelCh chan struct{}) {
			//for {
			//	if isCanceledByChannel(cancelCh) {
			//		break
			//	}
			//	time.Sleep(5 * time.Millisecond)
			//}

			// 等价改法，提升效率和可读性
			for !isCanceledByChannel(cancelCh) {
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is canceled (cancel1)")
		}(i, cancelChan)
	}
	cancel1(cancelChan) // 只能取消一个
	time.Sleep(time.Second * 1)
	fmt.Println("cancel1: now n =", runtime.NumGoroutine())
}

func TestCancel2(t *testing.T) {
	cancelChan := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int, cancelCh chan struct{}) {
			for !isCanceledByChannel(cancelCh) {
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is canceled (cancel2)")
		}(i, cancelChan)
	}
	cancel2(cancelChan) // 广播取消
	time.Sleep(time.Second * 1)
	fmt.Println("cancel2: now n =", runtime.NumGoroutine())
}
