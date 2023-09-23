package main

import (
	"bytes"
	"fmt"
)

// SplitAfterN 在每个 sep 实例之后将 s 分割成子片段，并返回这些子片段的片段。如果 sep 为空，SplitAfterN 会在每个 UTF-8 序列后进行分割。count 决定返回子片段的数量：
func main() {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte("m"), 2)) // ["a,b,c"]
}