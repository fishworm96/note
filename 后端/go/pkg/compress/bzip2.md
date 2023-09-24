## package bzip2

import "compress/bzip2"

软件 bzip2 实现了 bzip2 解压缩。

## Index

### func NewReader(r io.Reader) io.Reader

如果 r 没有同时实现 io.ByteReader，解压程序从 r 读取的数据可能会超过需要。

### type StructuralError

```go
type StructuralError string
```

当发现 bzip2 数据在语法上无效时，将返回 StructuralError。

  #### func (s StructuralError) Error() string