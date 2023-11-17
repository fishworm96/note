## package mache

包 macho 实现了对 Mach-O 对象文件的访问。

### Security

该软件包的设计并不针对对抗性输入进行加固，也不在 https://go.dev/security/policy 的范围之内。特别是在解析对象文件时，只进行了基本的验证。因此，在解析不受信任的输入时应小心谨慎，因为解析畸形文件可能会耗费大量资源或导致恐慌。

## Index

### Constants

```go
const (
	Magic32  uint32 = 0xfeedface
	Magic64  uint32 = 0xfeedfacf
	MagicFat uint32 = 0xcafebabe
)
```

```go
const (
	FlagNoUndefs              uint32 = 0x1
	FlagIncrLink              uint32 = 0x2
	FlagDyldLink              uint32 = 0x4
	FlagBindAtLoad            uint32 = 0x8
	FlagPrebound              uint32 = 0x10
	FlagSplitSegs             uint32 = 0x20
	FlagLazyInit              uint32 = 0x40
	FlagTwoLevel              uint32 = 0x80
	FlagForceFlat             uint32 = 0x100
	FlagNoMultiDefs           uint32 = 0x200
	FlagNoFixPrebinding       uint32 = 0x400
	FlagPrebindable           uint32 = 0x800
	FlagAllModsBound          uint32 = 0x1000
	FlagSubsectionsViaSymbols uint32 = 0x2000
	FlagCanonical             uint32 = 0x4000
	FlagWeakDefines           uint32 = 0x8000
	FlagBindsToWeak           uint32 = 0x10000
	FlagAllowStackExecution   uint32 = 0x20000
	FlagRootSafe              uint32 = 0x40000
	FlagSetuidSafe            uint32 = 0x80000
	FlagNoReexportedDylibs    uint32 = 0x100000
	FlagPIE                   uint32 = 0x200000
	FlagDeadStrippableDylib   uint32 = 0x400000
	FlagHasTLVDescriptors     uint32 = 0x800000
	FlagNoHeapExecution       uint32 = 0x1000000
	FlagAppExtensionSafe      uint32 = 0x2000000
)
```

### Variables

```go
var ErrNotFat = &FormatError{0, "not a fat Mach-O file", nil}
```

当文件不是通用二进制文件，但可能是瘦二进制文件时，NewFatFile 或 OpenFat 会根据文件的魔数返回 ErrNotFat。

### type Cpu

```go
type Cpu uint32
```

Cpu 是一种 Mach-O cpu 类型。

```go
const (
	Cpu386   Cpu = 7
	CpuAmd64 Cpu = Cpu386 | cpuArch64
	CpuArm   Cpu = 12
	CpuArm64 Cpu = CpuArm | cpuArch64
	CpuPpc   Cpu = 18
	CpuPpc64 Cpu = CpuPpc | cpuArch64
)
```

#### func (i Cpu) GoString() string

#### func (i Cpu) String() string

### type Dylib

```go
type Dylib struct {
  LoadBytes
  Name string
  Time uint32
  CurrentVersion uint32
  CompatVersion uint32
}
```

Dylib 表示 Mach-O 加载动态链接库命令。

### type DylibCmd

```go
type DylibCmd struct {
  Cmd LoadCmd
  Len uint32
  Name uint32
  Time uint32
  CurrentVersion uint32
  CompatVersion uint32
}
```

DylibCmd 是 Mach-O 加载动态链接库命令。

### type Dysymtab

```go
type Dysymtab struct {
  LoadBytes
  DysymtabCmd
  IndirectSyms []uint32 // 索引进入 Symtab.Syms
}
```

Dysymtab 表示 Mach-O 动态符号表命令。

### type DysymtabCmd

```go
type DysymtabCmd struct {
	Cmd            LoadCmd
	Len            uint32
	Ilocalsym      uint32
	Nlocalsym      uint32
	Iextdefsym     uint32
	Nextdefsym     uint32
	Iundefsym      uint32
	Nundefsym      uint32
	Tocoffset      uint32
	Ntoc           uint32
	Modtaboff      uint32
	Nmodtab        uint32
	Extrefsymoff   uint32
	Nextrefsyms    uint32
	Indirectsymoff uint32
	Nindirectsyms  uint32
	Extreloff      uint32
	Nextrel        uint32
	Locreloff      uint32
	Nlocrel        uint32
}
```

DysymtabCmd 是 Mach-O 动态符号表命令。

### type FatArch 添加于1.3

```go
type FatArch struct {
	FatArchHeader
	*File
}
```

FatArch 是 FatFile 中的一个 Mach-O 文件。

### type FatArchHeader 添加于1.3

