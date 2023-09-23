package main

import (
	"bytes"
	"fmt"
)

// Cap 返回缓冲区底层字节片的容量，即分配给缓冲区数据的总空间。
func main() {
	buf1 := bytes.NewBuffer(make([]byte, 10))
	buf2 := bytes.NewBuffer(make([]byte, 0, 10))
	fmt.Println(buf1.Cap()) // 10
	fmt.Println(buf2.Cap()) // 10
}