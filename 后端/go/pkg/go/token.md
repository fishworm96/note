## package token

包 token 定义了代表 Go 编程语言词性标记和标记基本操作（打印、谓词）的常量。

## Index

### Constants

```go
const (
	LowestPrec = 0 // 非运算符
	UnaryPrec = 6
	HighetPrec Prec = 7
)
```

一组常量，用于基于优先级的表达式解析。非运算符的优先级最低，然后是优先级从 1 开始的运算符，直到一元运算符。最高优先级是选择符、索引及其他运算符和分隔符标记的 "通用 "优先级。

### func IsExported(name string) bool 添加于1.13

IsExported 报告名称是否以大写字母开头。

### func IsIdentifier(name string) bool 添加于1.13

IsIdentifier 报告 name 是否为 Go 标识符，即由字母、数字和下划线组成的非空字符串，其中第一个字符不是数字。关键字不是标识符。

### func IsKeyword(name string) bool 添加于1.13

IsKeyword 报告 name 是否为 Go 关键字，如 "func "或 "return"。

### type File

```go
type File struct {
  // 包含未导出或已过滤的字段
}
```

文件是属于 FileSet 的文件句柄。文件有名称、大小和行偏移表。

#### func (f *File) AddLine(offset int)

AddLine 添加新行的行偏移量。行偏移量必须大于前一行的偏移量，并且小于文件大小；否则，行偏移量将被忽略。

#### func (f *File) AddLineColumnInfo(offset int, filename string, line, column int) 添加于1.11

AddLineColumnInfo 为给定的文件偏移量添加替代文件、行和列号信息。偏移量必须大于之前添加的替代行信息的偏移量，并且小于文件大小，否则该信息将被忽略。

AddLineColumnInfo 通常用于注册行指令的替代位置信息，如 //line filename:line:column。

#### func (f *File) AddLineInfo(offset int, filename string, line int)

AddLineInfo 与 AddLineColumnInfo 类似，都有一个列 = 1 参数。它的存在是为了向后兼容 Go 1.11 之前的代码。

#### func (f *File) Base() int

Base 返回用 AddFile 注册的文件 f 的基准偏移量。

#### func (f *File) Line(p Pos) int

Line 返回给定文件位置 p 的行号；p 必须是该文件中的 Pos 值或 NoPos。

#### func (f *File) LineCount() int

LineCount 返回文件 f 的行数。

#### func (f *File) LineStart(line int) Pos 添加于1.12

LineStart 返回指定行开始位置的 Pos 值。它会忽略使用 AddLineColumnInfo 设置的任何其他位置。如果以 1 为基础的行号无效，LineStart 就会崩溃。

#### func (f *File) Lines() []int 添加于1.21.0

Lines 返回 SetLines 所描述形式的有效行偏移表。调用者不得更改结果。

#### func (f *File) MergeLine(line int) 添加于1.2

MergeLine 将一行与下一行合并。这类似于用空格替换行尾的换行符（不改变其余偏移量）。要获得行号，请参考 Position.Line 等。如果给定的行号无效，MergeLine 就会出错。

#### func (f *File) Name() string

Name 返回用 AddFile 注册的文件 f 的文件名。

#### func (f *File) Offset(p Pos) int

f.Offset(f.Pos(offset)) == offset.Offset 返回给定文件位置 p 的偏移量；p 必须是该文件中有效的 Pos 值。

#### func (f *File) Pos(offset int) Pos

Pos 返回给定文件偏移量的 Pos 值；偏移量必须 <= f.Size()。

#### func (f *File) Position(p Pos) (pos Position)

调用 f.Position(p) 相当于调用 f.PositionFor(p, true)。

#### func (f *File) PositionFor(p Pos, adjusted bool) (pos Position) 添加于1.4

PositionFor 返回给定文件位置 p 的位置值。如果设置了 adjusted，则可以通过位置更改 //line 注释来调整位置；否则将忽略这些注释。

#### func (f *File) SetLines(line []int) bool

SetLines 设置文件的行偏移，并报告是否成功。行偏移量是每行第一个字符的偏移量；例如，对于内容为 "ab\nc\n "的文件，行偏移量为{0, 3}。空文件的行偏移表为空。每一行的偏移量必须大于前一行的偏移量，并且小于文件大小；否则 SetLines 将失败并返回 false。在 SetLines 返回后，调用者不得更改所提供的片段。

#### func (f *File) SetLinesForContent(content []byte)

SetLinesForContent 设置给定文件内容的行偏移。它忽略改变位置的 // 行注释。

#### func (f *File) Size() int

Size 返回用 AddFile 注册的文件 f 的大小。

### type FileSet

```go
type FileSet struct {
  // 包含未导出或已过滤的字段
}
```

文件集代表一组源文件。文件集的方法是同步的，多个程序可以同时调用这些方法。

文件集中每个文件的字节偏移被映射为不同的（整数）区间，每个文件一个区间 [base, base+size] 。base 表示文件中的第一个字节，size 表示相应的文件大小。Pos 值就是这种区间中的一个值。通过确定 Pos 值所属的区间，可以计算出该文件、其文件基数以及 Pos 值所代表的字节偏移量（位置）。

添加新文件时，必须提供文件基数。文件基数可以是超过文件集中已存在文件的任何区间末尾的任何整数值。为方便起见，FileSet.Base 提供了这样一个值，即最近添加的文件的 Pos 间隔的结束值加一。除非以后需要延长间隔，否则应使用 FileSet.Base 作为 FileSet.AddFile 的参数。

文件不再需要时，可以从 FileSet 中删除。这可以减少长期运行应用程序的内存使用量。

#### func NewFileSet() *FileSet

NewFileSet 创建一个新的文件集。

