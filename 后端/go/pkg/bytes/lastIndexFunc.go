package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// LastIndexFunc 将 s 解释为 UTF-8 编码的码位序列。它返回满足 f(c) 的最后一个 Unicode 代码点在 s 中的字节索引，如果不满足则返回-1。
func main() {
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter)) // 8
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct)) // 9
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber)) // -1
}