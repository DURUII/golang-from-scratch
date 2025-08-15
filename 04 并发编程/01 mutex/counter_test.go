//go:build !race

package mutex

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// 数据不一致问题
func AsyncCounter(numGoRoutine, numAddition int) int {
	var count = 0
	// var mu sync.Mutex // 地道的用法是使用零值
	var wg sync.WaitGroup

	for i := 0; i < numGoRoutine; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < numAddition; j++ {
				// mu.Lock()   // 不支持在锁的情况下再次请求锁
				count++ // 计数器加 1 不是一个原子操作
				// mu.Unlock() // 有加锁，就有释放锁，在函数中可以 Lock 了之后，立刻 defer Unlock
			}
			wg.Done()
		}()
	}

	// 等待 goroutine 执行完成
	wg.Wait()
	return count
}

func TestAsyncCounter(t *testing.T) {
	const (
		numGoRoutine = 5000
		numAddition  = 1
	)
	assert.Equal(t, numGoRoutine*numAddition, AsyncCounter(numGoRoutine, numAddition))
}
