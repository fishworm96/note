## package base32

包 base32 实现了 RFC 4648 规定的 base32 编码。

## Index

### Constants

```go
const (
	StdPadding rune = '=' // Standard padding character
	NoPadding  rune = -1  // No padding
)
```

### Variables

```go
var HexEncoding = NewEncoding(encodeHex)
```

HexEncoding 是 RFC 4648 中定义的 "扩展十六进制字母"。它通常用于 DNS。

```go
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding 是 RFC 4648 中定义的标准 base32 编码。

### func NewDecoder(enc *Encoding, r io.Reader) io.Reader

NewDecoder 构造一个新的 base32 流解码器。

### func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

NewEncoder 返回一个新的 base32 流编码器。写入返回写入器的数据将使用 enc 进行编码，然后写入 w。Base32 编码以 5 字节块为单位运行；写入完成后，调用者必须关闭返回的编码器，以清除任何部分写入的块。

### type CorruptInputError

```go
type CorruptInputError int64
```

#### func (e CorruptInputError) Error() string

### type Encoding

```go
type Encoding struct {
  // 包含已筛选或未导出字段
}
```

编码是一种弧度为 32 的编码/解码方案，由 32 个字符的字母表定义。最常用的是为 SASL GSSAPI 引入并在 RFC 4648 中标准化的 "base32 "编码。DNSSEC 中使用另一种 "base32hex "编码。

#### func NewEncoding(encoder string) *Encoding

NewEncoding 返回由给定字母定义的新编码，给定字母必须是 32 字节字符串。字母表被视为字节值序列，不会对多字节 UTF-8 进行任何特殊处理。

#### func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

Decode 使用编码 enc 对 src 进行解码。它最多向 dst 写入 DecodedLen(len(src)) 字节，并返回写入的字节数。如果 src 包含无效的 base32 数据，它会返回成功写入的字节数和 CorruptInputError。新行字符（\r 和 \n）将被忽略。

#### func (enc *Encoding) DecodeString(s string) ([]byte, error)

DecodeString 返回 base32 字符串 s 所代表的字节。

#### func (enc *Encoding) DecodedLen(n int) int

DecodedLen 返回与 n 字节基 32 编码数据相对应的解码数据的最大长度（以字节为单位）。

#### func (enc *Encoding) Encode(dst, src[]byte)

使用编码 enc 对 src 进行编码，将 EncodedLen(len(src)) 字节写入 dst。

编码会将输出填充为 8 字节的倍数，因此 Encode 不适合用于大型数据流中的单个数据块。请使用 NewEncoder() 代替。

#### func (enc *Encoding) EncodeToString(src []byte) string

EncodeToString 返回 src 的 base32 编码。

#### func (enc *Encoding) EncodedLen(n int) int

EncodedLen 返回长度为 n 的输入缓冲区的 base32 编码长度（以字节为单位）。

#### func (enc Encoding) WithPadding(padding rune) *Encoding 添加于1.9

WithPadding 会创建一个与 enc 相同的新编码，但会添加一个指定的填充字符；而 NoPadding 则会禁用填充字符。填充字符不能是"\r "或"\n"，不能包含在编码的字母表中，并且必须是等于或低于"\xff "的符文。高于"\x7f "的填充字符将以其确切的字节值进行编码，而不是使用编码点的 UTF-8 表示。