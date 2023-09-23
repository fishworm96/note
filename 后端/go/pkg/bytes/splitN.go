package main

import (
	"bytes"
	"fmt"
)

// SplitN 将 s 分割成以 sep 分隔的子片段，并返回这些分隔符之间的子片段。如果 sep 为空，SplitN 会在每个 UTF-8 序列后分割。
func main() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2)) // ["a" "b,c"]
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
}