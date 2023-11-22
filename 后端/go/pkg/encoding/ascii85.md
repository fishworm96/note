## package ascii85

软件包 ascii85 实现了 btoa 工具和 Adobe PostScript 及 PDF 文档格式中使用的 ascii85 数据编码。

## Index

### func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)

解码将 src 解码为 dst，同时返回写入 dst 的字节数和从 src 消耗的字节数。如果 src 包含无效的 ascii85 数据，解码将返回成功写入的字节数和一个 CorruptInputError。解码会忽略 src 中的空格和控制字符。通常，ascii85 编码的数据会被 <~ 和 ~> 符号包裹。解码器希望这些字符已被调用者去除。

如果 flush 为真，解码器会认为 src 代表输入流的结束，并完全处理它，而不是等待另一个 32 位数据块的完成。

NewDecoder 围绕 Decode 封装了一个 io.Reader 接口。

### func Encode(dst, src []byte) int

Encode 最多将 src 编码为 dst 的 MaxEncodedLen(len(src)) 字节，并返回实际写入的字节数。

编码处理 4 字节块，对最后一个片段使用特殊编码，因此 Encode 不适合用于大型数据流中的单个数据块。请使用 NewEncoder() 代替。

通常情况下，ascii85 编码数据会用 <~ 和 ~> 符号包裹。Encode 不会添加这些符号。

### func MaxEncodedLen(n int) int

MaxEncodedLen 返回 n 个源字节编码的最大长度。

### func NewDecoder(r io.Reader) io.WriteCloser

NewDecoder 构造一个新的 ascii85 数据流解码器。

### func NewEncoder(w io.Reader) io.WriteCloser

NewEncoder 返回一个新的 ascii85 流编码器。写入返回写入器的数据将被编码，然后写入 w。Ascii85 编码以 32 位数据块为单位；写入完成后，调用者必须关闭返回的编码器，以清除尾部的部分数据块。

### type CorruptInputError

```go
type CorruptInputError int64
```

#### func (e CorruptInputError) Error() string