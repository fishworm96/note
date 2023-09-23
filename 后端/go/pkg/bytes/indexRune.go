package main

import (
	"bytes"
	"fmt"
)

// IndexRune 将 s 解释为 UTF-8 编码的码位序列。如果 r 为 utf8.RuneError，则返回给定符文在 s 中首次出现的字节索引。如果 r 为 utf8.RuneError，则返回任何无效 UTF-8 字节序列的第一个实例。
func main() {
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k')) // 4
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd')) // -1
}