package main

import (
	"bytes"
	"fmt"
)

// 如果 sep 不在 s 中出现，cut 返回 s、nil、false。
// Cut 返回原始片段 s 的片段，而不是副本。
func main() {
	show := func(s, sep string) {
		before, after, found := bytes.Cut([]byte(s), []byte(sep))
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go") // Cut("Gopher", "Go") = "", "pher", true
	show("Gopher", "ph") // Cut("Gopher", "ph") = "Go", "er", true
	show("Gopher", "er") // Cut("Gopher", "er") = "Goph", "", true
	show("Gopher", "Badger") // Cut("Gopher", "Badger") = "Gopher", "", false
}