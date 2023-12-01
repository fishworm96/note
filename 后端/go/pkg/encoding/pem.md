## package pem

软件包 pem 实现了源于隐私增强邮件（Privacy Enhanced Mail）的 PEM 数据编码。目前，PEM 编码最常用于 TLS 密钥和证书。请参见 RFC 1421。

## Index

### func Encode(out io.Writer, b *Block) error

Encode 将 b 的 PEM 编码写入 out。

### func EncodeToMemory(b *Block) []byte

EncodeToMemory 返回 b 的 PEM 编码。

如果 b 的头无效且无法编码，EncodeToMemory 将返回 nil。如果需要报告此错误的详细信息，请使用 Encode。

### type Block

```go
type Block struct {
	Type string // 类型，取自前言（即 "RSA PRIVATE KEY"）。
	Headers map[string]string // 可选标题。
	Bytes[]byte // 内容的解码字节。通常是 DER 编码的 ASN.1 结构。
}
```

区块表示 PEM 编码结构。

编码形式为:

```go
-----BEGIN Type-----
Headers
base64-encoded Bytes
-----END Type-----
```

其中 Headers 是一个可能为空的 Key：值的行序列。

#### func Decode(data []byte) (p *Block, rest []byte)

解码会在输入中找到下一个 PEM 格式块（证书、私钥等）。它将返回该数据块和其余输入内容。如果没有找到 PEM 数据，则 p 为空，并返回整个输入的其余部分。