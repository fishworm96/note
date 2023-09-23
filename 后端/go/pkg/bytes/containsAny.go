package main

import (
	"bytes"
	"fmt"
)

// ContainsAny 报告字符串中是否有任何 UTF-8 编码的码位在 b 范围内。
func main() {
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!")) // true
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的.")) // true
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "")) // false
	fmt.Println(bytes.ContainsAny([]byte(""), "")) // false
}