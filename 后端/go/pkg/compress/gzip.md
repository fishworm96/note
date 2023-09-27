## package gzip

import "compress/gzip"
包 gzip 实现了 RFC 1952 中规定的 gzip 格式压缩文件的读写。

## Index

### Constants

```go
const (
  NoCompression = flate.NoCompression
  BestSpeed = flate.BestSpeed
  BestCompression = flate.BestCompression
  DefaultCompression = flate.DefaultCompression
  HuffmanOnly = falte.HuffmanOnly
)
```

这些常量是从 flate 软件包中复制的，因此导入 "compress/gzip "的代码不必同时导入 "compress/flate"。

### Variables

```go
var (
  // 当读取校验和无效的 GZIP 数据时，将返回 ErrChecksum。
  ErrChecksum = errors.New("gzip: invalid checksum")
  // ErrHeader 会在读取标头无效的 GZIP 数据时返回。
  ErrHeader = errors.New("gzip: invalid header")
)
```

### type Header

```go
type Header struct {
  Comment string // comment
  Extra []byte // "extra data"
  ModTime time.Time // modification time
  Name string // file name
  OS byte // operating system type
}
```

gzip 文件会存储一个头文件，提供压缩文件的元数据。该标头会作为 Writer 和 Reader 结构体的字段显示出来。

由于 GZIP 文件格式的限制，字符串必须采用 UTF-8 编码，且只能包含 U+0001 至 U+00FF 的 Unicode 代码点。

### type Reader

```go
type Reader struct {
  Header // 在 NewReader 或 Reader.Reset 之后有效
  // 包含已过滤或未导出的字段
}
```

阅读器是一个 io.Reader，可以从 gzip 格式的压缩文件中读取未压缩的数据。

一般来说，一个 gzip 文件可以是多个 gzip 文件的连接，每个文件都有自己的头。从读取器读取的是每个文件的未压缩数据。阅读器字段中只记录第一个头。

Gzip 文件存储了未压缩数据的长度和校验和。如果未压缩数据的长度或校验和不符合预期，读取器在读取到未压缩数据的末尾时会返回 ErrChecksum。客户端应将 "读取 "返回的数据视为暂定数据，直到收到标记数据结束的 io.EOF。

#### func NewReader(r io.Reader) (*Reader, error)

NewReader 会创建一个新的阅读器，读取给定的阅读器。如果 r 没有同时实现 io.ByteReader，解压程序从 r 读取的数据可能会超过需要。

调用者有责任在完成后调用关闭读取器。

在返回的读取器中，Reader.Header 字段将是有效的。

#### func (z *Reader) Close() error

Close 关闭阅读器。它不会关闭底层的 io.Reader.为了验证 GZIP 校验和，读取器必须完全耗尽，直到 io.EOF.Close 结束。

#### func (z *Reader) Multistream(ok bool) 添加于1.4

多流控制阅读器是否支持多流文件。

如果启用（默认值），阅读器会将输入视为一连串单独压缩的数据流，每个数据流都有自己的头和尾，并以 EOF 结束。这样做的效果是，一连串 gzip 文件的串联被视为等同于序列串联的 gzip。这是 gzip 阅读器的标准行为。

调用 Multistream(false) 会禁用这种行为；禁用这种行为在读取区分单个 gzip 数据流或将 gzip 数据流与其他数据流混合的文件格式时非常有用。在此模式下，当阅读器到达数据流的末端时，Read 会返回 io.EOF。底层读取器必须实现 io.ByteReader，这样才能在 gzip 数据流之后留下位置。要启动下一个数据流，请调用 z.Reset(r)，然后再调用 z.Multistream(false)。如果没有下一个流，z.Reset(r) 将返回 io.EOF。

#### func (z *Reader) Read(p []byte) 9n int, err error

Read 实现了 io.Reader，从底层的 Reader 读取未压缩的字节。

#### func (z *Reader) Reset(r io.Reader) error 添加于1.3

func (z *Reader) Reset(r io.Reader) error
重置会丢弃阅读器 z 的状态，使其等同于从 NewReader 读取的原始状态，但会从 r 读取。这就允许重复使用一个阅读器，而不是分配一个新的。

### type Writer

```go
type Writer struct {
  Header // 在首次调用 "写入"、"刷新 "或 "关闭 "时写入
  // 包含已筛选或未导出字段
}
```

Writer 是一个 io.WriteCloser。写入 Writer 的内容会被压缩并写入 w。

#### func NewWriter(w io.Writer) *Writer

NewWriter 返回一个新的写入器。写入返回写入器的内容会被压缩并写入 w。

写入完成后，调用者有责任调用 Writer 上的 Close。写入的内容可能会被缓冲，在关闭之前不会被刷新。

调用者如果希望设置 Writer.Header 中的字段，必须在第一次调用 Write、Flush 或 Close 之前完成。

#### func NewWriterLevel(w io.Writer, level int) (*Writer, error)

NewWriterLevel 与 NewWriter 类似，但指定的是压缩级别，而不是默认的 DefaultCompression。

压缩级别可以是 DefaultCompression、NoCompression、HuffmanOnly 或 BestSpeed 和 BestCompression 之间的任意整数值。如果压缩级别有效，返回的错误信息将为零。

#### func (z *Writer) Close() error

关闭写入器时，会将未写入的数据刷新到底层的 io.Writer 中，并写入 GZIP 页脚。它不会关闭底层的 io.Writer。

#### func (z *Writer) Flush() error 添加于1.1

将任何待处理的压缩数据刷新到底层写入器。

它主要用于压缩网络协议，以确保远程读取器有足够的数据重建数据包。在数据写入之前，Flush 不会返回。如果底层写入器返回错误，Flush 也会返回该错误。

在 zlib 库的术语中，Flush 等同于 Z_SYNC_FLUSH。

#### func (z *Writer) Reset(w io.Writer) 添加于1.2

重置会丢弃 Writer z 的状态，使其等同于 NewWriter 或 NewWriterLevel 的原始状态，但会写入 w。这样就可以重复使用写入器，而不是分配一个新的写入器。

#### func (z *Writer) Write(p []byte) (int, error)

Write 将 p 的压缩形式写入底层的 io.Writer 中。在 Writer 关闭之前，压缩字节不一定会被刷新。