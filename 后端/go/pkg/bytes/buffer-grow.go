package main

import (
	"bytes"
	"fmt"
)

// 必要时，Grow 会增加缓冲区的容量，以保证能再容纳 n 个字节。在 Grow(n) 之后，至少有 n 个字节可以写入缓冲区而无需再次分配。如果 n 为负数，Grow 就会崩溃。如果缓冲区无法增长，则会出现 ErrTooLarge 异常。
func main() {
	var b bytes.Buffer
	b.Grow(64)
	bb := b.Bytes()
	b.Write([]byte("64 bytes or fewer"))
	fmt.Printf("%q", bb[:b.Len()]) // "64 bytes or fewer"
}