## package encoding

软件包编码定义了其他软件包共享的接口，这些接口可将数据转换为字节级和文本表示形式。检查这些接口的包包括 encoding/gob、encoding/json 和 encoding/xml。因此，只要实现一次接口，就能使一个类型在多种编码中发挥作用。实现这些接口的标准类型包括 time.Time 和 net.IP。这些接口是成对的，可以产生和消耗编码数据。

为现有类型添加编码/解码方法可能会构成破坏性更改，因为它们可在与使用不同版本库编写的程序通信时用于序列化。Go 项目维护的软件包的政策是，只有在没有现存合理的编码时，才允许添加编码函数。

## Index

### type BinaryMarshaler

```go
type BinaryMarshaler interface {
  MarshalBinary() (data []byte, err error)
}
```

BinaryMarshaler 是一个对象实现的接口，该对象可以将自己 marshal 成二进制形式。

MarshalBinary 将接收器编码为二进制形式，并返回结果。

### type BinaryUnmarshaler

```go
type BinaryUnmarshaler interface {
  UnmarshalBinary(data []byte) error
}
```

BinaryUnmarshaler 是一个对象实现的接口，该对象可以解除自身的二进制表示。

UnmarshalBinary 必须能够解码 MarshalBinary 生成的表单。如果 UnmarshalBinary 希望在返回后保留数据，则必须复制数据。

### type TextMarshaler

```go
type TextMarshaler interface {
  MarshalText() (text []byte, err error)
}
```

TextMarshaler 是一个对象实现的接口，该对象可以将自己 Marshal 成文本形式。

MarshalText 将接收器编码为 UTF-8 编码文本，并返回结果。

### type TextUnmarshaler

```go
type TextUnmarshaler interface {
  UnmarshalText(text []byte) error
}
```

TextUnmarshaler 是一个对象实现的接口，该对象可以解除自身的文本表示。

UnmarshalText 必须能够解码 MarshalText 生成的表单。如果 UnmarshalText 希望在返回后保留文本，则必须复制文本。