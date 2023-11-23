## package asn1

根据 ITU-T Rec X.690 的定义，软件包 asn1 实现了对 DER 编码 ASN.1 数据结构的解析。

另请参阅 "A Layman's Guide to a Subset of ASN.1、BER 和 DER"，http://luca.ntop.org/Teaching/Appunti/asn1.html。

## Index

### Constants

```go
const (
	TagBoolean         = 1
	TagInteger         = 2
	TagBitString       = 3
	TagOctetString     = 4
	TagNull            = 5
	TagOID             = 6
	TagEnum            = 10
	TagUTF8String      = 12
	TagSequence        = 16
	TagSet             = 17
	TagNumericString   = 18
	TagPrintableString = 19
	TagT61String       = 20
	TagIA5String       = 22
	TagUTCTime         = 23
	TagGeneralizedTime = 24
	TagGeneralString   = 27
	TagBMPString       = 30
)
```

ASN.1 标记表示以下对象的类型。

```go
const (
	ClassUniversal       = 0
	ClassApplication     = 1
	ClassContextSpecific = 2
	ClassPrivate         = 3
)
```

ASN.1 类类型代表标签的命名空间。

### Variables

```go
var NullBytes = []byte{TagNull, 0}
```

NullBytes 包含表示 DER 编码 ASN.1 NULL 类型的字节。

```go
var NullRawValue = RawValue{Tag: TagNull}
```

NullRawValue 是一个 RawValue，其 Tag 设置为 ASN.1 NULL 类型标记 (5)。

### func Marshal(val nay) ([]byte, error)

Marshal 返回 val 的 ASN.1 编码。

除了 Unmarshal 可以识别的 struct 标记外，还可以使用以下标记：

```go
ia5:         causes strings to be marshaled as ASN.1, IA5String values
omitempty:   causes empty slices to be skipped
printable:   causes strings to be marshaled as ASN.1, PrintableString values
utf8:        causes strings to be marshaled as ASN.1, UTF8String values
utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
```

### func MarshalWithParams(val any, params string) ([]byte, error) 添加于1.10

MarshalWithParams 允许为顶层元素指定字段参数。参数的形式与字段标记相同。

### func Unmarshal(b []byte, val any) (rest []byte, err error)

Unmarshal 会解析经 DER 编码的 ASN.1 数据结构 b，并使用 reflect 包填充由 val 指向的任意值。由于 Unmarshal 使用 reflect 包，被写入的结构体必须使用大写字段名。如果 val 为 nil 或不是指针，Unmarshal 将返回错误。

解析 b 后，未用于填充 val 的剩余字节将以 rest 的形式返回。在将 SEQUENCE 解析为 struct 时，SEQUENCE 的任何尾部元素，如果在 val 中没有匹配字段，都不会包含在 rest 中，因为这些元素被视为 SEQUENCE 的有效元素，而不是尾部数据。

ASN.1 INTEGER 可以写入 int、int32、int64 或 *big.Int（来自 math/big 软件包）。如果编码值不适合 Go 类型，Unmarshal 会返回解析错误。

ASN.1 BIT STRING 可以写入 BitString。

一个 ASN.1 OCTET STRING 可以写入一个 []字节。

ASN.1 OBJECT IDENTIFIER 可写入 ObjectIdentifier。

ASN.1 ENUMERATED 可以写入枚举。

可将 ASN.1 PrintableString、IA5String 或 NumericString 写入字符串。

上述任何 ASN.1 值都可以写入接口{}。存储在接口中的值具有相应的 Go 类型。对于整数，该类型为 int64。

如果 x 可以写入片段的元素类型，那么 ASN.1 序列 x 或集合 x 就可以写入片段。

如果序列中的每个元素都能写入结构体中的相应元素，那么 ASN.1 序列或 SET 就能写入结构体。

结构体字段上的下列标记对 Unmarshal 有特殊意义：

```go
application 指定使用 APPLICATION 标签
private 指定使用 PRIVATE 标记
default:x 设置可选整数字段的默认值（仅在可选字段也存在时使用）
explicit 指定使用附加的显式标记来封装隐式标记
optional 将字段标记为 ASN.1 OPTIONAL
set 表示 SET 类型，而不是 SEQUENCE 类型
tag:x指定 ASN.1 标记编号；意味着 ASN.1 CONTEXT SPECIFIC
```

将带有 IMPLICIT 标记的 ASN.1 值解码为字符串字段时，Unmarshal 将默认使用 PrintableString，它不支持"@"和"&"等字符。要强制使用其他编码，请使用以下标记：

```go
ia5 会将字符串解屏蔽为 ASN.1 IA5String 值
numeric 会将字符串解码为 ASN.1 NumericString 值
utf8 会将字符串解码为 ASN.1 UTF8String 值
```

如果结构体第一个字段的类型是 RawContent，那么结构体的原始 ASN1 内容将被存储在其中。

如果片段类型的名称以 "SET "结尾，那么它将被视为已设置了 "set "标记。这将导致该类型被解释为 x 的集合（SET OF x），而不是 x 的序列（SEQUENCE OF x）。

Unmarshal 不支持其他 ASN.1 类型；如果遇到这些类型，Unmarshal 会返回解析错误。

### func UnmarshalWithParams(b []byte, val any, params string) (rest []byte, err error)

UnmarshalWithParams 允许为顶层元素指定字段参数。参数的形式与字段标记相同。

### type BitString

```go
type BitString struct {
  Bytes []byte // 位打包成字节。
  BitLength int // 长度（比特）。
}
```

当需要 ASN.1 位字符串类型时，可以使用 BitString 结构。位字符串会被填充到内存中最接近的字节，并记录有效位的数量。填充位将为零。

#### func (b BitString) At(i int) int

At 返回给定索引处的位。如果索引超出范围，则返回 0。

#### func (b BitString) RightAlign() []byte

RightAlign 返回填充位在开头的片段。该片可能与 BitString 共享内存。

### type Enumerated

```go
type Enumerated int
```

枚举表示为普通 int。

### type Flag

```go
type Flag bool
```

Flag 接受任何数据，如果存在，则设置为 true。

### type ObjectIdentifier

```go
type ObjectIdentifier []int
```

ObjectIdentifier 表示 ASN.1 物件标识符。

#### func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool

Equal 报告 oi 和 other 是否代表相同的标识符。

#### func (oi ObjectIdentifier) String() string 添加于1.3

### type RawContent

```go
type RawContent []byte
```

RawContent 用于表示需要为结构体保留未解码的 DER 数据。要使用它，结构体的第一个字段必须具有这种类型。如果其他字段也使用这种类型，则属于错误。

### type RawValue

```go
type RawValue struct {
  Class, Tag int
  IsCompound bool
  Bytes []byte
  FullBytes []byte // 包括标签和长度
}
```

RawValue 表示未解码的 ASN.1 对象。

### type StructuralError

```go
type StructuralError struct {
  Msg string
}
```

结构错误（StructuralError）表明 ASN.1 数据有效，但接收数据的 Go 类型不匹配。

#### func (e StructuralError) Error() string

### type SyntaxError

```go
type SyntaxError struct {
  Msg string
}
```

语法错误提示 ASN.1 数据无效。

#### func (e SyntaxError) Error() string