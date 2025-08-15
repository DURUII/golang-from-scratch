package pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 例如数据库连接、网络连接
type ReusableObj struct {
}

// 用 buffered channel 实现对象池
type ObjPool struct {
	bufChan chan *ReusableObj
}

// 创建对象池
func NewObjPool(size int) *ObjPool {
	objPool := new(ObjPool)
	objPool.bufChan = make(chan *ReusableObj, size)
	for i := 0; i < size; i++ {
		// 预置
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (pool *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case obj := <-pool.bufChan:
		return obj, nil
	// A slow error is even worse than a fast error —— Google SRE
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (pool *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// 尝试释放一个对象，如果失败，说明对象池已满
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			// if err := pool.ReleaseObj(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
}
