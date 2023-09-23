package main

import (
	"bytes"
	"os"
)

// TrimSuffix 返回不含所提供尾缀字符串的 s。如果 s 不以后缀结尾，则返回的 s 将保持不变。
func main() {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(b) // Hello, world!
}