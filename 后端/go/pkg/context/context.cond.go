package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 此示例使用 AfterFunc 定义了一个在 sync.Cond 上等待的函数，当上下文取消时停止等待。
func main() {
	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
		stopf := context.AfterFunc(ctx, func() {
			// 我们需要在这里获取 cond.L，以确保广播
			// 下面的操作不会在调用 Wait 之前发生，这将导致
			// 导致信号丢失（和死锁）。
			cond.L.Lock()
			defer cond.L.Unlock()
	
			// 如果多个程序同时等待 cond,我们需要确保我们准确地唤醒了这个。
			// 这意味着我们需要广播所有的程序,这将唤醒他们。
	
			// 如果有 N 个并发调用来等待 OnCond，则每个 goroutines 都会错误地唤醒 O（N） 个尚未准备好的其他 goroutine，因此这将导致整体 CPU 开销为 O（N²）。
			cond.Broadcast()
		})
		defer stopf()

		// 由于唤醒使用的是广播而不是信号，因此由于其他一些 goroutine 的上下文正在完成，因此对 Wait 的调用可能会取消阻止，因此为了确保 ctx 确实已完成，我们需要在循环中检查它。
		for !conditionMet() {
			cond.Wait()
			if ctx.Err() != nil {
				return ctx.Err()
			}
		}

		return nil
	}

	cond := sync.NewCond(new(sync.Mutex))

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
			defer cancel()

			cond.L.Lock()
			defer cond.L.Unlock()

			err := waitOnCond(ctx, cond, func() bool { return false })
			fmt.Println(err)
			// context deadline exceeded
			// context deadline exceeded
			// context deadline exceeded
			// context deadline exceeded
		}()
	}
	wg.Wait()
}