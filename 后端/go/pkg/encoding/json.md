## package json

包 json 实现了 RFC 7159 中定义的 JSON 编码和解码。JSON 和 Go 值之间的映射关系在 Marshal 和 Unmarshal 函数的文档中有所描述。

有关此软件包的介绍，请参阅 "JSON 与 Go"： https://golang.org/doc/articles/json_and_go.html

## Index

### func Compact(dst *bytes.Buffer, src []byte) error

Compact 会将 JSON 编码的 src 添加到 dst 中，并省略不重要的空格字符。

### func HTMLeSCAPE(dst *bytes.Buffer, src []byte)

HTMLEscape 会在 dst 中添加 JSON 编码的 src，并将字符串字面内的 <、>、&、U+2028 和 U+2029 字符更改为 \u003c、\u003e、\u0026、\u2028、\u2029，以便 JSON 可以安全地嵌入 HTML <script> 标记中。由于历史原因，网络浏览器不支持 <script> 标记内的标准 HTML 转义，因此必须使用另一种 JSON 编码。

### func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error

缩进将 JSON 编码 src 的缩进形式附加到 dst。JSON 对象或数组中的每个元素都以新的缩进行开始，以前缀开始，后面根据缩进嵌套复制一份或多份缩进。附加到 dst 的数据不以前缀或任何缩进开始，以便于嵌入到其他格式化的 JSON 数据中。虽然 src 开头的前导空格字符（空格、制表符、回车符、换行符）会被删除，但 src 末尾的尾部空格字符会被保留并复制到 dst。例如，如果 src 没有尾部空格，dst 也不会有尾部空格；如果 src 以尾部换行结束，dst 也不会有尾部换行。

### func Marshal(v any) ([]byte, error)

Marshal 返回 v 的 JSON 编码。

Marshal 会递归遍历 v 值。如果遇到的值实现了 Marshaler 接口且不是 nil 指针，Marshal 会调用 MarshalJSON 方法生成 JSON。如果没有 MarshalJSON 方法，但值实现了 encoding.TextMarshaler 接口，Marshal 就会调用 MarshalText 方法，将结果编码为 JSON 字符串。严格来说，nil 指针异常并非必要，但它模仿了 UnmarshalJSON 行为中类似的必要异常。

否则，Marshal 会使用以下与类型相关的默认编码：

布尔值编码为 JSON 布尔值。

浮点数、整数和数字值编码为 JSON 数字。NaN 和 +/-Inf 值将返回 UnsupportedValueError。

字符串值以 JSON 字符串的形式编码，并强制转换为有效的 UTF-8，用 Unicode 替换符替换无效字节。为了使 JSON 可以安全地嵌入 HTML <script> 标记，字符串使用 HTMLEscape 编码，将"<"、">"、"&"、U+2028 和 U+2029 转义为"\u003c"、"\u003e"、"\u0026"、"\u2028 "和"\u2029"。在使用编码器时，可以通过调用 SetEscapeHTML(false) 来禁用这种替换。

数组和片段值编码为 JSON 数组，但 []byte 编码为 base64 编码字符串，nil 片段编码为 null JSON 值。

结构体值编码为 JSON 对象。每个导出的结构体字段都会成为对象的成员，并使用字段名作为对象键，除非该字段因以下原因而被省略。

每个 struct 字段的编码都可以通过存储在 struct 字段标记中 "json "键下的格式字符串进行自定义。格式字符串给出了字段的名称，后面可能还有一个以逗号分隔的选项列表。名称可以为空，以便在不覆盖默认字段名称的情况下指定选项。

如果字段的值为空（定义为 false、0、nil 指针、nil 接口值，以及任何空数组、片、映射或字符串），则 "omitempty "选项指定从编码中省略该字段。

作为特例，如果字段标记为"-"，则字段总是被省略。请注意，名称为"-"的字段仍可使用标签"-, "生成。

struct 字段标记及其含义示例：

```go
// 字段在 JSON 中显示为关键字 "myName"。
Field int `json: "myName"`

// 字段作为键 "myName "出现在 JSON 中。
// 如果字段的值为空，对象中将省略该字段、
// 如上定义。
Field int `json: "myName,omitempty"`

// 字段在 JSON 中显示为关键字 "Field"（默认），但 // 如果字段为空，则跳过该字段。
// 如果为空，则跳过该字段。
// 注意前导逗号。
Field int `json:",omitempty"`

// 本软件包忽略 Field。
Field int `json:"-"`

// 字段在 JSON 中显示为关键字"-"。
Field int `json:"-,"`
```

