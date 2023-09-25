## package flate

import "compress/flate"

软件 flate 实现了 RFC 1951 中描述的 DEFLATE 压缩数据格式。gzip 和 zlib 包实现了对基于 DEFLATE 的文件格式的访问。

## Index

### Constants

```go
const (
	NoCompression      = 0
	BestSpeed          = 1
	BestCompression    = 9
	DefaultCompression = -1

	// HuffmanOnly 禁用 Lempel-Ziv 匹配搜索，只执行 Huffman
	// 熵编码。该模式适用于压缩已使用
	// 已使用 LZ 类型算法（如 Snappy 或 LZ4）压缩的数据时，该模式非常有用。
	// 缺乏熵编码器。当
	// 输入流中某些字节的出现频率高于其他字节。
	//
	// 请注意，HuffmanOnly 产生的压缩输出符合
	// 符合 RFC 1951 标准。也就是说，任何有效的 DEFLATE 解压器都将
	// 继续解压此输出。
	HuffmanOnly = -2
)
```

### func NewReader(r io.Reader) io.ReadCloser

如果 r 没有同时实现 io.ByteReader，解压程序从 r 中读取的数据可能会多于需要读取的数据。最后一个数据块后的任何尾随数据都将被忽略。

NewReader 返回的 ReadCloser 也实现了 Resetter。

### func NewReaderDict(r io.Reader, dict []byte) io.ReaderCloser

NewReaderDict 与 NewReader 类似，但会使用预设字典初始化阅读器。返回的阅读器的行为就好像未压缩的数据流是从已读取的给定字典开始的一样。NewReaderDict 通常用于读取由 NewWriterDict 压缩的数据。

NewReader 返回的 ReadCloser 也实现了重置。

### type CorruptInputError

```go
type CorruptInputError int64
```

CorruptInputError 报告在给定偏移量处出现了损坏的输入。

#### func (e CorruptInputError) Error() string

### type InternalError

```go
type InternalError string
```

内部错误（InternalError）会报告 flate 代码本身的错误。

#### func (e InternalError) Error() string

### type ReadError 废弃

#### func (e *ReadError) Error() string

### type Reader

```go
type Reader interface {
  io.Reader
  io.ByteReader
}
```

NewReader 实际需要的读取接口。如果传递进来的 io.Reader 没有 ReadByte，NewReader 将引入自己的缓冲。

### type Resetter

```go
type Resetter interface {
  // 重置会丢弃任何缓冲数据，并重置重置器，就好像它是用给定的读取器从新初始化的一样。
	// 使用给定的读取器重新初始化。
  Reset(r io.Reader, dict []byte) error
}
```

### type WriteError 废弃

#### func (e *WriteError) Error() string

### Writer

```go
type Writer struct {
  // 包含已过滤或未导出的字段
}
```

写入器接收写入的数据，并将数据的压缩形式写入底层写入器（见 NewWriter）。

#### func NewWriter(w io.Writer, level int) (*Writer, error)

NewWriter 返回一个新的 Writer，以给定的级别压缩数据。按照 zlib，级别范围从 1（BestSpeed）到 9（BestCompression）；级别越高，运行速度越慢，但压缩效果越好。0 级（NoCompression）不尝试任何压缩；只添加必要的 DEFLATE 框架。第 1 级（DefaultCompression）使用默认压缩级别。第 -2 级（HuffmanOnly）只使用哈夫曼压缩，对所有类型的输入都能提供非常快的压缩速度，但会牺牲相当高的压缩效率。

如果级别在 [-2, 9] 范围内，则返回的错误信息为零。否则，返回的错误值将为非零。

#### func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)

NewWriterDict 与 NewWriter 类似，但使用预设字典初始化新 Writer。返回的 Writer 就像字典已被写入一样，不会产生任何压缩输出。写入 w 的压缩数据只能由使用相同字典初始化的阅读器解压缩。

#### func (w *Writer) Close() error

  关闭冲洗并关闭写入器。
#### func (w *Writer) Flush() error

刷新（Flush）会将任何待处理数据刷新到底层写入器。它主要用于压缩网络协议，以确保远程读取器有足够的数据重建数据包。在数据写入之前，Flush 不会返回。在没有待处理数据时调用 Flush 仍会导致写入器发出至少 4 字节的同步标记。如果底层写入器返回错误，Flush 也会返回该错误。

在 zlib 库的术语中，Flush 等同于 Z_SYNC_FLUSH。

#### func (w *Writer) Reset(dst io.Writer)

重置会丢弃写入器的状态，使其等同于使用 dst 和 w 的级别和字典调用 NewWriter 或 NewWriterDict 的结果。

#### func (w *Writer) Write(data []byte) (n int, err error)

Write 向 w 写入数据，w 最终会将压缩后的数据写入底层写入器。