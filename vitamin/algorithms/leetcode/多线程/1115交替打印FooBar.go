package main

import (
	"fmt"
	"sync"
)

type FooBar struct {
	n      int
	fooSig chan struct{}
	barSig chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:      n,
		fooSig: make(chan struct{}, 1),
		barSig: make(chan struct{}, 1),
	}
	fb.fooSig <- struct{}{}
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooSig
		printFoo()
		fb.barSig <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barSig
		printBar()
		fb.fooSig <- struct{}{}
	}
}

func main() {
	n := 5
	foobar := NewFooBar(n)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		foobar.Bar(func() { fmt.Print("bar") })
	}()

	go func() {
		defer wg.Done()
		foobar.Foo(func() { fmt.Print("foo") })
	}()

	wg.Wait()
}
