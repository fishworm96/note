package main

import (
	"bytes"
	"fmt"
)

// TrimSpace 按照 Unicode 的定义，通过切掉所有前导和尾部空白，返回 s 的子片段。
func main() {
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n"))) // a lone gopher
}