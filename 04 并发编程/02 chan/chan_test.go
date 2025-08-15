package _chan

import (
	"fmt"
	"testing"
	"time"
)

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(1 * time.Second)
	fmt.Println("task finished")
}

func service() string {
	time.Sleep(2 * time.Second) // 模拟耗时操作
	return "done"
}

// 类似于 Java 中的 FutureTask
func asyncService() chan string {
	// channel 确保发送方和接收方在同一时间点同步
	// 发送操作会阻塞发送方，直到有接收方读取该值
	// 同样，接收操作也会阻塞，直到有发送方提供值
	// retChan := make(chan string) 

	// buffered channel 需要指定容量，允许发送方和接收方在时间上解耦
	// 发送操作在通道未满时不会阻塞，只有当缓冲区满时才会阻塞
	// 接收操作在通道非空时不会阻塞，只有当缓冲区为空时才会阻塞
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retChan <- ret // channel 放数据，只要没有接受消息，会被阻塞
		fmt.Println("service exited")
	}()
	return retChan
}

// 串行
func TestSync(t *testing.T) {
	t.Log(service()) // 2s
	otherTask()      // 1s
}

// 并发
func TestAsync(t *testing.T) {
	retChan := asyncService()
	otherTask()
	fmt.Println(<-retChan) // 取数据
}
