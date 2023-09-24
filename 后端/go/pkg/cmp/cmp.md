## package cmp
包 cmp 提供了与比较有序值相关的类型和函数。

## Index

### func Compare[T Ordered](x, y T) int

Compare 返回

```go
如果 x 小于 y -1，
如果 x 等于y 0，
如果 x 大于 y +1
```

对于浮点类型，NaN 表示小于任何非 NaN，NaN 表示等于 NaN，-0.0 表示等于 0.0。

### func Less[T Ordered](x, y T) bool

对于浮点类型，NaN 被认为小于任何非 NaN，而 -0.0 不小于（等于）0.0。

### type Ordered

```go
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
```

Ordered 是一个允许任何有序类型的约束：任何支持操作符 < <= >= > 的类型。如果 Go 的未来版本添加了新的有序类型，该约束将被修改以包含它们。

请注意，浮点类型可能包含 NaN（"非数字"）值。在将 NaN 值与任何其他值（无论是否为 NaN）进行比较时，操作符（如 == 或 <）将始终报错。有关比较 NaN 值的一致方法，请参阅比较函数。