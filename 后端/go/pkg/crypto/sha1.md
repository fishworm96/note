## package sha1

import "crypto/sha1"

包 sha1 实现了 RFC 3174 中定义的 SHA-1 哈希算法。

SHA-1 算法在密码学上已被破解，不应被用于安全应用。

## Index

### Constants

```go
const BlockSize = 64
```

SHA-1 的块大小（字节）。

```go
const Size = 20
```

SHA-1 校验和的大小（字节）。

### func New() hash.Hash

New 返回一个新的哈希值，计算 SHA1 校验和。哈希散列还实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler，以对哈希散列的内部状态进行存档和解档。

### func Sum(data []byte) [Size]byte 添加于1.2

Sum 返回数据的 SHA-1 校验和。