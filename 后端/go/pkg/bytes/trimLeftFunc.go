package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// TrimLeftFunc 将 s 视为 UTF-8 编码字节，并通过切掉满足 f(c) 的所有前导 UTF-8 编码码位 c 来返回 s 的子片段。
func main() {
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter))) // -gopher
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct))) // go-gopher!
	fmt.Println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber))) // go-gopher!567
}