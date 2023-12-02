## package xml

包 xml 实现了一个简单的 XML 1.0 解析器，它能理解 XML 名称空间。

## Index

### Constants

```go
const (
  // Header 是一个通用的 XML 标头，适合与 Marshal 的输出一起使用。它不会自动添加到本软件包的任何输出中，而是作为一种便利提供。
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)
```

### Variables

```go
var HTMLAutoClose []string = htmlAutoClose
```

HTMLAutoClose 是应考虑自动关闭的 HTML 元素集。

请参阅 Decoder.Strict 和 Decoder.Entity 字段的文档。

```go
var HTMLEntity map[string]string = htmlEntity
```

HTMLEntity 是一个实体图，包含标准 HTML 实体字符的翻译。

请参阅 Decoder.Strict 和 Decoder.Entity 字段的文档。

### func Escape(w io.Writer, s []byte)

Escape 与 EscapeText 类似，但省略了错误返回值。提供它是为了向后兼容 Go 1.0。Go 1.1 或更高版本的代码应使用 EscapeText。

### func EscapeText(w io.Writer, s []byte) error 添加于1.1

EscapeText 会向 w 写入经适当转义的 XML 等效纯文本数据 s。

### func Marshal(v any) ([]byte, error)

Marshal 返回 v 的 XML 编码。

Marshal 处理数组或片段时，会对每个元素进行编译。Marshal 处理指针时，会对指针指向的值进行标注，如果指针为 nil，则不会写入任何内容。Marshal 处理接口值时，会对其包含的值进行编译，如果接口值为 nil，则不会写入任何内容。通过写入一个或多个包含数据的 XML 元素，Marshal 可以处理所有其他数据。

XML 元素的名称按优先顺序取自以下内容

- XMLName 字段的标记（如果数据是结构体
- XMLName 字段中 Name 类型的值
- 用于获取数据的结构体字段的标记
- 用于获取数据的结构体字段名
- 被修饰类型的名称
- 结构体的 XML 元素包含结构体每个导出字段的修饰元素，但以下情况除外：

- 省略上述 XMLName 字段。
- 省略标记为"-"的字段。
- 带有标记 "name,attr "的字段将成为 XML 元素中带有给定名称的属性。
- 带标记",attr "的字段将成为 XML 元素中带有字段名称的属性。
- 带有标记",chardata "的字段将作为字符数据而不是 XML 元素写入。
- 带标记",cdata "的字段以字符数据的形式写入，用一个或多个 <![CDATA[ ... ]]> 标记包裹，而不是 XML 元素。
- 标记为",innerxml "的字段是逐字写入的，不采用通常的编排程序。
- 标记为",comment "的字段是作为 XML 注释写入的，不受通常的编排程序限制。其中不得包含"--"字符串。
- 如果字段值为空，则省略包含 "omitempty "选项的标记字段。空值包括 false、0、任何 nil 指针或接口值，以及长度为零的任何数组、片、映射或字符串。
- 匿名结构体字段的处理方式与外层结构体字段的处理方式相同。
- 通过调用 MarshalXML 方法来编写实现 Marshaler 的字段。
- 实现 encoding.TextMarshaler 的字段是通过将其 MarshalText 方法的结果编码为文本来编写的。

如果一个字段使用 "a>b>c "标签，那么元素 c 将嵌套在父元素 a 和 b 内。

如果结构体字段的 XML 名称是由字段标记和结构体的 XMLName 字段定义的，那么这两个名称必须匹配。

有关示例，请参见 MarshalIndent。

如果要求 Marshal marshal 一个通道、函数或映射，Marshal 将返回错误。

### func MarshalIndent(v any, prefix, indent string) ([]byte, error)

MarshalIndent 的工作原理与 Marshal 类似，但每个 XML 元素都以新的缩进行开始，新的缩进行以前缀开头，并根据嵌套深度跟随一个或多个缩进副本。

### func Unmarshal(data []byte, v any) error

Unmarshal 会解析 XML 编码数据，并将结果存储到 v 指向的值中，v 必须是任意结构体、片段或字符串。不适合 v 的格式良好的数据将被丢弃。

由于 Unmarshal 使用 reflect 包，因此只能对导出（大写）字段进行赋值。Unmarshal 使用大小写敏感比较法，将 XML 元素名称与标记值和结构字段名匹配。

Unmarshal 使用以下规则将 XML 元素映射到结构体。在这些规则中，字段的标记指的是与 struct 字段标记中的键 "xml "相关联的值（见上面的示例）。

