package main

import (
	"bytes"
	"fmt"
)

// HasPrefix 测试字节片 s 是否以前缀开头。
func main() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go"))) // true
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C"))) // false
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte(""))) // true
}