字符串 "选项表示字段以 JSON 编码字符串的形式存储。它只适用于字符串、浮点、整数或布尔类型的字段。在与 JavaScript 程序通信时，有时会用到这种额外的编码级别：

```go
Int64String int64 `json:",string"`
```

如果键名是由 Unicode 字母、数字和 ASCII 标点符号（引号、反斜杠和逗号除外）组成的非空字符串，则将使用该键名。

匿名结构体字段通常被当作外层结构体中的字段来编排，就像其内部导出字段一样，并遵守下一段所述的经过修订的 Go 可见性规则。在 JSON 标记中给出名称的匿名结构体字段将被视为具有该名称，而不是匿名。接口类型的匿名结构体字段与匿名结构体字段的名称相同。

在决定对哪个字段进行标记（marshal）或取消标记（unmarshal）时，针对 struct 字段的 Go 可见性规则将针对 JSON 进行修改。如果在同一层次上有多个字段，且该层次的嵌套最少（因此是通常 Go 规则所选择的嵌套层次），则适用以下额外规则：

1) 在这些字段中，如果有任何字段带有 JSON 标记，则只考虑带有标记的字段，即使有多个未标记的字段会发生冲突。

2) 如果正好有一个字段（根据第一条规则是否已标记），则选择该字段。

3) 否则有多个字段，所有字段都会被忽略；不会发生错误。

处理匿名 struct 字段是 Go 1.1 的新特性。在 Go 1.1 之前，匿名 struct 字段被忽略。要在当前版本和早期版本中强制忽略匿名 struct 字段，请给该字段加上"-"的 JSON 标记。

映射值编码为 JSON 对象。映射的键类型必须是字符串、整数类型或实现 encoding.TextMarshaler。通过应用以下规则对映射键进行排序并将其用作 JSON 对象键，但须遵守上文所述的字符串值 UTF-8 强制规则：

- 直接使用任何字符串类型的键
- 编码.TextMarshalers 被 marshaled
- 整数键转换为字符串

指针值编码为所指向的值。nil 指针编码为 JSON 空值。

接口值编码为接口中包含的值。接口值为零，则编码为空 JSON 值。

通道、复数和函数值不能用 JSON 编码。如果尝试对此类值进行编码，Marshal 将返回 UnsupportedTypeError（不支持类型错误）。

JSON 无法表示循环数据结构，因此 Marshal 无法处理它们。向 Marshal 传递循环结构将导致错误。

### func MarshalIndent([]byte, error)

MarshalIndent 与 Marshal 类似，但使用缩进来格式化输出。输出中的每个 JSON 元素都将以新行开始，以前缀开头，然后根据缩进嵌套复制一份或多份缩进。

### func Unmarshal(data []byte, v any) error

如果 v 为 nil 或不是指针，Unmarshal 会返回 InvalidUnmarshalError。

Unmarshal 使用 Marshal 所用编码的逆编码，根据需要分配映射、切片和指针，并遵循以下附加规则：

要将 JSON 解 Marshal 成一个指针，Unmarshal 首先要处理 JSON 字面为空的情况。在这种情况下，Unmarshal 会将指针设置为 nil。否则，Unmarshal 会将 JSON 分解为指针指向的值。如果指针为 nil，Unmarshal 会分配一个新值供其指向。

要将 JSON 解链到实现 Unmarshaler 接口的值中，Unmarshal 会调用该值的 UnmarshalJSON 方法，包括在输入为 JSON null 时。否则，如果值实现了 encoding.TextUnmarshaler，且输入是带引号的 JSON 字符串，则 Unmarshal 会调用该值的 UnmarshalText 方法，并使用未加引号的字符串形式。

要将 JSON 解 Marshal 到结构体中，Unmarshal 会将输入的对象键与 Marshal 使用的键（结构体字段名或其标记）进行匹配，优先选择精确匹配，但也接受不区分大小写的匹配。默认情况下，没有对应结构字段的对象键会被忽略（请参阅 Decoder.DisallowUnknownFields（解码器不允许未知字段）以获取替代方法）。

要解码 JSON 到接口值中，Unmarshal 会在接口值中存储其中一个：

```go
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
```

要将一个 JSON 数组解链为一个片段，Unmarshal 会将片段长度重置为零，然后将每个元素追加到片段中。作为一种特例，要将一个空的 JSON 数组解链为一个片段，Unmarshal 会用一个新的空片段替换该片段。

要将 JSON 数组解链为 Go 数组，Unmarshal 会将 JSON 数组元素解码为相应的 Go 数组元素。如果 Go 数组小于 JSON 数组，则会丢弃额外的 JSON 数组元素。如果 JSON 数组小于 Go 数组，则会将额外的 Go 数组元素设置为零值。