- 如果结构体中有[]字节或字符串类型的字段，且标记为",innerxml"，则 Unmarshal 会在该字段中累积嵌套在元素内部的原始 XML。其余规则仍然适用。

- 如果结构体有一个名为 XMLName 的 Name 类型字段，Unmarshal 会在该字段中记录元素名称。

- 如果 XMLName 字段具有 "name "或 "namespace-URL name "形式的关联标记，则 XML 元素必须具有给定的名称（以及可选的名称空间），否则 Unmarshal 将返回错误信息。

- 如果 XML 元素有一个属性，其名称与结构字段名相匹配，而结构字段名的关联标记包含",attr"，或者与结构字段标记形式为 "name,attr "的显式名称相匹配，则 Unmarshal 会在该字段中记录属性值。

- 如果 XML 元素具有前一条规则未处理的属性，且结构体具有关联标记为",any,attr "的字段，则 Unmarshal 会在第一个此类字段中记录属性值。

- 如果 XML 元素包含字符数据，则该数据会累积到第一个标记为",chardata "的 struct 字段中。结构字段的类型可以是[]字节或字符串。如果没有这样的字段，字符数据将被丢弃。

- 如果 XML 元素中包含注释，这些注释会被累加到第一个标记为",comment "的 struct 字段中。结构字段的类型可以是[]字节或字符串。如果没有这样的字段，注释将被丢弃。

- 如果 XML 元素包含一个子元素，其名称与格式为 "a "或 "a>b>c "的标记前缀相匹配，unmarshal 将深入 XML 结构，查找具有给定名称的元素，并将最内层的元素映射到该结构字段。以">"开头的标签等同于以字段名开头、后跟">"的标签。

- 如果 XML 元素包含一个子元素，其名称与结构字段的 XMLName 标记相匹配，而结构字段没有前一条规则中明确的名称标记，则 unmarshal 会将该子元素映射到该结构字段。

- 如果 XML 元素包含一个子元素，而该子元素的名称与一个没有任何模式标记（",attr"、",chardata "等）的字段相匹配，则 Unmarshal 会将该子元素映射到该 struct 字段。

- 如果 XML 元素包含的子元素不符合上述任何规则，且结构体有一个标记为",any "的字段，则解屏蔽会将该子元素映射到该结构体字段。

- 匿名结构体字段的处理方式与外部结构体字段的处理方式相同。

- 标记为"-"的结构体字段永远不会被解除映射。

如果 Unmarshal 遇到实现 Unmarshaler 接口的字段类型，Unmarshal 会调用其 UnmarshalXML 方法从 XML 元素中生成值。否则，如果值实现了 encoding.TextUnmarshaler，Unmarshal 会调用该值的 UnmarshalText 方法。

Unmarshal 将 XML 元素映射到字符串或[]字节，方法是将该元素的字符数据连接保存到字符串或[]字节中。保存的 []byte 绝不为空。

通过将属性值保存在字符串或片段中，将属性值映射为字符串或[]字节。

Unmarshal 将属性值映射到 Attr 时，会将属性（包括其名称）保存在 Attr 中。

Unmarshal 通过扩展片段长度并将元素或属性映射到新创建的值，将 XML 元素或属性值映射到片段。

Unmarshal 通过将 XML 元素或属性值设置为字符串所代表的布尔值，将其映射为布尔值。空白会被修剪并忽略。

Unmarshal 将 XML 元素或属性值映射到整数或浮点型字段，方法是将字段设置为十进制字符串值的解释结果。没有溢出检查。空白会被修剪并忽略。

Unmarshal 通过记录元素名称将 XML 元素映射为名称。

Unmarshal 通过将指针设置为新分配的值，然后将元素映射到该值，从而将 XML 元素映射到指针。

缺失的元素或空属性值将作为零值解除映射。如果字段是片段，零值将被附加到字段中。否则，字段将被设置为零值。

### type Attr

```go
type Attr struct {
  Name Name
  Value string
}
```

Attr 表示 XML 元素中的一个属性（Name=Value）。

### type CharData

```go
type CharData []byte
```

CharData 表示 XML 字符数据（原始文本），其中 XML 转义序列已被其所代表的字符替换。

#### func (c CharData) Copy() CharData

复制会创建 CharData 的新副本。

### type Comment

```go
type Comment []byte
```

