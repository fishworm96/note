package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// TrimRightFunc 通过切掉所有满足 f(c) 的尾部 UTF-8 编码码点 c，返回 s 的子片段。
func main() {
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter))) // go-
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct))) // go-gopher
	fmt.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber))) // 1234go-gopher!
}