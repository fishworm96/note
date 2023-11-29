## package hex

包 hex 实现十六进制编码和解码。

## Index

### Variables

```go
var ErrLength = errors.New("encoding/hex: odd length hex string")
```

ErrLength 会报告使用 Decode 或 DecodeString 对奇数长度输入进行解码的尝试。基于流的解码器返回 io.ErrUnexpectedEOF 而不是 ErrLength。

### func Decode(dst, src []byte) (int, error)

Decode 将 src 解码为 DecodedLen(len(src)) 字节，返回写入 dst 的实际字节数。

Decode 希望 src 只包含十六进制字符，并且 src 长度为偶数。如果输入是畸形的，解码会返回错误前已解码的字节数。

### func DecodeString(s string) ([]byte, error)

DecodeString 返回十六进制字符串 s 所代表的字节。

DecodeString 希望 src 只包含十六进制字符，并且 src 长度为偶数。如果输入是畸形的，DecodeString 会返回错误前已解码的字节。

### func DecodedLen(x int) int

DecodedLen 返回 x 个源字节的解码长度。具体来说，它返回 x / 2。

### func Dump(data []byte) string

Dump 返回一个字符串，其中包含给定数据的十六进制转储。十六进制转储的格式与命令行中 `hexdump -C` 的输出一致。

### func Dumper(w io.Writer) io.WriteCloser

Dumper 返回一个 WriteCloser，将所有写入数据的十六进制转储写入 w。

### func Encode(dst, src []byte) int

Encode 将 src 编码为 dst 的 EncodedLen(len(src)) 字节。为方便起见，它返回写入 dst 的字节数，但该值始终是 EncodedLen(len(src)) 。Encode 实现十六进制编码。

### func EncodeToString(src []byte) string

EncodeToString 返回 src 的十六进制编码。

### func EncodedLen(n int) int

EncodedLen 返回 n 个源字节的编码长度。具体来说，它返回 n * 2。

### func NewDecoder(r io.Reader) io.Reader 添加于1.10

NewDecoder 希望 r 只包含偶数个十六进制字符。

### func NewEncoder(w io.Writer) io.Writer 添加于1.10

NewEncoder 返回一个 io.Writer，用于向 w 中写入小写十六进制字符。

### type InvalidByteError

InvalidByteError 值描述十六进制字符串中无效字节导致的错误。

#### func (e InvalidByteError) Error() string