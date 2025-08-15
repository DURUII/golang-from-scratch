package _select

import (
	"fmt"
	"testing"
)

func selectRandom() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch <- 1
	ch2 <- 2

	// 如果 select 语句在某一时刻有多个 case 同时就绪
	// 它会伪随机地选择其中一个执行
	select {
	case <-ch:
		fmt.Println("ch")
	case <-ch2:
		fmt.Println("ch2")
	}

}

func TestSelectRandom(t *testing.T) {
	// 不要依赖 select 的选择顺序
	for range 5 {
		selectRandom()
	}
	// 永久阻塞当前 goroutine
	select {}
}
