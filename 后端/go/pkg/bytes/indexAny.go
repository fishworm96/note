package main

import (
	"bytes"
	"fmt"
)

// IndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。它返回 chars 中任何一个 Unicode 代码点在 s 中首次出现的字节索引。如果 chars 为空或没有共同的码位，则返回-1。
func main() {
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy")) // 2
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy")) // -1
}