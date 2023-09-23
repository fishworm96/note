package main

import (
	"bytes"
	"fmt"
)

// CutPrefix 返回不含前缀字节片段的 s，并报告是否找到前缀。如果 s 不以前缀开头，CutPrefix 会返回 s，false。如果前缀是空字节片，CutPrefix 返回 s，true。
// CutPrefix 返回原始片段 s 的片段，而不是副本。
func main() {
	show := func(s, sep string) {
		after, found := bytes.CutPrefix([]byte(s), []byte(sep))
		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, sep, after, found)
	}
	show("Gopher", "Go") // CutPrefix("Gopher", "Go") = "pher", true
	show("Gopher", "ph") // CutPrefix("Gopher", "ph") = "Gopher", false
}