package all

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("Result from %d", id)
}

// CSP 模式实现所有任务完成
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	// 等所有结果返回
	var sb strings.Builder
	for j := 0; j < numOfRunner; j++ {
		// unordered aggregation
		sb.WriteString(<-ch + "\n")
	}
	return sb.String()
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(2 * time.Second)
	t.Log("After: ", runtime.NumGoroutine())
}
