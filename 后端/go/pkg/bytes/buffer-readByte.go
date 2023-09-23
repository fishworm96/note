package main

import (
	"bytes"
	"fmt"
)

// ReadByte 从缓冲区读取并返回下一个字节。如果没有可用字节，则返回错误 io.EOF。
func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	c, err := b.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Println(c) // 97
	fmt.Println(b.String()) // bcde
}