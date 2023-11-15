## package elf

包 elf 实现了对 ELF 对象文件的访问。

### Security

该软件包的设计并不针对对抗性输入进行加固，也不在 https://go.dev/security/policy 的范围之内。特别是在解析对象文件时，只进行了基本的验证。因此，在解析不受信任的输入时应小心谨慎，因为解析畸形文件可能会耗费大量资源或导致恐慌。

## Index

### Constants

```go
const (
	EI_CLASS      = 4  /* 机器类别 */
	EI_DATA       = 5  /* 数据格式 */
	EI_VERSION    = 6  /* ELF 格式版本 */
	EI_OSABI      = 7  /* 操作系统 / ABI 标识 */
	EI_ABIVERSION = 8  /* ABI 版本 */
	EI_PAD        = 9  /* 填充开始（根据 SVR4 ABI） */
	EI_NIDENT     = 16 /* e_ident 数组的大小 */
)
```

Header.Ident 数组的索引。

```go
const ARM_MAGIC_TRAMP_NUMBER = 0x5c000003
```

ELF跳板的魔术数，被明确地选择为一个固定值。

```go
const ELFMAG = "\177ELF"
```

ELF 文件的初始魔数。

```go
const Sym32Size = 16
```

```go
const Sym64Size = 24
```

### Variables

```go
var ErrNoSymbols = errors.New("no symbol section")
```

如果文件中没有此类部分，File.Symbols 和 File.DynamicSymbols 将返回 ErrNoSymbols。

### func R_INFO(sym, typ uint32) uint32

### func R_INFO32(sym, typ uint32) uint32

### func R_SYM32(info uint32) uint32

### func R_SYM64(info uint64) uint32

### func R_TYPE32(info uint32) uint32

### func R_TYPE64(info uint64) uint32

### ST_INFO(bind SymBind, typ SymType) uint8

### type Chdr32 添加于1.6

```go
type Chdr32 struct {
  Type uint32
  Size uint32
  Addralign uint32
}
```

ELF32 压缩头。

### type Chdr64 添加于1.6

```go
type Chdr64 struct {
  Type uint32

  Size uint32
  Addralign uint64
  // 包含已筛选或未导出字段
}
```

ELF64 压缩头。

### type Class

```go
type Class byte
```

类存在于 Header.Ident[EI_CLASS] 和 Header.Class 中。

```go
const (
	ELFCLASSNONE Class = 0 /*  未知类 */
	ELFCLASS32   Class = 1 /* 32 位架构 */
	ELFCLASS64   Class = 2 /* 64 位架构 */
)
```

#### func (i Class) GoString() string

#### func (i Class) String() string

### type CompressionType 添加于1.6

```go
type CompressionType int
```

科室压缩类型。

```go
const (
	COMPRESS_ZLIB CompressionType = 1 /* ZLIB 压缩。*/
	COMPRESS_ZSTD CompressionType = 2 /* ZSTD 压缩。*/
	COMPRESS_LOOS CompressionType = 0x60000000 /* 第一操作系统专用。*/
	COMPRESS_HIOS CompressionType = 0x6fffffff /* 最后一个操作系统专用。*/
	COMPRESS_LOPROC CompressionType = 0x70000000 /* 第一个特定于处理器的类型。*/
	COMPRESS_HIPROC CompressionType = 0x7fffffff /* 最后一种特定于处理器的类型。*/
)
```

#### func (i CompressionType) GoString() string 添加于1.6

#### func (i CompressionType) String() string 添加于1.6

### type Data

```go
type Data byte
```

数据可在 Header.Ident[EI_DATA] 和 Header.Data 中找到。

```go
const (
	ELFDATANONE Data = 0 /* 未知数据格式。*/
	ELFDATA2LSB Data = 1 /* 2 的补码小二进制 */
	ELFDATA2MSB Data = 2 /* 2 的补码大双位 */
)
```

#### func (i Data) GoString() string

#### func (i Data) String() string

### type Dyn32

```go
type Dyn32 struct {
  Tag int32 /* 入口类型。*/
  Val uint32 /* 整数/地址值。*/
}
```

ELF32 动态结构。.dynamic "部分包含一个数组。

### type Dyn64

```go
类型 Dyn64 结构 {
	Tag int64 /* 入口类型。*/
	Val uint64 /* 整数/地址值 */
}
```

ELF64 动态结构。.dynamic "部分包含一个数组。

### type DynFlag

```go
type DynFlag int
```

DT_FLAGS 值。

```go
const (
	DF_ORIGIN DynFlag = 0x0001 /* 表示正在加载的对象可能会对其进行引用。
	   引用
	   $ORIGIN 替换字符串 */
	DF_SYMBOLIC DynFlag = 0x0002 /* 表示 "符号 "链接。*/
	DF_TEXTREL DynFlag = 0x0004 /* 表示在非可写段中可能存在重定位。*/
	DF_BIND_NOW DynFlag = 0x0008 /* 表示动态链接器应处理对象的所有重定位。
	   处理包含此条目对象的所有重定位
	   在将控制权转移给程序之前
	   将控制权转移给程序。*/
	DF_STATIC_TLS DynFlag = 0x0010 /* 表示共享对象或可执行文件包含使用静态链接器的代码。
	   可执行文件包含使用静态
	   线程本地存储方案的代码。*/
)
```

#### func (i DynFlag) GoString() string

#### func (i DynFlag) String() string

### type DynFlag1 添加于1.21.0

```go
type DynFlag1 uint32
```

DT_FLAGS_1 值

```go
const (
	// 表示在将控制权返回给程序之前，必须处理完该对象的所有重定位。
	DF_1_NOW DynFlag1 = 0x00000001
	// 未使用。
	DF_1_GLOBAL DynFlag1 = 0x00000002
	// 表示对象是某个组的成员。
	DF_1_GROUP DynFlag1 = 0x00000004
	// 表示不能从进程中删除对象。
	DF_1_NODELETE DynFlag1 = 0x00000008
	// 仅对筛选器有效。表示立即处理所有相关申请。
	DF_1_LOADFLTR DynFlag1 = 0x00000010
	// 表示在加载任何其他对象之前运行该对象的初始化部分。
	DF_1_INITFIRST DynFlag1 = 0x00000020
	// 表示不能使用 dlopen 将对象添加到正在运行的进程中。
	DF_1_NOOPEN DynFlag1 = 0x00000040
	// 表示对象需要 $ORIGIN 处理。
	DF_1_ORIGIN DynFlag1 = 0x00000080
	// 表示对象应使用直接绑定信息。
	DF_1_DIRECT DynFlag1 = 0x00000100
	// 未使用。
	DF_1_TRANS DynFlag1 = 0x00000200
	// 表示除主要加载对象（通常是可执行文件）外，对象符号表将插在所有符号之前。
	DF_1_INTERPOSE DynFlag1 = 0x00000400
	// 表示搜索该对象的依赖关系时忽略任何默认的库搜索路径。
	DF_1_NODEFLIB DynFlag1 = 0x00000800
	// 表示该对象不会被 dldump 转储。候选对象是没有重定位的对象，可能会在生成替代对象时被包含在内。
	DF_1_NODUMP DynFlag1 = 0x00001000
	// 将此对象标识为由 crle 生成的配置替代对象。触发运行时链接器搜索配置文件 $ORIGIN/ld.config.app-name。
	DF_1_CONFALT DynFlag1 = 0x00002000
	// 仅对申报人有效。终止对任何其他申报人的筛选搜索。
	DF_1_ENDFILTEE DynFlag1 = 0x00004000
	// 表示该对象已应用位移重定位。
	DF_1_DISPRELDNE DynFlag1 = 0x00008000
	// 表示该对象正在进行位移重定位。
	DF_1_DISPRELPND DynFlag1 = 0x00010000
	// 表示该对象包含无法直接绑定的符号。
	DF_1_NODIRECT DynFlag1 = 0x00020000
	// 保留给内核运行时链接器内部使用。
	DF_1_IGNMULDEF DynFlag1 = 0x00040000
	// 保留给内核运行时链接器内部使用。
	DF_1_NOKSYMS DynFlag1 = 0x00080000
	// 保留给内核运行时链接器内部使用。
	DF_1_NOHDR DynFlag1 = 0x00100000
	// 表示该对象已被链接编辑器编辑或修改。
	DF_1_EDITED DynFlag1 = 0x00200000
	// 保留给内核运行时链接器内部使用。
	DF_1_NORELOC DynFlag1 = 0x00400000
	// 表示对象包含个别符号，这些符号应插接在除主要加载对象（通常是可执行文件）之外的所有符号之前。
	DF_1_SYMINTPOSE DynFlag1 = 0x00800000
	// 表示可执行文件需要全局审计。
	DF_1_GLOBAUDIT DynFlag1 = 0x01000000
	// 表示对象定义或引用了单例符号。
	DF_1_SINGLETON DynFlag1 = 0x02000000
	// 表示该对象是一个存根。
	DF_1_STUB DynFlag1 = 0x04000000
	// 表示该对象是与位置无关的可执行文件。
	DF_1_PIE DynFlag1 = 0x08000000
	// 表示对象是内核模块。
	DF_1_KMOD DynFlag1 = 0x10000000
	// 表示该对象是弱标准过滤器。
	DF_1_WEAKFILTER DynFlag1 = 0x20000000
	// 未使用。
	DF_1_NOCOMMON DynFlag1 = 0x40000000
)
```

#### func (i DynFlag1) GoString() string 添加于1.21.0

#### func (i DynFlag1) String() string 添加于1.21.0

### type DynTag 

```go
type DynTag int
```

Dyn.Tag

```go
const (
  DT_NULL DynTag = 0 /* 终止条目。*/
  DT_NEEDED DynTag = 1 /* 所需共享库的字符串表偏移。*/
  DT_PLTRELSZ DynTag = 2 /* PLT 重定位的总大小（以字节为单位）。*/
  DT_PLTGOT DynTag = 3 /* 与处理器相关的地址。*/
  DT_HASH DynTag = 4 /* 符号哈希表地址。*/
  DT_STRTAB DynTag = 5 /* 字符串表地址。*/
  DT_SYMTAB DynTag = 6 /* 符号表地址。*/
  DT_RELA DynTag = 7 /* ElfNN_Rela 重定位地址。*/
  DT_RELASZ DynTag = 8 /* ElfNN_Rela 重定位的总大小。*/
  DT_RELAENT DynTag = 9 /* 每个 ElfNN_Rela 重定位条目的大小。*/
  DT_STRSZ DynTag = 10 /* 字符串表的大小。*/
	DT_SYMENT DynTag = 11 /* 每个符号表条目的大小。*/
	DT_INIT DynTag = 12 /* 初始化函数的地址。*/
	DT_FINI DynTag = 13 /* 最终确定函数的地址。*/
	DT_SONAME DynTag = 14 /* 共享对象名称的字符串表偏移量。*/
	DT_RPATH DynTag = 15 /* 字符串表中库路径的偏移量。[sup] */
	DT_SYMBOLIC DynTag = 16 /* 表示 "符号 "链接。[sup] */
	DT_REL DynTag = 17 /* ElfNN_Rel 重定位地址。*/
	DT_RELSZ DynTag = 18 /* ElfNN_Rel 重定位的总大小。*/
	DT_RELENT DynTag = 19 /* 每个 ElfNN_Rel 重定位的大小。*/
	DT_PLTREL DynTag = 20 /* 用于 PLT 的重定位类型。*/
	DT_DEBUG DynTag = 21 /* 保留（未使用）。*/
	DT_TEXTREL DynTag = 22 /* 表示在非可写段中可能存在重定位。[sup] */
	DT_JMPREL DynTag = 23 /* PLT 重定位的地址。*/
	DT_BIND_NOW DynTag = 24 /* [sup] */
	DT_INIT_ARRAY DynTag = 25 /* 初始化函数指针数组的地址 */
	DT_FINI_ARRAY DynTag = 26 /* 终止函数指针数组的地址 */
	DT_INIT_ARRAYSZ DynTag = 27 /* 以字节为单位的初始化函数数组大小。*/
	DT_FINI_ARRAYSZ DynTag = 28 /* 终止函数数组的大小（以字节为单位）。*/
	DT_RUNPATH DynTag = 29 /* 空尾库搜索路径字符串的字符串表偏移量。*/
	DT_FLAGS DynTag = 30 /* 特定于对象的标志值。*/
	DT_ENCODING     DynTag = 32 /* 大于或等于 DT_ENCODING、小于 DT_LOOS 的值按照以下规则解释 d_un 联盟：偶 == 'd_ptr'、偶 == 'd_val' 或无*/
	DT_PREINIT_ARRAY DynTag = 32 /* 预初始化函数指针数组的地址。*/
	DT_PREINIT_ARRAYSZ DynTag = 33 /* 预初始化函数数组的大小（以字节为单位）。*/
	DT_SYMTAB_SHNDX DynTag = 34 /* SHT_SYMTAB_SHNDX 部分的地址。*/

	DT_LOOS DynTag = 0x6000000d /* 第一个操作系统专用 */
	DT_HIOS DynTag = 0x6ffff000 /* 最后一个操作系统专用 */

	DT_VALRNGLO       DynTag = 0x6ffffd00
	DT_GNU_PRELINKED  DynTag = 0x6ffffdf5
	DT_GNU_CONFLICTSZ DynTag = 0x6ffffdf6
	DT_GNU_LIBLISTSZ  DynTag = 0x6ffffdf7
	DT_CHECKSUM       DynTag = 0x6ffffdf8
	DT_PLTPADSZ       DynTag = 0x6ffffdf9
	DT_MOVEENT        DynTag = 0x6ffffdfa
	DT_MOVESZ         DynTag = 0x6ffffdfb
	DT_FEATURE        DynTag = 0x6ffffdfc
	DT_POSFLAG_1      DynTag = 0x6ffffdfd
	DT_SYMINSZ        DynTag = 0x6ffffdfe
	DT_SYMINENT       DynTag = 0x6ffffdff
	DT_VALRNGHI       DynTag = 0x6ffffdff

	DT_ADDRRNGLO    DynTag = 0x6ffffe00
	DT_GNU_HASH     DynTag = 0x6ffffef5
	DT_TLSDESC_PLT  DynTag = 0x6ffffef6
	DT_TLSDESC_GOT  DynTag = 0x6ffffef7
	DT_GNU_CONFLICT DynTag = 0x6ffffef8
	DT_GNU_LIBLIST  DynTag = 0x6ffffef9
	DT_CONFIG       DynTag = 0x6ffffefa
	DT_DEPAUDIT     DynTag = 0x6ffffefb
	DT_AUDIT        DynTag = 0x6ffffefc
	DT_PLTPAD       DynTag = 0x6ffffefd
	DT_MOVETAB      DynTag = 0x6ffffefe
	DT_SYMINFO      DynTag = 0x6ffffeff
	DT_ADDRRNGHI    DynTag = 0x6ffffeff

	DT_VERSYM     DynTag = 0x6ffffff0
	DT_RELACOUNT  DynTag = 0x6ffffff9
	DT_RELCOUNT   DynTag = 0x6ffffffa
	DT_FLAGS_1    DynTag = 0x6ffffffb
	DT_VERDEF     DynTag = 0x6ffffffc
	DT_VERDEFNUM  DynTag = 0x6ffffffd
	DT_VERNEED    DynTag = 0x6ffffffe
	DT_VERNEEDNUM DynTag = 0x6fffffff

	DT_LOPROC DynTag = 0x70000000 /* 第一种特定于处理器的类型。*/

	DT_MIPS_RLD_VERSION           DynTag = 0x70000001
	DT_MIPS_TIME_STAMP            DynTag = 0x70000002
	DT_MIPS_ICHECKSUM             DynTag = 0x70000003
	DT_MIPS_IVERSION              DynTag = 0x70000004
	DT_MIPS_FLAGS                 DynTag = 0x70000005
	DT_MIPS_BASE_ADDRESS          DynTag = 0x70000006
	DT_MIPS_MSYM                  DynTag = 0x70000007
	DT_MIPS_CONFLICT              DynTag = 0x70000008
	DT_MIPS_LIBLIST               DynTag = 0x70000009
	DT_MIPS_LOCAL_GOTNO           DynTag = 0x7000000a
	DT_MIPS_CONFLICTNO            DynTag = 0x7000000b
	DT_MIPS_LIBLISTNO             DynTag = 0x70000010
	DT_MIPS_SYMTABNO              DynTag = 0x70000011
	DT_MIPS_UNREFEXTNO            DynTag = 0x70000012
	DT_MIPS_GOTSYM                DynTag = 0x70000013
	DT_MIPS_HIPAGENO              DynTag = 0x70000014
	DT_MIPS_RLD_MAP               DynTag = 0x70000016
	DT_MIPS_DELTA_CLASS           DynTag = 0x70000017
	DT_MIPS_DELTA_CLASS_NO        DynTag = 0x70000018
	DT_MIPS_DELTA_INSTANCE        DynTag = 0x70000019
	DT_MIPS_DELTA_INSTANCE_NO     DynTag = 0x7000001a
	DT_MIPS_DELTA_RELOC           DynTag = 0x7000001b
	DT_MIPS_DELTA_RELOC_NO        DynTag = 0x7000001c
	DT_MIPS_DELTA_SYM             DynTag = 0x7000001d
	DT_MIPS_DELTA_SYM_NO          DynTag = 0x7000001e
	DT_MIPS_DELTA_CLASSSYM        DynTag = 0x70000020
	DT_MIPS_DELTA_CLASSSYM_NO     DynTag = 0x70000021
	DT_MIPS_CXX_FLAGS             DynTag = 0x70000022
	DT_MIPS_PIXIE_INIT            DynTag = 0x70000023
	DT_MIPS_SYMBOL_LIB            DynTag = 0x70000024
	DT_MIPS_LOCALPAGE_GOTIDX      DynTag = 0x70000025
	DT_MIPS_LOCAL_GOTIDX          DynTag = 0x70000026
	DT_MIPS_HIDDEN_GOTIDX         DynTag = 0x70000027
	DT_MIPS_PROTECTED_GOTIDX      DynTag = 0x70000028
	DT_MIPS_OPTIONS               DynTag = 0x70000029
	DT_MIPS_INTERFACE             DynTag = 0x7000002a
	DT_MIPS_DYNSTR_ALIGN          DynTag = 0x7000002b
	DT_MIPS_INTERFACE_SIZE        DynTag = 0x7000002c
	DT_MIPS_RLD_TEXT_RESOLVE_ADDR DynTag = 0x7000002d
	DT_MIPS_PERF_SUFFIX           DynTag = 0x7000002e
	DT_MIPS_COMPACT_SIZE          DynTag = 0x7000002f
	DT_MIPS_GP_VALUE              DynTag = 0x70000030
	DT_MIPS_AUX_DYNAMIC           DynTag = 0x70000031
	DT_MIPS_PLTGOT                DynTag = 0x70000032
	DT_MIPS_RWPLT                 DynTag = 0x70000034
	DT_MIPS_RLD_MAP_REL           DynTag = 0x70000035

	DT_PPC_GOT DynTag = 0x70000000
	DT_PPC_OPT DynTag = 0x70000001

	DT_PPC64_GLINK DynTag = 0x70000000
	DT_PPC64_OPD   DynTag = 0x70000001
	DT_PPC64_OPDSZ DynTag = 0x70000002
	DT_PPC64_OPT   DynTag = 0x70000003

	DT_SPARC_REGISTER DynTag = 0x70000001

	DT_AUXILIARY DynTag = 0x7ffffffd
	DT_USED      DynTag = 0x7ffffffe
	DT_FILTER    DynTag = 0x7fffffff

	DT_HIPROC DynTag = 0x7fffffff /* 最后一种特定于处理器的类型。*/
)
```

