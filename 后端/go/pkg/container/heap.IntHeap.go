// 此示例向 IntHeap 中插入多个输入值，检查最小值，然后按优先级顺序将其移除。
// 本示例演示了使用堆接口构建的整数堆。
package main

import (
	"container/heap"
	"fmt"
)

// IntHeap 是一个 int 的最小堆。
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i , j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	// Push 和 Pop 使用指针接收器，因为它们会修改切片的长度、
	// 不仅仅是其内容。
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

// 本例将多个 int 插入 IntHeap，检查最小值并按优先顺序删除它们。
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0]) // minimum: 1
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 1 2 3 5
	}
}