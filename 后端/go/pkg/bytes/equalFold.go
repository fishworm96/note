package main

import (
	"bytes"
	"fmt"
)

// EqualFold 报告解释为 UTF-8 字符串的 s 和 t 在简单的 Unicode 大小写折叠（一种更普遍的大小写不敏感形式）下是否相等。
func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}