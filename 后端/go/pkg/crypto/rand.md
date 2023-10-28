## package rand

import "crypto/rand"

软件包 rand 实现了一个加密安全的随机数生成器。

## Index

### Variables

```go
var Reader io.Reader
```

读取器是一个加密安全随机数生成器的全局共享实例。

在 Linux、FreeBSD、Dragonfly、NetBSD 和 Solaris 上，如果可用，Reader 会使用 getrandom(2)，否则会使用 /dev/urandom。在 OpenBSD 和 macOS 上，Reader 使用 getentropy(2)。在其他类 Unix 系统上，Reader 从 /dev/urandom 读取数据。在 Windows 系统上，Reader 使用 RtlGenRandom API。在 JS/Wasm 系统上，Reader 使用 Web Crypto API。在 WASIP1/Wasm 系统上，Reader 使用 wasi_snapshot_preview1 中的 random_get。

### func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)

Int 返回 [0, max] 范围内的一个均匀随机值。如果 max <= 0，它就会崩溃。

### func Prime(rand io.Reader, bits int) (*big.Int, error)

Prime 返回给定位长的数字，该数字很可能是质数。如果 rand.Read 返回错误或比特数小于 2，Prime 将返回错误信息。

### func Read(b []byte) (n int, err error)

Read 是一个辅助函数，使用 io.ReadFull 调用 Reader.Read。返回时，如果且仅当 err == nil 时，n == len(b)。