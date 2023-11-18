## package pe

软件包 pe 实现了对 PE（微软视窗可移植可执行文件）文件的访问。

### Security

该软件包的设计并不针对对抗性输入进行加固，也不在 https://go.dev/security/policy 的范围之内。特别是在解析对象文件时，只进行了基本的验证。因此，在解析不受信任的输入时应小心谨慎，因为解析畸形文件可能会耗费大量资源或导致恐慌。

## Index

### Constants

```go
const (
	IMAGE_FILE_MACHINE_UNKNOWN     = 0x0
	IMAGE_FILE_MACHINE_AM33        = 0x1d3
	IMAGE_FILE_MACHINE_AMD64       = 0x8664
	IMAGE_FILE_MACHINE_ARM         = 0x1c0
	IMAGE_FILE_MACHINE_ARMNT       = 0x1c4
	IMAGE_FILE_MACHINE_ARM64       = 0xaa64
	IMAGE_FILE_MACHINE_EBC         = 0xebc
	IMAGE_FILE_MACHINE_I386        = 0x14c
	IMAGE_FILE_MACHINE_IA64        = 0x200
	IMAGE_FILE_MACHINE_LOONGARCH32 = 0x6232
	IMAGE_FILE_MACHINE_LOONGARCH64 = 0x6264
	IMAGE_FILE_MACHINE_M32R        = 0x9041
	IMAGE_FILE_MACHINE_MIPS16      = 0x266
	IMAGE_FILE_MACHINE_MIPSFPU     = 0x366
	IMAGE_FILE_MACHINE_MIPSFPU16   = 0x466
	IMAGE_FILE_MACHINE_POWERPC     = 0x1f0
	IMAGE_FILE_MACHINE_POWERPCFP   = 0x1f1
	IMAGE_FILE_MACHINE_R4000       = 0x166
	IMAGE_FILE_MACHINE_SH3         = 0x1a2
	IMAGE_FILE_MACHINE_SH3DSP      = 0x1a3
	IMAGE_FILE_MACHINE_SH4         = 0x1a6
	IMAGE_FILE_MACHINE_SH5         = 0x1a8
	IMAGE_FILE_MACHINE_THUMB       = 0x1c2
	IMAGE_FILE_MACHINE_WCEMIPSV2   = 0x169
	IMAGE_FILE_MACHINE_RISCV32     = 0x5032
	IMAGE_FILE_MACHINE_RISCV64     = 0x5064
	IMAGE_FILE_MACHINE_RISCV128    = 0x5128
)
```

```go
const (
	IMAGE_DIRECTORY_ENTRY_EXPORT         = 0
	IMAGE_DIRECTORY_ENTRY_IMPORT         = 1
	IMAGE_DIRECTORY_ENTRY_RESOURCE       = 2
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      = 3
	IMAGE_DIRECTORY_ENTRY_SECURITY       = 4
	IMAGE_DIRECTORY_ENTRY_BASERELOC      = 5
	IMAGE_DIRECTORY_ENTRY_DEBUG          = 6
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   = 7
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      = 8
	IMAGE_DIRECTORY_ENTRY_TLS            = 9
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    = 10
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   = 11
	IMAGE_DIRECTORY_ENTRY_IAT            = 12
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   = 13
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR = 14
)
```

IMAGE_DIRECTORY_ENTRY 常量

```go
const (
	IMAGE_FILE_RELOCS_STRIPPED         = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE        = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED      = 0x0004
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     = 0x0008
	IMAGE_FILE_AGGRESIVE_WS_TRIM       = 0x0010
	IMAGE_FILE_LARGE_ADDRESS_AWARE     = 0x0020
	IMAGE_FILE_BYTES_REVERSED_LO       = 0x0080
	IMAGE_FILE_32BIT_MACHINE           = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED          = 0x0200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP = 0x0400
	IMAGE_FILE_NET_RUN_FROM_SWAP       = 0x0800
	IMAGE_FILE_SYSTEM                  = 0x1000
	IMAGE_FILE_DLL                     = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY          = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI       = 0x8000
)
```

IMAGE_FILE_HEADER.Characteristics 的值。这些值可以组合在一起。

