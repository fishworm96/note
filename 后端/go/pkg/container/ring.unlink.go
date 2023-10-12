package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建大小为 5 的环
	r := ring.New(6)

	// 获得环长度
	n := r.Len()

	// 用整数初始化环的值
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// 从 r.Next() 开始，解除 r 中三个元素的链接
	r.Unlink(3)

	// 遍历环并打印其内容
	r.Do(func(p any) {
		fmt.Println(p.(int))
		// 0
		// 4
		// 5
	})
}