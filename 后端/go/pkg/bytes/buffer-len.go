package main

import (
	"bytes"
	"fmt"
)

// Len 返回缓冲区未读部分的字节数；b.Len() == len(b.Bytes())。
func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%d", b.Len()) // 5
}