#### func (i DynTag) GoString() string

#### func (i DynTag) String() string

### type File

```go
type File struct {
  FIleHeader
  Sections []*Section
  Progs []*Prog
  // 包含已筛选或未导出字段
}
```

文件表示打开的 ELF 文件。

#### func NewFile(r io.ReaderAt) (*File, error)

NewFile 创建一个新文件，用于访问底层阅读器中的 ELF 二进制文件。ELF 二进制文件应从读取器位置 0 开始。

#### func Open(name string) (*File, error)

使用 os.Open 打开命名的文件，并准备将其用作 ELF 二进制文件。

#### func (f *File) Close() error

关闭会关闭文件。如果文件是直接使用 NewFile 而不是 Open 创建的，则关闭没有任何作用。

#### func (f *File) DWARF() (*dwarf.Data, error)

#### func (f *File) DynString(tag DynTag) ([]string, error) 添加于1.1

DynString 返回文件动态部分中为给定标记列出的字符串。

标签必须是可取字符串值的标签：DT_NEEDED、DT_SONAME、DT_RPATH 或 DT_RUNPATH。

#### func (f *File) DynValue(tag DynTag) ([]uint64, error) 添加于1.21.0

DynValue 返回文件动态部分中给定标记的值。

#### func (f *File) DynamicSymbols() ([]Symbol, error) 添加于1.4

DynamicSymbols 返回 f 的动态符号表。符号将按照在 f 中出现的顺序排列。

如果 f 有一个符号版本表，返回的符号将有初始化的版本和库字段。

为了与符号表兼容，DynamicSymbols 会省略索引 0 处的空符号。以 symtab 的形式检索符号后，外部提供的索引 x 将对应于 symtab[x-1]，而不是 symtab[x]。

#### func (f *File) ImportedLibraries() ([]string, error)

ImportedLibraries 返回二进制文件 f 调用的所有库的名称，这些库在动态链接时应与二进制文件链接。

#### func (f *File) ImportedSymbols() ([]ImportedSymbol, error)

ImportedSymbols 返回二进制 f 所引用的所有符号的名称，这些符号在动态加载时应满足其他库的要求。它不会返回弱符号。

#### func (f *File) Section(name string) *Section

Section 返回具有给定名称的部分，如果不存在该部分，则返回 nil。

#### func (f *File) SectionByType(typ SectionType) *Section

SectionByType 返回 f 中第一个具有给定类型的部分，如果没有这样的部分，则返回 nil。

#### func (f *File) Symbols() ([]Symbol, error)

符号将按照在 f 中出现的顺序排列。

为了与 Go 1.0 兼容，Symbols 省略了位于索引 0 的空符号。以 symtab 的形式获取符号后，外部提供的索引 x 将对应 symtab[x-1]，而不是 symtab[x]。

### type FileHeader

```go
type FileHeader struct {
  Class Class
  Data Data
  Version Version
  OSABI OSABI
  ABIVersion uint8
  ByteOrder binary.ByteOrder
  Type Type
  Machine Machine
  Entry uint64
}
```

FileHeader 表示 ELF 文件头。

### type FormatError

```go
type FormatError struct {
  // 包含已筛选或未导出字段
}
```

#### func (e *FormatError) Error() string

### type Header32

```go
type Header32 struct {
	Ident [EI_NIDENT]byte /* 文件标识。*/
	Type uint16 /* 文件类型。*/
	Machine uint16 /* 机器架构。*/
	版本 uint32 /* ELF 格式版本。*/
	入口 uint32 /* 入口点。*/
	Phoff uint32 /* 程序头文件偏移量。*/
	Shoff uint32 /* 段落头文件偏移量。*/
	Flags uint32 /* 架构专用标志。*/
	Ehsize uint16 /* ELF 头文件的大小（字节）。*/
	Phentsize uint16 /* 程序头条目大小。*/
	Phnum uint16 /* 程序头条目数量。*/
	Shentsize uint16 /* 小节头条目大小。*/
	Shnum uint16 /* 节标题条目数。*/
	Shstrndx uint16 /* 部分名称字符串部分。*/
}
```

ELF32 文件头。

### type Header64

```go
type Heaer struct {
	Ident [EI_NIDENT]byte /* 文件标识。*/
	Type uint16 /* 文件类型。*/
	Machine uint16 /* 机器架构。*/
	版本 uint32 /* ELF 格式版本。*/
	入口 uint64 /* 入口点。*/
	Phoff uint64 /* 程序头文件偏移量。*/
	Shoff uint64 /* 段落头文件偏移量。*/
	Flags uint32 /* 架构专用标志。*/
	Ehsize uint16 /* ELF 头文件的大小（字节）。*/
	Phentsize uint16 /* 程序头条目大小。*/
	Phnum uint16 /* 程序头条目数量。*/
	Shentsize uint16 /* 小节头条目大小。*/
	Shnum uint16 /* 节标题条目数。*/
	Shstrndx uint16 /* 部分名称字符串部分。*/
}
```

ELF64 文件头。

### type ImportedSymbol

```go
type ImportedSymbol struct {
  Name string
  Version string
  Library string
}
```

### type Machine

```go
type Machine uint16
```

在 Machine 中找到 Header.Machine。

```go
const (
	EM_NONE Machine = 0 /* 未知机器。*/
	EM_M32 Machine = 1 /* AT&T WE32100.*/
	EM_SPARC Machine = 2 /* Sun SPARC.*/
	EM_386 机器 = 3 /* Intel i386.*/
	EM_68K Machine = 4 /* Motorola 68000.*/
	EM_88K Machine = 5 /* Motorola 88000.*/
	EM_860 机器 = 7 /* 英特尔 i860。*/
	EM_MIPS Machine = 8 /* MIPS R3000 Big-Endian only.*/
	EM_S370 Machine = 9 /* IBM System/370.*/
	EM_MIPS_RS3_LE Machine = 10 /* MIPS R3000 Little-Endian.*/
	EM_PARISC Machine = 15 /* HP PA-RISC。*/
	EM_VPP500 Machine = 17 /* Fujitsu VPP500.*/
	EM_SPARC32PLUS Machine = 18 /* SPARC v8plus.*/
	EM_960 Machine = 19 /* Intel 80960.
	EM_PPC Machine = 20 /* PowerPC 32 位。*/
	EM_PPC64 Machine = 21 /* PowerPC 64 位。*/
	EM_S390 机器 = 22 /* IBM System/390.
	EM_V800 机器 = 36 /* NEC V800。*/
	EM_FR20 机器 = 37 /* 富士通 FR20。*/
	EM_RH32 机器 = 38 /* TRW RH-32.
	EM_RCE 机器 = 39 /* Motorola RCE.*/
	EM_ARM 机器 = 40 /* ARM.*/
	EM_SH 机器 = 42 /* 日立 SH。*/
	EM_SPARCV9 Machine = 43 /* SPARC v9 64 位。*/
	EM_TRICORE Machine = 44 /* Siemens TriCore 嵌入式处理器。*/
	EM_ARC Machine = 45 /* Argonaut RISC Core。*/
	EM_H8_300 机器 = 46 /* 日立 H8/300。*/
	EM_H8_300H 机器 = 47 /* 日立 H8/300H。*/
	EM_H8S 机器 = 48 /* 日立 H8S。*/
	EM_H8_500 机器 = 49 /* 日立 H8/500。*/
	EM_IA_64 机器 = 50 /* 英特尔 IA-64 处理器。*/
	EM_MIPS_X Machine = 51 /* Stanford MIPS-X.*/
	EM_COLDFIRE Machine = 52 /* Motorola ColdFire.*/
	EM_68HC12 Machine = 53 /* Motorola M68HC12.*/
	EM_MMA Machine = 54 /* Fujitsu MMA.*/
	EM_PCP Machine = 55 /* Siemens PCP.*/
	EM_NCPU Machine = 56 /* Sony nCPU.*/
	EM_NDR1 机器 = 57 /* 电装 NDR1 微处理器。*/
	EM_STARCORE Machine = 58 /* Motorola Star*Core processor.*/
	EM_ME16 机器 = 59 /* 丰田 ME16 处理器。*/
	EM_ST100 Machine = 60 /* STMicroelectronics ST100 处理器。*/
	EM_TINYJ Machine = 61 /* Advanced Logic Corp.TinyJ 处理器。*/
	EM_X86_64 机器 = 62 /* Advanced Micro Devices x86-64 */
	EM_PDSP 机器 = 63 /* 索尼 DSP 处理器 */
	EM_PDP10 机器 = 64 /* Digital Equipment Corp.PDP-10 */
	EM_PDP11 Machine = 65 /* Digital Equipment Corp.PDP-11 */
	EM_FX66 机器 = 66 /* 西门子 FX66 微控制器 */
	EM_ST9PLUS 机器 = 67 /* 意法半导体 ST9+ 8/16 位微控制器 */
	EM_ST7 机器 = 68 /* 意法半导体 ST7 8 位微控制器 */
	EM_68HC16 机器 = 69 /* 摩托罗拉 MC68HC16 微控制器 */
	EM_68HC11 机器 = 70 /* 摩托罗拉 MC68HC11 微控制器 */
	EM_68HC08 机器 = 71 /* 摩托罗拉 MC68HC08 微控制器 */
	EM_68HC05 机器 = 72 /* 摩托罗拉 MC68HC05 微控制器 */
	EM_SVX 机器 = 73 /* Silicon Graphics SVx */
	EM_ST19 机器 = 74 /* 意法半导体 ST19 8 位微控制器 */
	EM_VAX 机器 = 75 /* 数字 VAX */
	EM_CRIS Machine = 76 /* Axis Communications 32 位嵌入式处理器 */
	EM_JAVELIN 机器 = 77 /* 英飞凌科技公司 32 位嵌入式处理器 */
	EM_FIREPATH Machine = 78 /* Element 14 64 位 DSP 处理器 */
	EM_ZSP 机器 = 79 /* LSI Logic 16 位 DSP 处理器 */
	EM_MMIX Machine = 80 /* Donald Knuth's educational 64 位处理器 */
	EM_HUANY Machine = 81 /* 哈佛大学机器独立对象文件 */
	EM_PRISM Machine = 82 /* SiTera Prism */
	EM_AVR Machine = 83 /* Atmel AVR 8 位微控制器 */
	EM_FR30 机器 = 84 /* 富士通 FR30 */
	EM_D10V 机器 = 85 /* Mitsubishi D10V */
	EM_D30V 机器 = 86 /* Mitsubishi D30V */
	EM_V850 机器 = 87 /* NEC v850 */
	EM_M32R 机器 = 88 /* Mitsubishi M32R */
	EM_MN10300 机器 = 89 /* Matsushita MN10300 */
	EM_MN10200 机器 = 90 /* Matsushita MN10200 */
	EM_PJ 机器 = 91 /* picoJava */
	EM_OPENRISC Machine = 92 /* OpenRISC 32 位嵌入式处理器 */
	EM_ARC_COMPACT Machine = 93 /* ARC International ARCompact 处理器（旧拼写/同义词：EM_ARC_A5） */
	EM_XTENSA 机器 = 94 /* Tensilica Xtensa 架构 */
	EM_VIDEOCORE 机器 = 95 /* Alphamosaic VideoCore 处理器 */
	EM_TMM_GPP Machine = 96 /* Thompson Multimedia General Purpose Processor */
	EM_NS32K 机器 = 97 /* 国家半导体 32000 系列 */
	EM_TPC 机器 = 98 /* 特诺网络 TPC 处理器 */
	EM_SNP1K 机器 = 99 /* Trebia SNP 1000 处理器 */
	EM_ST200 机器 = 100 /* STMicroelectronics (www.st.com) ST200 微控制器 */
	EM_IP2K Machine = 101 /* Ubicom IP2xxx 微控制器系列 */
	EM_MAX 机器 = 102 /* MAX 处理器 */
	EM_CR Machine = 103 /* National Semiconductor CompactRISC 微处理器 */
	EM_F2MC16 机器 = 104 /* 富士通 F2MC16 */
	EM_MSP430 机器 = 105 /* 德州仪器公司嵌入式微控制器 msp430 */
	EM_BLACKFIN Machine = 106 /* Analog Devices Blackfin (DSP) 处理器 */
	EM_SE_C33 机器 = 107 /* 精工爱普生处理器 S1C33 系列 */
	EM_SEP 机器 = 108 /* 夏普嵌入式微处理器 */
	EM_ARCA 机器 = 109 /* Arca RISC 微处理器 */
	EM_UNICORE 机器 = 110 /* PKU-Unity 有限公司和北京大学 MPRC 的微处理器系列 */
	EM_EXCESS 机器 = 111 /* eXcess：16/32/64 位可配置嵌入式 CPU */
	EM_DXP Machine = 112 /* Icera Semiconductor Inc.深度执行处理器 */
	EM_ALTERA_NIOS2 机器 = 113 /* Altera Nios II 软核处理器 */
	EM_CRX 机器 = 114 /* 国家半导体 CompactRISC CRX 微处理器 */
	EM_XGATE 机器 = 115 /* 摩托罗拉 XGATE 嵌入式处理器 */
	EM_C166 机器 = 116 /* 英飞凌 C16x/XC16x 处理器 */
	EM_M16C 机器 = 117 /* 瑞萨 M16C 系列微处理器 */
	EM_DSPIC30F 机器 = 118 /* Microchip Technology dsPIC30F 数字信号控制器 */
	EM_CE 机器 = 119 /* 飞思卡尔通信引擎 RISC 内核 */
	EM_M32C Machine = 120 /* Renesas M32C 系列微处理器 */
	EM_TSK3000 机器 = 131 /* Altium TSK3000 内核 */
	EM_RS08 Machine = 132 /* Freescale RS08 嵌入式处理器 */
	EM_SHARC Machine = 133 /* Analog Devices SHARC 系列 32 位 DSP 处理器 */
	EM_ECOG2 机器 = 134 /* Cyan Technology eCOG2 微处理器 */
	EM_SCORE7 机器 = 135 /* 凌阳 S+core7 RISC 处理器 */
	EM_DSP24 机器 = 136 /* 新日本电台（NJR）24 位 DSP 处理器 */
	EM_VIDEOCORE3 机器 = 137 /* Broadcom VideoCore III 处理器 */
	EM_LATTICEMICO32 机器 = 138 /* 莱迪思 FPGA 架构的 RISC 处理器 */
	EM_SE_C17 机器 = 139 /* 精工爱普生 C17 系列 */
	EM_TI_C6000 机器 = 140 /* 德州仪器 TMS320C6000 DSP 系列 */
	EM_TI_C2000 机器 = 141 /* 德州仪器 TMS320C2000 DSP 系列 */
	EM_TI_C5500 机器 = 142 /* 德州仪器 TMS320C55x DSP 系列 */
	EM_TI_ARP32 机器 = 143 /* 德州仪器特定应用 RISC 处理器，32 位取值 */
	EM_TI_PRU 机器 = 144 /* 德州仪器可编程实时单元 */
	EM_MMDSP_PLUS 机器 = 160 /* 意法半导体 64 位 VLIW 数据信号处理器 */
	EM_CYPRESS_M8C 机器 = 161 /* 赛普拉斯 M8C 微处理器 */
	EM_R32C Machine = 162 /* Renesas R32C 系列微处理器 */
	EM_TRIMEDIA 机器 = 163 /* 恩智浦半导体 TriMedia 架构系列 */
	EM_QDSP6 机器 = 164 /* QUALCOMM DSP6 处理器 */
	EM_8051 机器 = 165 /* 英特尔 8051 及其变体 */
	EM_STXP7X 机器 = 166 /* 意法半导体 STxP7x 系列可配置和可扩展 RISC 处理器 */
	EM_NDS32 Machine = 167 /* Andes Technology 紧凑型代码嵌入式 RISC 处理器系列 */
	EM_ECOG1 机器 = 168 /* Cyan Technology eCOG1X 系列 */
	EM_ECOG1X 机器 = 168 /* Cyan Technology eCOG1X 系列 */
	EM_MAXQ30 Machine = 169 /* Dallas Semiconductor MAXQ30 Core 微控制器 */
	EM_XIMO16 机器 = 170 /* 新日本无线电（NJR）16 位 DSP 处理器 */
	EM_MANIK 机器 = 171 /* M2000 可重构 RISC 微处理器 */
	EM_CRAYNV2 机器 = 172 /* Cray Inc.NV2 矢量架构 */
	EM_RX 机器 = 173 /* Renesas RX 系列 */
	EM_METAG Machine = 174 /* Imagination Technologies META 处理器架构 */
	EM_MCST_ELBRUS 机器 = 175 /* MCST Elbrus 通用硬件架构 */
	EM_ECOG16 机器 = 176 /* Cyan Technology eCOG16 系列 */
	EM_CR16 Machine = 177 /* National Semiconductor CompactRISC CR16 16 位微处理器 */
	EM_ETPU 机器 = 178 /* 飞思卡尔扩展时间处理单元 */
	EM_SLE9X 机器 = 179 /* 英飞凌科技公司 SLE9X 内核 */
	EM_L10M 机器 = 180 /* 英特尔 L10M */
	EM_K10M 机器 = 181 /* 英特尔 K10M */
	EM_AARCH64 机器 = 183 /* ARM 64 位架构 (AArch64) */
	EM_AVR32 机器 = 185 /* Atmel 公司 32 位微处理器系列 */
	EM_STM8 机器 = 186 /* STMicroeletronics STM8 8 位微控制器 */
	EM_TILE64 Machine = 187 /* Tilera TILE64 多核架构系列 */
	EM_TILEPRO 机器 = 188 /* Tilera TILEPro 多核架构系列 */
	EM_MICROBLAZE Machine = 189 /* Xilinx MicroBlaze 32 位 RISC 软处理器内核 */
	EM_CUDA Machine = 190 /* NVIDIA CUDA 架构 */
	EM_TILEGX Machine = 191 /* Tilera TILE-Gx 多核架构系列 */
	EM_CLOUDSHIELD Machine = 192 /* CloudShield 架构系列 */
	EM_COREA_1ST Machine = 193 /* KIPO-KAIST Core-A 第一代处理器系列 */
	EM_COREA_2ND 机器 = 194 /* KIPO-KAIST Core-A 第二代处理器系列 */
	EM_ARC_COMPACT2 Machine = 195 /* Synopsys ARCompact V2 */
	EM_OPEN8 Machine = 196 /* Open8 8 位 RISC 软处理器内核 */
	EM_RL78 机器 = 197 /* Renesas RL78 系列 */
	EM_VIDEOCORE5 Machine = 198 /* Broadcom VideoCore V 处理器 */
	EM_78KOR Machine = 199 /* Renesas 78KOR 系列 */
	EM_56800EX 机器 = 200 /* Freescale 56800EX 数字信号控制器 (DSC) */
	EM_BA1 Machine = 201 /* Beyond BA1 CPU 架构 */
	EM_BA2 机器 = 202 /* 超越 BA2 CPU 架构 */
	EM_XCORE Machine = 203 /* XMOS xCORE 处理器系列 */
	EM_MCHP_PIC 机器 = 204 /* Microchip 8 位 PIC(r) 系列 */
	EM_INTEL205 机器 = 205 /* 英特尔保留 */
	EM_INTEL206 机器 = 206 /* 由英特尔保留 */
	EM_INTEL207 Machine = 207 /* Reserved by Intel */
	EM_INTEL208 Machine = 208 /* Reserved by Intel */
	EM_INTEL209 机器 = 209 /* 被英特尔公司保留 */
	EM_KM32 机器 = 210 /* KM211 KM32 32 位处理器 */
	EM_KMX32 机器 = 211 /* KM211 KMX32 32 位处理器 */
	EM_KMX16 机器 = 212 /* KM211 KMX16 16 位处理器 */
	EM_KMX8 机器 = 213 /* KM211 KMX8 8 位处理器 */
	EM_KVARC 机器 = 214 /* KM211 KVARC 处理器 */
	EM_CDP Machine = 215 /* Paneve CDP 架构系列 */
	EM_COGE Machine = 216 /* Cognitive Smart Memory Processor */
	EM_COOL Machine = 217 /* Bluechip Systems CoolEngine */
	EM_NORC Machine = 218 /* Nanoradio Optimized RISC */
	EM_CSR_KALIMBA Machine = 219 /* CSR Kalimba 架构系列 */
	EM_Z80 机器 = 220 /* Zilog Z80 */
	EM_VISIUM Machine = 221 /* Controls and Data Services VISIUMcore 处理器 */
	EM_FT32 机器 = 222 /* FTDI 芯片 FT32 高性能 32 位 RISC 架构 */
	EM_MOXIE Machine = 223 /* Moxie 处理器系列 */
	EM_AMDGPU Machine = 224 /* AMD GPU 架构 */
	EM_RISCV 机器 = 243 /* RISC-V */
	EM_LANAI Machine = 244 /* Lanai 32 位处理器 */
	EM_BPF Machine = 247 /* Linux BPF - 内核虚拟机 */
	EM_LOONGARCH Machine = 258 /* LoongArch */
  /* 非标准或过时。*/
	EM_486 Machine = 6 /* Intel i486。*/
	EM_MIPS_RS4_BE 机器 = 10 /* MIPS R4000 Big-Endian */
	EM_ALPHA_STD Machine = 41 /* Digital Alpha（标准值）。*/
	EM_ALPHA Machine = 0x9026 /* Alpha (在没有 ABI 的情况下写入) */
)
```

