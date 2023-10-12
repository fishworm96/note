package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建大小为 5 的环
	r := ring.New(5)

	// 获得环长度
	n := r.Len()

	// 用整数初始化环的值
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// 遍历环并打印其内容
	for j := 0; j < n; j++ {
		// Prev 返回上一个环元素。
		r = r.Prev()
		fmt.Println(r.Value)
		// 4
		// 3
		// 2
		// 1
		// 0
	}
}