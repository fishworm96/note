## package scanner

Package scanner 实现了 Go 源文本的扫描。它将 []byte 作为源文本，然后通过重复调用 Scan 方法对源文本进行标记化。

## Index

### func PrintError(w io.Writer, err error)

PrintError 是一个实用程序，如果 err 参数是 ErrorList，它将向 w 打印错误列表，每行一个错误。否则，它将打印 err 字符串。

### type Error

```go
type Error struct {
  Pos token.Position
  Msg string
}
```

在 ErrorList 中，错误由 *Error 表示。位置 Pos（如果有效）指向错误标记的开头，错误条件由 Msg.

#### func (e Error) Error() string

Error 实现了错误接口。

### type ErrorHandler

可向 Scanner.Init 提供 ErrorHandler。如果遇到语法错误且已安装了处理程序，则会调用该处理程序并给出一个位置和一条错误信息。位置指向错误标记的起始位置。

### type ErrorList

```go
type ErrorList []*Error
```

ErrorList 是 *Errors 的列表。ErrorList 的零值是一个可随时使用的空 ErrorList。

#### func (p *ErrorList) Add(pos token.Position, msg string)

Add 将给定位置和错误信息的错误添加到 ErrorList 中。

#### func (p ErrorList) Err() error

Err 返回一个与错误列表等价的错误。如果列表为空，Err 返回 nil。

#### func (p ErrorList) Error() string

ErrorList 实现了错误接口。

#### func (p ErrorList) Len() int

ErrorList 实现了排序接口。

#### func (p ErrorList) Less(i, j int) bool

#### func (p *ErrorList) RemoveMultiples()

RemoveMultiples 对 ErrorList 进行排序，并删除每行中除第一个错误之外的所有错误。

#### func (p *ErrorList) Reset()

重置会将 ErrorList 重置为无错误。

#### func (p ErrorList) Sort()

排序可对 ErrorList 进行排序。*Error 条目按位置排序，其他错误按错误信息排序，并且排在任何 *Error 条目之前。

#### func (p ErrorList) Swap(i, j int)

### type Mode

```go
type Mode uint
```

模式值是一组标志（或 0）。它们控制扫描仪的行为。

```go
const (
  ScanComments Mode = 1 << iota // 将注释作为 COMMENT 标记返回
)
```

### type Scanner

```go
type Scanner struct {

	// 公共状态 - 可以修改
	ErrorCount int // 遇到的错误数
	// 包含已过滤或未导出字段
}
```

扫描仪在处理给定文本时保存扫描仪的内部状态。它可以作为其他数据结构的一部分进行分配，但在使用前必须通过 Init 进行初始化。

#### func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode)

Init 将扫描器 s 设置在 src 的开头，为标记化文本 src 做准备。扫描器使用文件设置文件获取位置信息，并为每一行添加行信息。在重新扫描同一文件时，可以重复使用同一文件，因为已经存在的行信息会被忽略。如果文件大小与 src 文件大小不一致，Init 就会引起恐慌。

如果遇到语法错误且 err 不是 nil，对 Scan 的调用将调用错误处理程序 err。此外，每遇到一个错误，Scanner 字段 ErrorCount 就会递增一个。模式参数决定如何处理注释。

请注意，如果文件的第一个字符出现错误，Init 可能会调用 err。

#### func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)

Scan 扫描下一个标记，并返回标记位置、标记及其字面字符串（如果适用）。token.EOF 表示源结束。

如果返回的标记是字面字符串（token.IDENT、token.INT、token.FLOAT、token.IMAG、token.CHAR、token.STRING）或 token.COMMENT，则字面字符串具有相应的值。

如果返回的标记是关键字，则字面字符串就是关键字。

如果返回的标记是 token.SEMICOLON，那么如果源代码中存在分号，则相应的字面字符串为";"；如果分号是由于换行或在 EOF 处插入的，则字面字符串为"\n"。

如果返回的标记为 token.ILLEGAL，则字面字符串为违规字符。

在所有其他情况下，Scan 返回空字面字符串。

为了提高解析的容忍度，即使遇到语法错误，Scan 也会尽可能返回有效的标记。因此，即使生成的标记序列不包含非法标记，客户端也不能认为没有错误发生。相反，它必须检查扫描仪的 ErrorCount 或错误处理程序的调用次数（如果安装了错误处理程序）。

扫描会将行信息添加到用 Init 添加到文件集的文件中。标记位置相对于该文件，因此也相对于文件集。