package ch00

import (
	"sync"
	"testing"
)

// for range 原生支持 整型（1.22起）、数组、切片、字符串、map、channel
func TestIterationOverIntAndChan(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Go 1.22 解决了循环变量重用的问题 并引入了 for range 对 int 的迭代支持
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- i
		}()
	}

	// closing is necessary when the receiver
	// must be told there are no more values coming
	go func() {
		wg.Wait()
		close(ch)
	}()
	// otherwise, oops, deadlock
	for i := range ch {
		t.Log(i)
	}
}

func Test(t *testing.T) {

}
