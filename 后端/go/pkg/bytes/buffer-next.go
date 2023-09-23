package main

import (
	"bytes"
	"fmt"
)

// Next 返回一个片段，其中包含缓冲区中下一个 n 字节，并将缓冲区向前推进，就像读取已返回的字节一样。如果缓冲区中的字节少于 n 个，Next 会返回整个缓冲区。该分片仅在下一次调用读或写方法之前有效。
func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%s\n", string(b.Next(2))) // ab
	fmt.Printf("%s\n", string(b.Next(2))) // cd
	fmt.Printf("%s", string(b.Next(2))) // e
}