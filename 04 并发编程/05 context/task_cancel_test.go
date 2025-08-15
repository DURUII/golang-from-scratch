package _context

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

var ErrTooHot = errors.New("太热了，机器热化了")

// 关联任务的取消（级联退出）
// context 用于管理相关任务的上下文，包含了共享值的传递，超时，取消通知
func isCanceledByContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	// 父节点
	ctx, cancel := context.WithCancelCause(context.Background())
	for i := 0; i < 5; i++ {
		go func(ctx context.Context, i int) {
			for {
				if isCanceledByContext(ctx) {
					fmt.Println(i, "下班了，因为", context.Cause(ctx))
					break
				}
				fmt.Println(i, "在工作")
				time.Sleep(5 * time.Millisecond)
			}
		}(ctx, i)
	}
	// 广播机制
	cancel(ErrTooHot)
	time.Sleep(time.Second * 1)
	// 确保没有协程泄漏
	assert.Equal(t, runtime.NumGoroutine(), 2)
}
