package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// 切片现在应包含随机字节，而不是只有 0。
	fmt.Println(bytes.Equal(b, make([]byte, c))) // false
}