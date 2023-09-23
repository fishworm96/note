package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// IndexFunc 将 s 解释为 UTF-8 编码的码位序列。它返回满足 f(c) 的第一个 Unicode 代码点在 s 中的字节索引，如果不满足则返回-1。
func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f)) // 7
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f)) // -1
}