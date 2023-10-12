package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建大小为 4 的环
	r := ring.New(4)

	// 打印出长度
	fmt.Println(r.Len())
}