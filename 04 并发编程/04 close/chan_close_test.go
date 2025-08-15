package _close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
// 矛盾 1：Receiver 不知道 Producer 放多少数据，是否放完
// 矛盾 2：假如放 -1 作为结束标志，但是 Producer 不知道放几个结束标志（有几个 Receiver，每个 Receiver 都需要一个结束标志）
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭通道，否则两个消费者都会死等数据，避免 dead lock
		// 向关闭的 channel 发送数据，会导致 panic
		close(ch)
		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// ok 用于退出信号的广播机制
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				// 如果 channel 被关闭，立即返回零值
				fmt.Println("channel closed, with zero value", data)
				break
			}
		}
		wg.Done()
	}()
}

func TestChanClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Wait()
}
