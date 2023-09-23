package main

import (
	"bytes"
	"fmt"
)

// ToValidUTF8 将 s 作为 UTF-8 编码字节处理，并返回一个副本，其中代表无效 UTF-8 的每一行字节都被替换为替换中的字节，替换中的字节可能为空。
func main() {
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("abc"), []byte("\uFFFD"))) // abc
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("a\xffb\xC0\xAFc\xff"), []byte(""))) // abc
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("\xed\xa0\x80"), []byte("abc"))) // abc
}