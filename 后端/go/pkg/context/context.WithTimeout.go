package main

import (
	"context"
	"time"
	"fmt"
)

// 此示例传递了一个带有超时的上下文，告诉阻塞函数在超时后应放弃工作。
func main() {
	shortDurant := 1 * time.Second
	neverReady := make(chan struct{})
	// 传递具有超时的上下文，以告知阻塞函数应在超时过后放弃其工作。
	ctx, cancel := context.WithTimeout(context.Background(), shortDurant)
	defer cancel()

	select {
	case <-neverReady:
		fmt.Println("read")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}