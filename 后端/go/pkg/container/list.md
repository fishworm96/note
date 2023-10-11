## package list

import "container/list"

包列表实现了双链路列表。

## Index

### type Element

```go
type Element struct {
  // 该元素存储的值。
  Value any
  // 包含已筛选或未导出字段
}
```

Element 是链表中的一个元素。

#### func (e *Element) Next() *Element

Next 返回下一个列表元素或 nil。

#### func (e *Element) Prev() *Element

Prev 返回前一个列表元素或 nil。

### type List

```go
type List struct {
  // 包含已经过滤或未导出字段
}
```

List 表示一个双链表。List 的零值是一个可随时使用的空 list。

#### func New() *List

New 返回一个初始化列表。

#### func (l *List) Back() *Element

Back 返回列表 l 的最后一个元素，如果列表为空，则返回 nil。

#### func (l *List) Front() *Element

Front 返回列表 l 的第一个元素，如果列表为空，则返回 nil。

#### func (l *List) Init() *List

初始化或清除列表 l。

#### func (l *List) InsertAfter(v any, mark *Element) *Element

InsertAfter 在 mark 后面插入一个值为 v 的新元素 e，并返回 e。如果 mark 不是 l 的元素，则不修改列表。标记不能为 nil。

#### func (l *List) InsertBefore(v any, mark *Element) *Element

InsertBefore 在紧靠 mark 之前插入一个值为 v 的新元素 e，并返回 e。如果 mark 不是 l 的元素，则不修改列表。标记不能为 nil。

#### func (l *List) Len() int

Len 返回列表 l 的元素个数。复杂度为 O(1)。

#### func (l *List) MoveAfter(e, mark *Element) 添加于1.2

MoveAfter 将元素 e 移动到 mark 之后的新位置。如果 e 或 mark 不是 l 的元素，或者 e == mark，则不会修改列表。元素和 mark 不能为空。

#### func (l *List) MoveBefore(e, mark *Element) 添加于1.2

MoveBefore 将元素 e 移动到新位置 mark 之前。如果 e 或 mark 不是 l 的元素，或者 e == mark，列表不会被修改。元素和 mark 不能为空。

#### func (l *List) MoveToBack(e *Element)

MoveToBack 将元素 e 移到 list l 的后面。如果 e 不是 list l 的元素，则不会修改 list。元素不能为 nil。

#### func (l *List) MoveToFront(e *Element)

MoveToFront 将元素 e 移到 list l 的前面。如果 e 不是 list l 的元素，则不修改 list。元素不能为 nil。

#### func (l *List) PushBack(v any) *Element

PushBack 在 list l 的后面插入一个值为 v 的新元素 e，并返回 e。

#### func (l *List) PushBackList(other *List)

PushBackList 在列表 l 的后面插入另一个列表的副本。它们不能为零。

#### func (l *List) PushFront(v any) *Element

PushFront 在 list l 的前面插入一个值为 v 的新元素 e，并返回 e。

#### func (l *List) PushFrontList(other *List)

PushFrontList 在列表 l 的前面插入另一个列表的副本。它们不能为空。

#### func (l *List) Remove(e *Element) any

如果 e 是 list l 的一个元素，Remove 会将 e 从 l 中移除，并返回元素值 e.Value。元素不能为 nil。