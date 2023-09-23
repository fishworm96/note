package main

import (
	"bytes"
	"fmt"
)

// Len 返回片段未读部分的字节数。
func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len()) // 3
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len()) // 16
}