#### func (i Machine) GoString() string

#### func (i Machine) String() string

### type NType

```go
type NType int
```

NType 值；用于核心文件。

```go
const (
	NT_PRSTATUS NType = 1 /* 处理状态。*/
	NT_FPREGSET NType = 2 /* 浮点寄存器。*/
	NT_PRPSINFO NType = 3 /* 进程状态信息。*/
)
```

#### func (i NType) GoString() string

#### func (i NType) String() string

### type OSABI

```go
type OSABI byte
```

OSABI 位于 Header.Ident[EI_OSABI] 和 Header.OSABI 中。

```go
const (
	ELFOSABI_NONE OSABI = 0 /* UNIX System V ABI */
	ELFOSABI_HPUX OSABI = 1 /* HP-UX 操作系统 */
	ELFOSABI_NETBSD OSABI = 2 /* NetBSD */
	ELFOSABI_LINUX OSABI = 3 /* Linux */
	ELFOSABI_HURD OSABI = 4 /* Hurd */ ELFOSABI_HURD OSABI = 4 /* Hurd */
	ELFOSABI_86OPEN OSABI = 5 /* 86Open 通用 IA32 ABI */
	ELFOSABI_SOLARIS OSABI = 6 /* Solaris */
	ELFOSABI_AIX OSABI = 7 /* AIX */
	elfosabi_irix osabi = 8 /* irix */
	ELFOSABI_FREEBSD OSABI = 9 /* FreeBSD */
	elfosabi_tru64 osabi = 10 /* tru64 unix */
	ELFOSABI_MODESTO OSABI = 11 /* Novell Modesto */
	ELFOSABI_OPENBSD OSABI = 12 /* OpenBSD */
	ELFOSABI_OPENVMS OSABI = 13 /* Open VMS */
	ELFOSABI_NSK OSABI = 14 /* HP Non-Stop Kernel */
	ELFOSABI_AROS OSABI = 15 /* Amiga 研究操作系统 */
	ELFOSABI_FENIXOS OSABI = 16 /* FenixOS 高扩展性多核操作系统 */
	ELFOSABI_CLOUDABI OSABI = 17 /* Nuxi CloudABI */
	ELFOSABI_ARM OSABI = 97 /* ARM */
	ELFOSABI_STANDALONE OSABI = 255 /* 独立（嵌入式）应用程序 */
)
```

#### func (i OSABI) GoString() string

#### func (i OSABI) String() string

### type Prog

```go
type Prog struct {
  ProgHeader

  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

一个 Prog 表示 ELF 二进制文件中的一个 ELF 程序头。

#### func (p *Prog) Open() io.ReadSeeker

Open 返回一个新的 ReadSeeker，读取 ELF 程序体。

### type Prog32

```go
type Prog32 struct {
	Type uint32 /* 输入类型。*/
	Off uint32 /* 内容的文件偏移量。*/
	Vaddr uint32 /* 内存映像中的虚拟地址。*/
	Paddr uint32 /* 物理地址（未使用）。*/
	Filesz uint32 /* 文件内容大小。*/
	Memsz uint32 /* 内存中内容的大小。*/
	Flags uint32 /* 访问权限标志。*/
	Align uint32 /* 内存和文件中的对齐方式。*/
}
```

ELF32 程序头。

### type Prog64

```go
type Prog64 struct {
  Type uint32 /* 输入类型。*/
	Flags uint32 /* 访问权限标志。*/
	Off uint64 /* 内容的文件偏移量。*/
	Vaddr uint64 /* 内存映像中的虚拟地址。*/
	Paddr uint64 /* 物理地址（未使用）。*/
	Filesz uint64 /* 文件内容大小。*/
	Memsz uint64 /* 内存中内容的大小。*/
	Align uint64 /* 内存和文件中的对齐方式。*/
}
```

ELF64 程序头。

### type ProgFlag

```go
type ProgFlag uint32
```

Prog.Flag

```go
const (
	PF_X ProgFlag = 0x1 /* 可执行文件。*/
	PF_W ProgFlag = 0x2 /* 可写入。*/
	PF_R ProgFlag = 0x4 /* 可读。*/
	PF_MASKOS ProgFlag = 0x0ff00000 /* 特定于操作系统。*/
	PF_MASKPROC ProgFlag = 0xf0000000 /* 特定于处理器。*/
)
```

#### func (i ProgFlag) GoString() string

#### func (i ProgFlag) String() string

### type ProgHeader

```go
type ProgHeader struct {
  Type ProgType
  Flags ProgFlag
  Off uint64
  Vaddr uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}