```go
type FatArchHeader struct {
  Cpu Cpu
  SubCpu uint32
  Offset uint32
  Size uint32
  Align uint32
}
```

FatArchHeader 表示特定图像架构的胖头文件。

### type FatFile 添加于1.3

```go
type FatFile struct {
  Magic uint32
  Arches []FatArch
  // 包含已筛选或未导出字段
}
```

#### func NewFatFile(r io.ReaderAt) (*FatFile, error) 添加于1.3

NewFatFile 创建一个新的 FatFile，用于访问通用二进制文件中的所有 Mach-O 映像。Mach-O 二进制文件将从 ReaderAt 中的位置 0 开始。

#### func OpenFat(name string) (*FatFile, error) 添加于1.3

OpenFat 使用 os.Open 打开命名的文件，并准备将其用作 Mach-O 通用二进制文件。

#### func (ff *FatFile) Close() error 添加于1.3

### type File

```go
type File struct {
  FileHeader
  ByteOrder binary.ByteOrder
  Loads []Load
  Sections []*Section

  Symtab *Symtab
  Dysymtab *Dysymtab
  // 包含已筛选或未导出字段
}
```

文件表示打开的 Mach-O 文件。

#### func NewFile(r io.ReaderAt) (*File, error)

NewFile 创建一个新文件，用于访问底层阅读器中的 Mach-O 二进制文件。Mach-O 二进制文件预计将从 ReaderAt 中的位置 0 开始。

#### func Open(name string) (*File, error)

Open 使用 os.Open 打开命名的文件，并准备将其用作 Mach-O 二进制文件。

#### func (*File) Close() error

Close 会关闭文件。如果文件是直接使用 NewFile 而不是 Open 创建的，则关闭没有任何
作用。

#### func (f *File) DWARF() (*dwarf.Data, error)

DWARF 返回 Mach-O 文件的 DWARF 调试信息。

#### func (f *File) ImportedLibraries() ([]string, error)

ImportedLibraries 返回二进制文件 f 调用的所有库的路径，这些库预计将在动态链接时与二进制文件链接。

#### func (f *File) ImportedSymbols() ([]string, error)

ImportedSymbols 返回二进制 f 调用的所有符号的名称，这些符号在动态加载时应由其他库满足。

#### func (f *File) Section(name string) *Section

Section 返回具有给定名称的第一个小节，如果不存在该小节，则返回 nil。

#### func (f *File) Segment(name string) *Segment

Segment 返回具有给定名称的第一个 Segment，如果不存在该 Segment，则返回 nil。

### type FileHeader

```go
type FileHeader struct {
	Magic  uint32
	Cpu    Cpu
	SubCpu uint32
	Type   Type
	Ncmd   uint32
	Cmdsz  uint32
	Flags  uint32
}
```

FileHeader 表示 Mach-O 文件头。

### type FormatError

```go
type FormatError struct {
  // 包含已筛选或未导出字段
}
```

如果数据不具备对象文件的正确格式，某些操作会返回 FormatError。

#### func (e *FormatError) Error() string

### type Load

```go
type Load interface {
  Raw() []byte
}
```

加载表示任何 Mach-O 加载命令。

### type LoadBytes

LoadBytes 是 Mach-O 加载命令的未解释字节。

#### func (b LoadBytes) Raw() []byte

### type LoadCmd

```go
type LoadCmd uint32
```

LoadCmd 是 Mach-O 加载命令。

```go
const (
	LoadCmdSegment LoadCmd = 0x1
	LoadCmdSymtab LoadCmd = 0x2
	LoadCmdThread LoadCmd = 0x4
	LoadCmdUnixThread LoadCmd = 0x5 // 线程+栈
	LoadCmdDysymtab LoadCmd = 0xb
	LoadCmdDylib LoadCmd = 0xc // 加载 dylib 命令
	LoadCmdDylinker LoadCmd = 0xf // id dylinker 命令（非加载 dylinker 命令）
	LoadCmdSegment64 LoadCmd = 0x19
	LoadCmdRpath LoadCmd = 0x8000001c
)
```

#### func (i LoadCmd) GoString() string

#### func (i LoadCmd) String() string

### type Nlist32

```go
type Nlist32 struct {
  Name uint32
  Type uint8
  Sect uint8
  Desc uint16
  Value uint32
}
```

Nlist32 是 Mach-O 32 位符号表项。

### type Nlist64

```go
type Nlist64 struct {
	Name  uint32
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}
```

Nlist64 是 Mach-O 64 位符号表项。

### type Regs386

