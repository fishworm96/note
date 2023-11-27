## package csv

软件包 csv 读写逗号分隔值（CSV）文件。CSV 文件有很多种，本软件包支持 RFC 4180 中描述的格式。

csv 文件包含零条或多条记录，每条记录包含一个或多个字段。每条记录之间用换行符隔开。最后一条记录后面可以选择加换行符。

```go
field1,field2,field3
```

空白被视为字段的一部分。

换行符前的回车符会被静默删除。

空行将被忽略。只有空白字符（不包括换行结束符）的行不视为空行。

以引号字符""开始和结束的字段称为引号字段。开始和结束的引号不是字段的一部分。

资料来源:

```go
normal string,"quoted-field"
```

领域的成果

```go
{`normal string`, `quoted-field`}
```

在引号字段内，一个引号字符后跟的第二个引号字符被视为单引号。

```go
"the ""word"" is true","a ""quoted-field"""
```

结果是

```go
{`the "word" is true`, `a "quoted-field"`}
```

引号字段中可以包含换行符和逗号

```go
"Multi-line
field","comma is ,"
```

结果是

```go
{`Multi-line
field`, `comma is ,`}
```

## Index

### Variables

```go
var (
  ErrBareQuote = errors.New("bare \" in non-quoted-field")
	ErrQuote      = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")

	// 已废弃：ErrTrailingComma 已不再使用。
	ErrTrailingComma = errors.New("extra delimiter at end of line")
)
```

这些是 ParseError.Err 中可以返回的错误。

### type ParseError

```go
type ParseError struct {
	StartLine int // 记录开始的行
	Line int // 发生错误的行
	Column int // 发生错误的列（基于 1 的字节索引
	Err error // 实际错误
}
```

解析错误会返回 ParseError。行号以 1 为索引，列以 0 为索引。

#### func (e *ParseError) Error() string

#### func (e *ParseError) Unwrap() Error 添加于1.13

### type Reader

```go
type Reader struct {
  // 逗号是字段分隔符。NewReader 将其设置为逗号（','）。逗号必须是有效的符码，不能是 \r、\n 或 Unicode 替换字符（0xFFFD）。
  Comma rune

  // 如果注释字符不为 0，则为注释字符。以注释字符开头的行如果前面没有空白，将被忽略。如果有前导空白，即使 TrimLeadingSpace 为 true，Comment 字符也会成为字段的一部分。注释必须是有效的符码，不能是 \r、\n 或 Unicode 替换字符 (0xFFFD)。它也不能等于逗号。
  Comment rune

  // FieldsPerRecord 是每条记录预期字段的数量。如果 FieldsPerRecord 为正数，则读取要求每条记录都有给定的字段数。如果 FieldsPerRecord 为 0，Read 会将其设置为第一条记录中的字段数，这样以后的记录就必须有相同的字段数。如果 FieldsPerRecord 为负数，则不做任何检查，记录的字段数可能不固定。
  FieldsPerRecord int

  // 如果 LazyQuotes 为 true，引号可能出现在无引号字段中，而非双引号可能出现在有引号字段中。
  LazyQuotes bool

  // 如果 TrimLeadingSpace 为 true，字段中的前导空白将被忽略。即使字段分隔符逗号是空白，也会这样做。
  TrimLeadingSpace bool

  // ReuseRecord 控制对 "读取 "的调用是否可以返回共享前次调用返回片段的后备数组的片段，以提高性能。默认情况下，每次调用 Read 都会返回调用者新分配的内存。
  ReuseRecord bool

  // 过时：TrailingComma 不再使用。
	trailingComma bool
	// 包含已过滤或未导出的字段
}
```

阅读器从 CSV 编码文件中读取记录。

正如 NewReader 所返回的那样，Reader 希望输入的内容符合 RFC 4180 标准。在第一次调用 Read 或 ReadAll 之前，可以更改导出字段以自定义详细信息。

阅读器会将其输入中的所有 （\r\n）序列转换为普通 （\n），包括多行字段值，这样返回的数据就不会依赖于输入文件所使用的行结束约定。

#### func NewReader(r io.Reader) *Reader

NewReader 返回一个从 r 读取数据的新阅读器。

#### func (r *Reader) FieldPos(field int) (line, column int) 添加于1.17

FieldPos 返回与最近由读取返回的片段中给定索引的字段起始行和列相对应的行和列。行和列的编号从 1 开始；列以字节而不是符为单位计算。

如果在调用此函数时索引超出范围，该函数将崩溃。

#### func (r *Reader) InputOffset() int64 添加于1.19

InputOffset 返回当前读取器位置的输入流字节偏移量。偏移量给出了最近读取行的结束位置和下一行的开始位置。

#### func (r *Reader) Read() (record []string, err error)

如果记录的字段数量出乎意料，则读取会返回该记录，并提示错误 ErrFieldCount。如果记录中包含一个无法解析的字段，Read 会返回部分记录和解析错误信息。部分记录包含出错前读取的所有字段。如果没有数据可读取，Read 会返回 nil，即 io.EOF。如果 ReuseRecord 为 true，返回的片段可在多次调用 Read 时共享。

#### func (r *Reader) ReadAll() (records [][]string, err error)

ReadAll 会读取 r 中的所有剩余记录。成功调用会返回 err == nil，而不是 err == io.EOF。因为 ReadAll 的定义是读取直到 EOF，所以它不会将文件结束作为要报告的错误。

### type Writer

```go
type Writer struct {
  Comma rune // 字段分隔符（由 NewWriter 设置为','）
	UseCRLF bool // 使用 True 作为行结束符
	// 包含已过滤或未导出的字段
}
```

写入器使用 CSV 编码写入记录。

正如 NewWriter 所返回的，Writer 写入的记录以换行结束，并使用", "作为字段分隔符。在首次调用 Write 或 WriteAll 之前，可以更改导出字段以自定义详细信息。

逗号是字段分隔符。

如果 UseCRLF 为 true，Writer 会以 \r\n 代替 \n 结束每一行输出。

单条记录的写入会被缓冲。写入所有数据后，客户端应调用 Flush 方法，以确保所有数据都已转发到底层的 io.Writer 中。任何错误都应通过调用 Error 方法来检查。

#### func NewWriter(w io.Writer) *Writer

NewWriter 返回一个写入 w 的新 Writer。

#### func (w *Writer) Error() error 添加于1.1

错误会报告上一次写入或刷新时发生的任何错误。

#### func (w *Writer) Flush()

刷新会将缓冲数据写入底层的 io.Writer 中。要检查刷新过程中是否发生错误，请调用 Error.Writer。

#### func (w *Writer) Write(record []string) error

Write 将一条 CSV 记录写入 w，并加上必要的引号。记录是字符串的片段，每个字符串是一个字段。写入是缓冲的，因此最终必须调用 Flush，以确保记录被写入底层的 io.Writer.Write 文件。

#### func (w *Writer) WriteAll(records [][]string) error

WriteAll 使用 Write 将多条 CSV 记录写入 w，然后调用 Flush，并返回 Flush 的任何错误信息。