package main

import (
	"bytes"
	"fmt"
)

// Join 将 s 的元素连接起来，创建一个新的字节片段。分隔符 sep 放在结果片段的元素之间。
func main() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", "))) // foo, bar, baz
}