注释表示形式为 <!--comment--> 的 XML 注释。字节不包括 <!-- 和 --> 注释标记。

#### func (c Comment) Copy() Comment

复制会创建一个新的 "注释 "副本。

### type Decoder

```go
type Decoder struct {
  // 严格默认为 true，强制执行 XML 规范的要求。如果设置为 false，解析器允许输入包含常见
  //  错误：
  // 	* 如果元素缺少结束标记，解析器会根据需要创建结束标记，以保持 Token 返回值的适当平衡。
  //  * 在属性值和字符数据中，未知或畸形的字符实体（以 & 开头的序列）不会被处理。
  // 
  // 设置:
  // 
	//	d.Strict = false
	//	d.AutoClose = xml.HTMLAutoClose
	//	d.Entity = xml.HTMLEntity
  // 
  // 创建了一个可以处理典型 HTML 的解析器。
  // 严格模式不执行 XML 名称空间 TR 的要求。特别是，它不拒绝使用未定义前缀的名称空间标记。此类标记会以未知前缀作为名称空间 URL 记录。
  Strict bool

  // 严格模式不执行 XML 名称空间 TR 的要求。特别是，它不拒绝使用未定义前缀的名称空间标记。此类标记会以未知前缀作为名称空间 URL 记录。
  AutoClose []string

  // // Entity 可用于将非标准实体名称映射为字符串替换。无论映射表的实际内容如何，解析器都会像映射表中存在这些标准映射一样进行处理：
	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,
	Entity map[string]string

  // CharsetReader（若非为零）定义了一个函数，用于生成字符集转换阅读器，将所提供的非 UTF-8 字符集转换为 UTF-8 字符集。如果 CharsetReader 为 nil 或返回错误，解析过程将以错误停止。CharsetReader 的结果值之一必须为非 nil。
  CharsetReader func(charset string, input io.Reader) (io.Reader, error)

  // DefaultSpace 设置用于未加修饰标记的默认名称空间，就像整个 XML 流被包裹在包含属性 xmlns="DefaultSpace" 的元素中一样。
  DefaultSpace string
  // 包含已过滤或未导出的字段
}
```

解码器表示读取特定输入流的 XML 解析器。该解析器假定其输入是以 UTF-8 编码的。

#### func NewDecoder(r io.Reader) *Decoder

如果 r 没有实现 io.ByteReader，NewDecoder 将自己进行缓冲。

#### func NewTokenDecoder(t TokenReader) *Decoder 添加于1.10

NewTokenDecoder 使用底层标记流创建新的 XML 解析器。

#### func (d *Decoder) Decoder(v any) error

解码器的工作原理与 Unmarshal 类似，只不过它通过读取解码器数据流来查找起始元素。

#### func (d *Decoder) DecodeElement(v any, start *StartElement) error

DecodeElement 的工作方式与 Unmarshal 类似，但它需要一个指向要解码为 v 的 XML 元素起始部分的指针。当客户端自己读取一些原始 XML 标记，但又想将某些元素延迟到 Unmarshal 时，DecodeElement 就会派上用场。

#### func (d *Decoder) InputOffset() int64 添加于1.4

InputOffset 返回当前解码器位置的输入流字节偏移量。偏移量给出了最近返回的标记的结束位置和下一个标记的开始位置。

#### func (d *Decoder) InputPos() (line, column int) 添加于1.19

InputPos 返回当前解码器位置的行以及该行基于 1 的输入位置。位置给出了最近返回的标记的结束位置。

#### func (d *Decoder) RawToken() (Token, error)

RawToken 与 Token 类似，但不会验证开始和结束元素是否匹配，也不会将名称空间前缀转换为相应的 URL。

#### func (d *Decoder) Skip() error

Skip 会读取标记，直到读完与已读取的最新起始元素相匹配的结束元素为止，同时跳过嵌套结构。如果找到与起始元素匹配的结束元素，则返回 nil；否则返回错误信息，说明问题所在。

#### func (d *Decoder) Token() (Token, error)

Token 返回输入流中的下一个 XML 标记。在输入流结束时，Token 会返回 nil，即 io.EOF。

返回的标记数据中的字节片段指的是解析器的内部缓冲区，仅在下一次调用 Token 之前有效。要获取字节的副本，请调用 CopyToken 或令牌的 Copy 方法。

令牌会将 <br> 等自闭合元素扩展为连续调用返回的独立的开始和结束元素。

