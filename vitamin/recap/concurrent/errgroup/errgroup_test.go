package errgroup

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

// errgroup.Group 提供了类似 WaitGroup 的功能，但它额外处理了错误传播和基于 context 的取消机制
func TestErrGroup(t *testing.T) {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	tasks := []func(context.Context) error{
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				fmt.Println("task 1 cancelled")
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				fmt.Println("task 1 completed")
				return nil
			}
		},
		func(ctx context.Context) error {
			// 对一个值为 nil 的 channel 进行发送或接收操作，会永久阻塞，同样导致死锁
			//
			var ch <-chan struct{}

			select {
			case <-ctx.Done():
				fmt.Println("task 2 cancelled")
				return ctx.Err()
			case <-ch: // 对 nil channel 进行读或写操作的 case 永远不会被选中
				fmt.Println("task 2 completed")
				return nil
			case <-time.After(200 * time.Millisecond):
				return context.DeadlineExceeded
			}
		},
		func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				fmt.Println("task 3 cancelled")
				return ctx.Err()
			case <-time.After(300 * time.Millisecond):
				fmt.Println("task 3 completed")
				return nil
			}
		},
	}

	for _, task := range tasks {
		// 如果其中一个任务返回错误（或者传递给 errgroup.WithContext 的 context 被取消）
		// 那么与该 group 关联的 context 也会被取消
		// 可以通过这一机制通知相关任务尽快退出
		g.Go(func() error {
			return task(ctx)
		})
	}

	err := g.Wait()
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
