package main

import (
	"bytes"
	"fmt"
)

// Equal 报告 a 和 b 的长度和字节数是否相同。nil 参数等同于空片段。
func main() {
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go"))) // true
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++"))) // false
}