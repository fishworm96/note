package main

import (
	"bytes"
	"fmt"
)

// 重复会返回一个由 b 的计数副本组成的新字节片。

// 如果 count 为负数或 (len(b) * count) 的结果溢出，程序就会崩溃。
func main() {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2)) // banana
}