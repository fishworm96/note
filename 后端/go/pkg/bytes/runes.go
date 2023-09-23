package main

import (
	"bytes"
	"fmt"
)

// Runes 将 s 解释为 UTF-8 编码的码位序列。它会返回与 s 相同的符文（Unicode 代码点）片段。
func main() {
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\b", r) // U+0067 'gU+006F 'oU+0020 ' U+0067 'gU+006F 'oU+0070 'pU+0068 'hU+0065 'eU+0072 'r'
	}
}