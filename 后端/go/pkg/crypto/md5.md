## package md5

import "crypto/md5"

软件包 md5 实现了 RFC 1321 中定义的 MD5 哈希算法。

MD5 在密码学上已被破解，不应在安全应用程序中使用。

## Index

### Constants

```go
const BlockSize = 64
```

MD5 的块大小（字节）。

```go
const Size = 16
```

MD5 校验和的大小（字节）。

### func New() hash.Hash

New 返回一个新的哈希值，计算 MD5 校验和散列。还实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler，以对散列的内部状态进行存档和解档。

### func Sum(data []byte) [Size]byte 添加于 1.2

Sum 返回 MD5 校验数据。