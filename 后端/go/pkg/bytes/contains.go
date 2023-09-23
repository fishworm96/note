package main

import (
	"bytes"
	"fmt"
)

// 包含子片是否在 b 范围内的报告。
func main() {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo"))) // true
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar"))) // false
	fmt.Println(bytes.Contains([]byte("seafood"), []byte(""))) // true
	fmt.Println(bytes.Contains([]byte(""), []byte(""))) // true
}