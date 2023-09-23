package main

import (
	"bytes"
	"fmt"
)

// Trim 通过切掉 cutset 中包含的所有前端和后端 UTF-8 编码点，返回 s 的子片段。
func main() {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! ")) // ["Achtung! Achtung"]
}