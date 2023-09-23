package main

import (
	"bytes"
	"fmt"
)

// HasSuffix 测试字节片 s 是否以后缀结尾。
func main() {
	fmt.Println(bytes.HasPrefix([]byte("Amigo"), []byte("go"))) // true
	fmt.Println(bytes.HasPrefix([]byte("Amigo"), []byte("O"))) // false
	fmt.Println(bytes.HasPrefix([]byte("Amigo"), []byte("Ami"))) // false
	fmt.Println(bytes.HasPrefix([]byte("Amigo"), []byte(""))) // true
}