package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建大小为 5 的环
	r := ring.New(5)

	// 获得环的长度
	n := r.Len()

	// 用一些整数值对环进行初始化
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// 将指针向前移动三步
	r = r.Move(3)

	// 遍历组合环并打印其内容
	r.Do(func(p any) {
		fmt.Println(p.(int))
		// 3
		// 4
		// 0
		// 1
		// 2
	})
}