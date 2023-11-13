## package dwarf

按照 http://dwarfstd.org/doc/dwarf-2.0.0.pdf 上 DWARF 2.0 标准的定义，包 dwarf 可访问从可执行文件加载的 DWARF 调试信息。

### Security

该软件包的设计并不针对对抗性输入进行加固，也不在 https://go.dev/security/policy 的范围之内。特别是在解析对象文件时，只进行了基本的验证。因此，在解析不受信任的输入时应小心谨慎，因为解析畸形文件可能会耗费大量资源或导致恐慌。

## Variables

```go
var ErrUnknownPC = errors.New("ErrUnknownPc")
```

ErrUnknownPC 是 LineReader.ScanPC 在行表中的任何条目都未覆盖寻道 PC 时返回的错误。

### type AddrType

```go
type AddrType struct {
  BasicType
}
```

AddrType 表示机器地址类型。

### type ArrayType

```go
type ArrayType struct {
  CommonType
  Type Type
  StrideBitSize int64 // 如果每个元素的位数大于 0
  Count int64 // 如果等于 -1 表示不完整的数组，如 char x[]。
}
```

ArrayType 表示大小固定的数组类型。

#### func (t *ArrayType) Size() int64

#### func (t *ArrayType) String() string

### type Attr

Attr 用于标识 DWARF 输入字段中的属性类型。

