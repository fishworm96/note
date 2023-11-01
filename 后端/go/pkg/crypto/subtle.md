## package subtle

软件包 subtle 实现了一些在加密代码中非常有用的功能，但要正确使用这些功能还需要仔细斟酌。

## Index

### func ConstantTimeByteEq(x, y uint8) int

如果 x == y，ConstantTimeByteEq 返回 1，否则返回 0。

### func ConstantTimeCompare(x, y []byte) int

如果 x 和 y 这两个片段的内容相等，ConstantTimeCompare 将返回 1，否则返回 0。所用时间是切片长度的函数，与内容无关。如果 x 和 y 的长度不匹配，则立即返回 0。

### func ConstantTimeCopy(v int, x, y []byte)

如果 v == 1，ConstantTimeCopy 会将 y 的内容复制到 x（等长的片段）中。如果 v 取其他值，则其行为未定义。

### func ConstantTImeEq(x, y int32) int

如果 x == y，ConstantTimeEq 返回 1，否则返回 0。

### func ConstantTimeLessOrEq(x, y int) int 添加于1.2

如果 x <= y，ConstantTimeLessOrEq 返回 1，否则返回 0。如果 x 或 y 为负数或 > 2**31 - 1，则其行为未定义。

### func ConstantTimeSelect(v, x, y int) int

如果 v == 1，ConstantTimeSelect 将返回 x；如果 v == 0，ConstantTimeSelect 将返回 y。

### func XORBytes(dst, x, y []byte) int 添加于1.20

对于所有 i < n = min(len(x)，len(y))，XORBytes 设置 dst[i] = x[i] ^ y[i]，返回 n，即写入 dst 的字节数。如果 dst 的长度不至少为 n，XORBytes 将停止向 dst 写入任何内容。