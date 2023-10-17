package main

import (
	"context"
	"time"
	"fmt"
)

// 此示例传递了一个带有任意截止日期的上下文，告诉阻塞函数一旦到达截止日期就应放弃工作。
func main() {
	shortDuration := 1 * time.Second
	neverReady := make(chan struct{})
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 即使 ctx 将过期，最好将其调用在任何情况下都有取消功能。如果不这样做，可能会使上下文及其父级的寿命超过必要的时间。
	defer cancel()

	select {
	case <-neverReady:
		fmt.Println("read")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}