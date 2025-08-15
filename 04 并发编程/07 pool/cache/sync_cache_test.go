package cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type ReusableObj struct {
	Value int
}

func (r *ReusableObj) String() string {
	return fmt.Sprintf("%d", r.Value)
}

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return &ReusableObj{Value: 100}
		},
	}

	pool.Put(&ReusableObj{Value: 1})
	pool.Put(&ReusableObj{Value: 2})
	pool.Put(&ReusableObj{Value: 3})
	runtime.GC()

	// 第一个 Get取出的是私有池中的，私有池仅可放一个对象
	// 后面的Get取出的则是共享池中的，而共享池的是后进先出的。
	fmt.Println(pool.Get().(*ReusableObj))
	fmt.Println(pool.Get().(*ReusableObj))
	fmt.Println(pool.Get().(*ReusableObj))
	fmt.Println(pool.Get().(*ReusableObj))
}
