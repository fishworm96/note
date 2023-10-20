## package des

import "crypto/des"

数据包 des 实现了美国联邦信息处理标准出版物 46-3 中定义的数据加密标准 (DES) 和三重数据加密算法 (TDEA)。

DES 在密码学上已被破解，不应在安全应用程序中使用。

## Index

### Constants

```go
const BlockSize = 8
```

以字节为单位的 DES 数据块大小。

### func NewCipher(key []byte) (cipher.Block, error)

NewCipher 创建并返回一个新的密码块。

### func NewTripleDESCipher(key []byte) (cipher.Block, error)

NewTripleDESCipher 创建并返回一个新的密码块。

### type KeySizeError

```go
type KeySizeError int
```

#### func (k KeySizeError) Error() string