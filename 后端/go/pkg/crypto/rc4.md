## package rc4

import "crypto/cr4"

软件包 rc4 实现了布鲁斯-施奈尔（Bruce Schneier）的《应用密码学》（Applied Cryptography）一书中定义的 RC4 加密。

RC4 在密码学上已被破解，不应被用于安全应用。

## Index

### type Cipher

```go
type Cipher struct {
  // 包含过滤或未导出的字段
}
```

密码是使用特定密钥的 RC4 实例。

#### func NewCipher(key []byte) (*Cipher, error)

NewCipher 创建并返回一个新密码。密钥参数应为 RC4 密钥，至少 1 字节，最多 256 字节。

#### func (c *Cipher) Reset() 废除

#### func (c *Cipher) XORKeyStream(dst, src []byte)

XORKeyStream 将 dst 设置为 src 与密钥流 XOR 后的结果。Dst 和 src 必须完全重叠或完全不重叠。

### type KeySizeError

```go
type KeySizeError int
```

#### func (k KeySizeError) Error() string