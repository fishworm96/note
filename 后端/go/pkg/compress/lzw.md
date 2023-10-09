## package lzw

import "compress/lzw"

软件包 lzw 实现了 Lempel-Ziv-Welch 压缩数据格式，该格式在 T. A. Welch 的《高性能数据压缩技术》（"A Technique for High-Performance Data Compression"）一文中有所描述，《计算机》（Computer），17（6）（1984 年 6 月），第 8-19 页。

特别是，它实现了 GIF 和 PDF 文件格式所使用的 LZW，这意味着可变宽度代码最高可达 12 位，并且前两个非字面代码是清晰代码和 EOF 代码。

TIFF 文件格式使用类似但不兼容的 LZW 算法版本。请参阅 golang.org/x/image/tiff/lzw 软件包以获取实现。

## Index

### func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser

NewReader 会创建一个新的 io.ReadCloser。如果 r 没有同时实现 io.ByteReader，解压程序从 r 中读取的数据可能会超过需要。用于字面代码的位数 litWidth 必须在 [2,8] 范围内，通常为 8。 它必须等于压缩时使用的 litWidth。

保证返回的 io.ReadCloser 的底层类型为 *Reader。

### func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser

NewWriter 会创建一个新的 io.WriteCloser。向返回的 io.WriteCloser 写入的内容会被压缩并写入 w。写入完成后，调用者有责任调用关闭 WriteCloser。用于字面代码的位数 litWidth 必须在 [2,8] 范围内，通常为 8。 输入字节必须小于 1<<litWidth。

保证返回的 io.WriteCloser 的底层类型是 *Writer。

### type Order

```go
type Order int
```

Order 指定 LZW 数据流中的比特排序。

```go
const (
  // LSB 指 GIF 文件格式中使用的最小有效位（Least Significant Bits first）。
  LSB Order = iota
  // MSB 表示最重要位在前，如 TIFF 和 PDF 中使用的那样
  // 文件格式。
  MSB
)
```

### type Reader

```go
// 添加于1.17
type Reader struct {
  // 包含已筛选或未导出字段
}
```

Reader 是一个 io.Reader，可用于读取 LZW 格式的压缩数据。

#### func (r *Reader) Close() error 添加于1.17

Close 关闭读取器，并对今后的读取操作返回错误信息。它不会关闭底层的 io.Reader.Viewer 阅读器。

#### func (r *Reader) Read(b []byte) (int, error) 添加于1.17

Read 实现了 io.Reader，从底层的 Reader 读取未压缩的字节。

#### func (r *Reader) Reset(src io.Reader, order Order, litWidth int) 添加于1.17

重置会清除阅读器的状态，使其可以作为新的阅读器再次使用。

### type Writer 添加于1.17

```go
type Writer struct {
  // 包含已筛选或未导出字段
}
```

Writer 是一个 LZW 压缩器。它将压缩后的数据写入底层写入器（见 NewWriter）。

#### func (w *Writer) Close() 添加于1.17

Close 关闭写入器，清除所有待处理的输出。它不会关闭 w 的底层写入器。

#### func (w *Writer) Reset(dst io.Writer, order Order, litWidth int) 添加于1.17

重置会清除写入器的状态，使其可以作为新的写入器再次使用。

#### func (w *Writer) Write(p []byte) (n int, err error) 添加于1.17

Write 将 p 的压缩表示写入 w 的底层写入器。