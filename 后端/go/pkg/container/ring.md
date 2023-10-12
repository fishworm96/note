## package ring

import "container/ring"

包环实现了对循环列表的操作。

## Index

### type Ring

```go
type Ring struct {
  Value any // 供客户端使用；本库未触及
  // 包含已经过滤或未导出字段
}
```

环是循环列表或环的一个元素。环没有开始或结束；指向任何环元素的指针都是对整个环的引用。空环用 nil Ring 指针表示。环的零值是一个具有 nil 值的单元素环。

#### func New(n int) *Ring

新建一个包含 n 个元素的环。

#### func(r *Ring) Do(f func(any))

Do 按正向顺序调用环中每个元素的函数 f。如果 f 改变了 *r，Do 的行为将是未定义的。

#### func (r *Ring) Len() int

Len 计算环 r 中的元素个数，执行时间与元素个数成正比。

#### func (r *Ring) Link(s *Ring) *Ring

Link 将环 r 与环 s 连接起来，这样 r.Next() 就会变成 s，并返回 r.Next() 的原始值。

如果 r 和 s 指向同一个环，则链接会从环中移除 r 和 s 之间的元素。移除的元素形成一个子环，结果是对该子环的引用（如果没有移除任何元素，结果仍然是 r.Next() 的原始值，而不是 nil）。

如果 r 和 s 指向不同的环，将它们连接起来会创建一个单一的环，并在 r 之后插入 s 的元素。

#### func (r *Ring) Move(n int) *Ring

Move 将 n % r.Len() 个元素向后移动（n < 0）或向前移动（n >= 0），并返回该环形元素。
#### func (r *Ring) Next() *Ring

Next 返回下一个环元素。r 不能为空

#### func (r *Ring) Prev() *Ring

Prev 返回上一个环元素。r 不能为空

#### func (r *Ring) Unlink(n int) *Ring

Unlink 从 r.Next() 开始，删除 r 环中的 n % r.Len() 个元素。如果 n % r.Len() == 0，则 r 保持不变。结果就是被移除的子环。r 不能为空