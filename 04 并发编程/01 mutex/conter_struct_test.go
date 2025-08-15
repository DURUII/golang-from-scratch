package mutex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// 线程安全的计数器（共享内存并发机制）
type Counter struct {
	Name  string
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) GetCount() uint64 {
	c.mu.Lock()
	// 在设计函数的时候就想清楚收尾工作，释放资源
	defer c.mu.Unlock()
	return c.count
}

func worker(id int, c *Counter, wg *sync.WaitGroup) {
	// 在设计函数的时候就想清楚收尾工作，释放资源
	defer wg.Done()
	c.Increment()
	fmt.Println(id, "=>", c.GetCount())
}

func TestCounterStruct(t *testing.T) {
	const numGoRoutines = 10000
	var counter Counter
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go worker(i, &counter, &wg)
	}
	// 检查点，等待 goroutine 都完成任务
	wg.Wait()
	assert.Equal(t, counter.GetCount(), uint64(numGoRoutines))
}
