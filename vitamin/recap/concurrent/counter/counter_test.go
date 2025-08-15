package counter

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestSimpleCounter(t *testing.T) {
	var counter SimpleCounter // 计数值
	var wg sync.WaitGroup
	// Go 1.22 语法糖
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				counter.Inc() // 计数器自增
			}
		}()
	}
	wg.Wait() // 等待所有协程完成
	assert.Equal(t, int32(10000), counter.Value())
}

// 不运行任何 Test 函数，只运行 Benchmark 函数
// go test -bench=Bench -run=^$
func BenchmarkCounterMutex(b *testing.B) {
	for b.Loop() {
		var counter MutexCounter // 锁计数器
		var wg sync.WaitGroup
		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 100 {
					counter.Inc() // 计数器自增
				}
			}()
		}
		wg.Wait() // 等待所有协程完成
		assert.Equal(b, int32(10000), counter.Value())
	}
}

func BenchmarkCounterAtomic(b *testing.B) {
	for b.Loop() {
		var counter AtomicCounter // 原子计数器
		var wg sync.WaitGroup
		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 100 {
					counter.Inc() // 计数器自增
				}
			}()
		}
		wg.Wait() // 等待所有协程完成
		assert.Equal(b, int32(10000), counter.Value())
	}
}