要将 JSON 对象解映射到映射表中，Unmarshal 首先要建立一个映射表。如果映射为空，Unmarshal 会分配一个新映射。否则，Unmarshal 会重用现有映射，保留现有条目。然后，Unmarshal 会将 JSON 对象中的键值对存储到映射中。映射的键类型必须是任意字符串类型、整数、实现 json.Unmarshaler 或实现 encoding.TextUnmarshaler。

如果 JSON 编码数据包含语法错误，Unmarshal 会返回一个 SyntaxError。

如果某个 JSON 值不适合给定的目标类型，或者某个 JSON 数字溢出了目标类型，Unmarshal 会跳过该字段，尽可能完成解屏蔽。如果没有遇到更严重的错误，Unmarshal 会返回一个 UnmarshalTypeError，描述最早出现的此类错误。在任何情况下，都不能保证有问题字段之后的所有剩余字段都能解锁到目标对象中。

JSON null 值通过将 Go 值设置为 nil 来解码到接口、映射、指针或片段中。由于 JSON 中的 null 通常表示 "不存在"，因此将 JSON null 解映射到任何其他 Go 类型都不会影响该值，也不会产生错误。

在解除引号字符串的标记时，无效的 UTF-8 或无效的 UTF-16 代理对不会被视为错误。相反，它们会被 Unicode 替换字符 U+FFFD 代替。

### func Valid(data []byte) bool 添加于1.9

Valid 报告数据是否为有效的 JSON 编码。

### type Decoder

```go
type Decoder struct {
  // 包含已筛选或未导出字段
}
```

解码器从输入流中读取并解码 JSON 值。

#### func NewDecoder(r io.Reader) *Decoder

NewDecoder 返回一个从 r 读取数据的新解码器。

解码器会引入自己的缓冲，并可能从 r 中读取超出所请求的 JSON 值的数据。

#### func (dec *Decoder) Buffer() io.Reader 添加于1.1

Buffered 返回解码器缓冲区中剩余数据的读取器。该读取器在下一次调用解码器之前一直有效。

#### func (dec *Decoder) Decode(v any) error

解码从输入中读取下一个 JSON 编码值，并将其存储到 v 指向的值中。

有关将 JSON 转换为 Go 值的详情，请参阅 Unmarshal 文档。

#### func (dec *Decoder) DisallowUnknownFields() 添加于1.10

当目标是结构体，而输入包含的对象键与目标中任何未忽略的导出字段不匹配时，DisallowUnknownFields 会导致解码器返回错误。

#### func (dec *Decoder) InputOffset() int64 添加于1.14

InputOffset 返回当前解码器位置的输入流字节偏移量。偏移量给出了最近返回的标记的结束位置和下一个标记的开始位置。

#### func (dec *Decoder) More() bool 添加于1.5

更多地报告当前数组或正在解析的对象中是否有其他元素。

#### func (dec *Decoder) Token() (TOken, error) 添加于1.5

Token 返回输入流中的下一个 JSON 标记。在输入流结束时，Token 会返回 nil，即 io.EOF。

Token 保证它返回的分隔符 [ ] { } 是正确嵌套和匹配的：如果 Token 在输入中遇到意外的分隔符，它会返回错误信息。

输入流包括基本的 JSON 值--bool、string、number 和 null--以及 Delim 类型的分隔符 [ ] { }，用于标记数组和对象的开始和结束。逗号和冒号被省略。

#### func (dec *Decoder) UseNumber() 添加于1.1

UseNumber 会使解码器将一个数字作为 Number（而不是 float64）解码到接口{}中。

### type Delim 添加于1.5

```go
type Delim rune
```

Delim 是 JSON 数组或对象的定界符，可以是 [ ] { 或 } 之一。

#### func (d Delim) String() string 添加于1.5

### type Encoder

```go
type Encoder struct {
  // 包含已筛选或未导出字段
}
```

编码器将 JSON 值写入输出流。

#### func NewEncoder(w io.Writer) *Encoder

NewEncoder 返回写入 w 的新编码器。

#### func (enc *Encoder) Encode(v any) error

Encode 会将 v 的 JSON 编码写入数据流，之后是换行符。

有关将 Go 值转换为 JSON 的详细信息，请参阅 Marshal 文档。

#### func (enc *Encoder) SetEscapeHTML(on bool) 添加于1.7

SetEscapeHTML 用于指定是否要在 JSON 引号字符串中转义有问题的 HTML 字符。默认行为是将 &、< 和 > 转义为 \u0026、\u003c 和 \u003e，以避免在 HTML 中嵌入 JSON 时可能出现的某些安全问题。