```go
const (
	AttrSibling        Attr = 0x01
	AttrLocation       Attr = 0x02
	AttrName           Attr = 0x03
	AttrOrdering       Attr = 0x09
	AttrByteSize       Attr = 0x0B
	AttrBitOffset      Attr = 0x0C
	AttrBitSize        Attr = 0x0D
	AttrStmtList       Attr = 0x10
	AttrLowpc          Attr = 0x11
	AttrHighpc         Attr = 0x12
	AttrLanguage       Attr = 0x13
	AttrDiscr          Attr = 0x15
	AttrDiscrValue     Attr = 0x16
	AttrVisibility     Attr = 0x17
	AttrImport         Attr = 0x18
	AttrStringLength   Attr = 0x19
	AttrCommonRef      Attr = 0x1A
	AttrCompDir        Attr = 0x1B
	AttrConstValue     Attr = 0x1C
	AttrContainingType Attr = 0x1D
	AttrDefaultValue   Attr = 0x1E
	AttrInline         Attr = 0x20
	AttrIsOptional     Attr = 0x21
	AttrLowerBound     Attr = 0x22
	AttrProducer       Attr = 0x25
	AttrPrototyped     Attr = 0x27
	AttrReturnAddr     Attr = 0x2A
	AttrStartScope     Attr = 0x2C
	AttrStrideSize     Attr = 0x2E
	AttrUpperBound     Attr = 0x2F
	AttrAbstractOrigin Attr = 0x31
	AttrAccessibility  Attr = 0x32
	AttrAddrClass      Attr = 0x33
	AttrArtificial     Attr = 0x34
	AttrBaseTypes      Attr = 0x35
	AttrCalling        Attr = 0x36
	AttrCount          Attr = 0x37
	AttrDataMemberLoc  Attr = 0x38
	AttrDeclColumn     Attr = 0x39
	AttrDeclFile       Attr = 0x3A
	AttrDeclLine       Attr = 0x3B
	AttrDeclaration    Attr = 0x3C
	AttrDiscrList      Attr = 0x3D
	AttrEncoding       Attr = 0x3E
	AttrExternal       Attr = 0x3F
	AttrFrameBase      Attr = 0x40
	AttrFriend         Attr = 0x41
	AttrIdentifierCase Attr = 0x42
	AttrMacroInfo      Attr = 0x43
	AttrNamelistItem   Attr = 0x44
	AttrPriority       Attr = 0x45
	AttrSegment        Attr = 0x46
	AttrSpecification  Attr = 0x47
	AttrStaticLink     Attr = 0x48
	AttrType           Attr = 0x49
	AttrUseLocation    Attr = 0x4A
	AttrVarParam       Attr = 0x4B
	AttrVirtuality     Attr = 0x4C
	AttrVtableElemLoc  Attr = 0x4D
	// DWARF 3 新增了以下功能。
	AttrAllocated     Attr = 0x4E
	AttrAssociated    Attr = 0x4F
	AttrDataLocation  Attr = 0x50
	AttrStride        Attr = 0x51
	AttrEntrypc       Attr = 0x52
	AttrUseUTF8       Attr = 0x53
	AttrExtension     Attr = 0x54
	AttrRanges        Attr = 0x55
	AttrTrampoline    Attr = 0x56
	AttrCallColumn    Attr = 0x57
	AttrCallFile      Attr = 0x58
	AttrCallLine      Attr = 0x59
	AttrDescription   Attr = 0x5A
	AttrBinaryScale   Attr = 0x5B
	AttrDecimalScale  Attr = 0x5C
	AttrSmall         Attr = 0x5D
	AttrDecimalSign   Attr = 0x5E
	AttrDigitCount    Attr = 0x5F
	AttrPictureString Attr = 0x60
	AttrMutable       Attr = 0x61
	AttrThreadsScaled Attr = 0x62
	AttrExplicit      Attr = 0x63
	AttrObjectPointer Attr = 0x64
	AttrEndianity     Attr = 0x65
	AttrElemental     Attr = 0x66
	AttrPure          Attr = 0x67
	AttrRecursive     Attr = 0x68
	// DWARF 4 新增了以下功能。
	AttrSignature      Attr = 0x69
	AttrMainSubprogram Attr = 0x6A
	AttrDataBitOffset  Attr = 0x6B
	AttrConstExpr      Attr = 0x6C
	AttrEnumClass      Attr = 0x6D
	AttrLinkageName    Attr = 0x6E
	// DWARF 5 新增了以下功能。
	AttrStringLengthBitSize  Attr = 0x6F
	AttrStringLengthByteSize Attr = 0x70
	AttrRank                 Attr = 0x71
	AttrStrOffsetsBase       Attr = 0x72
	AttrAddrBase             Attr = 0x73
	AttrRnglistsBase         Attr = 0x74
	AttrDwoName              Attr = 0x76
	AttrReference            Attr = 0x77
	AttrRvalueReference      Attr = 0x78
	AttrMacros               Attr = 0x79
	AttrCallAllCalls         Attr = 0x7A
	AttrCallAllSourceCalls   Attr = 0x7B
	AttrCallAllTailCalls     Attr = 0x7C
	AttrCallReturnPC         Attr = 0x7D
	AttrCallValue            Attr = 0x7E
	AttrCallOrigin           Attr = 0x7F
	AttrCallParameter        Attr = 0x80
	AttrCallPC               Attr = 0x81
	AttrCallTailCall         Attr = 0x82
	AttrCallTarget           Attr = 0x83
	AttrCallTargetClobbered  Attr = 0x84
	AttrCallDataLocation     Attr = 0x85
	AttrCallDataValue        Attr = 0x86
	AttrNoreturn             Attr = 0x87
	AttrAlignment            Attr = 0x88
	AttrExportSymbols        Attr = 0x89
	AttrDeleted              Attr = 0x8A
	AttrDefaulted            Attr = 0x8B
	AttrLoclistsBase         Attr = 0x8C
)
```

#### func (a Attr) GoString() string

#### func (i Attr) String() string

```go
type BasicType struct {
	CommonType
	BitSize       int64
	BitOffset     int64
	DataBitOffset int64
}
```

一个 BasicType 包含所有基本类型共有的字段。

有关 BitSize/BitOffset/DataBitOffset 字段解释的更多信息，请参阅 StructField 文档。

#### func (b *BasicType) Basic() *BasicType

#### func (t *BasicType) String() string

### type BoolType

```go
type BoolType struct {
  BasicType
}
```

BoolType 表示布尔类型。

### type CharType

```go
type CharType struct {
  BasicType
}
```

CharType 表示带符号的字符类型。

### type Class

```go
type Class int
```

类是属性值的 DWARF 4 类。

