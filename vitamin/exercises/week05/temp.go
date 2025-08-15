package main

import "time"

var c = make(chan int, 1)

var a int

func f() {
	a = 1
	<-c
}

// for i in {1..10000}; do go run temp.go; done | grep -E '^1$'
func main() {
	go f()
	time.Sleep(3500 * time.Nanosecond)
	c <- 0
	print(a)
}
