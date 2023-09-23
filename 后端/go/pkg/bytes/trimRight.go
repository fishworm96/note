package main

import (
	"bytes"
	"fmt"
)

// TrimRight 通过切掉 cutset 中包含的所有尾部 UTF-8 编码点，返回 s 的子片段。
func main() {
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789"))) // 453gopher
}