package main

import (
	"context"
	"fmt"
)

// 该示例演示了如何将一个值传递给上下文，以及如何在该值存在的情况下获取它。
func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k) // // found value: Go
	f(ctx, favContextKey("color")) // key not found: color
}