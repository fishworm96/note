package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个新列表，在其中输入一些数字。
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// 遍历列表并打印其内容。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		// 1
		// 2
		// 3
		// 4
	}
}