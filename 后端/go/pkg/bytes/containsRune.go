package main

import (
	"bytes"
	"fmt"
)

// ContainsRune 报告符文是否包含在 UTF-8 编码的字节片段 b 中。
func main()  {
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'f')) // true
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö')) // false
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大')) // true
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!')) // true
	fmt.Println(bytes.ContainsRune([]byte(""), '@')) // false
}