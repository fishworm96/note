package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建大小为 2 的两个环 r 和 s
	r := ring.New(2)
	s := ring.New(2)

	// 获取到环的长度
	lr := r.Len()
	ls := r.Len()

	// 用 0 初始化 r
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// 用 1 初始化 s
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	// 链接环 r 和环 s
	rs := r.Link(s)

	// 遍历组合环并打印其内容
	rs.Do(func(p any) {
		fmt.Println(p.(int))
		// 0
		// 0
		// 1
		// 1
	})
}