```go
type Regs386 struct {
	AX    uint32
	BX    uint32
	CX    uint32
	DX    uint32
	DI    uint32
	SI    uint32
	BP    uint32
	SP    uint32
	SS    uint32
	FLAGS uint32
	IP    uint32
	CS    uint32
	DS    uint32
	ES    uint32
	FS    uint32
	GS    uint32
}
```

Regs386 是 Mach-O 386 寄存器结构。

### type RegsAMD64

```go
type RegsAMD64 struct {
	AX    uint64
	BX    uint64
	CX    uint64
	DX    uint64
	DI    uint64
	SI    uint64
	BP    uint64
	SP    uint64
	R8    uint64
	R9    uint64
	R10   uint64
	R11   uint64
	R12   uint64
	R13   uint64
	R14   uint64
	R15   uint64
	IP    uint64
	FLAGS uint64
	CS    uint64
	FS    uint64
	GS    uint64
}
```

RegsAMD64 是 Mach-O AMD64 寄存器结构。

### type Reloc 添加于1.10

```go
类型 Reloc 结构 {
	Addr uint32
	值 uint32
	// 当 Scattered == false && Extern == true 时，Value 是符号编号。
	// 当 Scattered == false && Extern == false 时，Value 是段号。
	// 当 Scattered == true 时，Value 是此 reloc 指向的值。
	类型 uint8
	Len uint8 // 0=byte, 1=word, 2=long, 3=quad
	Pcrel bool
	Extern bool // 在 Scattered == false 时有效
	散布 bool
}
```

Reloc 表示 Mach-O 迁移。

### type RelocTypeARM 添加于1.10

```go
type RelocTypeARM int
```

```go
const (
	ARM_RELOC_VANILLA        RelocTypeARM = 0
	ARM_RELOC_PAIR           RelocTypeARM = 1
	ARM_RELOC_SECTDIFF       RelocTypeARM = 2
	ARM_RELOC_LOCAL_SECTDIFF RelocTypeARM = 3
	ARM_RELOC_PB_LA_PTR      RelocTypeARM = 4
	ARM_RELOC_BR24           RelocTypeARM = 5
	ARM_THUMB_RELOC_BR22     RelocTypeARM = 6
	ARM_THUMB_32BIT_BRANCH   RelocTypeARM = 7
	ARM_RELOC_HALF           RelocTypeARM = 8
	ARM_RELOC_HALF_SECTDIFF  RelocTypeARM = 9
)
```

#### func (r RelocTypeARM) GoString() string 添加于1.10

#### func (r RelocTypeARM) String() string 添加于1.10

### type RelocTypeARM64 添加于1.10

```go
type RelocTypeARM64 int
```

```go
const (
	ARM64_RELOC_UNSIGNED            RelocTypeARM64 = 0
	ARM64_RELOC_SUBTRACTOR          RelocTypeARM64 = 1
	ARM64_RELOC_BRANCH26            RelocTypeARM64 = 2
	ARM64_RELOC_PAGE21              RelocTypeARM64 = 3
	ARM64_RELOC_PAGEOFF12           RelocTypeARM64 = 4
	ARM64_RELOC_GOT_LOAD_PAGE21     RelocTypeARM64 = 5
	ARM64_RELOC_GOT_LOAD_PAGEOFF12  RelocTypeARM64 = 6
	ARM64_RELOC_POINTER_TO_GOT      RelocTypeARM64 = 7
	ARM64_RELOC_TLVP_LOAD_PAGE21    RelocTypeARM64 = 8
	ARM64_RELOC_TLVP_LOAD_PAGEOFF12 RelocTypeARM64 = 9
	ARM64_RELOC_ADDEND              RelocTypeARM64 = 10
)
```

#### func (r RelocTypeARM64) GoString() string 添加于1.10

#### func (r RelocTypeARM64) String() string 添加于1.10

### type RelocTypeGeneric 添加于1.10

```go
type RelocTypeGeneric int
```

```go
const (
	GENERIC_RELOC_VANILLA        RelocTypeGeneric = 0
	GENERIC_RELOC_PAIR           RelocTypeGeneric = 1
	GENERIC_RELOC_SECTDIFF       RelocTypeGeneric = 2
	GENERIC_RELOC_PB_LA_PTR      RelocTypeGeneric = 3
	GENERIC_RELOC_LOCAL_SECTDIFF RelocTypeGeneric = 4
	GENERIC_RELOC_TLV            RelocTypeGeneric = 5
)
```

#### func (r RelocTypeGeneric) GoString() string 添加于1.10

#### func (i RelocTypeGeneric) String() string 添加于1.10

### type RelocTypeX86_64 添加于1.10

```go
type RelocTypeX86_64 int
```