令牌保证它返回的 StartElement 和 EndElement 令牌是正确嵌套和匹配的：如果令牌遇到意外的结束元素或在所有预期的结束元素之前遇到 EOF，它将返回错误。

如果 CharsetReader 被调用并返回错误，该错误将被封装并返回。

Token 实现了 https://www.w3.org/TR/REC-xml-names/ 所描述的 XML 名称空间。令牌中包含的每个名称结构都将空间（Space）设置为识别其已知名称空间的 URL。如果令牌遇到未识别的名称空间前缀，它会使用该前缀作为空间，而不是报告错误。

### type Directive

```go
type Directive []byte
```

指令表示形式为 <!text> 的 XML 指令。字节不包括 <!

#### func (d Directive) Copy() Directive

复制 "会创建一个新的 "指令 "副本。

### type Encoder

```go
type Encoder struct {
  // 包含已过滤或未导出的字段
}
```

编码器将 XML 数据写入输出流。

#### func NewEncoder(w io.Writer) *Encoder

NewEncoder 返回写入 w 的新编码器。

#### func (enc *Encoder) Close() error 添加于1.20

关闭编码器，表示不再写入数据。它将缓冲的 XML 清除到底层写入器，如果写入的 XML 无效（例如包含未关闭的元素），则返回错误信息。

#### func (enc *Encoder) Encode(v any) error

Encode 将 v 的 XML 编码写入数据流。

有关将 Go 值转换为 XML 的详细信息，请参阅 Marshal 文档。

在返回之前，Encode 会调用 Flush。

#### func (enc *Encoder) EncodeElement(v any, start StartElement) error 添加于1.2

EncodeElement 将 v 的 XML 编码写入数据流，在编码中使用 start 作为最外层标签。

有关将 Go 值转换为 XML 的详细信息，请参阅 Marshal 文档。

在返回之前，EncodeElement 会调用 Flush。

#### func (enc *Encoder) EncodeToken(t Token) error 添加于1.2

EncodeToken 将给定的 XML 标记写入数据流。如果 StartElement 和 EndElement 标记不匹配，则返回错误信息。

EncodeToken 不会调用 "刷新"（Flush），因为它通常是 Encode 或 EncodeElement（或在这些操作中调用的自定义 MarshalXML）等较大操作的一部分，这些操作完成后会调用 "刷新"（Flush）。如果调用者创建了编码器，然后直接调用 EncodeToken，而没有使用 Encode 或 EncodeElement，则需要在完成后调用 Flush，以确保 XML 被写入底层写入器。

EncodeToken 只允许将目标设置为 "xml "的 ProcInst 作为流中的第一个标记写入。

#### func (enc *Encoder) Flush() error 添加于1.2

将任何缓冲的 XML 冲洗到底层写入器。有关何时需要这样做的详细信息，请参阅 EncodeToken 文档。

#### func (enc *Encoder) Indent(prefix, indent string) 添加于1.1

缩进（Indent）设置编码器生成 XML 时，每个元素都以新的缩进行开始，新缩进行以前缀开头，并根据嵌套深度在其后添加一份或多份缩进。

### type EndElement

```go
type EndElement struct {
  Name Name
}
```

EndElement 表示一个 XML 结束元素。

### type Marshaler 添加于1.2

```go
type Marshaler interface {
	MarshalXML(e *Encoder, start StartElement) error
}
```

Marshaler 是对象实现的接口，这些对象可以将自己 Marshal 成有效的 XML 元素。

MarshalXML 将接收者编码为零个或多个 XML 元素。按照惯例，数组或片段通常编码为元素序列，每个条目一个元素。使用 start 作为元素标记不是必须的，但这样做可以使 Unmarshal 将 XML 元素与正确的 struct 字段相匹配。一种常见的实施策略是构建一个单独的值，其布局与所需的 XML 相对应，然后使用 e.EncodeElement 对其进行编码。另一种常见策略是重复调用 e.EncodeToken，一次生成一个 XML 标记。编码标记的序列必须由零个或多个有效的 XML 元素组成。

### type MarshalerAttr 添加于1.2

```go
type MarshalerAttr interface {
	MarshalXMLAttr(name Name) (Attr, error)
}
```

MarshalerAttr 是对象实现的接口，这些对象可以将自己 Marshal 成有效的 XML 属性。

MarshalXMLAttr 返回带有接收器编码值的 XML 属性。虽然不需要使用 name 作为属性名，但这样做可以使 Unmarshal 将属性与正确的结构字段匹配起来。如果 MarshalXMLAttr 返回零属性 Attr{}，输出中将不会生成任何属性。MarshalXMLAttr 仅用于字段标记中带有 "attr "选项的 struct 字段。