一般来说，一个给定属性的值可能属于 DWARF 定义的几种可能类别之一，每种类别对属性的解释略有不同。

与 DWARF 以前的版本相比，DWARF 第 4 版对属性值类别的区分更加细致。读者可以将 DWARF 早期版本中较粗略的类别区分为相应的 DWARF 4 类别。例如，DWARF 2 使用 "常量 "来表示常量和所有类型的段偏移量，但读者会将 DWARF 2 文件中引用段偏移量的属性规范化为 Class*Ptr 类之一，尽管这些类仅在 DWARF 3 中定义。

```go
const (
  // ClassUnknown 表示未知 DWARF 类的值。
	ClassUnknown Class = iota

	// ClassAddress 表示 uint64 类型的值，即目标计算机上的地址。
	ClassAddress

	// ClassBlock 表示 []byte 类型的值，其解释取决于属性。
	ClassBlock

	// ClassConstant 表示 int64 类型的常量值。该常量的解释取决于属性。
	ClassConstant

	// ClassExprLoc 表示 []byte 类型的值，其中包含已编码的 DWARF 表达式或位置描述。
	ClassExprLoc

	// ClassFlag 表示 bool 类型的值。
	ClassFlag

	// ClassLinePtr 表示 int64 偏移到 "行 "部分的值。
	ClassLinePtr

	// ClassLocListPtr 表示 "loclist "部分的 int64 偏移值。
	ClassLocListPtr

	// ClassMacPtr 表示偏移到 "mac "部分的 int64 值。
	ClassMacPtr

	// ClassRangeListPtr 表示 int64 偏移到 "rangelist "部分的值。
	ClassRangeListPtr

	// ClassReference 表示信息部分条目的偏移值（与 Reader.Seek 一起使用）。DWARF 规范将 ClassReference 和 ClassReferenceSig 合并为 "引用 "类。
	ClassReference

	// ClassReferenceSig 表示引用类型 Entry 的 uint64 类型特征值。
	ClassReferenceSig

	// ClassString 表示字符串值。如果编译单元指定了 AttrUseUTF8 标志（强烈推荐），字符串值将以 UTF-8 编码。否则，将不指定编码。
	ClassString

	// ClassReferenceAlt 表示 int64 类型的值，这些值是替代对象文件 DWARF "信息 "部分的偏移量。
	ClassReferenceAlt

	// ClassStringAlt 表示 int64 类型的值，这些值是替代对象文件 DWARF 字符串部分的偏移量。
	ClassStringAlt

	// ClassAddrPtr 表示 int64 偏移到 "addr "部分的值。
	ClassAddrPtr

	// ClassLocList 表示 "loclists "部分的 int64 偏移值。
	ClassLocList

	// ClassRngList 表示从 "rnglists "部分的基数偏移 uint64 的值。
	ClassRngList

	// ClassRngListsPtr 表示 "rnglists "部分的 int64 偏移值。这些值将作为 ClassRngList 值的基础。
	ClassRngListsPtr

	// ClassStrOffsetsPtr 表示在 "str_offsets "部分的 int64 偏移值。
	ClassStrOffsetsPtr
)
```

#### func (i Class) GoString() string 添加于1.5

#### func (i Class) String() string 添加于1.5

### type CommonType

```go
type CommonType struct {
  ByteSize int64 // 该类型值的大小，以字节为单位
  Name string // name 可用于指代类型
}
```

CommonType 包含多个类型共有的字段。如果某个字段未知或不适用于给定类型，则使用零值。

#### func (c *CommonType) Common() *CommonType

#### func (c *CommonType) Size() int64

### type ComplexType

```go
type ComplexType struct {
  BasicType
}
```

ComplexType 表示复杂浮点类型。

### type Data

```go
type Data struct {
  // 包含已筛选或未导出字段
}
```

数据代表从可执行文件（如 ELF 或 Mach-O 可执行文件）加载的 DWARF 调试信息。

#### func New(abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) (*Data, error)

