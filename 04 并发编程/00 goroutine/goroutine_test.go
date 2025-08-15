package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	fmt.Println("Before:", runtime.NumGoroutine()) // 2
	for i := 0; i < 10; i++ {
		// 在1.22版本修复前，修改 go.mod 重新运行，有竞争关系（https://go.dev/blog/loopvar-preview）
		//go func() {
		//	time.Sleep(500 * time.Millisecond)
		//	fmt.Println(i)
		//}()

		// 无竞争关系，传入副本
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
		}(i)
	}
	fmt.Println("Middle:", runtime.NumGoroutine()) // 12
	// 等待协程执行结束
	time.Sleep(1 * time.Second)
	fmt.Println("After:", runtime.NumGoroutine()) // 2
}
