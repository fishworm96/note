// 本例演示了使用堆接口构建优先级队列。
package main

import (
	"container/heap"
	"fmt"
)

// Item 是我们在优先级队列中管理的东西。
type Item struct {
	value string // item 的值；任意。
	priority int // 队列中 Item 的优先级。
	// 更新需要索引，索引由 heap.Interface 方法维护。
	index int // 堆中 Item 的索引。
}

// PriorityQueue 实现 heap.Interface 并持有 Items。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 给我们最高而不是最低的优先级，所以我们在这里使用大于。
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n - 1]
	// 避免内存泄漏
	old[n-1] = nil
	// 为安全起见
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update 会修改队列中某个项目的优先级和值。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// 此示例创建了一个包含若干项目的 PriorityQueue，添加并操作了一个项目、然后按优先级顺序删除项目。
func main() {
	// 一些项目及其优先次序。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建优先队列，将项目放入其中，并建立优先队列（堆）不变式。
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value: value,
			priority: priority,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新 Item，然后修改其优先级
	item :=&Item{
		value: "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// 将物品取出；它们按优先级递减的顺序送达。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value) // 05:orange 04:pear 03:banana 02:apple
	}
}