### type Name

```go
type Name struct {
	Space, Local string
}
```

一个名称代表一个 XML 名称（本地），并附有一个名称空间标识符（空间）。在 Decoder.Token 返回的标记中，Space 标识符以规范 URL 的形式给出，而不是解析文档中使用的短前缀。

### type ProcInst

```go
type ProcInst struct {
	Target string
	Inst   []byte
}
```

ProcInst 表示格式为 <?target inst?> 的 XML 处理指令。

#### func (p ProcInst) Copy() ProcInst

副本创建 ProcInst 的新副本。

### type StartElement

```go
type StartElement struct {
	Name Name
	Attr []Attr
}
```

StartElement 表示一个 XML 开始元素。

#### func (e StartElement) Copy() StartElement

复制会创建 StartElement 的新副本。

#### func (e StartElement) End() EndElement 添加于1.2

End 返回相应的 XML 结束元素。

### type SyntaxError

```go
type SyntaxError struct {
	Msg  string
	Line int
}
```

SyntaxError 表示 XML 输入流中的语法错误。

#### func (e *SyntaxError) Error() string

### type TagPathError

```go
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}
```

标签路径错误（TagPathError）表示由于使用了路径冲突的字段标签而导致的解屏蔽过程中的错误。

#### func (e *TagPathError) Error() string

### type Token

```go
type Token any
```

令牌是持有令牌类型之一的接口：StartElement、EndElement、CharData、Comment、ProcInst 或 Directive。

#### func CopyToken(t Token) Token

CopyToken 返回令牌的副本。

### type TokenReader 添加于1.10

```go
type TokenReader interface {
	Token() (Token, error)
}
```

令牌阅读器是指任何可以解码 XML 令牌流的设备，包括解码器。

当令牌成功读取令牌后遇到错误或文件结束条件时，它会返回令牌。它可以在同一次调用中返回错误（非零），也可以在后续调用中返回错误（和一个零标记）。这种一般情况的一个例子是，令牌阅读器在令牌流结束时返回一个非零令牌，可能会返回 io.EOF 或一个零错误。下一次读取应返回 nil 或 io.EOF。

不鼓励令牌的实现在返回 nil 错误时返回 nil 令牌。调用者应将返回 nil、nil 表示什么都没发生，尤其是不表示 EOF。

### type UnmarshalError

```go
type UnmarshalError string
```

UnmarshalError 表示解分级过程中出现错误。

#### func (e UnmarshalError) Error() string

### type Unmarshaler 添加于1.2

```go
type Unmarshaler interface {
	UnmarshalXML(d *Decoder, start StartElement) error
}
```

Unmarshaler 是一个接口，由可以解除 XML 元素描述的对象实现。

UnmarshalXML 从给定的起始元素开始解码单个 XML 元素。如果它返回错误，外部调用 Unmarshal 的过程就会停止，并返回错误信息。UnmarshalXML 必须恰好消耗一个 XML 元素。一种常见的实现策略是使用 d.DecodeElement 将 Unmarshal 转换为一个单独的值，该值的布局与预期的 XML 相匹配，然后将数据从该值复制到接收器中。另一种常见策略是使用 d.Token 一次处理一个 XML 对象。UnmarshalXML 可能不会使用 d.RawToken。

### type UnmarshalerAttr 添加于1.2

```go
type UnmarshalerAttr interface {
	UnmarshalXMLAttr(attr Attr) error
}
```

UnmarshalerAttr 是一个接口，由可以解除自身 XML 属性描述的对象实现。

UnmarshalXMLAttr 对单个 XML 属性进行解码。如果它返回一个错误，外部对 Unmarshal 的调用就会停止并返回该错误。UnmarshalXMLAttr 仅用于字段标记中带有 "attr "选项的 struct 字段。

### type UnsupportedTypeError

```go
type UnsupportedTypeError struct {
	Type reflect.Type
}
```

当 Marshal 遇到无法转换为 XML 的类型时，将返回 UnsupportedTypeError。

#### func (e *UnsupportedTypeError) Error() string

### Bugs

XML 元素和数据结构之间的映射本身就存在缺陷：XML 元素是匿名值的有序集合，而数据结构是命名值的无序集合。有关更适合数据结构的文本表示法，请参阅软件包 json。