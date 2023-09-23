package main

import (
	"bytes"
	"fmt"
)

// CutSuffix 返回不含提供的后缀结尾字节片段的 s，并报告是否找到了后缀。如果 s 不以 suffix 结尾，CutSuffix 会返回 s，false。如果后缀是空字节片段，CutSuffix 返回 s，true。
// CutSuffix 返回原始片段 s 的片段，而不是副本。
func main() {
	show := func(s, sep string) {
		before, found := bytes.CutSuffix([]byte(s), []byte(sep))
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
	}
	show("Gopher", "Go") // CutSuffix("Gopher", "Go") = "Gopher", false
	show("Gopher", "er") // CutSuffix("Gopher", "er") = "Goph", true
}