```

一个 ProgHeader 表示一个 ELF 程序头。

### type ProgType

```go
type ProgType int
```

Prog.Type

```go
const (
	PT_NULL ProgType = 0 /* 未使用的条目。*/
	PT_LOAD ProgType = 1 /* 可加载段。*/
	PT_DYNAMIC ProgType = 2 /* 动态链接信息段。*/
	PT_INTERP ProgType = 3 /* 解释器的路径名。*/
	PT_NOTE ProgType = 4 /* 辅助信息。*/
	PT_SHLIB ProgType = 5 /* 保留（未使用）。*/
	PT_PHDR ProgType = 6 /* 程序头本身的位置。*/
	PT_TLS ProgType = 7 /* 线程本地存储段 */

	PT_LOOS ProgType = 0x60000000 /* 第一个操作系统专用。*/

	PT_GNU_EH_FRAME ProgType = 0x6474e550 /* 帧展开信息 */
	PT_GNU_STACK ProgType = 0x6474e551 /* 堆栈标志 */
	PT_GNU_RELRO ProgType = 0x6474e552 /* 重定位后只读 */
	PT_GNU_PROPERTY ProgType = 0x6474e553 /* GNU 属性 */
	PT_GNU_MBIND_LO ProgType = 0x6474e555 /* Mbind 段开始 */
	PT_GNU_MBIND_HI ProgType = 0x6474f554 /* 绑定段结束 */

	PT_PAX_FLAGS ProgType = 0x65041580 /* PAX 标志 */
	PT_OPENBSD_RANDOMIZE ProgType = 0x65a3dbe6 /* 随机数据 */
	PT_OPENBSD_WXNEEDED ProgType = 0x65a3dbe7 /* 违反 W^X */
	PT_OPENBSD_BOOTDATA ProgType = 0x65a41be6 /* 启动参数 */

	PT_SUNW_EH_FRAME ProgType = 0x6474e550 /* 帧展开信息 */
	PT_SUNWSTACK ProgType = 0x6ffffffb /* 堆栈段 */

	PT_HIOS ProgType = 0x6fffffff /* 最后操作系统专用。*/

	PT_LOPROC ProgType = 0x70000000 /* 第一个特定于处理器的类型。*/

	PT_ARM_ARCHEXT ProgType = 0x70000000 /* 体系结构兼容性 */
	PT_ARM_EXIDX ProgType = 0x70000001 /* 异常解卷表 */

	PT_AARCH64_ARCHEXT ProgType = 0x70000000 /* 体系结构兼容性 */
	PT_AARCH64_UNWIND ProgType = 0x70000001 /* 异常解卷表 */

	PT_MIPS_REGINFO ProgType = 0x70000000 /* 寄存器使用 */
	PT_MIPS_RTPROC ProgType = 0x70000001 /* 运行时程序 */
	PT_MIPS_OPTIONS ProgType = 0x70000002 /* 选项 */
	PT_MIPS_ABIFLAGS ProgType = 0x70000003 /* ABI 标志 */

	PT_S390_PGSTE ProgType = 0x70000000 /* 4k 页表大小 */

	PT_HIPROC ProgType = 0x7fffffff /* 最后一种特定于处理器的类型。*/
)
```

#### func (i ProgType) GoString() string

#### func (i ProgType) String() string

### type R_386

```go
type R_386 int
```

386 的迁移类型。

```go
const (
	R_386_NONE R_386 = 0 /* 没有重新定位。*/
	R_386_32 R_386 = 1 /* 添加符号值。*/
	R_386_PC32 R_386 = 2 /* 添加 PC 相关符号值。*/
	R_386_GOT32 R_386 = 3 /* 添加 PC 相对 GOT 偏移量。*/
	R_386_PLT32 R_386 = 4 /* 添加 PC 相对 PLT 偏移量。*/
	R_386_COPY R_386 = 5 /* 从共享对象复制数据。*/
	R_386_GLOB_DAT R_386 = 6 /* 将 GOT 条目设置为数据地址。*/
	R_386_JMP_SLOT R_386 = 7 /* 将 GOT 条目设置为代码地址。*/
	R_386_RELATIVE R_386 = 8 /* 添加共享对象的加载地址。*/
	R_386_GOTOFF R_386 = 9 /* 添加 GOT 相关符号地址。*/
	R_386_GOTPC R_386 = 10 /* 添加 PC 相关的 GOT 表地址。*/
	R_386_32PLT R_386 = 11
	R_386_TLS_TPOFF R_386 = 14 /* 静态 TLS 块中的负偏移 */
	R_386_TLS_IE R_386 = 15 /* 负静态 TLS 的 GOT 的绝对地址 */
	R_386_TLS_GOTIE R_386 = 16 /* 负静态 TLS 块的 GOT 条目 */
	R_386_TLS_LE R_386 = 17 /* 相对于静态 TLS 的负偏移 */
	R_386_TLS_GD R_386 = 18 /* GOT（index,off）对的 32 位偏移 */
	R_386_TLS_LDM R_386 = 19 /* 32 位偏移到 GOT（index,0）对 */
  R_386_16            R_386 = 20
	R_386_PC16          R_386 = 21
	R_386_8             R_386 = 22
	R_386_PC8           R_386 = 23
	R_386_TLS_GD_32 R_386 = 24 /* GOT（index,off）对的 32 位偏移 */
	R_386_TLS_GD_PUSH R_386 = 25 /* 用于 Sun ABI GD 序列的 pushl 指令 */
	R_386_TLS_GD_CALL R_386 = 26 /* Sun ABI GD 序列的调用指令 */
	R_386_TLS_GD_POP R_386 = 27 /* 用于 Sun ABI GD 序列的 popl 指令 */
	R_386_TLS_LDM_32 R_386 = 28 /* 32 位偏移到 GOT (index,zero) 对 */
	R_386_TLS_LDM_PUSH R_386 = 29 /* 用于 Sun ABI LD 序列的 pushl 指令 */
	R_386_TLS_LDM_CALL R_386 = 30 /* Sun ABI LD 序列的调用指令 */
	R_386_TLS_LDM_POP R_386 = 31 /* 用于 Sun ABI LD 序列的 popl 指令 */
	R_386_TLS_LDO_32 R_386 = 32 /* 从 TLS 块开始的 32 位偏移 */
	R_386_TLS_IE_32 R_386 = 33 /* 到 GOT 静态 TLS 偏移量条目的 32 位偏移 */
	R_386_TLS_LE_32 R_386 = 34 /* 静态 TLS 块内的 32 位偏移 */
	R_386_TLS_DTPMOD32 R_386 = 35 /* 包含 TLS 索引的 GOT 条目 */
	R_386_TLS_DTPOFF32 R_386 = 36 /* 包含 TLS 偏移量的 GOT 条目 */
	R_386_TLS_TPOFF32 R_386 = 37 /* -ve 静态 TLS 偏移量的 GOT 条目 */
	R_386_size32 R_386 = 38
	r_386_tls_gotdesc r_386 = 39
	r_386_tls_desc_call r_386 = 40
	r_386_tls_desc r_386 = 41
	r_386_irelative r_386 = 42
	r_386_got32x r_386 = 43
)
```

#### func (i R_386) GoString() string

#### func (i R_386) String() string

### type R_390 添加于1.7

```go
type R_390 int
```

s390x 处理器的重定位类型。

```go
const (
	R_390_NONE        R_390 = 0
	R_390_8           R_390 = 1
	R_390_12          R_390 = 2
	R_390_16          R_390 = 3
	R_390_32          R_390 = 4
	R_390_PC32        R_390 = 5
	R_390_GOT12       R_390 = 6
	R_390_GOT32       R_390 = 7
	R_390_PLT32       R_390 = 8
	R_390_COPY        R_390 = 9
	R_390_GLOB_DAT    R_390 = 10
	R_390_JMP_SLOT    R_390 = 11
	R_390_RELATIVE    R_390 = 12
	R_390_GOTOFF      R_390 = 13
	R_390_GOTPC       R_390 = 14
	R_390_GOT16       R_390 = 15
	R_390_PC16        R_390 = 16
	R_390_PC16DBL     R_390 = 17
	R_390_PLT16DBL    R_390 = 18
	R_390_PC32DBL     R_390 = 19
	R_390_PLT32DBL    R_390 = 20
	R_390_GOTPCDBL    R_390 = 21
	R_390_64          R_390 = 22
	R_390_PC64        R_390 = 23
	R_390_GOT64       R_390 = 24
	R_390_PLT64       R_390 = 25
	R_390_GOTENT      R_390 = 26
	R_390_GOTOFF16    R_390 = 27
	R_390_GOTOFF64    R_390 = 28
	R_390_GOTPLT12    R_390 = 29
	R_390_GOTPLT16    R_390 = 30
	R_390_GOTPLT32    R_390 = 31
	R_390_GOTPLT64    R_390 = 32
	R_390_GOTPLTENT   R_390 = 33
	R_390_GOTPLTOFF16 R_390 = 34
	R_390_GOTPLTOFF32 R_390 = 35
	R_390_GOTPLTOFF64 R_390 = 36
	R_390_TLS_LOAD    R_390 = 37
	R_390_TLS_GDCALL  R_390 = 38
	R_390_TLS_LDCALL  R_390 = 39
	R_390_TLS_GD32    R_390 = 40
	R_390_TLS_GD64    R_390 = 41
	R_390_TLS_GOTIE12 R_390 = 42
	R_390_TLS_GOTIE32 R_390 = 43
	R_390_TLS_GOTIE64 R_390 = 44
	R_390_TLS_LDM32   R_390 = 45
	R_390_TLS_LDM64   R_390 = 46
	R_390_TLS_IE32    R_390 = 47
	R_390_TLS_IE64    R_390 = 48
	R_390_TLS_IEENT   R_390 = 49
	R_390_TLS_LE32    R_390 = 50
	R_390_TLS_LE64    R_390 = 51
	R_390_TLS_LDO32   R_390 = 52
	R_390_TLS_LDO64   R_390 = 53
	R_390_TLS_DTPMOD  R_390 = 54
	R_390_TLS_DTPOFF  R_390 = 55
	R_390_TLS_TPOFF   R_390 = 56
	R_390_20          R_390 = 57
	R_390_GOT20       R_390 = 58
	R_390_GOTPLT20    R_390 = 59
	R_390_TLS_GOTIE20 R_390 = 60
)
```

### func R_AARCH64 添加于1.7

#### func (i R_390) GoString() string 添加于1.7

#### func (i R_390) String() string 添加于1.4

```go
type R_AARCH64 int
```

AArch64（又称 arm64）的重定位类型

```go
const (
	R_AARCH64_NONE                            R_AARCH64 = 0
	R_AARCH64_P32_ABS32                       R_AARCH64 = 1
	R_AARCH64_P32_ABS16                       R_AARCH64 = 2
	R_AARCH64_P32_PREL32                      R_AARCH64 = 3
	R_AARCH64_P32_PREL16                      R_AARCH64 = 4
	R_AARCH64_P32_MOVW_UABS_G0                R_AARCH64 = 5
	R_AARCH64_P32_MOVW_UABS_G0_NC             R_AARCH64 = 6
	R_AARCH64_P32_MOVW_UABS_G1                R_AARCH64 = 7
	R_AARCH64_P32_MOVW_SABS_G0                R_AARCH64 = 8
	R_AARCH64_P32_LD_PREL_LO19                R_AARCH64 = 9
	R_AARCH64_P32_ADR_PREL_LO21               R_AARCH64 = 10
	R_AARCH64_P32_ADR_PREL_PG_HI21            R_AARCH64 = 11
	R_AARCH64_P32_ADD_ABS_LO12_NC             R_AARCH64 = 12
	R_AARCH64_P32_LDST8_ABS_LO12_NC           R_AARCH64 = 13
	R_AARCH64_P32_LDST16_ABS_LO12_NC          R_AARCH64 = 14
	R_AARCH64_P32_LDST32_ABS_LO12_NC          R_AARCH64 = 15
	R_AARCH64_P32_LDST64_ABS_LO12_NC          R_AARCH64 = 16
	R_AARCH64_P32_LDST128_ABS_LO12_NC         R_AARCH64 = 17
	R_AARCH64_P32_TSTBR14                     R_AARCH64 = 18
	R_AARCH64_P32_CONDBR19                    R_AARCH64 = 19
	R_AARCH64_P32_JUMP26                      R_AARCH64 = 20
	R_AARCH64_P32_CALL26                      R_AARCH64 = 21
	R_AARCH64_P32_GOT_LD_PREL19               R_AARCH64 = 25
	R_AARCH64_P32_ADR_GOT_PAGE                R_AARCH64 = 26
	R_AARCH64_P32_LD32_GOT_LO12_NC            R_AARCH64 = 27
	R_AARCH64_P32_TLSGD_ADR_PAGE21            R_AARCH64 = 81
	R_AARCH64_P32_TLSGD_ADD_LO12_NC           R_AARCH64 = 82
	R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21   R_AARCH64 = 103
	R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC R_AARCH64 = 104
	R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19    R_AARCH64 = 105
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G1         R_AARCH64 = 106
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0         R_AARCH64 = 107
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC      R_AARCH64 = 108
	R_AARCH64_P32_TLSLE_ADD_TPREL_HI12        R_AARCH64 = 109
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12        R_AARCH64 = 110
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC     R_AARCH64 = 111
	R_AARCH64_P32_TLSDESC_LD_PREL19           R_AARCH64 = 122
	R_AARCH64_P32_TLSDESC_ADR_PREL21          R_AARCH64 = 123
	R_AARCH64_P32_TLSDESC_ADR_PAGE21          R_AARCH64 = 124
	R_AARCH64_P32_TLSDESC_LD32_LO12_NC        R_AARCH64 = 125
	R_AARCH64_P32_TLSDESC_ADD_LO12_NC         R_AARCH64 = 126
	R_AARCH64_P32_TLSDESC_CALL                R_AARCH64 = 127
	R_AARCH64_P32_COPY                        R_AARCH64 = 180
	R_AARCH64_P32_GLOB_DAT                    R_AARCH64 = 181
	R_AARCH64_P32_JUMP_SLOT                   R_AARCH64 = 182
	R_AARCH64_P32_RELATIVE                    R_AARCH64 = 183
	R_AARCH64_P32_TLS_DTPMOD                  R_AARCH64 = 184
	R_AARCH64_P32_TLS_DTPREL                  R_AARCH64 = 185
	R_AARCH64_P32_TLS_TPREL                   R_AARCH64 = 186
	R_AARCH64_P32_TLSDESC                     R_AARCH64 = 187
	R_AARCH64_P32_IRELATIVE                   R_AARCH64 = 188
	R_AARCH64_NULL                            R_AARCH64 = 256
	R_AARCH64_ABS64                           R_AARCH64 = 257
	R_AARCH64_ABS32                           R_AARCH64 = 258
	R_AARCH64_ABS16                           R_AARCH64 = 259
	R_AARCH64_PREL64                          R_AARCH64 = 260
	R_AARCH64_PREL32                          R_AARCH64 = 261
	R_AARCH64_PREL16                          R_AARCH64 = 262
	R_AARCH64_MOVW_UABS_G0                    R_AARCH64 = 263
	R_AARCH64_MOVW_UABS_G0_NC                 R_AARCH64 = 264
	R_AARCH64_MOVW_UABS_G1                    R_AARCH64 = 265
	R_AARCH64_MOVW_UABS_G1_NC                 R_AARCH64 = 266
	R_AARCH64_MOVW_UABS_G2                    R_AARCH64 = 267
	R_AARCH64_MOVW_UABS_G2_NC                 R_AARCH64 = 268
	R_AARCH64_MOVW_UABS_G3                    R_AARCH64 = 269
	R_AARCH64_MOVW_SABS_G0                    R_AARCH64 = 270
	R_AARCH64_MOVW_SABS_G1                    R_AARCH64 = 271
	R_AARCH64_MOVW_SABS_G2                    R_AARCH64 = 272
	R_AARCH64_LD_PREL_LO19                    R_AARCH64 = 273
	R_AARCH64_ADR_PREL_LO21                   R_AARCH64 = 274
	R_AARCH64_ADR_PREL_PG_HI21                R_AARCH64 = 275
	R_AARCH64_ADR_PREL_PG_HI21_NC             R_AARCH64 = 276
	R_AARCH64_ADD_ABS_LO12_NC                 R_AARCH64 = 277
	R_AARCH64_LDST8_ABS_LO12_NC               R_AARCH64 = 278
	R_AARCH64_TSTBR14                         R_AARCH64 = 279
	R_AARCH64_CONDBR19                        R_AARCH64 = 280
	R_AARCH64_JUMP26                          R_AARCH64 = 282
	R_AARCH64_CALL26                          R_AARCH64 = 283
	R_AARCH64_LDST16_ABS_LO12_NC              R_AARCH64 = 284
	R_AARCH64_LDST32_ABS_LO12_NC              R_AARCH64 = 285
	R_AARCH64_LDST64_ABS_LO12_NC              R_AARCH64 = 286
	R_AARCH64_LDST128_ABS_LO12_NC             R_AARCH64 = 299
	R_AARCH64_GOT_LD_PREL19                   R_AARCH64 = 309
	R_AARCH64_LD64_GOTOFF_LO15                R_AARCH64 = 310
	R_AARCH64_ADR_GOT_PAGE                    R_AARCH64 = 311
	R_AARCH64_LD64_GOT_LO12_NC                R_AARCH64 = 312
	R_AARCH64_LD64_GOTPAGE_LO15               R_AARCH64 = 313
	R_AARCH64_TLSGD_ADR_PREL21                R_AARCH64 = 512
	R_AARCH64_TLSGD_ADR_PAGE21                R_AARCH64 = 513
	R_AARCH64_TLSGD_ADD_LO12_NC               R_AARCH64 = 514
	R_AARCH64_TLSGD_MOVW_G1                   R_AARCH64 = 515
	R_AARCH64_TLSGD_MOVW_G0_NC                R_AARCH64 = 516
	R_AARCH64_TLSLD_ADR_PREL21                R_AARCH64 = 517
	R_AARCH64_TLSLD_ADR_PAGE21                R_AARCH64 = 518
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G1          R_AARCH64 = 539
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC       R_AARCH64 = 540
	R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21       R_AARCH64 = 541
	R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC     R_AARCH64 = 542
	R_AARCH64_TLSIE_LD_GOTTPREL_PREL19        R_AARCH64 = 543
	R_AARCH64_TLSLE_MOVW_TPREL_G2             R_AARCH64 = 544
	R_AARCH64_TLSLE_MOVW_TPREL_G1             R_AARCH64 = 545
	R_AARCH64_TLSLE_MOVW_TPREL_G1_NC          R_AARCH64 = 546
	R_AARCH64_TLSLE_MOVW_TPREL_G0             R_AARCH64 = 547
	R_AARCH64_TLSLE_MOVW_TPREL_G0_NC          R_AARCH64 = 548
	R_AARCH64_TLSLE_ADD_TPREL_HI12            R_AARCH64 = 549
	R_AARCH64_TLSLE_ADD_TPREL_LO12            R_AARCH64 = 550
	R_AARCH64_TLSLE_ADD_TPREL_LO12_NC         R_AARCH64 = 551
	R_AARCH64_TLSDESC_LD_PREL19               R_AARCH64 = 560
	R_AARCH64_TLSDESC_ADR_PREL21              R_AARCH64 = 561
	R_AARCH64_TLSDESC_ADR_PAGE21              R_AARCH64 = 562
	R_AARCH64_TLSDESC_LD64_LO12_NC            R_AARCH64 = 563
	R_AARCH64_TLSDESC_ADD_LO12_NC             R_AARCH64 = 564
	R_AARCH64_TLSDESC_OFF_G1                  R_AARCH64 = 565
	R_AARCH64_TLSDESC_OFF_G0_NC               R_AARCH64 = 566
	R_AARCH64_TLSDESC_LDR                     R_AARCH64 = 567
	R_AARCH64_TLSDESC_ADD                     R_AARCH64 = 568
	R_AARCH64_TLSDESC_CALL                    R_AARCH64 = 569
	R_AARCH64_TLSLE_LDST128_TPREL_LO12        R_AARCH64 = 570
	R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC     R_AARCH64 = 571
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12       R_AARCH64 = 572
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC    R_AARCH64 = 573
	R_AARCH64_COPY                            R_AARCH64 = 1024
	R_AARCH64_GLOB_DAT                        R_AARCH64 = 1025
	R_AARCH64_JUMP_SLOT                       R_AARCH64 = 1026
	R_AARCH64_RELATIVE                        R_AARCH64 = 1027
	R_AARCH64_TLS_DTPMOD64                    R_AARCH64 = 1028
	R_AARCH64_TLS_DTPREL64                    R_AARCH64 = 1029
	R_AARCH64_TLS_TPREL64                     R_AARCH64 = 1030
	R_AARCH64_TLSDESC                         R_AARCH64 = 1031
	R_AARCH64_IRELATIVE                       R_AARCH64 = 1032
)
```

#### func (i R_AARCH64) GoString() string 添加于1.4

#### func (i R_AARCH64) String() string 添加于1.4

### type R_ALPHA

```go
type R_ALPHA int
```

阿尔法的迁移类型。

```go
const (
	R_ALPHA_NONE R_ALPHA = 0 /* No reloc */
	R_ALPHA_REFLONG R_ALPHA = 1 /* 直接 32 位 */
	R_ALPHA_REFQUAD R_ALPHA = 2 /* 直接 64 位 */
	R_ALPHA_GPREL32 R_ALPHA = 3 /* GP 相对 32 位 */
	R_ALPHA_LITERAL R_ALPHA = 4 /* GP 相对 16 位/优化 */
	R_ALPHA_LITUSE R_ALPHA = 5 /* LITERAL 的优化提示 */
	R_ALPHA_GPDISP R_ALPHA = 6 /* 为 GP 添加位移 */
	R_ALPHA_BRADDR R_ALPHA = 7 /* PC+4 相对 23 位移位 */
	R_ALPHA_HINT R_ALPHA = 8 /* PC+4 相对 16 位移位 */
	R_ALPHA_SREL16 R_ALPHA = 9 /* PC 相对 16 位 */
	R_ALPHA_SREL32 R_ALPHA = 10 /* PC 相对 32 位 */
	R_ALPHA_SREL64 R_ALPHA = 11 /* PC 相对 64 位 */
	R_ALPHA_OP_PUSH R_ALPHA = 12 /* OP 堆栈推送 */
	R_ALPHA_OP_STORE R_ALPHA = 13 /* OP 堆栈弹出和存储 */
	R_ALPHA_OP_PSUB R_ALPHA = 14 /* OP 栈减法 */
	R_ALPHA_OP_PRSHIFT R_ALPHA = 15 /* OP 栈右移 */
	r_alpha_gpvalue r_alpha = 16
	r_alpha_gprelhigh r_alpha = 17
	r_alpha_gprellow r_alpha = 18
	r_alpha_immed_gp_16 r_alpha = 19
	r_alpha_immed_gp_hi32 r_alpha = 20
	r_alpha_immed_scn_hi32 r_alpha = 21
	r_alpha_immed_br_hi32 r_alpha = 22
	r_alpha_immed_lo32 r_alpha = 23
	R_ALPHA_COPY R_ALPHA = 24 /* 运行时复制符号 */
	R_ALPHA_GLOB_DAT R_ALPHA = 25 /* 创建 GOT 条目 */
	R_ALPHA_JMP_SLOT R_ALPHA = 26 /* 创建 PLT 条目 */
	R_ALPHA_RELATIVE R_ALPHA = 27 /* 按程序基数调整 */
)
```

#### func (i R_ALPHA) GoString() string

#### func (i R_ALPHA) String() string

### type R_ARM

```go
type R_ARM int
```

ARM 的迁移类型。

```go
const (
	R_ARM_NONE R_ARM = 0 /* 没有重新定位。*/
	R_ARM_PC24               R_ARM = 1
	R_ARM_ABS32              R_ARM = 2
	R_ARM_REL32              R_ARM = 3
	R_ARM_PC13               R_ARM = 4
	R_ARM_ABS16              R_ARM = 5
	R_ARM_ABS12              R_ARM = 6
	R_ARM_THM_ABS5           R_ARM = 7
	R_ARM_ABS8               R_ARM = 8
	R_ARM_SBREL32            R_ARM = 9
	R_ARM_THM_PC22           R_ARM = 10
	R_ARM_THM_PC8            R_ARM = 11
	R_ARM_AMP_VCALL9         R_ARM = 12
	R_ARM_SWI24              R_ARM = 13
	R_ARM_THM_SWI8           R_ARM = 14
	R_ARM_XPC25              R_ARM = 15
	R_ARM_THM_XPC22          R_ARM = 16
	R_ARM_TLS_DTPMOD32       R_ARM = 17
	R_ARM_TLS_DTPOFF32       R_ARM = 18
	R_ARM_TLS_TPOFF32        R_ARM = 19
	R_ARM_COPY R_ARM = 20 /* 从共享对象复制数据。*/
	R_ARM_GLOB_DAT R_ARM = 21 /* 将 GOT 条目设置为数据地址。*/
	R_ARM_JUMP_SLOT R_ARM = 22 /* 将 GOT 条目设置为代码地址。*/
	R_ARM_RELATIVE R_ARM = 23 /* 添加共享对象的加载地址。*/
	R_ARM_GOTOFF R_ARM = 24 /* 添加 GOT 相关符号地址。*/
	R_ARM_GOTPC R_ARM = 25 /* 添加 PC 相关的 GOT 表地址。*/
	R_ARM_GOT32 R_ARM = 26 /* 添加 PC 相对 GOT 偏移量。*/
	R_ARM_PLT32 R_ARM = 27 /* 添加相对于 PC 的 PLT 偏移量。*/
	R_ARM_CALL               R_ARM = 28
	R_ARM_JUMP24             R_ARM = 29
	R_ARM_THM_JUMP24         R_ARM = 30
	R_ARM_BASE_ABS           R_ARM = 31
	R_ARM_ALU_PCREL_7_0      R_ARM = 32
	R_ARM_ALU_PCREL_15_8     R_ARM = 33
	R_ARM_ALU_PCREL_23_15    R_ARM = 34
	R_ARM_LDR_SBREL_11_10_NC R_ARM = 35
	R_ARM_ALU_SBREL_19_12_NC R_ARM = 36
	R_ARM_ALU_SBREL_27_20_CK R_ARM = 37
	R_ARM_TARGET1            R_ARM = 38
	R_ARM_SBREL31            R_ARM = 39
	R_ARM_V4BX               R_ARM = 40
	R_ARM_TARGET2            R_ARM = 41
	R_ARM_PREL31             R_ARM = 42
	R_ARM_MOVW_ABS_NC        R_ARM = 43
	R_ARM_MOVT_ABS           R_ARM = 44
	R_ARM_MOVW_PREL_NC       R_ARM = 45
	R_ARM_MOVT_PREL          R_ARM = 46
	R_ARM_THM_MOVW_ABS_NC    R_ARM = 47
	R_ARM_THM_MOVT_ABS       R_ARM = 48
	R_ARM_THM_MOVW_PREL_NC   R_ARM = 49
	R_ARM_THM_MOVT_PREL      R_ARM = 50
	R_ARM_THM_JUMP19         R_ARM = 51
	R_ARM_THM_JUMP6          R_ARM = 52
	R_ARM_THM_ALU_PREL_11_0  R_ARM = 53
	R_ARM_THM_PC12           R_ARM = 54
	R_ARM_ABS32_NOI          R_ARM = 55
	R_ARM_REL32_NOI          R_ARM = 56
	R_ARM_ALU_PC_G0_NC       R_ARM = 57
	R_ARM_ALU_PC_G0          R_ARM = 58
	R_ARM_ALU_PC_G1_NC       R_ARM = 59
	R_ARM_ALU_PC_G1          R_ARM = 60
	R_ARM_ALU_PC_G2          R_ARM = 61
	R_ARM_LDR_PC_G1          R_ARM = 62
	R_ARM_LDR_PC_G2          R_ARM = 63
	R_ARM_LDRS_PC_G0         R_ARM = 64
	R_ARM_LDRS_PC_G1         R_ARM = 65
	R_ARM_LDRS_PC_G2         R_ARM = 66
	R_ARM_LDC_PC_G0          R_ARM = 67
	R_ARM_LDC_PC_G1          R_ARM = 68
	R_ARM_LDC_PC_G2          R_ARM = 69
	R_ARM_ALU_SB_G0_NC       R_ARM = 70
	R_ARM_ALU_SB_G0          R_ARM = 71
	R_ARM_ALU_SB_G1_NC       R_ARM = 72
	R_ARM_ALU_SB_G1          R_ARM = 73
	R_ARM_ALU_SB_G2          R_ARM = 74
	R_ARM_LDR_SB_G0          R_ARM = 75
	R_ARM_LDR_SB_G1          R_ARM = 76
	R_ARM_LDR_SB_G2          R_ARM = 77
	R_ARM_LDRS_SB_G0         R_ARM = 78
	R_ARM_LDRS_SB_G1         R_ARM = 79
	R_ARM_LDRS_SB_G2         R_ARM = 80
	R_ARM_LDC_SB_G0          R_ARM = 81
	R_ARM_LDC_SB_G1          R_ARM = 82
	R_ARM_LDC_SB_G2          R_ARM = 83
	R_ARM_MOVW_BREL_NC       R_ARM = 84
	R_ARM_MOVT_BREL          R_ARM = 85
	R_ARM_MOVW_BREL          R_ARM = 86
	R_ARM_THM_MOVW_BREL_NC   R_ARM = 87
	R_ARM_THM_MOVT_BREL      R_ARM = 88
	R_ARM_THM_MOVW_BREL      R_ARM = 89
	R_ARM_TLS_GOTDESC        R_ARM = 90
	R_ARM_TLS_CALL           R_ARM = 91
	R_ARM_TLS_DESCSEQ        R_ARM = 92
	R_ARM_THM_TLS_CALL       R_ARM = 93
	R_ARM_PLT32_ABS          R_ARM = 94
	R_ARM_GOT_ABS            R_ARM = 95
	R_ARM_GOT_PREL           R_ARM = 96
	R_ARM_GOT_BREL12         R_ARM = 97
	R_ARM_GOTOFF12           R_ARM = 98
	R_ARM_GOTRELAX           R_ARM = 99
	R_ARM_GNU_VTENTRY        R_ARM = 100
	R_ARM_GNU_VTINHERIT      R_ARM = 101
	R_ARM_THM_JUMP11         R_ARM = 102
	R_ARM_THM_JUMP8          R_ARM = 103
	R_ARM_TLS_GD32           R_ARM = 104
	R_ARM_TLS_LDM32          R_ARM = 105
	R_ARM_TLS_LDO32          R_ARM = 106
	R_ARM_TLS_IE32           R_ARM = 107
	R_ARM_TLS_LE32           R_ARM = 108
	R_ARM_TLS_LDO12          R_ARM = 109
	R_ARM_TLS_LE12           R_ARM = 110
	R_ARM_TLS_IE12GP         R_ARM = 111
	R_ARM_PRIVATE_0          R_ARM = 112
	R_ARM_PRIVATE_1          R_ARM = 113
	R_ARM_PRIVATE_2          R_ARM = 114
	R_ARM_PRIVATE_3          R_ARM = 115
	R_ARM_PRIVATE_4          R_ARM = 116
	R_ARM_PRIVATE_5          R_ARM = 117
	R_ARM_PRIVATE_6          R_ARM = 118
	R_ARM_PRIVATE_7          R_ARM = 119
	R_ARM_PRIVATE_8          R_ARM = 120
	R_ARM_PRIVATE_9          R_ARM = 121
	R_ARM_PRIVATE_10         R_ARM = 122
	R_ARM_PRIVATE_11         R_ARM = 123
	R_ARM_PRIVATE_12         R_ARM = 124
	R_ARM_PRIVATE_13         R_ARM = 125
	R_ARM_PRIVATE_14         R_ARM = 126
	R_ARM_PRIVATE_15         R_ARM = 127
	R_ARM_ME_TOO             R_ARM = 128
	R_ARM_THM_TLS_DESCSEQ16  R_ARM = 129
	R_ARM_THM_TLS_DESCSEQ32  R_ARM = 130
	R_ARM_THM_GOT_BREL12     R_ARM = 131
	R_ARM_THM_ALU_ABS_G0_NC  R_ARM = 132
	R_ARM_THM_ALU_ABS_G1_NC  R_ARM = 133
	R_ARM_THM_ALU_ABS_G2_NC  R_ARM = 134
	R_ARM_THM_ALU_ABS_G3     R_ARM = 135
	R_ARM_IRELATIVE          R_ARM = 160
	R_ARM_RXPC25             R_ARM = 249
	R_ARM_RSBREL32           R_ARM = 250
	R_ARM_THM_RPC22          R_ARM = 251
	R_ARM_RREL32             R_ARM = 252
	R_ARM_RABS32             R_ARM = 253
	R_ARM_RPC24              R_ARM = 254
	R_ARM_RBASE              R_ARM = 255
)
```

#### func (i R_ARM) GoString() string

#### func (i R_ARM) String() string

### type R_LARCH 添加于1.19

```go
type R_LARCH int
```

LoongArch 的迁移类型。

```go
const (
	R_LARCH_NONE                       R_LARCH = 0
	R_LARCH_32                         R_LARCH = 1
	R_LARCH_64                         R_LARCH = 2
	R_LARCH_RELATIVE                   R_LARCH = 3
	R_LARCH_COPY                       R_LARCH = 4
	R_LARCH_JUMP_SLOT                  R_LARCH = 5
	R_LARCH_TLS_DTPMOD32               R_LARCH = 6
	R_LARCH_TLS_DTPMOD64               R_LARCH = 7
	R_LARCH_TLS_DTPREL32               R_LARCH = 8
	R_LARCH_TLS_DTPREL64               R_LARCH = 9
	R_LARCH_TLS_TPREL32                R_LARCH = 10
	R_LARCH_TLS_TPREL64                R_LARCH = 11
	R_LARCH_IRELATIVE                  R_LARCH = 12
	R_LARCH_MARK_LA                    R_LARCH = 20
	R_LARCH_MARK_PCREL                 R_LARCH = 21
	R_LARCH_SOP_PUSH_PCREL             R_LARCH = 22
	R_LARCH_SOP_PUSH_ABSOLUTE          R_LARCH = 23
	R_LARCH_SOP_PUSH_DUP               R_LARCH = 24
	R_LARCH_SOP_PUSH_GPREL             R_LARCH = 25
	R_LARCH_SOP_PUSH_TLS_TPREL         R_LARCH = 26
	R_LARCH_SOP_PUSH_TLS_GOT           R_LARCH = 27
	R_LARCH_SOP_PUSH_TLS_GD            R_LARCH = 28
	R_LARCH_SOP_PUSH_PLT_PCREL         R_LARCH = 29
	R_LARCH_SOP_ASSERT                 R_LARCH = 30
	R_LARCH_SOP_NOT                    R_LARCH = 31
	R_LARCH_SOP_SUB                    R_LARCH = 32
	R_LARCH_SOP_SL                     R_LARCH = 33
	R_LARCH_SOP_SR                     R_LARCH = 34
	R_LARCH_SOP_ADD                    R_LARCH = 35
	R_LARCH_SOP_AND                    R_LARCH = 36
	R_LARCH_SOP_IF_ELSE                R_LARCH = 37
	R_LARCH_SOP_POP_32_S_10_5          R_LARCH = 38
	R_LARCH_SOP_POP_32_U_10_12         R_LARCH = 39
	R_LARCH_SOP_POP_32_S_10_12         R_LARCH = 40
	R_LARCH_SOP_POP_32_S_10_16         R_LARCH = 41
	R_LARCH_SOP_POP_32_S_10_16_S2      R_LARCH = 42
	R_LARCH_SOP_POP_32_S_5_20          R_LARCH = 43
	R_LARCH_SOP_POP_32_S_0_5_10_16_S2  R_LARCH = 44
	R_LARCH_SOP_POP_32_S_0_10_10_16_S2 R_LARCH = 45
	R_LARCH_SOP_POP_32_U               R_LARCH = 46
	R_LARCH_ADD8                       R_LARCH = 47
	R_LARCH_ADD16                      R_LARCH = 48
	R_LARCH_ADD24                      R_LARCH = 49
	R_LARCH_ADD32                      R_LARCH = 50
	R_LARCH_ADD64                      R_LARCH = 51
	R_LARCH_SUB8                       R_LARCH = 52
	R_LARCH_SUB16                      R_LARCH = 53
	R_LARCH_SUB24                      R_LARCH = 54
	R_LARCH_SUB32                      R_LARCH = 55
	R_LARCH_SUB64                      R_LARCH = 56
	R_LARCH_GNU_VTINHERIT              R_LARCH = 57
	R_LARCH_GNU_VTENTRY                R_LARCH = 58
	R_LARCH_B16                        R_LARCH = 64
	R_LARCH_B21                        R_LARCH = 65
	R_LARCH_B26                        R_LARCH = 66
	R_LARCH_ABS_HI20                   R_LARCH = 67
	R_LARCH_ABS_LO12                   R_LARCH = 68
	R_LARCH_ABS64_LO20                 R_LARCH = 69
	R_LARCH_ABS64_HI12                 R_LARCH = 70
	R_LARCH_PCALA_HI20                 R_LARCH = 71
	R_LARCH_PCALA_LO12                 R_LARCH = 72
	R_LARCH_PCALA64_LO20               R_LARCH = 73
	R_LARCH_PCALA64_HI12               R_LARCH = 74
	R_LARCH_GOT_PC_HI20                R_LARCH = 75
	R_LARCH_GOT_PC_LO12                R_LARCH = 76
	R_LARCH_GOT64_PC_LO20              R_LARCH = 77
	R_LARCH_GOT64_PC_HI12              R_LARCH = 78
	R_LARCH_GOT_HI20                   R_LARCH = 79
	R_LARCH_GOT_LO12                   R_LARCH = 80
	R_LARCH_GOT64_LO20                 R_LARCH = 81
	R_LARCH_GOT64_HI12                 R_LARCH = 82
	R_LARCH_TLS_LE_HI20                R_LARCH = 83
	R_LARCH_TLS_LE_LO12                R_LARCH = 84
	R_LARCH_TLS_LE64_LO20              R_LARCH = 85
	R_LARCH_TLS_LE64_HI12              R_LARCH = 86
	R_LARCH_TLS_IE_PC_HI20             R_LARCH = 87
	R_LARCH_TLS_IE_PC_LO12             R_LARCH = 88
	R_LARCH_TLS_IE64_PC_LO20           R_LARCH = 89
	R_LARCH_TLS_IE64_PC_HI12           R_LARCH = 90
	R_LARCH_TLS_IE_HI20                R_LARCH = 91
	R_LARCH_TLS_IE_LO12                R_LARCH = 92
	R_LARCH_TLS_IE64_LO20              R_LARCH = 93
	R_LARCH_TLS_IE64_HI12              R_LARCH = 94
	R_LARCH_TLS_LD_PC_HI20             R_LARCH = 95
	R_LARCH_TLS_LD_HI20                R_LARCH = 96
	R_LARCH_TLS_GD_PC_HI20             R_LARCH = 97
	R_LARCH_TLS_GD_HI20                R_LARCH = 98
	R_LARCH_32_PCREL                   R_LARCH = 99
	R_LARCH_RELAX                      R_LARCH = 100
)
```

#### func (i R_LARCH) GoString() string 添加于1.19

#### func (i R_LARCH) String() string 添加于1.19

### type R_MIPS 添加于1.6

```go
type R_MIPS int
```

MIPS 的迁移类型。

```go
const (
	r_mips_none r_mips = 0
	r_mips_16 r_mips = 1
	r_mips_32 r_mips = 2
	r_mips_rel32 r_mips = 3
	r_mips_26 r_mips = 4
	R_MIPS_HI16 R_MIPS = 5 /* 符号值的高 16 位 */
	R_MIPS_LO16 R_MIPS = 6 /* 低 16 位符号值 */
	R_MIPS_GPREL16 R_MIPS = 7 /* GP 相对引用 */
	R_MIPS_LITERAL R_MIPS = 8 /* 文字部分引用 */
	R_MIPS_GOT16 R_MIPS = 9 /* 对全局偏移表的引用 */
	R_MIPS_PC16 R_MIPS = 10 /* 16 位 PC 相对引用 */
	R_MIPS_CALL16 R_MIPS = 11 /* 通过 glbl 偏移 tbl 调用 16 位 */
	R_MIPS_GPREL32       R_MIPS = 12
	R_MIPS_SHIFT5        R_MIPS = 16
	R_MIPS_SHIFT6        R_MIPS = 17
	R_MIPS_64            R_MIPS = 18
	R_MIPS_GOT_DISP      R_MIPS = 19
	R_MIPS_GOT_PAGE      R_MIPS = 20
	R_MIPS_GOT_OFST      R_MIPS = 21
	R_MIPS_GOT_HI16      R_MIPS = 22
	R_MIPS_GOT_LO16      R_MIPS = 23
	R_MIPS_SUB           R_MIPS = 24
	R_MIPS_INSERT_A      R_MIPS = 25
	R_MIPS_INSERT_B      R_MIPS = 26
	R_MIPS_DELETE        R_MIPS = 27
	R_MIPS_HIGHER        R_MIPS = 28
	R_MIPS_HIGHEST       R_MIPS = 29
	R_MIPS_CALL_HI16     R_MIPS = 30
	R_MIPS_CALL_LO16     R_MIPS = 31
	R_MIPS_SCN_DISP      R_MIPS = 32
	R_MIPS_REL16         R_MIPS = 33
	R_MIPS_ADD_IMMEDIATE R_MIPS = 34
	R_MIPS_PJUMP         R_MIPS = 35
	R_MIPS_RELGOT        R_MIPS = 36
	R_MIPS_JALR          R_MIPS = 37

	R_MIPS_TLS_DTPMOD32 R_MIPS = 38 /* 模块编号 32 位 */
	R_MIPS_TLS_DTPREL32 R_MIPS = 39 /* 模块相对偏移 32 位 */
	R_MIPS_TLS_DTPMOD64 R_MIPS = 40 /* 模块号 64 位 */
	R_MIPS_TLS_DTPREL64 R_MIPS = 41 /* 模块相对偏移 64 位 */
	R_MIPS_TLS_GD R_MIPS = 42 /* GD 的 16 位 GOT 偏移量 */
	R_MIPS_TLS_LDM R_MIPS = 43 /* LDM 的 16 位 GOT 偏移量 */
	R_MIPS_TLS_DTPREL_HI16 R_MIPS = 44 /* 模块相关偏移，高 16 位 */
	R_MIPS_TLS_DTPREL_LO16 R_MIPS = 45 /* 模块相关偏移，低 16 位 */
	R_MIPS_TLS_GOTTPREL R_MIPS = 46 /* 16 位 IE 的 GOT 偏移量 */
	R_MIPS_TLS_TPREL32 R_MIPS = 47 /* TP 相关偏移，32 位 */
	R_MIPS_TLS_TPREL64 R_MIPS = 48 /* TP 相对偏移，64 位 */
	R_MIPS_TLS_TPREL_HI16 R_MIPS = 49 /* TP 相对偏移，高 16 位 */
	R_MIPS_TLS_TPREL_LO16 R_MIPS = 50 /* TP 相对偏移，低 16 位 */
)
```

#### func (i R_MIPS) GoString() string 添加于1.6

#### func (i R_MIPS) String() string 添加于1.6

### type R_PPC

```go
type R_PPC int
```

PowerPC 的重定位类型

在 ELF 标准中，R_PPC 和 R_PPC64 共享的值以 R_POWERPC_ 为前缀。对于 R_PPC 类型，相关的共享重定位已用前缀 R_PPC_ 重新命名。原始名称以注释的形式跟在值后面。

```go
const (
	R_PPC_NONE            R_PPC = 0  // R_POWERPC_NONE
	R_PPC_ADDR32          R_PPC = 1  // R_POWERPC_ADDR32
	R_PPC_ADDR24          R_PPC = 2  // R_POWERPC_ADDR24
	R_PPC_ADDR16          R_PPC = 3  // R_POWERPC_ADDR16
	R_PPC_ADDR16_LO       R_PPC = 4  // R_POWERPC_ADDR16_LO
	R_PPC_ADDR16_HI       R_PPC = 5  // R_POWERPC_ADDR16_HI
	R_PPC_ADDR16_HA       R_PPC = 6  // R_POWERPC_ADDR16_HA
	R_PPC_ADDR14          R_PPC = 7  // R_POWERPC_ADDR14
	R_PPC_ADDR14_BRTAKEN  R_PPC = 8  // R_POWERPC_ADDR14_BRTAKEN
	R_PPC_ADDR14_BRNTAKEN R_PPC = 9  // R_POWERPC_ADDR14_BRNTAKEN
	R_PPC_REL24           R_PPC = 10 // R_POWERPC_REL24
	R_PPC_REL14           R_PPC = 11 // R_POWERPC_REL14
	R_PPC_REL14_BRTAKEN   R_PPC = 12 // R_POWERPC_REL14_BRTAKEN
	R_PPC_REL14_BRNTAKEN  R_PPC = 13 // R_POWERPC_REL14_BRNTAKEN
	R_PPC_GOT16           R_PPC = 14 // R_POWERPC_GOT16
	R_PPC_GOT16_LO        R_PPC = 15 // R_POWERPC_GOT16_LO
	R_PPC_GOT16_HI        R_PPC = 16 // R_POWERPC_GOT16_HI
	R_PPC_GOT16_HA        R_PPC = 17 // R_POWERPC_GOT16_HA
	R_PPC_PLTREL24        R_PPC = 18
	R_PPC_COPY            R_PPC = 19 // R_POWERPC_COPY
	R_PPC_GLOB_DAT        R_PPC = 20 // R_POWERPC_GLOB_DAT
	R_PPC_JMP_SLOT        R_PPC = 21 // R_POWERPC_JMP_SLOT
	R_PPC_RELATIVE        R_PPC = 22 // R_POWERPC_RELATIVE
	R_PPC_LOCAL24PC       R_PPC = 23
	R_PPC_UADDR32         R_PPC = 24 // R_POWERPC_UADDR32
	R_PPC_UADDR16         R_PPC = 25 // R_POWERPC_UADDR16
	R_PPC_REL32           R_PPC = 26 // R_POWERPC_REL32
	R_PPC_PLT32           R_PPC = 27 // R_POWERPC_PLT32
	R_PPC_PLTREL32        R_PPC = 28 // R_POWERPC_PLTREL32
	R_PPC_PLT16_LO        R_PPC = 29 // R_POWERPC_PLT16_LO
	R_PPC_PLT16_HI        R_PPC = 30 // R_POWERPC_PLT16_HI
	R_PPC_PLT16_HA        R_PPC = 31 // R_POWERPC_PLT16_HA
	R_PPC_SDAREL16        R_PPC = 32
	R_PPC_SECTOFF         R_PPC = 33 // R_POWERPC_SECTOFF
	R_PPC_SECTOFF_LO      R_PPC = 34 // R_POWERPC_SECTOFF_LO
	R_PPC_SECTOFF_HI      R_PPC = 35 // R_POWERPC_SECTOFF_HI
	R_PPC_SECTOFF_HA      R_PPC = 36 // R_POWERPC_SECTOFF_HA
	R_PPC_TLS             R_PPC = 67 // R_POWERPC_TLS
	R_PPC_DTPMOD32        R_PPC = 68 // R_POWERPC_DTPMOD32
	R_PPC_TPREL16         R_PPC = 69 // R_POWERPC_TPREL16
	R_PPC_TPREL16_LO      R_PPC = 70 // R_POWERPC_TPREL16_LO
	R_PPC_TPREL16_HI      R_PPC = 71 // R_POWERPC_TPREL16_HI
	R_PPC_TPREL16_HA      R_PPC = 72 // R_POWERPC_TPREL16_HA
	R_PPC_TPREL32         R_PPC = 73 // R_POWERPC_TPREL32
	R_PPC_DTPREL16        R_PPC = 74 // R_POWERPC_DTPREL16
	R_PPC_DTPREL16_LO     R_PPC = 75 // R_POWERPC_DTPREL16_LO
	R_PPC_DTPREL16_HI     R_PPC = 76 // R_POWERPC_DTPREL16_HI
	R_PPC_DTPREL16_HA     R_PPC = 77 // R_POWERPC_DTPREL16_HA
	R_PPC_DTPREL32        R_PPC = 78 // R_POWERPC_DTPREL32
	R_PPC_GOT_TLSGD16     R_PPC = 79 // R_POWERPC_GOT_TLSGD16
	R_PPC_GOT_TLSGD16_LO  R_PPC = 80 // R_POWERPC_GOT_TLSGD16_LO
	R_PPC_GOT_TLSGD16_HI  R_PPC = 81 // R_POWERPC_GOT_TLSGD16_HI
	R_PPC_GOT_TLSGD16_HA  R_PPC = 82 // R_POWERPC_GOT_TLSGD16_HA
	R_PPC_GOT_TLSLD16     R_PPC = 83 // R_POWERPC_GOT_TLSLD16
	R_PPC_GOT_TLSLD16_LO  R_PPC = 84 // R_POWERPC_GOT_TLSLD16_LO
	R_PPC_GOT_TLSLD16_HI  R_PPC = 85 // R_POWERPC_GOT_TLSLD16_HI
	R_PPC_GOT_TLSLD16_HA  R_PPC = 86 // R_POWERPC_GOT_TLSLD16_HA
	R_PPC_GOT_TPREL16     R_PPC = 87 // R_POWERPC_GOT_TPREL16
	R_PPC_GOT_TPREL16_LO  R_PPC = 88 // R_POWERPC_GOT_TPREL16_LO
	R_PPC_GOT_TPREL16_HI  R_PPC = 89 // R_POWERPC_GOT_TPREL16_HI
	R_PPC_GOT_TPREL16_HA  R_PPC = 90 // R_POWERPC_GOT_TPREL16_HA
	R_PPC_EMB_NADDR32     R_PPC = 101
	R_PPC_EMB_NADDR16     R_PPC = 102
	R_PPC_EMB_NADDR16_LO  R_PPC = 103
	R_PPC_EMB_NADDR16_HI  R_PPC = 104
	R_PPC_EMB_NADDR16_HA  R_PPC = 105
	R_PPC_EMB_SDAI16      R_PPC = 106
	R_PPC_EMB_SDA2I16     R_PPC = 107
	R_PPC_EMB_SDA2REL     R_PPC = 108
	R_PPC_EMB_SDA21       R_PPC = 109
	R_PPC_EMB_MRKREF      R_PPC = 110
	R_PPC_EMB_RELSEC16    R_PPC = 111
	R_PPC_EMB_RELST_LO    R_PPC = 112
	R_PPC_EMB_RELST_HI    R_PPC = 113
	R_PPC_EMB_RELST_HA    R_PPC = 114
	R_PPC_EMB_BIT_FLD     R_PPC = 115
	R_PPC_EMB_RELSDA      R_PPC = 116
)
```

#### func (i R_PPC) GoString() string

#### func (i R_PPC) String() string

### type R_PPC64 添加于1.5

```go
type R_PPC64 int
```

64 位 PowerPC 或 Power 架构处理器的重定位类型。

在 ELF 标准中，R_PPC 和 R_PPC64 共享的值均以 R_POWERPC_ 作为前缀。对于 R_PPC64 类型，相关的共享重定位已用前缀 R_PPC64_ 重新命名。原始名称跟在注释值后面。

```go
const (
	R_PPC64_NONE               R_PPC64 = 0  // R_POWERPC_NONE
	R_PPC64_ADDR32             R_PPC64 = 1  // R_POWERPC_ADDR32
	R_PPC64_ADDR24             R_PPC64 = 2  // R_POWERPC_ADDR24
	R_PPC64_ADDR16             R_PPC64 = 3  // R_POWERPC_ADDR16
	R_PPC64_ADDR16_LO          R_PPC64 = 4  // R_POWERPC_ADDR16_LO
	R_PPC64_ADDR16_HI          R_PPC64 = 5  // R_POWERPC_ADDR16_HI
	R_PPC64_ADDR16_HA          R_PPC64 = 6  // R_POWERPC_ADDR16_HA
	R_PPC64_ADDR14             R_PPC64 = 7  // R_POWERPC_ADDR14
	R_PPC64_ADDR14_BRTAKEN     R_PPC64 = 8  // R_POWERPC_ADDR14_BRTAKEN
	R_PPC64_ADDR14_BRNTAKEN    R_PPC64 = 9  // R_POWERPC_ADDR14_BRNTAKEN
	R_PPC64_REL24              R_PPC64 = 10 // R_POWERPC_REL24
	R_PPC64_REL14              R_PPC64 = 11 // R_POWERPC_REL14
	R_PPC64_REL14_BRTAKEN      R_PPC64 = 12 // R_POWERPC_REL14_BRTAKEN
	R_PPC64_REL14_BRNTAKEN     R_PPC64 = 13 // R_POWERPC_REL14_BRNTAKEN
	R_PPC64_GOT16              R_PPC64 = 14 // R_POWERPC_GOT16
	R_PPC64_GOT16_LO           R_PPC64 = 15 // R_POWERPC_GOT16_LO
	R_PPC64_GOT16_HI           R_PPC64 = 16 // R_POWERPC_GOT16_HI
	R_PPC64_GOT16_HA           R_PPC64 = 17 // R_POWERPC_GOT16_HA
	R_PPC64_COPY               R_PPC64 = 19 // R_POWERPC_COPY
	R_PPC64_GLOB_DAT           R_PPC64 = 20 // R_POWERPC_GLOB_DAT
	R_PPC64_JMP_SLOT           R_PPC64 = 21 // R_POWERPC_JMP_SLOT
	R_PPC64_RELATIVE           R_PPC64 = 22 // R_POWERPC_RELATIVE
	R_PPC64_UADDR32            R_PPC64 = 24 // R_POWERPC_UADDR32
	R_PPC64_UADDR16            R_PPC64 = 25 // R_POWERPC_UADDR16
	R_PPC64_REL32              R_PPC64 = 26 // R_POWERPC_REL32
	R_PPC64_PLT32              R_PPC64 = 27 // R_POWERPC_PLT32
	R_PPC64_PLTREL32           R_PPC64 = 28 // R_POWERPC_PLTREL32
	R_PPC64_PLT16_LO           R_PPC64 = 29 // R_POWERPC_PLT16_LO
	R_PPC64_PLT16_HI           R_PPC64 = 30 // R_POWERPC_PLT16_HI
	R_PPC64_PLT16_HA           R_PPC64 = 31 // R_POWERPC_PLT16_HA
	R_PPC64_SECTOFF            R_PPC64 = 33 // R_POWERPC_SECTOFF
	R_PPC64_SECTOFF_LO         R_PPC64 = 34 // R_POWERPC_SECTOFF_LO
	R_PPC64_SECTOFF_HI         R_PPC64 = 35 // R_POWERPC_SECTOFF_HI
	R_PPC64_SECTOFF_HA         R_PPC64 = 36 // R_POWERPC_SECTOFF_HA
	R_PPC64_REL30              R_PPC64 = 37 // R_POWERPC_ADDR30
	R_PPC64_ADDR64             R_PPC64 = 38
	R_PPC64_ADDR16_HIGHER      R_PPC64 = 39
	R_PPC64_ADDR16_HIGHERA     R_PPC64 = 40
	R_PPC64_ADDR16_HIGHEST     R_PPC64 = 41
	R_PPC64_ADDR16_HIGHESTA    R_PPC64 = 42
	R_PPC64_UADDR64            R_PPC64 = 43
	R_PPC64_REL64              R_PPC64 = 44
	R_PPC64_PLT64              R_PPC64 = 45
	R_PPC64_PLTREL64           R_PPC64 = 46
	R_PPC64_TOC16              R_PPC64 = 47
	R_PPC64_TOC16_LO           R_PPC64 = 48
	R_PPC64_TOC16_HI           R_PPC64 = 49
	R_PPC64_TOC16_HA           R_PPC64 = 50
	R_PPC64_TOC                R_PPC64 = 51
	R_PPC64_PLTGOT16           R_PPC64 = 52
	R_PPC64_PLTGOT16_LO        R_PPC64 = 53
	R_PPC64_PLTGOT16_HI        R_PPC64 = 54
	R_PPC64_PLTGOT16_HA        R_PPC64 = 55
	R_PPC64_ADDR16_DS          R_PPC64 = 56
	R_PPC64_ADDR16_LO_DS       R_PPC64 = 57
	R_PPC64_GOT16_DS           R_PPC64 = 58
	R_PPC64_GOT16_LO_DS        R_PPC64 = 59
	R_PPC64_PLT16_LO_DS        R_PPC64 = 60
	R_PPC64_SECTOFF_DS         R_PPC64 = 61
	R_PPC64_SECTOFF_LO_DS      R_PPC64 = 62
	R_PPC64_TOC16_DS           R_PPC64 = 63
	R_PPC64_TOC16_LO_DS        R_PPC64 = 64
	R_PPC64_PLTGOT16_DS        R_PPC64 = 65
	R_PPC64_PLTGOT_LO_DS       R_PPC64 = 66
	R_PPC64_TLS                R_PPC64 = 67 // R_POWERPC_TLS
	R_PPC64_DTPMOD64           R_PPC64 = 68 // R_POWERPC_DTPMOD64
	R_PPC64_TPREL16            R_PPC64 = 69 // R_POWERPC_TPREL16
	R_PPC64_TPREL16_LO         R_PPC64 = 70 // R_POWERPC_TPREL16_LO
	R_PPC64_TPREL16_HI         R_PPC64 = 71 // R_POWERPC_TPREL16_HI
	R_PPC64_TPREL16_HA         R_PPC64 = 72 // R_POWERPC_TPREL16_HA
	R_PPC64_TPREL64            R_PPC64 = 73 // R_POWERPC_TPREL64
	R_PPC64_DTPREL16           R_PPC64 = 74 // R_POWERPC_DTPREL16
	R_PPC64_DTPREL16_LO        R_PPC64 = 75 // R_POWERPC_DTPREL16_LO
	R_PPC64_DTPREL16_HI        R_PPC64 = 76 // R_POWERPC_DTPREL16_HI
	R_PPC64_DTPREL16_HA        R_PPC64 = 77 // R_POWERPC_DTPREL16_HA
	R_PPC64_DTPREL64           R_PPC64 = 78 // R_POWERPC_DTPREL64
	R_PPC64_GOT_TLSGD16        R_PPC64 = 79 // R_POWERPC_GOT_TLSGD16
	R_PPC64_GOT_TLSGD16_LO     R_PPC64 = 80 // R_POWERPC_GOT_TLSGD16_LO
	R_PPC64_GOT_TLSGD16_HI     R_PPC64 = 81 // R_POWERPC_GOT_TLSGD16_HI
	R_PPC64_GOT_TLSGD16_HA     R_PPC64 = 82 // R_POWERPC_GOT_TLSGD16_HA
	R_PPC64_GOT_TLSLD16        R_PPC64 = 83 // R_POWERPC_GOT_TLSLD16
	R_PPC64_GOT_TLSLD16_LO     R_PPC64 = 84 // R_POWERPC_GOT_TLSLD16_LO
	R_PPC64_GOT_TLSLD16_HI     R_PPC64 = 85 // R_POWERPC_GOT_TLSLD16_HI
	R_PPC64_GOT_TLSLD16_HA     R_PPC64 = 86 // R_POWERPC_GOT_TLSLD16_HA
	R_PPC64_GOT_TPREL16_DS     R_PPC64 = 87 // R_POWERPC_GOT_TPREL16_DS
	R_PPC64_GOT_TPREL16_LO_DS  R_PPC64 = 88 // R_POWERPC_GOT_TPREL16_LO_DS
	R_PPC64_GOT_TPREL16_HI     R_PPC64 = 89 // R_POWERPC_GOT_TPREL16_HI
	R_PPC64_GOT_TPREL16_HA     R_PPC64 = 90 // R_POWERPC_GOT_TPREL16_HA
	R_PPC64_GOT_DTPREL16_DS    R_PPC64 = 91 // R_POWERPC_GOT_DTPREL16_DS
	R_PPC64_GOT_DTPREL16_LO_DS R_PPC64 = 92 // R_POWERPC_GOT_DTPREL16_LO_DS
	R_PPC64_GOT_DTPREL16_HI    R_PPC64 = 93 // R_POWERPC_GOT_DTPREL16_HI
	R_PPC64_GOT_DTPREL16_HA    R_PPC64 = 94 // R_POWERPC_GOT_DTPREL16_HA
	R_PPC64_TPREL16_DS         R_PPC64 = 95
	R_PPC64_TPREL16_LO_DS      R_PPC64 = 96
	R_PPC64_TPREL16_HIGHER     R_PPC64 = 97
	R_PPC64_TPREL16_HIGHERA    R_PPC64 = 98
	R_PPC64_TPREL16_HIGHEST    R_PPC64 = 99
	R_PPC64_TPREL16_HIGHESTA   R_PPC64 = 100
	R_PPC64_DTPREL16_DS        R_PPC64 = 101
	R_PPC64_DTPREL16_LO_DS     R_PPC64 = 102
	R_PPC64_DTPREL16_HIGHER    R_PPC64 = 103
	R_PPC64_DTPREL16_HIGHERA   R_PPC64 = 104
	R_PPC64_DTPREL16_HIGHEST   R_PPC64 = 105
	R_PPC64_DTPREL16_HIGHESTA  R_PPC64 = 106
	R_PPC64_TLSGD              R_PPC64 = 107
	R_PPC64_TLSLD              R_PPC64 = 108
	R_PPC64_TOCSAVE            R_PPC64 = 109
	R_PPC64_ADDR16_HIGH        R_PPC64 = 110
	R_PPC64_ADDR16_HIGHA       R_PPC64 = 111
	R_PPC64_TPREL16_HIGH       R_PPC64 = 112
	R_PPC64_TPREL16_HIGHA      R_PPC64 = 113
	R_PPC64_DTPREL16_HIGH      R_PPC64 = 114
	R_PPC64_DTPREL16_HIGHA     R_PPC64 = 115
	R_PPC64_REL24_NOTOC        R_PPC64 = 116
	R_PPC64_ADDR64_LOCAL       R_PPC64 = 117
	R_PPC64_ENTRY              R_PPC64 = 118
	R_PPC64_PLTSEQ             R_PPC64 = 119
	R_PPC64_PLTCALL            R_PPC64 = 120
	R_PPC64_PLTSEQ_NOTOC       R_PPC64 = 121
	R_PPC64_PLTCALL_NOTOC      R_PPC64 = 122
	R_PPC64_PCREL_OPT          R_PPC64 = 123
	R_PPC64_REL24_P9NOTOC      R_PPC64 = 124
	R_PPC64_D34                R_PPC64 = 128
	R_PPC64_D34_LO             R_PPC64 = 129
	R_PPC64_D34_HI30           R_PPC64 = 130
	R_PPC64_D34_HA30           R_PPC64 = 131
	R_PPC64_PCREL34            R_PPC64 = 132
	R_PPC64_GOT_PCREL34        R_PPC64 = 133
	R_PPC64_PLT_PCREL34        R_PPC64 = 134
	R_PPC64_PLT_PCREL34_NOTOC  R_PPC64 = 135
	R_PPC64_ADDR16_HIGHER34    R_PPC64 = 136
	R_PPC64_ADDR16_HIGHERA34   R_PPC64 = 137
	R_PPC64_ADDR16_HIGHEST34   R_PPC64 = 138
	R_PPC64_ADDR16_HIGHESTA34  R_PPC64 = 139
	R_PPC64_REL16_HIGHER34     R_PPC64 = 140
	R_PPC64_REL16_HIGHERA34    R_PPC64 = 141
	R_PPC64_REL16_HIGHEST34    R_PPC64 = 142
	R_PPC64_REL16_HIGHESTA34   R_PPC64 = 143
	R_PPC64_D28                R_PPC64 = 144
	R_PPC64_PCREL28            R_PPC64 = 145
	R_PPC64_TPREL34            R_PPC64 = 146
	R_PPC64_DTPREL34           R_PPC64 = 147
	R_PPC64_GOT_TLSGD_PCREL34  R_PPC64 = 148
	R_PPC64_GOT_TLSLD_PCREL34  R_PPC64 = 149
	R_PPC64_GOT_TPREL_PCREL34  R_PPC64 = 150
	R_PPC64_GOT_DTPREL_PCREL34 R_PPC64 = 151
	R_PPC64_REL16_HIGH         R_PPC64 = 240
	R_PPC64_REL16_HIGHA        R_PPC64 = 241
	R_PPC64_REL16_HIGHER       R_PPC64 = 242
	R_PPC64_REL16_HIGHERA      R_PPC64 = 243
	R_PPC64_REL16_HIGHEST      R_PPC64 = 244
	R_PPC64_REL16_HIGHESTA     R_PPC64 = 245
	R_PPC64_REL16DX_HA         R_PPC64 = 246 // R_POWERPC_REL16DX_HA
	R_PPC64_JMP_IREL           R_PPC64 = 247
	R_PPC64_IRELATIVE          R_PPC64 = 248 // R_POWERPC_IRELATIVE
	R_PPC64_REL16              R_PPC64 = 249 // R_POWERPC_REL16
	R_PPC64_REL16_LO           R_PPC64 = 250 // R_POWERPC_REL16_LO
	R_PPC64_REL16_HI           R_PPC64 = 251 // R_POWERPC_REL16_HI
	R_PPC64_REL16_HA           R_PPC64 = 252 // R_POWERPC_REL16_HA
	R_PPC64_GNU_VTINHERIT      R_PPC64 = 253
	R_PPC64_GNU_VTENTRY        R_PPC64 = 254
)
```

#### func (i R_PPC64) GoString() string 添加于1.5

#### func (i R_PPC64) String() string 添加于1.5

### type R_SPARC 添加于1.11

```go
type R_RISCV int
```

RISC-V 处理器的重定位类型

```go
const (
	R_RISCV_NONE R_RISCV = 0 /* 没有重新定位。*/
	R_RISCV_32 R_RISCV = 1 /* 添加 32 位零扩展符号值 */
	R_RISCV_64 R_RISCV = 2 /* 添加 64 位符号值。*/
	R_RISCV_RELATIVE R_RISCV = 3 /* 添加共享对象的加载地址。*/
	R_RISCV_COPY R_RISCV = 4 /* 从共享对象复制数据。*/
	R_RISCV_JUMP_SLOT R_RISCV = 5 /* 将 GOT 条目设置为代码地址。*/
	R_RISCV_TLS_DTPMOD32 R_RISCV = 6 /* 包含符号的模块的 32 位 ID */
	R_RISCV_TLS_DTPMOD64 R_RISCV = 7 /* 包含符号的模块 ID */
	R_RISCV_TLS_DTPREL32 R_RISCV = 8 /* TLS 块中的 32 位相对偏移 */
	R_RISCV_TLS_DTPREL64 R_RISCV = 9 /* TLS 块中的相对偏移 */
	R_RISCV_TLS_TPREL32 R_RISCV = 10 /* 静态 TLS 块中的 32 位相对偏移 */
	R_RISCV_TLS_TPREL64 R_RISCV = 11 /* 静态 TLS 块中的相对偏移 */
	R_RISCV_BRANCH R_RISCV = 16 /* PC 相对分支 */
	R_RISCV_JAL R_RISCV = 17 /* PC 相对跳转 */
	R_RISCV_CALL R_RISCV = 18 /* PC 相对调用 */
	R_RISCV_CALL_PLT R_RISCV = 19 /* PC 相对调用（PLT） */
	R_RISCV_GOT_HI20 R_RISCV = 20 /* PC 相对 GOT 参考 */
	R_RISCV_TLS_GOT_HI20 R_RISCV = 21 /* PC 相对 TLS IE GOT 偏移量 */
	R_RISCV_TLS_GD_HI20 R_RISCV = 22 /* PC 相对 TLS GD 引用 */
	R_RISCV_PCREL_HI20 R_RISCV = 23 /* PC 相对引用 */
	R_RISCV_PCREL_LO12_I R_RISCV = 24 /* PC 相对引用 */
	R_RISCV_PCREL_LO12_S R_RISCV = 25 /* PC 相对基准 */
	R_RISCV_HI20 R_RISCV = 26 /* 绝对地址 */
	R_RISCV_LO12_I R_RISCV = 27 /* 绝对地址 */
	R_RISCV_LO12_S R_RISCV = 28 /* 绝对地址 */
	R_RISCV_TPREL_HI20 R_RISCV = 29 /* TLS LE 线程偏移 */
	R_RISCV_TPREL_LO12_I R_RISCV = 30 /* TLS LE 线程偏移 */
	R_RISCV_TPREL_LO12_S R_RISCV = 31 /* TLS LE 线程偏移 */
	R_RISCV_TPREL_ADD R_RISCV = 32 /* TLS LE 线程使用 */
	R_RISCV_ADD8 R_RISCV = 33 /* 8 位标签加法 */
	R_RISCV_ADD16 R_RISCV = 34 /* 16 位标签加法 */
	R_RISCV_ADD32 R_RISCV = 35 /* 32 位标签加法 */
	R_RISCV_ADD64 R_RISCV = 36 /* 64 位标签加法 */
	R_RISCV_SUB8 R_RISCV = 37 /* 8 位标签减法 */
	R_RISCV_SUB16 R_RISCV = 38 /* 16 位标签减法 */
	R_RISCV_SUB32 R_RISCV = 39 /* 32 位标签减法 */
	R_RISCV_SUB64 R_RISCV = 40 /* 64 位标签减法 */
	R_RISCV_GNU_VTINHERIT R_RISCV = 41 /* GNU C++ vtable 层次结构 */
	R_RISCV_GNU_VTENTRY R_RISCV = 42 /* GNU C++ vtable 成员用法 */
	R_RISCV_ALIGN R_RISCV = 43 /* 对齐声明 */
	R_RISCV_RVC_BRANCH R_RISCV = 44 /* PC 相对分支偏移 */
	R_RISCV_RVC_JUMP R_RISCV = 45 /* PC 相对跳转偏移 */
	R_RISCV_RVC_LUI R_RISCV = 46 /* 绝对地址 */
	R_RISCV_GPREL_I R_RISCV = 47 /* GP 相对引用 */
	R_RISCV_GPREL_S R_RISCV = 48 /* GP 相对引用 */
	R_RISCV_TPREL_I R_RISCV = 49 /* TP 相对 TLS LE 负载 */
	R_RISCV_TPREL_S R_RISCV = 50 /* TP 相对 TLS LE 存储 */
	R_RISCV_RELAX R_RISCV = 51 /* 指令对可以放宽 */
	R_RISCV_SUB6 R_RISCV = 52 /* 本地标签减法 */
	R_RISCV_SET6 R_RISCV = 53 /* 本地标签减法 */
	R_RISCV_SET8 R_RISCV = 54 /* 本地标签减法 */
	R_RISCV_SET16 R_RISCV = 55 /* 本地标签减法 */
	R_RISCV_SET32 R_RISCV = 56 /* 本地标签减法 */
	R_RISCV_32_PCREL R_RISCV = 57 /* 32 位 PC 相对 */
)
```

#### func (i R_PISCV) GoString() string 添加于1.11

#### func (i R_RISCV) String() string 添加于1.11

### type R_SPARC

```go
type R_SPARC int
```

SPARC 迁移类型。

```go
const (
	R_SPARC_NONE     R_SPARC = 0
	R_SPARC_8        R_SPARC = 1
	R_SPARC_16       R_SPARC = 2
	R_SPARC_32       R_SPARC = 3
	R_SPARC_DISP8    R_SPARC = 4
	R_SPARC_DISP16   R_SPARC = 5
	R_SPARC_DISP32   R_SPARC = 6
	R_SPARC_WDISP30  R_SPARC = 7
	R_SPARC_WDISP22  R_SPARC = 8
	R_SPARC_HI22     R_SPARC = 9
	R_SPARC_22       R_SPARC = 10
	R_SPARC_13       R_SPARC = 11
	R_SPARC_LO10     R_SPARC = 12
	R_SPARC_GOT10    R_SPARC = 13
	R_SPARC_GOT13    R_SPARC = 14
	R_SPARC_GOT22    R_SPARC = 15
	R_SPARC_PC10     R_SPARC = 16
	R_SPARC_PC22     R_SPARC = 17
	R_SPARC_WPLT30   R_SPARC = 18
	R_SPARC_COPY     R_SPARC = 19
	R_SPARC_GLOB_DAT R_SPARC = 20
	R_SPARC_JMP_SLOT R_SPARC = 21
	R_SPARC_RELATIVE R_SPARC = 22
	R_SPARC_UA32     R_SPARC = 23
	R_SPARC_PLT32    R_SPARC = 24
	R_SPARC_HIPLT22  R_SPARC = 25
	R_SPARC_LOPLT10  R_SPARC = 26
	R_SPARC_PCPLT32  R_SPARC = 27
	R_SPARC_PCPLT22  R_SPARC = 28
	R_SPARC_PCPLT10  R_SPARC = 29
	R_SPARC_10       R_SPARC = 30
	R_SPARC_11       R_SPARC = 31
	R_SPARC_64       R_SPARC = 32
	R_SPARC_OLO10    R_SPARC = 33
	R_SPARC_HH22     R_SPARC = 34
	R_SPARC_HM10     R_SPARC = 35
	R_SPARC_LM22     R_SPARC = 36
	R_SPARC_PC_HH22  R_SPARC = 37
	R_SPARC_PC_HM10  R_SPARC = 38
	R_SPARC_PC_LM22  R_SPARC = 39
	R_SPARC_WDISP16  R_SPARC = 40
	R_SPARC_WDISP19  R_SPARC = 41
	R_SPARC_GLOB_JMP R_SPARC = 42
	R_SPARC_7        R_SPARC = 43
	R_SPARC_5        R_SPARC = 44
	R_SPARC_6        R_SPARC = 45
	R_SPARC_DISP64   R_SPARC = 46
	R_SPARC_PLT64    R_SPARC = 47
	R_SPARC_HIX22    R_SPARC = 48
	R_SPARC_LOX10    R_SPARC = 49
	R_SPARC_H44      R_SPARC = 50
	R_SPARC_M44      R_SPARC = 51
	R_SPARC_L44      R_SPARC = 52
	R_SPARC_REGISTER R_SPARC = 53
	R_SPARC_UA64     R_SPARC = 54
	R_SPARC_UA16     R_SPARC = 55
)
```

#### func (i R_SPARC) GoString() string

#### func (i R_SPARC) String() string

### type R_X86_64

```go
type R_X86 int
```

x86-64 迁移类型

```go
const (
	R_X86_64_NONE R_X86_64 = 0 /* 没有重新定位。*/
	R_X86_64_64 R_X86_64 = 1 /* 添加 64 位符号值。*/
	R_X86_64_PC32 R_X86_64 = 2 /* PC 相对 32 位符号值。*/
	R_X86_64_GOT32 R_X86_64 = 3 /* PC 相对的 32 位 GOT 偏移量。*/
	R_X86_64_PLT32 R_X86_64 = 4 /* PC 相对的 32 位 PLT 偏移量。*/
	R_X86_64_COPY R_X86_64 = 5 /* 从共享对象复制数据。*/
	R_X86_64_GLOB_DAT R_X86_64 = 6 /* 将 GOT 条目设置为数据地址。*/
	R_X86_64_JMP_SLOT R_X86_64 = 7 /* 将 GOT 条目设置为代码地址。*/
	R_X86_64_RELATIVE R_X86_64 = 8 /* 添加共享对象的加载地址。*/
	R_X86_64_GOTPCREL R_X86_64 = 9 /* 添加 32 位有符号 pcrel 偏移量到 GOT。*/
	R_X86_64_32 R_X86_64 = 10 /* 添加 32 位零扩展符号值 */
	R_X86_64_32S R_X86_64 = 11 /* 添加 32 位符号扩展符号值 */
	R_X86_64_16 R_X86_64 = 12 /* 添加 16 位零扩展符号值 */
	R_X86_64_PC16 R_X86_64 = 13 /* 添加 16 位有符号扩展 pc 相对符号值 */
	R_X86_64_8 R_X86_64 = 14 /* 添加 8 位零扩展符号值 */
	R_X86_64_PC8 R_X86_64 = 15 /* 添加 8 位有符号扩展 pc 相对符号值 */
	R_X86_64_DTPMOD64 R_X86_64 = 16 /* 包含符号的模块 ID */
	R_X86_64_DTPOFF64 R_X86_64 = 17 /* TLS 块中的偏移 */
	R_X86_64_TPOFF64 R_X86_64 = 18 /* 静态 TLS 块中的偏移量 */
	R_X86_64_TLSGD R_X86_64 = 19 /* PC 相对于 GD GOT 条目的偏移 */
	R_X86_64_TLSLD R_X86_64 = 20 /* PC 相对于 LD GOT 条目的偏移 */
	R_X86_64_DTPOFF32 R_X86_64 = 21 /* TLS 块中的偏移量 */
	R_X86_64_GOTTPOFF R_X86_64 = 22 /* PC 相对于 IE GOT 条目的偏移 */
	R_X86_64_TPOFF32 R_X86_64 = 23 /* 静态 TLS 块中的偏移 */
	R_X86_64_PC64 R_X86_64 = 24 /* PC 相对 64 位符号扩展符号值。*/
	R_X86_64_GOTOFF64        R_X86_64 = 25
	R_X86_64_GOTPC32         R_X86_64 = 26
	R_X86_64_GOT64           R_X86_64 = 27
	R_X86_64_GOTPCREL64      R_X86_64 = 28
	R_X86_64_GOTPC64         R_X86_64 = 29
	R_X86_64_GOTPLT64        R_X86_64 = 30
	R_X86_64_PLTOFF64        R_X86_64 = 31
	R_X86_64_SIZE32          R_X86_64 = 32
	R_X86_64_SIZE64          R_X86_64 = 33
	R_X86_64_GOTPC32_TLSDESC R_X86_64 = 34
	R_X86_64_TLSDESC_CALL    R_X86_64 = 35
	R_X86_64_TLSDESC         R_X86_64 = 36
	R_X86_64_IRELATIVE       R_X86_64 = 37
	R_X86_64_RELATIVE64      R_X86_64 = 38
	R_X86_64_PC32_BND        R_X86_64 = 39
	R_X86_64_PLT32_BND       R_X86_64 = 40
	R_X86_64_GOTPCRELX       R_X86_64 = 41
	R_X86_64_REX_GOTPCRELX   R_X86_64 = 42
)
```

#### func (i R_X86_64) GoString() string

#### func (i R_X86_64) String() string

### type Rel32

```go
type Rel32 struct {
	Off uint32 /* 要重新定位的位置。*/
	Info uint32 /* 重定位类型和符号索引。*/
}
```

ELF32 不需要附加字段的重定位。

### type Rel64

```go
type Rel64 struct {
	Off uint64 /* 要重新定位的位置。*/
	Info uint64 /* 重定位类型和符号索引。*/
}
```

不需要 addend 字段的 ELF64 重新定位。

### type Rela32

```go
type Rela32 struct {
  Off uint32 /* 要重新定位的位置。*/
	Info uint32 /* 重定位类型和符号索引。*/
	Addend int32 /* 添加。*/
}
```

需要附加字段的 ELF32 重定位。

### type Rela64

```go
type Rela64 struct {
	Off uint64 /* 要重新定位的位置。*/
	Info uint64 /* 重定位类型和符号索引。*/
	Addend int64 /* 添加。*/
}
```

需要 addend 字段的 ELF64 重定位。

### type Section

```go
type Section struct {
  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  // 
  // 为 ReadAt 方法嵌入 ReaderAt。不要直接嵌入 SectionReader，以避免出现读取和查找。如果客户端需要 "读取和查找"，则必须使用 Open()，以避免与其他客户端争夺查找偏移。
  io.ReaderAt
  // 包含已筛选或未导出字段
}
```

一个小节代表 ELF 文件中的一个单独小节。

#### func (s *Section) Data() ([]byte, error)

Data 读取并返回 ELF 部分的内容。即使 ELF 文件中存储的部分是压缩的，Data 也会返回未压缩的数据。

对于 SHT_NOBITS 部分，Data 总是返回非零错误。

#### func (s *Section) Open() io.ReadSeeker

Open 返回一个读取 ELF 部分的新 ReadSeeker。即使该部分已压缩存储在 ELF 文件中，ReadSeeker 也会读取未压缩的数据。

对于 SHT_NOBITS 部分，所有对已打开读取器的调用都将返回非零错误。

### type Section32

```go
type Section32 struct {
	名称 uint32 /* 部分名称（部分标题字符串表的索引）。*/
	类型 uint32 /* 部分类型。*/
	Flags uint32 /* 科标志。*/
	Addr uint32 /* 内存图像中的地址。*/
	Off uint32 /* 文件中的偏移量。*/
	Size uint32 /* 大小（字节）。*/
	链接 uint32 /* 相关部分的索引。*/
	Info uint32 /* 取决于章节类型。*/
	Addralign uint32 /* 以字节为单位对齐。*/
	Entsize uint32 /* 部分中每个条目的大小。*/
}
```

ELF32 部分标题。

### type Section64

```go
type Section64 struct {
	名称 uint32 /* 部分名称（部分标题字符串表的索引）。*/
	类型 uint32 /* 部分类型。*/
	Flags uint64 /* 科标志。*/
	Addr uint64 /* 内存图像中的地址。*/
	Off uint64 /* 文件中的偏移量。*/
	Size uint64 /* 大小（字节）。*/
	链接 uint32 /* 相关部分的索引。*/
	Info uint32 /* 取决于章节类型。*/
	Addralign uint64 /* 以字节为单位的对齐方式。*/
	Entsize uint64 /* 部分中每个条目的大小。*/
}
```

ELF64 部分标题

### type SectionFlag

```go
type SectionFlag uint32
```

flags 部分

```go
const (
	SHF_WRITE SectionFlag = 0x1 /* 部分包含可写数据。*/
	SHF_ALLOC SectionFlag = 0x2 /* 部分占用内存。*/
	SHF_EXECINSTR SectionFlag = 0x4 /* 部分包含指令。*/
	SHF_MERGE SectionFlag = 0x10 /* 部分可能被合并。*/
	SHF_STRINGS SectionFlag = 0x20 /* 部分包含字符串。*/
	SHF_INFO_LINK SectionFlag = 0x40 /* sh_info 保存分节索引。*/
	SHF_LINK_ORDER SectionFlag = 0x80 /* 特殊排序要求。*/
	SHF_OS_NONCONFORMING SectionFlag = 0x100 /* 需要进行操作系统特定处理。*/
	SHF_GROUP SectionFlag = 0x200 /* 章节组的成员。*/
	SHF_TLS SectionFlag = 0x400 /* 部分包含 TLS 数据。*/
	SHF_COMPRESSED SectionFlag = 0x800 /* 部分已压缩。*/
	SHF_MASKOS SectionFlag = 0x0ff00000 /* 特定于操作系统的语义。*/
	SHF_MASKPROC SectionFlag = 0xf0000000 /* 特定于处理器的语义。*/
)
```

#### func (i SectionFlat) GoString() string

#### func (i SectionFlag) String() string

### type SectionHeader

```go
type SectionHeader struct {
	Name      string
	Type      SectionType
	Flags     SectionFlag
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64

  // FileSize 是该部分在文件中的大小（以字节为单位）。如果某部分已压缩，则 FileSize 是压缩数据的大小，而 Size（上文）是未压缩数据的大小。
  FileSize uint64
}
```

一个 SectionHeader 表示一个 ELF 段落标头。

### type SectionIndex

```go
type SectionIndex int
```

特别部分索引。

```go
const (
	SHN_UNDEF SectionIndex = 0 /* 未定义、缺失、无关。*/
	SHN_LORESERVE SectionIndex = 0xff00 /* 保留范围的第一个。*/
	SHN_LOPROC SectionIndex = 0xff00 /* 第一个特定处理器。*/
	SHN_HIPROC SectionIndex = 0xff1f /* 最后一个处理器专用。*/
	SHN_LOOS SectionIndex = 0xff20 /* 第一个操作系统专用。*/
	SHN_HIOS SectionIndex = 0xff3f /* 最后一个操作系统专用。*/
	SHN_ABS SectionIndex = 0xfff1 /* 绝对值。*/
	SHN_COMMON SectionIndex = 0xfff2 /* 常用数据。*/
	SHN_XINDEX SectionIndex = 0xffff /* Escape; index stored elsewhere.*/
	SHN_HIRESERVE SectionIndex = 0xffff /* 保留范围的最后一个。*/
)
```

#### func (i SectionIndex) GoString() string

#### func (i SectionIndex) String() string

### type SectionType

```go
type SectionType uint32
```

Section type.

```go
const (
	SHT_NULL SectionType = 0 /* 非活动 */
	SHT_PROGBITS SectionType = 1 /* 程序定义信息 */
	SHT_SYMTAB SectionType = 2 /* 符号表部分 */
	SHT_STRTAB SectionType = 3 /* string table section */
	SHT_RELA SectionType = 4 /* 带添加项的重定位部分 */
	SHT_HASH SectionType = 5 /* 符号哈希表部分 */
	SHT_DYNAMIC SectionType = 6 /* 动态部分 */
	SHT_NOTE SectionType = 7 /* 注释部分 */
	SHT_NOBITS SectionType = 8 /* 无空格部分 */
	SHT_REL SectionType = 9 /* relocation section - no addends */
	SHT_SHLIB SectionType = 10 /* 保留 - 用途不明 */
	SHT_DYNSYM SectionType = 11 /* 动态符号表部分 */
	SHT_INIT_ARRAY SectionType = 14 /* 初始化函数指针。*/
	SHT_FINI_ARRAY SectionType = 15 /* 终止函数指针。*/
	SHT_PREINIT_ARRAY SectionType = 16 /* 预初始化函数指针。*/
	SHT_GROUP SectionType = 17 /* 节段组。*/
	SHT_SYMTAB_SHNDX SectionType = 18 /* 科目索引（参见 SHN_XINDEX）。*/
	SHT_LOOS SectionType = 0x60000000 /* 第一个操作系统特定语义 */
	SHT_GNU_ATTRIBUTES SectionType = 0x6ffffff5 /* GNU 对象属性 */
	SHT_GNU_HASH SectionType = 0x6ffff6 /* GNU 哈希表 */
	SHT_GNU_LIBLIST SectionType = 0x6ffff7 /* GNU 预链接库列表 */
	SHT_GNU_VERDEF SectionType = 0x6ffffffd /* GNU 版本定义部分 */
	SHT_GNU_VERNEED SectionType = 0x6ffffffe /* GNU 版本需求部分 */
	SHT_GNU_VERSYM SectionType = 0x6fffffff /* GNU 版本符号表 */
	SHT_HIOS SectionType = 0x6ffffff /* 最后的操作系统特定语义 */
	SHT_LOPROC SectionType = 0x70000000 /* 为处理器保留的范围 */
	SHT_MIPS_ABIFLAGS SectionType = 0x7000002a /* .MIPS.abiflags */
	SHT_HIPROC SectionType = 0x7fffffff /* 特定部分头类型 */
	SHT_LOUSER SectionType = 0x80000000 /* 为应用程序保留的范围 */
	SHT_HIUSER SectionType = 0xffffffff /* 特定索引 */
)
```

#### func (i SectionType) GoString() string

#### func (i SectionType) String() string

### type Sym32

```go
type Sym32 struct {
  Name uint32
  Value uint32
  Size uint32
  Info uint8
  Other uint8
  Shndx uint16
}
```

ELF32 符号。

### type Sym64

```go
type Sym64 struct {
	名称 uint32 /* 名称的字符串表索引。*/
	Info uint8 /* 类型和绑定信息。*/
	Other uint8 /* 保留（未使用）。*/
	Shndx uint16 /* 符号段索引。*/
	Value uint64 /* 符号值。*/
	Size uint64 /* 关联对象的大小。*/
}
```

ELF64 符号表条目。

### type SymBind

```go
type SymBind int
```

符号绑定 - ELFNN_ST_BIND - st_info

```go
const (
	STB_LOCAL SymBind = 0 /* 本地符号 */
	STB_GLOBAL SymBind = 1 /* 全局符号 */
	STB_WEAK SymBind = 2 /* 类似全局 - 低优先级 */
	STB_LOOS SymBind = 10 /* 为操作系统保留范围 */
	STB_HIOS SymBind = 12 /* 特定语义。*/
	STB_LOPROC SymBind = 13 /* 为处理器保留的范围 */
	STB_HIPROC SymBind = 15 /* 特定语义。*/
)
```

#### func ST_BIND(info uint8) SymBind

#### func (i SymBind) GoString() string

#### func (i SymBind) String() string

### type SymType

```go
type SymType int
```

符号类型 - ELFNN_ST_TYPE - st_info

```go
const (
	STT_NOTYPE SymType = 0 /* 未指定类型。*/
	STT_OBJECT SymType = 1 /* 数据对象。*/
	STT_FUNC SymType = 2 /* 函数。*/
	STT_SECTION SymType = 3 /* Section.*/
	STT_FILE SymType = 4 /* 源文件。*/
	STT_COMMON SymType = 5 /* 未初始化的公共代码块。*/
	STT_TLS SymType = 6 /* TLS 对象。*/
	STT_LOOS SymType = 10 /* 为操作系统保留的范围 */
	STT_HIOS SymType = 12 /* 特定语义。*/
	STT_LOPROC SymType = 13 /* 为处理器保留的范围 */
	STT_HIPROC SymType = 15 /* 特定语义。*/
)
```

#### func ST_TYPE(info uint8) SymType

#### func (i SymType) GoString() string

#### func (i SymType) String() string

### type SymVis

```go
type SymVis int
```

符号可见性 - ELFNN_ST_VISIBILITY - st_other

```go
const (
	STV_DEFAULT SymVis = 0x0 /* 默认可见性（参见绑定）。*/
	STV_INTERNAL SymVis = 0x1 /* 在可重置对象中的特殊含义。*/
	STV_HIDDEN SymVis = 0x2 /* 不可见。*/
	STV_PROTECTED SymVis = 0x3 /* 可见但不可抢占。*/
)
```

#### func ST_VISIBILITY(other uint8) SymVis

#### func (i SymVis) GoString() string

#### func (i SymVis) String() string

### type Symbol

```go
type Symbol struct {
	Name        string
	Info, Other byte
	Section     SectionIndex
	Value, Size uint64

  // 版本和库仅存在于动态符号表中。
  Version string
  Library string
}
```

一个符号代表 ELF 符号表部分的一个条目。

### type Type

```go
type Type uint16
```

类型可在 Header.Type 中找到。

```go
const (
	ET_NONE Type = 0 /* 未知类型。*/
	ET_REL Type = 1 /* 可重新定位。*/
	ET_EXEC Type = 2 /* 可执行。*/
	ET_DYN 类型 = 3 /* 共享对象。*/
	ET_CORE 类型 = 4 /* 核心文件。*/
	ET_LOOS 类型 = 0xfe00 /* 第一个操作系统专用。*/
	ET_HIOS 类型 = 0xfeff /* 最后一个操作系统专用。*/
	ET_LOPROC 类型 = 0xff00 /* 第一个处理器专用。*/
	ET_HIPROC 类型 = 0xffff /* 最后一个处理器专用。*/
)
```

#### func (i Type) GoString() string

#### func (i Type) String() string

### type Version

```go
type Version byte
```

版本可在 Header.Ident[EI_VERSION] 和 Header.Version 中找到。

```go
const (
  EV_NONE = Version = 0
  EV_CURRENT Version = 1
)
```

#### func (i Version) GoString() string

#### func (i Version) String() string