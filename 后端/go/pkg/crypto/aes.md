## package aes

import "crypto/aes"

软件包 aes 实现了美国联邦信息处理标准 197 出版物中定义的 AES 加密（前身为 Rijndael）。

本软件包中的 AES 操作不是通过恒定时间算法实现的。但在启用了 AES 硬件支持的系统上运行时例外，因为硬件支持使这些操作成为恒定时间算法。例如，使用 AES-NI 扩展的 amd64 系统和使用消息安全辅助扩展的 s390x 系统。在此类系统上，当将 NewCipher 的结果传递给 cipher.NewGCM 时，GCM 使用的 GHASH 操作也是恒定时间的。

## Index

### Constants

```go
const BlockSize = 16
```

AES 数据块大小（字节）。

### func NewCipher(key []byte) (cipher.Block, error)

NewCipher 创建并返回一个新的 cipher.Block 密钥。密钥参数应为 AES 密钥，16、24 或 32 字节（可选择 AES-128、AES-192 或 AES-256）。

### type KeySizeError

```go
type KeySizeError int
```

#### func (k KeySizeError) Error() string