New 返回一个根据给定参数初始化的新数据对象。客户通常不应直接调用此函数，而应使用相应软件包 debug/elf、debug/macho 或 debug/pe 中文件类型的 DWARF 方法。

字节参数是对象文件中相应调试部分的数据；例如，对于 ELF 对象，缩写是".debug_abbrev "部分的内容。

#### func (d *Data) AddSection(name string, contents []byte) error 添加于1.14

AddSection 按名称添加另一个 DWARF 部分。名称应为 DWARF 部分名称，如".debug_addr"、".debug_str_offsets "等。这种方法用于在 DWARF 5 及更高版本中添加新的 DWARF 部分。

#### func (d *Data) AddTypes(name string, types []byte) error 添加于1.3

AddTypes 将在 DWARF 数据中添加一个 .debug_types 部分。一个包含 DWARF 第 4 版调试信息的典型对象将有多个 .debug_types 部分。该名称仅用于错误报告，并用于区分一个 .debug_types 部分和另一个 .debug_types 部分。

#### func (d *Data) LineReader(cu *Entry) (*LineReader, error) 添加于1.5

LineReader 返回编译单元 cu 行表的新阅读器，它必须是带有 TagCompileUnit 标记的条目。

如果该编译单元没有行表，则返回 nil、nil。

#### func (d *Data) Ranges(e *Entry) ([][2]uint64, error) 添加于1.7

Ranges 返回 e 所覆盖的 PC 范围，即 [low,high) 对的片段。只有某些条目类型（如 TagCompileUnit 或 TagSubprogram）具有 PC 范围；对于其他条目类型，Ranges 将返回 nil，且不会出错。

#### func (d *Data) Reader() *Reader

读取器返回一个新的数据读取器。读取器位于 DWARF "信息 "部分的字节偏移 0 处。

#### func (d *Data) Type(off Offset) (Type, error)

类型 "读取 DWARF "信息 "部分中处于关闭状态的类型。

### type DecodeError

```go
type DecodeError struct {
  Name string
  Offset Offset
  Err string
}
```

#### func (e DecodeError) error() string

### type DotDotDotType

```go
type DotDotDotType struct {
  CommonType
}
```

DotDotDotType 表示可变的...函数参数。

#### func (t *DotDotDotType) String() string

### type Entry

```go
type Entry struct {
  Offset Offset // DWARF 信息中条目偏移量
  Tag Tag // 标记（条目类型）
  Children bool // 是否有子跟随
  Field []Field
}
```

条目是属性/值对的序列。

#### func (e *Entry) AttrField(a Attr) *Field 添加于1.5

AttrField 返回与条目中属性 Attr 相关的字段，如果没有此类属性，则返回 nil。

#### func (e *Entry) Val(a Attr) any

Val 返回与 Entry 中属性 Attr 相关联的值，如果没有此类属性，则返回 nil。

一种常见的习惯做法是，将检查是否返回 nil 与检查值是否具有预期的动态类型合并起来，例如:

```go
v, ok := e.Val(AttrSibling).(int64)
```

### type EnumType

```go
type EnumType struct {
  CommonType
  EnumName string
  Val []*EnumValue
}
```

EnumType 表示枚举类型。其本地整数类型的唯一标识是字节大小（在 CommonType 中）。

#### func (t *EnumType) String() string

### type EnumValue

```go
type EnumValue struct {
  Name string
  Val int64
}
```

EnumValue 表示一个枚举值。

### type Field

```go
Attr Attr
Val any
Class Class
```

字段是条目中的单个属性/值对。

值可以是 DWARF 定义的几种 "属性类 "之一。每个类别对应的 Go 类型如下:

```go
DWARF class       Go type        Class
-----------       -------        -----
address           uint64         ClassAddress
block             []byte         ClassBlock
constant          int64          ClassConstant
flag              bool           ClassFlag
reference
  to info         dwarf.Offset   ClassReference
  to type unit    uint64         ClassReferenceSig
string            string         ClassString
exprloc           []byte         ClassExprLoc
lineptr           int64          ClassLinePtr
loclistptr        int64          ClassLocListPtr
macptr            int64          ClassMacPtr
rangelistptr      int64          ClassRangeListPtr
```

