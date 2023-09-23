package main

import (
	"bytes"
	"fmt"
)

// LastIndex 返回 s 中 sep 最后一个实例的索引，如果 s 中没有 sep，则返回-1。
func main() {
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go"))) // 0
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go"))) // 3
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent"))) // -1
}