package main

import (
	"context"
	"fmt"
)

// 本例演示了如何使用可取消上下文来防止 goroutine 泄漏。在示例函数结束时，由 gen 启动的 goroutine 将返回，不会发生泄漏。
func main() {
	// Gen 在单独的 goroutines 中生成整数，并将它们发送到返回的通道。gen 的调用者在使用完生成的整数后需要取消上下文，以免泄漏由 gen 启动的内部 goroutine。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回时不会泄漏 goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 消耗完整数后取消

	for n := range gen(ctx) {
		fmt.Println(n)
		// 1
		// 2
		// 3
		// 4
		// 5
		if n == 5 {
			break
		}
	}
}