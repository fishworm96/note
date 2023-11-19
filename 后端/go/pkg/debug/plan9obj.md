## package plan9obj

包 plan9obj 实现了对 Plan 9 a.out 对象文件的访问。

### Security

该软件包的设计并不针对对抗性输入进行加固，也不在 https://go.dev/security/policy 的范围之内。特别是在解析对象文件时，只进行了基本的验证。因此，在解析不受信任的输入时应小心谨慎，因为解析畸形文件可能会耗费大量资源或导致恐慌。

## Index

### Constants

```go
const (
  Magic64 = 0x8000 // 64 位扩展标头

  Magic386 = (4*11+0)*11 + 7
	MagicAMD64 = (4*26+0)*26 + 7 + Magic64
	MagicARM   = (4*20+0)*20 + 7
)
```

### type File

```go
var ErrNoSymbols = errors.New("no symbol section")
```

如果文件中没有此类部分，File.Symbols 将返回 ErrNoSymbols。

### type File

```go
type File struct {
  FileHeader
  Sections []*Section
  // 包含已筛选或未导出字段
}
```

文件表示打开的 Plan 9 a.out 文件。

#### func NewFile(r io.ReaderAt) (*File, error)

NewFile 创建一个新文件，用于访问底层阅读器中的 Plan 9 二进制文件。Plan 9 二进制文件预计将从 ReaderAt 中的位置 0 开始。

#### func Open(name string) (*File, error)

使用 os.Open 打开命名的文件，并准备将其用作 Plan 9 a.out 二进制文件。

#### func (f *File) Close() error

Close 关闭文件。如果文件是直接使用 NewFile 而不是 Open 创建的，则关闭没有任何作用。

#### func (f *File) Section(name string) *Section

Section 返回具有给定名称的部分，如果不存在该部分，则返回 nil。

#### func (f *File) Symbols() ([]Sym, error)

符号返回 f 的符号表。

### type FileHeader

```go
type FileHeader struct {
  Magic uint32
  Bss uint32
  Entry uint64
  PtrSize int
  LoadAddress uint64
  HdrSize uint64
}
```

FileHeader 表示 Plan 9 a.out 文件头。

### type Section

```go
type Section struct {
  SectionHeader

  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

一个小节代表 Plan 9 a.out 文件中的一个小节。

#### func (s *Section) Data() ([]byte, error)

Data 读取并返回 Plan 9 a.out 部分的内容。

#### func (s *Section) Open() io.ReadSeeker

打开后会返回一个新的 ReadSeeker，阅读 Plan 9 a.out 部分。

### type SectionHeader

```go
type SectionHeader struct {
  Name string
  Size uint32
  Offset uint32
}
```

SectionHeader 表示单个 Plan 9 a.out 章节标题。这种结构在磁盘上并不存在，但可以方便地浏览对象文件。

### type Sym

```go
type Sym struct {
  Value uint64
  Type rune
  Name string
}
```

一个符号代表 Plan 9 a.out 符号表部分的一个条目。