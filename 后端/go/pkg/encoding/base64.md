## package base64

包 base64 实现了 RFC 4648 规定的 base64 编码。

## Index

### Constants

```go
const (
	StdPadding rune = '=' // 标准填充字符
	NoPadding rune = -1 // 无填充字符
)
```

### Variables

```go
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
```

RawStdEncoding 是 RFC 4648 第 3.2 节定义的标准原始、无填充 base64 编码。它与 StdEncoding 相同，但省略了填充字符。


```go
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
```

RawURLEncoding 是 RFC 4648 中定义的无填充备用 base64 编码。它通常用于 URL 和文件名。它与 URLEncoding 相同，但省略了填充字符。

```go
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding 是 RFC 4648 中定义的标准 base64 编码。

```go
var URLEncoding = NewEncoding(encodeURL)
```

URLEncoding 是 RFC 4648 中定义的备用 base64 编码。它通常用于 URL 和文件名。

### func NewDecoder(enc *Encoding, r io.Reader) io.Reader

NewDecoder 构造一个新的 base64 流解码器。

### func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

NewEncoder 返回一个新的 base64 流编码器。写入返回写入器的数据将使用 enc 编码，然后写入 w。Base64 编码以 4 字节块为单位运行；写入完成后，调用者必须关闭返回的编码器，以清除写入的部分块。

### type CorruptInputError

```go
type CorruptInputError int64
```

#### func (e CorruptInputError) Error() string

### type Encoding

```go
type Encoding struct {
  // 包含已筛选或未导出的字段
}
```

编码是一种弧度为 64 的编码/解码方案，由 64 个字符的字母表定义。最常用的编码是 RFC 4648 中定义的 "base64 "编码，在 MIME (RFC 2045) 和 PEM (RFC 1421) 中使用。RFC 4648 还定义了另一种编码，即用 - 和 _ 代替 + 和 / 的标准编码。

#### func NewEncoding(encoder string) *Encoding

NewEncoding 返回由给定字母定义的新填充编码，该字母必须是不包含填充字符或 CR / LF（'\r'、'\n'）的 64 字节字符串。字母表会被视为字节值序列，不会对多字节 UTF-8 进行任何特殊处理。生成的编码使用默认的填充字符（'='），可以通过 WithPadding 更改或禁用。

#### func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

Decode 使用编码 enc 对 src 进行解码。它最多向 dst 写入 DecodedLen(len(src)) 字节，并返回写入的字节数。如果 src 包含无效的 base64 数据，它会返回成功写入的字节数和 CorruptInputError。新行字符（\r 和 \n）将被忽略。

#### func (enc *Encoding) DecodeString(s string) ([]byte, error)

DecodeString 返回 base64 字符串 s 所代表的字节。

#### func (enc *Encoding) DecodedLen(n int) int

DecodedLen 返回与 n 个字节的 base64 编码数据相对应的解码数据的最大长度（以字节为单位）。

#### func (enc *Encoding) Encode(dst, src[]byte)

使用编码 enc 对 src 进行编码，将 EncodedLen(len(src)) 字节写入 dst。

编码会将输出填充为 4 字节的倍数，因此 Encode 不适合用于大型数据流中的单个数据块。请使用 NewEncoder() 代替。

#### func (enc *Encoding) EncodeToString(src []byte) string

EncodeToString 返回 src 的 base64 编码。

#### func (enc *Encoding) EncodedLen()

EncodedLen 返回长度为 n 的输入缓冲区的 base64 编码长度（以字节为单位）。

#### func (enc Encoding) Strict() *Encoding 添加于1.8

严格模式创建的新编码与 enc 相同，只是启用了严格解码。在这种模式下，解码器要求尾部填充位为零，如 RFC 4648 第 3.5 节所述。

请注意，由于新行字符（CR 和 LF）仍会被忽略，因此输入仍具有可塑性。

#### func (enc Encoding) WithPadding(padding rune) *Encoding 添加于1.5

WithPadding 会创建一个与 enc 相同的新编码，但会添加一个指定的填充字符；而 NoPadding 则会禁用填充字符。填充字符不能是"\r "或"\n"，不能包含在编码的字母表中，并且必须是等于或低于"\xff "的符文。高于"\x7f "的填充字符将以其确切的字节值进行编码，而不是使用编码点的 UTF-8 表示。