package main

import (
	"bytes"
	"fmt"
)

// ToUpper 返回将所有 Unicode 字母映射为大写字母的字节片 s 的副本。
func main() {
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher"))) // GOPHER
}