## package bufio

包 bufio 实现缓冲 I/O。它包装了一个 io.Reader 或 io.Writer 对象，创建了另一个对象（Reader或Writer），该对象也实现了该接口，但为文本I/O提供了缓冲和一些帮助

## Index

[Constants](#constants)
[Variables](#variables)
    [func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)](#scan-bytes)
    [func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)](#scan-lines)
    [func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)](#scan-runes)
    [func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)](#scan-words)
[type ReadWriter](#read-writer)
    [func NewReadWriter(r *Reader, w *Writer) *ReadWriter](#new-read-writer)
[type Reader](#reader)
    [func NewReader(rd io.Reader) *Reader](#new-reader)
    [func New ReaderSize(rd io.Reader, size int) *Reader](#new-reader-size)
    [func (b* Reader) Buffered() int](#reader-buffered)
    [func (b* Reader) Discard(n int) (discarded int, err error)](#discard)
    [func (b* Reader) Peek(n int) ([]byte, error)](#peek)
    [func (b* Reader) Read(p []byte) (n int, err error)](#read)
    [func (b* Reader) ReadByte() (byte, error)](#read-byte)
    [func (b* Reader) ReadBytes(delim byte) ([]byte, error)](#read-bytes)
    [func (b* Reader) ReadLine() (line []byte, isPrefix bool, err error)](#read-line)
    [func (b* Reader) ReadRune() （r rune, size int, err error](#read-rune)
    [func (b* Reader) ReadSlice(delim byte) (line []byte, err error)](#read-slice)
    [func (b* Reader) ReadString(delim byte) (string, error)](#read-string)
    [func (b* Reader) Reset(r io.Reader)](#reader-reset)
    [func (b* Reader) Size() int](#size)
    [func (b* Reader) UnreadByte() error](#unread-byte)
    [func (b* Reader) UnreadRune() error](#unread-rune)
    [func (b* Reader) WriteTo(w io.Writer) (n int64, err error)](#write-to)
[type Scanner](#scanner)
    [func NewScanner(r io.Reader) *Scanner](#new-scanner)
    [func (s *Scanner) Buffer(buf []byte, max int)](#buffer)
    [func (s *Scanner) Bytes() []byte](#bytes)
    [func (s *Scanner) Err() error](#err)
    [func (s *Scanner) Scan() bool](#scan)
    [func (s *Scanner) Split(split SplitFunc)](#split)
    [func (s *Scanner) Text() string](#text)
[type SplitFunc](#split-func)
[type Writer](#writer)
    [func NewWriter(w io.Writer) *Writer](#new-writer)
    [func NewWriterSize(w io.Writer, size int) *Writer](#new-writer-size)
    [func (b *Writer) Available() int](#available)
    [func (b *Writer) AvailableBuffer() []byte](#available-buffer)
    [func (b *Writer) Buffered() int](#writer-buffered)
    [func (b *Writer) Flush() error](#flush)
    [func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)](#read-from)
    [func (b *Writer) Reset(w io.Writer)](#writer-reset)
    [func (b *Writer) Size() int](#size)
    [func (b *Writer) Write(p []byte) (nn int, err error)](#write)
    [func (b *Writer) WriteByte(c byte) error](#write-byte)
    [func (b *Writer) WriteRune(r rune) (size int, err error)](#write-rune)
    [func (b *Writer) WriteString(s string) (int, error)](#write-string)

## **<a id="constants">Constants</a>**

```go
const (
  // MaxScanTokenSize 是用于缓冲令牌的最大大小
  // 除非用户使用 Scanner.Buffer 提供明确的缓冲区。
  // 作为缓冲区，实际的最大标记大小可能更小
  // 可能需要包括，例如，换行符。
  MaxScanTokenSize = 64 * 1024
)
```

## **<a id="variables">Variables</a>**

```go
var (
  ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
  ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
```

```go
var (
	ErrTooLong         = errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
	ErrBadReadCount    = errors.New("bufio.Scanner: Read returned impossible count")
)
```

扫描程序返回的错误。

```go
var ErrFinalToken = errors.New("final token")
```

ErrFinalToken 是一个特殊的 sentinel 错误值。它旨在由 Split 函数返回，以指示发送错误的令牌是最后一个令牌，扫描应在此令牌之后停止。扫描收到 ErrFinalToken 后，扫描停止，没有错误。该值在早期停止处理或在需要交付最终空令牌时非常有用。可以使用自定义错误值实现相同的行为，但在这里提供一个更整洁。有关此值的用法，请参见 emptyFinalToken 示例。

## Functions

### **<a id="scan-bytes">func ScanBytes</a>**

```go
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

ScanBytes 是 Scanner 的拆分函数，它将每个字节作为令牌返回。

### **<a id="scan-lines">func ScanLines</a>**

```go
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
```

ScanLines 是 Scanner 的一个拆分函数，它返回每一行文本，去掉任何尾随的行尾标记。返回的行可能为空。行尾标记是一个可选的回车符，后跟一个强制性的换行符。在正则表达式表示法中，它是`\r？\ n`。输入的最后一个非空行将被返回，即使它没有换行符。

### **<a id="scan-runes">func ScanRunes</a>**

```go
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

ScanRunes 是一个 Scanner 的拆分函数，它返回每个 UTF-8 编码的符文作为令牌。返回的符文序列等同于作为字符串的输入范围循环，这意味着错误的 UTF-8 编码转换为 U+FFFD =“\xef\xbf\xbd”。由于扫描界面，这使得客户无法区分正确编码的替换符文和编码错误。

### **<a id="scan-words">func ScanWords</a>**

```go
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
```

ScanWords 是一个用于 Scanner 的拆分函数，它返回文本的每个空格分隔的单词，删除周围的空格。它永远不会返回空字符串。空间的定义由 unicode. IsSpace 设置。

## Types

### **<a id="read-writer">type ReadWriter</a>**

```go
type ReadWriter struct {
  *Reader
  *Writer
}
```

ReadWriter 存储指向 Reader 和 Writer 的指针。它实现了 io.ReadWriter。

#### **<a id="new-read-writer">func NewReadWriter</a>**

```go
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
```

NewReadWriter 分配一个新的 ReadWriter，该 ReadWriter 分派给 r 和 w。

### **<a id="reader">type Reader</a>**

```go
type Reader struct {
  // 包含经过筛选或未导出的字段
}
```

Reader 为 io.Reader 对象实现缓冲。

#### **<a id="new-reader">func NewReader</a>**

```go
func NewReader(rd io.Reader) *Reader
```

NewReader 返回一个缓冲区具有默认大小的新 Reader。

#### **<a id="new-reader-size">func NewReaderSize</a>**

```go
func NewReaderSize(rd io.Reader, size int) *Reader
```

NewReaderSize 返回一个新的 Reader，其缓冲区至少具有指定的大小。如果参数 io.Reader 已经是一个足够大的 Reader，它将返回基础 Reader。

#### **<a id="reader-buffered">func (*Reader) Buffered</a>**

```go
func (b *Reader) Buffered() int
```

Buffered 返回可从当前缓冲区读取的字节数。

#### **<a id="discard">func (*Reader) Discard</a>**

```go
func (b *Reader) Discard(n int) (discarded int, err error)
```

Discard 跳过接下来的 n 个字节，返回丢弃的字节数。

如果 Discard 跳过的字节少于 n 个，它也会返回错误。如果0 <= n <= B.Buffered（），则 Discard 保证在不从底层 io.Reader 阅读的情况下成功。

#### **<a id="Peek">func (*Reader) Peek</a>**

```go
func (b *Reader) Peek(n int) ([]byte, error)
```

Peek 返回接下来的 n 个字节，而不推进读取器。字节在下一次读取调用时不再有效。如果 Peek 返回少于 n 个字节，它也会返回一个错误，解释为什么读短。如果 n 大于 b 的缓冲区大小，则错误为 ErrBufferFull。

调用 Peek 可防止 UnreadByte 或 UnreadRune 调用成功，直到下一次读取操作。

#### **<a id="read">func (*Reader) Read</a>**

```go
func (b *Reader) Read(p []byte) (n int, err error)
```

Read 将数据读入p。它返回读取到p中的字节数。字节取自底层读取器上的至多一个读取，因此n可以小于 len（p）。要准确读取 len（p）字节，请使用 io.ReadFull（b，p）。如果底层 Reader 可以返回一个带有 io.EOF 的非零计数，那么这个 Read 方法也可以这样做;请参阅 i.Reader 文档。

#### **<a id="read-byte">func (*Reader( ReadByte</a>**

```go
func (b *Reader) ReadByte() (byte, error)
```

ReadByte 读取并返回单个字节。如果没有字节可用，则返回错误。

#### **<a id="read-bytes">func (*Reader) ReadBytes</a>**

```go
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
```

ReadBytes 一直读取到输入中第一次出现 delim，返回一个包含数据的切片，直到并包括分隔符。如果 ReadBytes 在找到分隔符之前遇到错误，它将返回在错误之前读取的数据和错误本身（通常是io.EOF）。ReadBytes 返回 err！= nil 当且仅当返回的数据不以 delim 结尾。对于简单的使用，扫描仪可能更方便。

#### **<a id="read-line">func (*Reader) ReadLine</a>**

```go
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
```

ReadLine 是一个低级的行读取原语。大多数调用方应使用 ReadBytes（'\n'）或ReadString（'\n'）或使用扫描程序。

ReadLine 尝试返回单行，不包括行尾字节。如果该行对于缓冲区来说太长，则设置 isPrefix 并返回该行的开头。线路的其余部分将从将来的调用中返回。返回行的最后一个片段时，isPrefix 将为 false。返回的缓冲区仅在下一次调用 ReadLine 之前有效。ReadLine 要么返回一个非 nil 的行，要么返回一个错误，永远不会同时返回。

从 ReadLine 返回的文本不包括行尾（“\r\n”或“\n”）。如果输入结束而没有最终行结束，则不会给出任何指示或错误。在 ReadLine 之后调用 UnreadByte 将始终取消读取最后一个字节（可能是属于行结束的字符），即使该字节不是ReadLine返回的行的一部分。

#### **<a id="read-rune">func (*Reader) ReadRune</a>**

```go
func (b *Reader) ReadRune() (r rune, size int, err error)
```

ReadRune 读取一个 UTF-8 编码的 Unicode 字符，并返回该字符及其字节大小。如果编码的符文无效，则消耗一个字节并返回大小为1的 unicode.ReplacementChar（U+FFFD）。

#### **<a id="read-slice">func (*Reader) ReadSlice</a>**

```go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

ReadSlice 一直读到输入中第一次出现 delim，返回一个指向缓冲区中字节的切片。字节在下一次读取时停止有效。如果 ReadSlice 在找到分隔符之前遇到错误，它将返回缓冲区中的所有数据和错误本身（通常是io.EOF）。如果缓冲区没有 delim 填充，ReadSlice 将失败并显示错误 ErrBufferFull。因为从 ReadSlice 返回的数据将被下一个 I/O 操作覆盖，所以大多数客户端应该使用 ReadBytes 或 ReadString。ReadSlice 返回错误！= nil当且仅当行不以 delim 结束。

#### **<a id="read-string">func (*Reader) ReadString</a>**

```go
func (b *Reader) ReadString(delim byte) (string, error)
```

ReadString 一直读取到delim在输入中第一次出现为止，返回一个字符串，其中包含的数据直到分隔符为止（包括分隔符）。如果 ReadString 在找到分隔符之前遇到错误，则返回在错误之前读取的数据和错误本身（通常为io.EOF）。ReadString 返回 err！= nil 当且仅当返回的数据不以 delim 结尾。对于简单的使用，扫描仪可能更方便。

#### **<a id="writer-reset">func (*Reader) Reset</a>**

```go
func (b *Reader) Reset(r io.Reader)
```

Reset 丢弃所有未刷新的缓冲数据，清除所有错误，并重置 b 以将其输出写入 w。对 Writer 的零值调用 Reset 会将内部缓冲区初始化为默认大小。调用 w.Reset（w）（即，将Writer重置为自身）不会执行任何操作。

#### **<a id="size">func (*Writer) Size</a>**

```go
func (b *Writer) Size() int
```

Size 返回基础缓冲区的大小（以字节为单位）。

#### **<a id="write">func (*Writer) Write</a>**

```go
func (b *Writer) Write(p []byte) (nn int, err error)
```

Write 将 p 的内容写入缓冲区。它返回写入的字节数。如果 nn < len（p），它还返回一个错误，解释为什么写是短的。

#### **<a id="write-byte">func (*Writer) WriteByte</a>**

```go
func (b *Writer) WriteByte(c byte) error
```

WriteByte 写入单个字节。

#### **<a id="write-rune">func (*Write) WriteRune</a>**

```go
func (b *Writer) WriterRune(r rune) (size int, err error)
```

WriteRune 写入单个 Unicode 代码点，返回写入的字节数和任何错误。

#### **<a id="write-string">func (*Writer) WriteString ¶</a>**

```go
func (b *Writer) WriteString(s string) (int, error)
```

WriteString 写入字符串。它返回写入的字节数。如果计数小于 len（s），它也会返回一个错误，解释为什么写操作很短。