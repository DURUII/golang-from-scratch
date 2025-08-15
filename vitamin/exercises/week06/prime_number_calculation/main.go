package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const numWorker = 4

var (
	start int = 2
	end   int = 1000000
)

var numTable = make([]int, 1e7)
var taskChan = make(chan int, 1000)
var wg sync.WaitGroup
var jobFinishMsg = make(chan struct{}, 1)

// isPrime 判断是否为素数
func isPrime(num uint) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := uint(3); i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// producer 生产者
func producer(start, end int) {
	for i := start; i <= end; i++ {
		taskChan <- i
	}
}

// worker 消费者
func worker() {
	for {
		select {
		case taskNum := <-taskChan: // 仍有任务
			// 判断是否为素数，并更新数表
			if isPrime(uint(taskNum)) {
				numTable[taskNum] = 1
			}
			wg.Done()
		case <-jobFinishMsg:
			// 全部完成，避免协程泄露
			return
		}
	}
}

// 更新数表
func run(start, end int) {
	if start > end {
		return
	}
	// 启动消费者
	wg.Add(end - start + 1)
	for i := 0; i < numWorker; i++ {
		go worker()
	}
	go producer(start, end)
	wg.Wait()
	close(jobFinishMsg)
}

func main() {
	if len(os.Args) == 3 {
		start, _ = strconv.Atoi(os.Args[1])
		end, _ = strconv.Atoi(os.Args[2])
	} else {
		log.Fatal("usage: prime_number_calculation <start> <end>")
	}
	tic := time.Now()
	run(start, end)
	count := 0
	f, err := os.Create(fmt.Sprintf("primes_%d_%d.txt", start, end))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for i := start; i <= end; i++ {
		if numTable[i] == 1 {
			fmt.Fprintln(writer, i)
			count++
		}
	}
	writer.Flush()
	toc := time.Now()
	fmt.Println("time: ", toc.Sub(tic).String(), "len: ", count)
}
