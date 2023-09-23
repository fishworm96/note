package main

import (
	"bytes"
	"fmt"
)

// ReplaceAll 会返回一个片段 s 的副本，其中所有不重叠的 old 实例都会被 new 替换。如果 old 为空，则会在片段开头和每个 UTF-8 序列后进行匹配，一个 k 符片段最多会有 k+1 个替换。
func main() {
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo"))) // moo moo moo
}