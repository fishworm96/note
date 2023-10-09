## package heap

import "container/heap"

包 heap 为任何实现 heap.Interface 的类型提供堆操作。堆是一棵树，其属性是每个节点都是其子树中最小值的节点。

树中的最小元素是索引为 0 的根节点。

堆是实现优先队列的常用方法。要建立一个优先级队列，可以用（负）优先级作为 Less 方法的排序来实现 Heap 接口，这样 Push 就可以添加项目，而 Pop 则可以从队列中删除优先级最高的项目。示例中就包含这样的实现；example_pq_test.go 文件中有完整的源代码。

## Index

### func Fix(h Interface, i int) 添加于1.2

在索引 i 处的元素改变其值后，Fix 会重新建立堆排序。改变索引 i 处元素的值然后调用 Fix，相当于调用 Remove(h, i)，然后推送新值，但成本更低。复杂度为 O(log n)，其中 n = h.Len()。

### func Init(h Interface)

Init 建立了本软件包中其他例程所需的堆不变式。Init 相对于堆不变式是幂等的，只要堆不变式可能失效，就可以调用 Init。复杂度为 O(n)，其中 n = h.Len()。

### func Pop(h Interface) any

Pop 删除并返回堆中最小的元素（根据 Less）。复杂度为 O(log n)，其中 n = h.Len()。Pop 等同于 Remove(h,0)。

### func Push(h Interface, x any)

Push 将元素 x 推到堆上。复杂度为 O(log n)，其中 n = h.Len()。

### func Remove(h Interface, i int) any

Remove 删除并返回堆中索引为 i 的元素。复杂度为 O(log n)，其中 n = h.Len()。

### type Interface

```go
type Interface Interface {
  sort.Interface
  Push(x any) // 添加 x 为元素 Len()
  Pop() any // 移除并返回元素 Len() - 1。
}
```

接口类型描述了对使用本软件包例程的类型的要求。任何实现该接口的类型都可用作迷你堆，并具有以下不变式（在调用 Init 后或数据为空或排序时建立）：

```go
!h.Less(j, i) for 0 < i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
```

请注意，该接口中的 Push 和 Pop 是供包堆的实现调用的。要从堆中添加和删除内容，请使用 heap.Push 和 heap.Pop。