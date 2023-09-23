package main

import (
	"bytes"
	"fmt"
)

// IndexByte 返回 b 中第一个 c 实例的索引，如果 b 中没有 c，则返回-1。
func main() {
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k'))) // 4
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g'))) // -1
}