```go
const (
	IMAGE_SUBSYSTEM_UNKNOWN                  = 0
	IMAGE_SUBSYSTEM_NATIVE                   = 1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              = 3
	IMAGE_SUBSYSTEM_OS2_CUI                  = 5
	IMAGE_SUBSYSTEM_POSIX_CUI                = 7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           = 8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           = 9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          = 10
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  = 11
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       = 12
	IMAGE_SUBSYSTEM_EFI_ROM                  = 13
	IMAGE_SUBSYSTEM_XBOX                     = 14
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION = 16
)
```

OptionalHeader64.Subsystem 和 OptionalHeader32.Subsystem 值。

```go
const (
	IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA       = 0x0020
	IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE          = 0x0040
	IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY       = 0x0080
	IMAGE_DLLCHARACTERISTICS_NX_COMPAT             = 0x0100
	IMAGE_DLLCHARACTERISTICS_NO_ISOLATION          = 0x0200
	IMAGE_DLLCHARACTERISTICS_NO_SEH                = 0x0400
	IMAGE_DLLCHARACTERISTICS_NO_BIND               = 0x0800
	IMAGE_DLLCHARACTERISTICS_APPCONTAINER          = 0x1000
	IMAGE_DLLCHARACTERISTICS_WDM_DRIVER            = 0x2000
	IMAGE_DLLCHARACTERISTICS_GUARD_CF              = 0x4000
	IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE = 0x8000
)
```

OptionalHeader64.DllCharacteristics 和 OptionalHeader32.DllCharacteristics 值。这些值可以组合在一起。

```go
const (
	IMAGE_SCN_CNT_CODE               = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA   = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA = 0x00000080
	IMAGE_SCN_LNK_COMDAT             = 0x00001000
	IMAGE_SCN_MEM_DISCARDABLE        = 0x02000000
	IMAGE_SCN_MEM_EXECUTE            = 0x20000000
	IMAGE_SCN_MEM_READ               = 0x40000000
	IMAGE_SCN_MEM_WRITE              = 0x80000000
)
```

科室特征标志。

```go
const (
	IMAGE_COMDAT_SELECT_NODUPLICATES = 1
	IMAGE_COMDAT_SELECT_ANY          = 2
	IMAGE_COMDAT_SELECT_SAME_SIZE    = 3
	IMAGE_COMDAT_SELECT_EXACT_MATCH  = 4
	IMAGE_COMDAT_SELECT_ASSOCIATIVE  = 5
	IMAGE_COMDAT_SELECT_LARGEST      = 6
)
```

这些常量构成了 AuxFormat5 中 "选择 "字段的可能值。

```go
const COFFSymbolSize = 18
```

### type COFFSymbol 添加于1.1

```go
type COFFSymbol struct {
  Name [8]uint8
  Value uint32
  SectionNumber int16
  Type uint16
  StorageClass uint8
  NumberOfAuxSymbols uint8
}
```

COFFSymbol 表示单个 COFF 符号表记录。

#### func (sym *COFFSymbol) FullName(st StringTable) (string, error) 添加于1.8

FullName 查找符号 sym 的真实名称。通常，名称存储在 sym.Name 中，但如果名称长度超过 8 个字符，则会存储在 COFF 字符串表 st 中。

### type COFFSymbolAuxFormat5 添加于1.19

```go
type COFFSymbolAuxFormat5 struct {
  Size uint32
  NumRelocs uint16
  NumLineNumbers uint16
  Checksum uint32
  SecNum uint16
  Selection uint8
  // 包含已筛选或未导出字段
}
```

COFFSymbolAuxFormat5 描述了连接到段定义符号的辅助符号的预期形式。PE 格式定义了多种不同的辅助符号格式：格式 1 用于函数定义，格式 2 用于 .be 和 .ef 符号，等等。格式 5 保存与段定义相关的额外信息，包括重定位次数和行号，以及 COMDAT 信息。有关更多信息，请参见 https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions。

### type DataDirectory 添加于1.3

```go
type DataDirectory struct {
  VirtualAddress uint32
  Size uint32
}
```

### type File

```go
type File struct {
  FileHeader
  OptionalHeader any // // 类型为 *OptionalHeader32 或 *OptionalHeader64 的文件
  Sections []*Section
  Symbols []*Symbol // 已删除辅助符号记录的 COFF 符号
  COFFSymbols []COFFSymbol // 所有 COFF 符号（包括辅助符号记录）
  StringTable StringTable
  // 包含已筛选或未导出字段
}
```

