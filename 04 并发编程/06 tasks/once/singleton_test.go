package once

import (
	"fmt"
	"sync"
	"testing"
)

type Object struct {
	id int
}

// Stringer 接口
func (o *Object) String() string {
	return fmt.Sprintf("Object: %p, Object{id: %v}", o, o.id)
}

var singleInstance *Object
var once sync.Once

// 单例模式（懒汉式，线程安全）
func GetSingletonObject() *Object {
	once.Do(func() {
		fmt.Println("Create Object")
		singleInstance = new(Object)
	})
	return singleInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObject()
			fmt.Println("Get SingletonObject:", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
