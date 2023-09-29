## package main

import "compress/zlib"

zlib 包实现了 RFC 1950 中规定的 zlib 格式压缩数据的读写。

该实现提供了过滤器，在读取时解压缩，在写入时压缩。例如，将压缩数据写入缓冲区：

## Index

### Constants

```go
const (
  NoCompression = flate.NoCompression
  BestSpeed = flate.BestSpeed
  BestCompression = falte.BestCompression
  DefaultCompression = falte.DefaultCompression
  HuffmanOnly = falte.HuffmanOnly
)
```

这些常量是从 flate 软件包中复制的，因此导入 "compress/zlib "的代码不必同时导入 "compress/flate"。

### Variables

```go
var (
  // 当读取校验和无效的 ZLIB 数据时，将返回 ErrChecksum。
  ErrChecksum = errors.New("zlib: invalid checksum")
  // 当读取字典无效的 ZLIB 数据时，将返回 ErrDictionary。
  ErrDictionary = errors.New("zlib: invalid dictionary")
  // ErrHeader 会在读取有无效标头的 ZLIB 数据时返回。
  ErrHeader = errors.New("zlib: invalid header")
)
```

### func NewReader(r io.Reader) (io.ReadeCloser, error)

NewReader 会创建一个新的 ReadCloser。如果 r 没有实现 io.ByteReader，解压程序从 r 中读取的数据可能会多于需要。

NewReader 返回的 ReadCloser 也实现了 Resetter。

### func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)

NewReaderDict 与 NewReader 类似，但使用的是预设字典。如果压缩数据未引用字典，NewReaderDict 会忽略该字典。如果压缩数据引用了不同的字典，NewReaderDict 会返回 ErrDictionary。

NewReaderDict 返回的 ReadCloser 也实现了重置。

### type Resetter

```go
// 添加于1.4
type Resetter interface {
  // 重置会丢弃任何缓冲数据，并重置重置器，就好像它是
  // 新初始化的阅读器。
  Reset(r io.Reader, dict []byte) error
}
```

重置器重置由 NewReader 或 NewReaderDict 返回的读取器，以切换到新的底层读取器。这样就可以重复使用一个 ReadCloser，而不是分配一个新的。

### type Writer

```go
type Writer struct {
  // 包含已筛选或未导出字段
}
```

写入器接收写入的数据，并将数据的压缩形式写入底层写入器（见 NewWriter）。

#### func NewWriter(w io.Writer) *Writer

NewWriter 会创建一个新的 Writer。写入返回 Writer 的内容会被压缩并写入 w。

写入完成后，调用者有责任调用关闭 Writer。写入的内容可能会被缓冲，在关闭之前不会被刷新。

#### func NewWriterLevel(w io.Writer, level int) (*Writer, error)

NewWriterLevel 与 NewWriter 类似，但指定的是压缩级别，而不是默认的 DefaultCompression。

压缩级别可以是 DefaultCompression、NoCompression、HuffmanOnly 或 BestSpeed 和 BestCompression 之间的任意整数值。如果压缩级别有效，返回的错误信息将为零。

#### func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)

NewWriterLevelDict 与 NewWriterLevel 类似，但指定了一个要压缩的字典。

字典可以为空。否则，在 Writer 关闭之前，其内容不应被修改。

#### func (z *Writer) Close() error

Close 关闭写入器，将任何未写入的数据刷新到底层 io.Writer，但不会关闭底层 io.Writer。

#### func (z *Writer) Flush() error

将 Writer 刷新到其底层 io.Writer。

#### func (z *Writer) Reset(w io.Writer) 添加于1.2

重置会清除写入器 z 的状态，使其等同于 NewWriterLevel 或 NewWriterLevelDict 的初始状态，但会写入 w。

#### func (z *Writer) Write(p []byte) (n int, err error)

Write 会将 p 的压缩形式写入底层的 io.Writer 中。在 Writer 关闭或显式刷新之前，压缩字节不一定会被刷新。