对于未识别或供应商定义的属性，Class 可以是 ClassUnknown。

### type FloatType

```go
type FloatType struct {
  BasicType
}
```

FloatType 表示浮点类型。

### type FuncType

```go
type FuncType struct {
  CommonType
  ReturnType Type
  ParamType []Type
}
```

FuncType 表示函数类型。

#### func (t *FuncType) String() string

### type IntType

```go
type IntType struct {
  BasicType
}
```

### type LineEntry 添加于1.5

```go
type LineEntry struct {
	// Address 是编译器生成的机器指令的程序计数器值。该 LineEntry 适用于从 Address 到下一个 LineEntry 的地址之前的每一条指令。
	Address uint64

	// OpIndex 是 VLIW 指令中操作的索引。第一个操作的索引为 0，对于非 VLIW 架构，该索引始终为 0。 地址和 OpIndex 共同构成一个操作指针，可以引用指令流中的任何单个操作。
	OpIndex int

	// 文件是与这些指令相对应的源文件。
	File *LineFile

	// 行是与这些指令相对应的源代码行号。如果这些指令无法归属于任何源代码行，则行号可以为 0。
	Line int

	// Column 是这些指令源代码行中的列编号。列号从 1 开始，也可以是 0，表示该行的 "左边缘"。
	Column int

	// IsStmt 表示 Address 是一个推荐的断点位置，如一行、语句或语句的一个独特子部分的开头。
	IsStmt bool

	// BasicBlock 表示 Address 是一个基本数据块的开头。
	BasicBlock bool

	// PrologueEnd 表示 Address 是一个（可能是多个）应暂停执行的 PC，以便在进入包含函数时设置断点。
	//
	// 在 DWARF 3 中添加。
	PrologueEnd bool

	// EpilogueBegin 表示 Address 是一个（可能是多个）PC，从该函数退出时应暂停执行以设置断点。
	//
	// 在 DWARF 3 中添加。
	EpilogueBegin bool

	// ISA 是这些指令的指令集架构。可能的 ISA 值应由适用的 ABI 规范定义。
	//
	// 在 DWARF 3 中添加。
	ISA int

	// 判别器是一个任意整数，表示这些指令所属的程序块。它的作用是区分可能具有相同源文件、行和列的多个程序块。如果给定的源位置只存在一个程序块，则它应为 0。
	//
	// 在 DWARF 3 中添加。
	Discriminator int

	// EndSequence 表示 Address 是目标机指令序列结束后的第一个字节。如果设置了该值，则只有该字段和地址字段有意义。行号表可能包含多个可能不相连的指令序列的信息。行号表中的最后一个条目应始终设置 EndSequence。
	EndSequence bool
}
```

LineEntry 是 DWARF 行表中的一行。

### type LineFile 添加于1.5

```go
type LineFile struct {
  Name string
  Mtime uint64 // 执行定义的修改时间，如果未知，则为 0
  Length int // 文件长度，未知时为 0
}
```

行文件是由 DWARF 行表条目引用的源文件。

### type LineReader 添加于1.14

```go
type LineReader struct {
  // 包含已筛选或未导出字段
}
```

LineReader 从 DWARF "行 "部分读取单个编译单元的 LineEntry 结构序列。LineEntry 按 PC 递增的顺序出现，每个 LineEntry 都提供了从该 LineEntry 的 PC 到下一个 LineEntry 的 PC 之前的指令元数据。最后一个条目将设置 EndSequence 字段。

#### func (r *LineReader) Files() []*LineFile 添加于1.5

Files 返回该编译单元在行表中当前位置的文件名表。文件名表可从该编译单元的属性（如 AttrDeclFile）中引用。

由于文件索引 0 表示 "无文件"，因此条目 0 始终为空。

编译单元的文件名表不是固定的。Files 返回行表中当前位置的文件表。它可能比行表中较早位置的文件表包含更多条目，但现有条目不会改变。

