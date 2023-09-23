package main

import (
	"bytes"
	"fmt"
)

// TrimPrefix 返回不含前缀字符串的 s。如果 s 不以前缀开头，返回的 s 将保持不变。
func main() {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b) // Hello, world!
}