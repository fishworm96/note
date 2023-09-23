package main

import (
	"bytes"
	"fmt"
)

// ToTitle 将 s 作为 UTF-8 编码字节处理，并返回一个将所有 Unicode 字母映射到标题大小写的副本。
func main() {
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
}