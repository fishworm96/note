package main

import (
	"bytes"
	"fmt"
)

// ToLower 返回将所有 Unicode 字母映射为小写字母的字节片 s 的副本。
func main() {
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher"))) // gopher
}