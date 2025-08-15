package counter

import (
	"sync"
	"sync/atomic"
)

type SimpleCounter struct {
	data int32
}

func (c *SimpleCounter) Inc() {
	c.data++
}

type MutexCounter struct {
	mu   sync.Mutex
	data int32
}

// 对临界区加锁
func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data++
}

type AtomicCounter struct {
	data atomic.Int32
}

func (c *AtomicCounter) Inc() {
	c.data.Add(1)
}

func (c *SimpleCounter) Value() int32 {
	return c.data
}

func (c *MutexCounter) Value() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data
}

func (c *AtomicCounter) Value() int32 {
	return c.data.Load()
}