在非 HTML 设置中，转义会影响输出的可读性，因此 SetEscapeHTML(false) 会禁用这一行为。

#### func (enc *Encoder) SetIndent(prefix, indent string) 添加于1.7

SetIndent 会指示编码器按照包级函数 Indent(dst, src, prefix, indent) 的缩进方式格式化每个后续编码值。调用 SetIndent("", "") 将禁用缩进。

### type InvalidUTF8Error 废除

#### func (e *InvalidUTF8Error) Error() string

### type InvalidUnmarshalError

```go
type InvalidUnmarshalError struct {
  Type reflect.Type
}
```

InvalidUnmarshalError 表示传递给 Unmarshal 的参数无效。(传给 Unmarshal 的参数必须是非零指针）。

#### func (e *InvalidUnmarshalError) Error() string

### type Marshaler

```go
type Marshaler interface {
  MarshalJSON() ([]byte, error)
}
```

Marshaler 是一种接口，由可以将自己 Marshal 成有效 JSON 的类型实现。

### type MarshalerError

```go
type Marshaler struct {
  Type reflect.Type
  Err error
  // 包含已筛选或未导出字段
}
```

MarshalerError 表示调用 MarshalJSON 或 MarshalText 方法时出错。

#### func (e *MarshalerError) Error() string

#### func (e *MarshalerError) Unwrap() error 添加于1.13

Unwrap 返回底层错误。

### type Number 添加于1.1

```go
type Number string
```

一个 Number 表示一个 JSON 数字字面。

#### func (n Number) Float64() (float64, error) 添加于1.1

Float64 返回 float64 格式的数字。

#### func (n Number) Int64() (int64, error) 添加于1.1

Int64 返回 int64 形式的数字。

#### func (n Number) String() string 添加于1.1

字符串返回数字的字面文本。

### type RawMessage

RawMessage 是原始编码 JSON 值。它实现了 Marshaler 和 Unmarshaler，可用于延迟 JSON 解码或预先计算 JSON 编码。

#### func (m RawMessage) MarshalJSON() ([]byte, error)

MarshalJSON 返回 m 的 JSON 编码。

#### func (m RawMessage) UnmarshalJSON(data []byte) error

UnmarshalJSON 将 *m 设置为数据副本。

### type SyntaxError

```go
type SyntaxError struct {
  Offset int64 // 读取偏移字节后发生错误
  // 包含已筛选或未导出字段
}
```

语法错误（SyntaxError）是对 JSON 语法错误的描述。如果 JSON 无法解析，Unmarshal 将返回语法错误。

#### func (e *SyntaxError) Error() string

### type Token 添加于1.5

```go
type Token any
```

令牌持有其中一种类型的值：

```go
Delim, for the four JSON delimiters [ ] { }
bool, for JSON booleans
float64, for JSON numbers
Number, for JSON numbers
string, for JSON string literals
nil, for JSON null
```

#### func UnmarshalFieldError 废除

### type UnmarshalTypeError

```go
type UnmarshalTypeError struct {
	Value string // JSON 值的描述 - "bool"、"array"、"number -5"
	Type reflect.Type // 无法赋值的 Go 值的类型
	Offset int64 // 读取 Offset 字节后发生错误
	Struct string // 包含字段的结构类型名称
	Field // 从根节点到字段的完整路径
}
```

UnmarshalTypeError 描述了不适合特定 Go 类型值的 JSON 值。

#### func (e *UnmarshalTypeError) Error() string

### type Unmarshaler

```go
type Unmarshaler interface {
  UnmarshalJSON([]byte) error
}
```

Unmarshaler 是一个接口，由可以解除对自身的 JSON 描述的类型实现。输入可以假定为 JSON 值的有效编码。如果 UnmarshalJSON 希望在返回后保留数据，则必须复制 JSON 数据。

按照惯例，为了近似 Unmarshal 本身的行为，Unmarshalers 将 UnmarshalJSON([]byte("null")) 作为无操作来实现。

### type UnsupportedTypeError

```go
type UnsupportedTypeError struct {
  Type reflect.Type
}
```

当 Marshal 尝试对不支持的值类型进行编码时，会返回 UnsupportedTypeError。

#### func (e *UnsupportedTypeError) Error() string

### type UnsupportedValueError

```go
type UnsupportedValueError struct {
  Value reflect.Value
  Str string
}
```

当 Marshal 尝试对不支持的值进行编码时，会返回 UnsupportedValueError。

#### func (e *UnsupportedValueError) Error() string