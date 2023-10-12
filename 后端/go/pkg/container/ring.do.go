package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个大小为 5 的环
	r := ring.New(5)

	// 获得环长度
	n := r.Len()

	// 用整数初始化环的值
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// 遍历环并打印其内容
	r.Do(func(p any) {
		fmt.Println(p.(int))
	// 0
	// 1
	// 2
	// 3
	// 4
	})
}