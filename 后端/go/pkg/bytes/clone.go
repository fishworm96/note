package main

import (
	"bytes"
	"fmt"
)

// // 克隆返回 b[:len(b)] 的副本。结果可能有额外的未使用容量。Clone(nil) 返回 nil。
func main() {
	b := []byte("abc")
	clone := bytes.Clone(b)
	fmt.Printf("%s\n", clone) // abc
	clone[0] = 'd'
	fmt.Printf("%s\n", b) // abc
	fmt.Printf("%s\n", clone) // dbc
}