#### func (r *LineReader) Next(entry *LineEntry) error 添加于1.5

下一步将 * 条目设置为行表中的下一行，并移动到下一行。如果没有其他条目，且行表已正确结束，则返回 io.EOF。

行总是按照 entry.Address 递增的顺序排列，但 entry.Line 可以向前或向后移动。

#### func (r *LineReader) Reset() 添加于1.5

重置）将行式表阅读器重新置于行式表的起始位置。

#### func (r *LineReader) Seek(pos LineReaderPos) 添加于1.5

查找将行表阅读器恢复到 Tell 返回的位置。

参数 pos 必须是在该行表上调用 Tell 时返回的。

#### func (r *LineReader) SeekPC(pc uint64, entry *LineEntry) error 添加于1.5

SeekPC 将 * 条目设置为包含 pc 的 LineEntry，并将阅读器定位在行表的下一个条目上。如有必要，它会向后查找 pc。

如果 pc 未被该行表中的任何条目覆盖，SeekPC 将返回 ErrUnknownPC。在这种情况下，*条目和最终查找位置是未指定的。

请注意，DWARF 行表只允许向前顺序扫描。因此，在最坏的情况下，所需的时间与行表的大小成线性关系。如果调用者希望重复快速查找 PC，则应在行表中建立适当的索引。

#### func (r *LineReader) Tell() LineReaderPos 添加于1.5

Tell 返回行表中的当前位置。

### type LineReaderPos 添加于1.5

```go
type LineReaderPos struct {
  // 包含已筛选或未导出字段
}
```

LineReaderPos 表示行表中的一个位置。

### type Offset

