package main

import (
	"bytes"
	"fmt"
)

// LastIndexByte 返回 s 中最后一个 c 实例的索引，如果 s 中没有 c，则返回-1。
func main() {
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "MüQp")) // 5
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠"), "地大")) // 3
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!.")) // -1
}