文件表示打开的 PE 文件。

#### func NewFile(r io.ReaderAt) (*File, error)

NewFile 创建一个新文件，用于访问底层阅读器中的 PE 二进制文件。

#### func Open(name string) (*File, error)

使用 os.Open 打开命名的文件，并准备将其用作 PE 二进制文件。

#### func (f *File) COFFSymbolReadSectionDefAux(idx int) (*COFFSymbolAuxFormat5, error) 添加于1.19

COFFSymbolReadSectionDefAux 返回段定义符号的辅助信息（包括 COMDAT 信息）。这里的 "idx "是文件的 COFFSymbol 主数组中分段符号的索引。返回值是指向相应辅助符号结构的指针。更多信息，请参阅

辅助符号：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-symbol-records COMDAT 部分：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#comdat-sections-object-only 部分定义的辅助信息：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions

#### func (f *File) Close() error

关闭会关闭文件。如果文件是直接使用 NewFile 而不是 Open 创建的，则关闭没有任何作用。

#### func (f *File) DWARF() (*dwarf.Data, error)

#### func (f *File) ImportedLibraries() ([]string, error)

ImportedLibraries 返回二进制文件 f 调用的所有库的名称，这些库在动态链接时应与二进制文件链接。

#### func (f *File) ImportedSymbols() ([]string, error)

ImportedSymbols 返回二进制 f 所引用的所有符号的名称，这些符号在动态加载时应满足其他库的要求。它不会返回弱符号。

#### func (f *File) Section(name string) *Section

Section 返回具有给定名称的第一个小节，如果不存在该小节，则返回 nil。

### type FileHeader

```go
type FileHeader struct {
  Machine uint16
  NumberOfSections uint16
  TimeDateStamp uint32
  PointerToSymbolTable uint32
  NumberOfSymbols uint32
  SizeOfOptionHeader uint16
  Characteristics uint16
}
```

### type FormatError

```go
type FormatError struct {}
```

FormatError 未使用。保留该类型是为了兼容。

#### func (e *FormatError) Error() string

### type ImportDirectory

```go
type ImportDirectory struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
  // 包含已筛选或未导出字段
}
```

### type OptionalHeader32 添加于1.3

```go
type OptionalHeader32 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]DataDirectory
}
```

### type OptionalHeader64 添加于1.3

```go
type OptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]DataDirectory
}
```

### type Reloc 添加于1.8

```go
type Reloc struct {
  VirtualAddress uint32
  SymbolTableIndex uint32
  Type uint16
}
```

Reloc 表示 PE COFF 重新定位。每个部分都包含自己的重新定位列表。

### Section

```go
type Section struct {
  SectionHeader
  Relocs []Reloc

  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

访问 PE COFF 部分。

#### func (s *Section) Data() ([]byte, error)

数据读取并返回 PE 部分 s 的内容。

如果 s.Offset 为 0，则该部分没有内容，Data 将始终返回一个非零错误。

#### func (s *Section) Open() io.ReadSeeker

Open 返回一个新的 ReadSeeker，读取 PE 部分 s。

如果 s.Offset 为 0，则该部分没有内容，所有对返回的读取器的调用都将返回非零错误。

### type SectionHeader

```go
type SectionHeader struct {
	Name                 string
	VirtualSize          uint32
	VirtualAddress       uint32
	Size                 uint32
	Offset               uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}
```

SectionHeader 与 SectionHeader32 类似，但名称字段由 Go 字符串代替。

### type SectionHeader32

```go
type SectionHeader32 struct {
	Name                 [8]uint8
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}
```

SectionHeader32 表示真正的 PE COFF 部分标头。

### type StringTable 添加于1.8

```go
type StringTable []byte
```

StringTable 是 COFF 字符串表。

#### func (st StringTable) String(start uint32) (string, error) 添加于1.8

从 COFF 字符串表 st 开始偏移处提取字符串。

### type Symbol 添加于1.1

```go
type Symbol struct {
  Name string
  Value uint32
  SectionNumber int16
  Type uint16
  StorageClass uint8
}
```

Symbol 与 COFFSymbol 类似，名称字段由 Go 字符串代替。Symbol 也没有 NumberOfAuxSymbols。