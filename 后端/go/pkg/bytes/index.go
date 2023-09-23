package main

import (
	"bytes"
	"fmt"
)

// Index 返回 s 中第一个 sep 实例的索引，如果 s 中没有 sep，则返回-1。
func main() {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken"))) // 4
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr"))) // -1
}