偏移表示条目在 DWARF 信息中的位置。(请参阅 "阅读器.查找"）。

### type PtrType

```go
type PtrType struct {
  CommonType
  Type Type
}
```

PtrType 表示指针类型。

#### func (t *PtrType) String() string

### type QualType

```go
type QualType struct {
  CommonType
  Qual string
  Type Type
}
```

QualType 表示具有 C/C++ "const"、"restrict "或 "volatile "限定符的类型。

#### func (t *QualType) Size() int64

#### func (t *QualType) String() string

### type Reader

```go
type Reader struct {
  // 包含已筛选或未导出字段
}
```

读取器允许从 DWARF "信息 "部分读取条目结构。条目结构以树状排列。读取器的 Next 函数从树状结构的预排序遍历中返回连续条目。如果一个条目有子条目，则其子条目字段将为 true，而子条目将以标记为 0 的条目为结束。

#### func (r *Reader) AddressSize() int  添加于1.5

AddressSize 返回当前编译单元中地址的大小（以字节为单位）。

#### func (r *Reader) ByteOrder() binary.ByteOrder 添加于1.14

ByteOrder 返回当前编译单元中的字节顺序。

#### func (r *Reader) Next() (*Entry, error)

Next 从编码条目流中读取下一个条目。当读取到该部分的末尾时，会返回 nil、nil。如果当前偏移无效或偏移处的数据无法解码为有效条目，则返回错误。

#### func (r *Reader) Seek(off Offset)

寻道将读取器定位在编码条目流的偏移量处。偏移量 0 可用来表示第一个条目。

#### func (r *Reader) SeekPC(pc uint64) (*Entry, error) 添加于1.7

SeekPC 返回包含 pc 的编译单元的条目，并定位阅读器以读取该单元的子单元。如果 pc 未被任何单元覆盖，SeekPC 将返回 ErrUnknownPC，并且读取器的位置未定义。

由于编译单元可以描述可执行文件的多个区域，因此在最坏的情况下，SeekPC 必须搜索所有编译单元中的所有范围。每次调用 SeekPC 都会从上次调用的编译单元开始搜索，因此一般来说，如果对一系列 PC 进行排序，查找速度会更快。如果调用者希望重复快速查找 PC，则应使用 Ranges 方法建立适当的索引。

#### func (r *Reader) SkipChildren()

SkipChildren 跳过与 Next 返回的最后一个条目相关的子条目。如果该条目没有子条目或 Next 未被调用，SkipChildren 就不起作用。

### type StructField

```go
type StrcutField struct {
  Name string
  Type Type
  ByteOffset int64
  ByteSize int64 // 通常为零；正常字段使用 Type.Size()。
  BitOffset int64
  DataBitOffset int64
  BitSize int64 // 如果不是位域则为零
}
```

StructField 表示 struct、union 或 C++ 类类型中的一个字段。

#### Bit Fields

BitSize、BitOffset 和 DataBitOffset 字段描述了在 C/C++ 结构/联盟/类类型中声明为位字段的数据成员的位大小和偏移量。

BitSize 是位字段的位数。

数据位偏移（DataBitOffset）（如果非零）是从外层实体（例如包含结构体/类/联合体）的起始位置到位字段起始位置的位数。这与 DWARF 4 中引入的 DW_AT_data_bit_offset DWARF 属性相对应。

如果 BitOffset 不为零，那么它就是持有比特字段的存储单元的最有效位到比特字段的最有效位之间的比特数。这里的 "存储单元 "是位字段前的类型名称（对于 "无符号 x:17 "字段，存储单元是 "无符号"）。BitOffset 的值可能因系统的字节序而异。BitOffset 与 DWARF 属性 DW_AT_bit_offset 相对应，该属性在 DWARF 4 中已被弃用，并在 DWARF 5 中被删除。

DataBitOffset 和 BitOffset 中最多有一个非零；只有当 BitSize 非零时，DataBitOffset/BitOffset 才会非零。C 编译器使用其中一个还是另一个取决于编译器版本和命令行选项。

下面是一个 C/C++ 位字段使用示例，以及 DWARF 位偏移信息的预期结果。请看这段代码:

```go
struct S {
	int q;
	int j:5;
	int k:6;
	int m:5;
	int n:8;
} s;
```

在上述代码中，我们可以看到以下 DW_AT_bit_offset 值（使用 GCC 8）：

```go
       Little   |     Big
       Endian   |    Endian
                |
"j":     27     |     0
"k":     21     |     5
"m":     16     |     11
"n":     8      |     16
```

请注意，上面的偏移量纯粹是相对于包含 j/k/m/n 的存储单元而言的--这些值不会因包含结构中先前数据成员的大小而变化。

如果编译器发出 DW_AT_data_bit_offset，预期的值将是:

```go
"j":     32
"k":     37
"m":     43
"n":     48
```

这里，"j "的值 32 反映了比特字段前面还有其他数据成员的事实（请注意，DW_AT_data_bit_offset 值是相对于包含结构体的起始值而言的）。因此，对于字段较多的结构体，DW_AT_data_bit_offset 值可能会很大。

DWARF 还允许基本类型具有非零位大小和位偏移，因此基本类型也会捕获这些信息，但值得注意的是，使用主流语言无法触发这种行为。

### type StructType

```go
type StructType struct {
  CommonType
  StructName string
  Kind string // "结构"、"联盟 "或 "类"。
  Field []*StructField
  Incomplete bool // 如果为 true，则声明了 struct、union、class，但未定义
}
```

StructType 表示结构体、联合体或 C++ 类类型。

#### func (t *StructType) Defn() string

#### func (t *StructType) String() string

### type Tag

```go
type Tag uint32
```

标签是条目的分类（类型）。

```go
const (
	TagArrayType              Tag = 0x01
	TagClassType              Tag = 0x02
	TagEntryPoint             Tag = 0x03
	TagEnumerationType        Tag = 0x04
	TagFormalParameter        Tag = 0x05
	TagImportedDeclaration    Tag = 0x08
	TagLabel                  Tag = 0x0A
	TagLexDwarfBlock          Tag = 0x0B
	TagMember                 Tag = 0x0D
	TagPointerType            Tag = 0x0F
	TagReferenceType          Tag = 0x10
	TagCompileUnit            Tag = 0x11
	TagStringType             Tag = 0x12
	TagStructType             Tag = 0x13
	TagSubroutineType         Tag = 0x15
	TagTypedef                Tag = 0x16
	TagUnionType              Tag = 0x17
	TagUnspecifiedParameters  Tag = 0x18
	TagVariant                Tag = 0x19
	TagCommonDwarfBlock       Tag = 0x1A
	TagCommonInclusion        Tag = 0x1B
	TagInheritance            Tag = 0x1C
	TagInlinedSubroutine      Tag = 0x1D
	TagModule                 Tag = 0x1E
	TagPtrToMemberType        Tag = 0x1F
	TagSetType                Tag = 0x20
	TagSubrangeType           Tag = 0x21
	TagWithStmt               Tag = 0x22
	TagAccessDeclaration      Tag = 0x23
	TagBaseType               Tag = 0x24
	TagCatchDwarfBlock        Tag = 0x25
	TagConstType              Tag = 0x26
	TagConstant               Tag = 0x27
	TagEnumerator             Tag = 0x28
	TagFileType               Tag = 0x29
	TagFriend                 Tag = 0x2A
	TagNamelist               Tag = 0x2B
	TagNamelistItem           Tag = 0x2C
	TagPackedType             Tag = 0x2D
	TagSubprogram             Tag = 0x2E
	TagTemplateTypeParameter  Tag = 0x2F
	TagTemplateValueParameter Tag = 0x30
	TagThrownType             Tag = 0x31
	TagTryDwarfBlock          Tag = 0x32
	TagVariantPart            Tag = 0x33
	TagVariable               Tag = 0x34
	TagVolatileType           Tag = 0x35
	// DWARF 3 新增了以下功能。
	TagDwarfProcedure  Tag = 0x36
	TagRestrictType    Tag = 0x37
	TagInterfaceType   Tag = 0x38
	TagNamespace       Tag = 0x39
	TagImportedModule  Tag = 0x3A
	TagUnspecifiedType Tag = 0x3B
	TagPartialUnit     Tag = 0x3C
	TagImportedUnit    Tag = 0x3D
	TagMutableType     Tag = 0x3E // 后从 DWARF 中删除。
	TagCondition       Tag = 0x3F
	TagSharedType      Tag = 0x40
	// DWARF 4 新增了以下功能。
	TagTypeUnit            Tag = 0x41
	TagRvalueReferenceType Tag = 0x42
	TagTemplateAlias       Tag = 0x43
	// DWARF 5 新增了以下功能。
	TagCoarrayType       Tag = 0x44
	TagGenericSubrange   Tag = 0x45
	TagDynamicType       Tag = 0x46
	TagAtomicType        Tag = 0x47
	TagCallSite          Tag = 0x48
	TagCallSiteParameter Tag = 0x49
	TagSkeletonUnit      Tag = 0x4A
	TagImmutableType     Tag = 0x4B
)
```

#### func (t Tag) GoString() string

#### func (t Tag) String() string

### type Type

```go
type Type interface {
  Common() *CommonType
  String() string
  Size() int64
}
```

类型通常表示指向任何特定类型结构（CharType、StructType 等）的指针。

### type TypedeType

```go
type TypedeType struct {
  CommonType
  Type Type
}
```

TypedefType 表示已命名的类型。

#### func (t *TypedeType) Size() int64

#### func (t *TypedeType) String() string

### type UcharType

```go
type UcharType struct {
  BasicType
}
```

UcharType 表示无符号字符类型。

### type UintType

```go
type UintType struct {
  BasicType
}
```

UintType 表示无符号整数类型。

### type UnspecifiedType 添加于1.4

```go
type UnspecifiedType struct {
  BasicType
}
```

UnspecifiedType 表示隐含、未知、模糊或不存在的类型。

### type UnsupportedType 添加于1.13

```go
type UnsupportedType struct {
  CommonType
  Tag Tag
}
```

UnsupportedType 是在遇到不支持的类型时返回的占位符。

#### func (t *UnsupportedType) String() string 添加于1.13

### type VoidType

```go
type VoidType struct {
  CommonType
}
```

VoidType 表示 C 语言的 void 类型。

#### func (t *VoidType) String() string