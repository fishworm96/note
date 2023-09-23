package main

import (
	"bytes"
	"fmt"
)

// SplitAfter 会在每个 sep 实例之后将 s 分割成所有子片段，并返回这些子片段的片段。如果 sep 为空，SplitAfter 会在每个 UTF-8 序列后进行分割。它等同于计数为-1的 SplitAfterN。
func main() {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(","))) // ["a," "b," "c"]
}