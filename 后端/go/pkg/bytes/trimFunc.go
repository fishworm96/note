package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// TrimFunc 通过切掉满足 f(c) 的所有前导和尾部 UTF-8 编码码点 c，返回 s 的子片段。
func main() {
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter))) // -gopher!
	fmt.Println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter))) // "go-gopher!"
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct))) // go-gopher
	fmt.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber))) // go-gopher!
}