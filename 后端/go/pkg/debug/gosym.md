## package gosym

软件包 gosym 实现了对 gc 编译器生成的 Go 二进制文件中嵌入的 Go 符号和行号表的访问。

## Index

### type DecodingError

```go
type DecodingError struct {
  // 包含已筛选或未导出字段
}
```

DecodingError 表示符号表解码过程中出现错误。

#### func (e *DecodingError) Error() string

### type func

```go
type Func struct {
  Entry uint64
  *Sym
  End uint64
  Params []*Sym // nil 适用于 Go 1.3 及更高版本的二进制文件
  Locals []*Sym // nil 适用于 Go 1.3 及更高版本的二进制文件
  FrameSize int
  LineTable *LineTable
  Obj *Obj
}
```

Func 收集有关单个函数的信息。

### type LineTable

```go
type LineTable struct {
  Data []byte
  PC uint64
  Line int
  // 包含已筛选或未导出字段
}
```

行表（LineTable）是一种将程序计数器映射到行号的数据结构。

在 Go 1.1 及更早版本中，每个函数（由 Func 表示）都有自己的 LineTable，行号对应于程序中所有文件的所有源代码行的编号。绝对行号必须分别转换为文件名和文件中的行号。

在 Go 1.2 中，数据格式发生了变化，整个程序只有一个 LineTable，由所有 Funcs 共享，没有绝对行号，只有特定文件中的行号。

在大多数情况下，LineTable 的方法应被视为包的内部细节；调用者应使用 Table 上的方法。

#### func NewLineTable(data []byte, text uint64) *LineTable

NewLineTable 返回与编码数据相对应的新 PC/行表。文本必须是相应文本段的起始地址。

#### func (t *LineTable) LineToPC(line int, maxpc uint64) uint64 废除

#### func (t *LineTable) PCToLine(pc uint64) int 废除

### type Obj

```go
type Obj struct {
  // Funcs 是 Obj 中的函数列表。
  Funcs []Func

  // 在 Go 1.1 及更早版本中，Paths 是一个符号列表，对应于产生 Obj.
	// 与产生 Obj 的源文件名相对应的符号列表。
	// 在 Go 1.2 中，Paths 为空。
	// 使用 Table.Files 的键获取源文件列表。
  Paths []Sym // 元
}
```

Obj 表示符号表中的函数集合。

将二进制分割成不同 Obj 的具体方法是符号表格式的内部细节。

在 Go 的早期版本中，每个源文件都是一个不同的 Obj。

在 Go 1 和 Go 1.1 中，每个软件包为所有 Go 源文件生成一个 Obj，为每个 C 源文件生成一个 Obj。

在 Go 1.2 中，整个程序只有一个 Obj。

### type Sym

```go
type Sym struct {
  Value uint64
  Type type
  Name string
  GoType uint64
  // 如果该符号是函数符号，则相应的 Func
  Func *Func
  // 包含已筛选或未导出字段
}
```

一个 Sym 表示一个符号表条目。

#### func (s *Sym) BaseName() string

BaseName 返回不含软件包或接收器名称的符号名称。

#### func (s *Sym) PackageName() string

PackageName 返回符号名称中的软件包部分，如果没有，则返回空字符串。

#### func (s *Sym) ReceiverName() string

ReceiverName 返回该符号的接收器类型名称，如果没有，则返回空字符串。只有在 s.Name 完全指定了软件包名称的情况下，才会检测到接收器名称。

#### func (s *Sym) Static() bool

Static 报告该符号是否静态（在文件外不可见）。

### type Table

```go
type Table struct {
  Syms []Sym // nil 适用于 Go 1.3 及更高版本的二进制文件
  Funcs []Func
  Files map[string]*Obj // 对于 Go 1.2 及更高版本，所有文件都映射到一个 Obj。
  Objs []Obj // 对于 Go 1.2 及更高版本，片中只有一个对象
  // 包含已筛选或未导出字段
}
```

Table 表示 Go 符号表。它存储了从程序中解码出来的所有符号，并提供了在符号、名称和地址之间进行转换的方法。

#### func NewTable(symtab []byte, pcln *LineTable) (*Table, Error)

NewTable 对 Go 符号表（ELF 中的".gosymtab "部分）进行解码，返回内存中的表示。从 Go 1.3 开始，Go 符号表不再包含符号数据。

#### func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)

LineToPC 查找指定文件中给定行的第一个程序计数器。如果查找该行时出现错误，则返回 UnknownPathError 或 UnknownLineError。

#### func (t *Table) LookupFunc(name string) *Func

LookupFunc 返回具有给定名称的文本、数据或 bss 符号，如果找不到，则返回 nil。

#### func (t *Table) LookupSym(name string) *Func

LookupSym 返回具有给定名称的文本、数据或 bss 符号，如果找不到，则返回 nil。

#### func (t *Table) PCToFunc(pc uint64) *Func

PCToFunc 返回包含程序计数器 pc 的函数，如果没有该函数，则返回 nil。

#### func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)

PCToLine 查找程序计数器的行号信息。如果没有信息，则返回 fn == nil。

#### func (t *Table) SymByAddr(addr uint64) *Sym

SymByAddr 返回从给定地址开始的文本、数据或 bss 符号。

### type UnknownFileError

UnknownFileError 表示无法在符号表中找到特定文件。

#### func (e UnknownFileError) Error() string

### type UnknownLineError

```go
type UnknownLineError struct {
  File string
  Line int
}
```

UnknownLineError 表示无法将某一行映射到程序计数器，原因可能是该行超出了文件的范围，或者是给定行上没有代码。

#### func (e *UnknownLineError) Error() string