```go
const (
	X86_64_RELOC_UNSIGNED   RelocTypeX86_64 = 0
	X86_64_RELOC_SIGNED     RelocTypeX86_64 = 1
	X86_64_RELOC_BRANCH     RelocTypeX86_64 = 2
	X86_64_RELOC_GOT_LOAD   RelocTypeX86_64 = 3
	X86_64_RELOC_GOT        RelocTypeX86_64 = 4
	X86_64_RELOC_SUBTRACTOR RelocTypeX86_64 = 5
	X86_64_RELOC_SIGNED_1   RelocTypeX86_64 = 6
	X86_64_RELOC_SIGNED_2   RelocTypeX86_64 = 7
	X86_64_RELOC_SIGNED_4   RelocTypeX86_64 = 8
	X86_64_RELOC_TLV        RelocTypeX86_64 = 9
)
```

#### func (r RelocTypeX86_64) GoString() string 添加于1.10

#### func (i RelocTypeX86_64) String() string 添加于1.10

### type Rpath 添加于1.10

```go
type Rpath struct {
  LoadBytes
  Path string
}
```

Rpath 代表 Mach-O rpath 命令。

### type RpathCmd 添加于1.10

```go
type RpathCmd struct {
  Cmd LoadCmd
  Len uint32
  Path uint32
}
```

RpathCmd 是 Mach-O rpath 命令。

### type Section

```go
type Section struct {
  SectionHeader
  Relocs []Reloc

  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

#### func (s *Section) Data() ([]byte, error)

数据读取并返回 Mach-O 部分的内容。

#### func (s *Section) Open() io.ReadSeeker

打开返回一个新的 ReadSeeker，阅读 Mach-O 部分。

### type Section32

```go
type Section32 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint32
	Size     uint32
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
}
```

### type Section64

Section32 是 32 位 Mach-O 小节标头。

### type Segment

```go
type Section64 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint64
	Size     uint64
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
	Reserve3 uint32
}
```

Section64 是 64 位 Mach-O 小节标头。

### type SectionHeader

```go
type SectionHeader struct {
	Name   string
	Seg    string
	Addr   uint64
	Size   uint64
	Offset uint32
	Align  uint32
	Reloff uint32
	Nreloc uint32
	Flags  uint32
}
```

### type Segment

```go
type Segment struct {
  LoadBytes
  SegmentHeader

  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

段代表 Mach-O 32 位或 64 位加载段命令。

#### func (s *Segment) Data() ([]byte, error)

数据读取并返回段的内容。

#### func (s *Segment) Open() io.ReadSeeker

Open 返回一个新的 ReadSeeker 读取片段。

### type Segment32

```go
type Segment32 struct {
	Cmd     LoadCmd
	Len     uint32
	Name    [16]byte
	Addr    uint32
	Memsz   uint32
	Offset  uint32
	Filesz  uint32
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

Segment32 是一条 32 位 Mach-O 段加载命令。

### type Segment64

```go
type Segment64 struct {
	Cmd     LoadCmd
	Len     uint32
	Name    [16]byte
	Addr    uint64
	Memsz   uint64
	Offset  uint64
	Filesz  uint64
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

Segment64 是 64 位 Mach-O 段加载命令。

### SegmentHeader

```go
type SegmentHeader struct {
	Cmd     LoadCmd
	Len     uint32
	Name    string
	Addr    uint64
	Memsz   uint64
	Offset  uint64
	Filesz  uint64
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

SegmentHeader 是 Mach-O 32 位或 64 位加载段命令的头。

### type Symbol

```go
type Symbol struct {
	Name  string
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}
```

符号是 Mach-O 32 位或 64 位符号表项。

### type Symtab

```go
type Symtab struct {
	LoadBytes
	SymtabCmd
	Syms []Symbol
}
```

Symtab 表示 Mach-O 符号表命令。

### type SymtabCmd

```go
type SymtabCmd struct {
	Cmd     LoadCmd
	Len     uint32
	Symoff  uint32
	Nsyms   uint32
	Stroff  uint32
	Strsize uint32
}
```

SymtabCmd 是 Mach-O 符号表命令。

### type Thread

```go
type Thread struct {
	Cmd  LoadCmd
	Len  uint32
	Type uint32
	Data []uint32
}
```

线程是一个 Mach-O 线程状态命令。

### type Type

```go
type Type uint32
```

类型是 Mach-O 文件类型，如对象文件、可执行文件或动态链接库。

```go
const (
  TypeObj Type = 1
  TypeExec Type = 2
  TypeDylib Type = 6
  TypeBundle Type = 8
)
```

#### func (t Type) GoString() string 添加于1.10

#### func (t Type) String() string 添加于1.10