#### func (s *FileSet) AddFile(filename string, base, size int) *File

AddFile 在文件集 s 中添加一个新文件，文件名、基本偏移量和文件大小均已给定，并返回该文件。多个文件可以使用相同的名称。基本偏移量不得小于文件集的 Base()，文件大小不得为负数。在特殊情况下，如果提供的基准偏移量为负数，则会使用 FileSet 的 Base() 的当前值。

添加文件会将文件集的 Base() 值设置为 base + size + 1，作为下一个文件的最小基数值。给定文件偏移量 offs 的 Pos 值 p 之间存在以下关系：

```go
int(p) = base + offs
```

的范围为 [0，size]，因此 p 的范围为 [base，base+size]。为方便起见，File.Pos 可用于从文件偏移量创建特定于文件的位置值。

#### func (s *FileSet) Base() int

Base 返回添加下一个文件时必须提供给 AddFile 的最小基准偏移量。

#### func (s *FileSet) File(p Pos) (f *File)

如果找不到此类文件（例如 p == NoPos），结果为 nil。

#### func (s *FileSet) Iterate(f func(*File) bool)

按照文件集中文件的添加顺序遍历调用 f，直到 f 返回 false。

#### func (s *FileSet) Position(p Pos) (pos Position)

Position 将文件集中的 Pos p 转换为 Position 值。调用 s.Position(p) 相当于调用 s.PositionFor(p,true)。

#### func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position) 添加于1.4

PositionFor 将文件集中的 Pos p 转换为 Position 值。如果设置了 adjusted，则可以通过位置改变 // 行注释来调整位置；否则将忽略这些注释。

#### func (s *FileSet) Read(decode func(any) error) error

读取时调用 decode 将文件集反序列化为 s；s 不得为空。

#### func (s *FileSet) RemoveFile(file *File) 添加于1.20

RemoveFile 会从 FileSet 中移除一个文件，这样后续对其 Pos 间隔的查询结果就会是负数。这样可以减少遇到无限制文件流的长寿命 FileSet 的内存使用量。

删除不属于文件集的文件不会产生任何影响。

#### func (s *FileSet) Write(encode func(any) error) error

Write 调用 encode 将文件集 s 序列化。

### type Pos

```go
type Pos int
```

Pos 是文件集中源位置的紧凑编码。它可以转换成 Position，以获得更方便但更大的表示。

给定文件的 Pos 值是 [base, base+size] 范围内的一个数字，其中 base 和 size 在文件添加到文件集时指定。Pos 值与相应文件基数之间的差值对应于该位置（由 Pos 值表示）与文件开头的字节偏移量。因此，文件基偏移量就是代表文件第一个字节的 Pos 值。

要为特定源偏移量（以字节为单位）创建 Pos 值，首先要使用 FileSet.AddFile 将相应文件添加到当前文件集中，然后为该文件调用 File.Pos(offset)。如果给定了特定文件集 fset 的 Pos 值 p，则可通过调用 fset.Position(p) 获得相应的 Position 值。

Pos 值可以直接使用常用的比较运算符进行比较：如果两个 Pos 值 p 和 q 位于同一文件中，比较 p 和 q 相当于比较各自的源文件偏移量。如果 p 和 q 位于不同的文件中，如果 p 所隐含的文件在 q 所隐含的文件之前添加到各自的文件集中，则 p < q 为真。

```go
const NoPos Pos = 0
```

Pos 的零值为 NoPos；没有与之相关的文件和行信息，NoPos.IsValid() 为假。NoPos 始终小于任何其他 Pos 值。NoPos 对应的 Position 值是 Position 的零值。

#### func (p Pos) IsValid() bool

IsValid 报告位置是否有效。

### type Position

```go
type Position struct {
	Filename string // 文件名（如有
	Offset int // 偏移量，从 0 开始
	Line int // 行号，从 1 开始
	Column int // 列号，从 1 开始（字节数）
}
```

位置描述了一个任意的源位置，包括文件、行和列的位置。如果行号大于 0，则位置有效。

#### func (pos *Position) IsValid() bool

IsValid 报告位置是否有效。

#### func (pos Position) String() string

String 返回多种形式之一的字符串：

```go
file:line:column    valid position with file name
file:line           valid position with file name but no column (column == 0)
line:column         valid position without file name
line                valid position without file name and no column (column == 0)
file                invalid position with file name
-                   invalid position without file name
```

### type Token

```go
type Token int
```

Token 是 Go 编程语言的词性标记集。

```go
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	// 标识符和基本类型文字
	// (这些符号代表字面意义的类别)
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"

	// 操作符和分隔符
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	// 关键字
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR

	// 以临时方式处理的附加标记
	TILDE
)
```

标记列表。

#### func Lookup(ident string) Token

查找将标识符映射到其关键字标记或 IDENT（如果不是关键字）。

#### func (tok Token) IsKeyword() bool

如果标记与关键字对应，IsKeyword 返回 true；否则返回 false。

#### func (tok Token) IsLiteral() bool

对于与标识符和基本类型文字相对应的标记，IsLiteral 返回 true；否则返回 false。

#### func (tok Token) IsOperator() bool

对于操作符和分隔符对应的标记，IsOperator 返回 true；否则返回 false。

#### func (op Token) Precedence() int

优先级返回二进制运算符 op 的运算符优先级。如果 op 不是二进制运算符，结果就是 LowestPrecedence。

#### func (tok Token) String() string

String 返回与 token tok 对应的字符串。对于运算符、分隔符和关键字，字符串是实际的标记字符序列（例如，对于 ADD 标记，字符串是 "+"）。对于所有其他标记，字符串与标记常量名称相对应（例如，对于 IDENT 标记，字符串为 "IDENT"）。