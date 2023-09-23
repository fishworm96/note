package main

import (
	"bytes"
	"fmt"
)

// 读取从缓冲区读取下一个 len(p) 字节或直到缓冲区耗尽。返回值 n 是读取的字节数。如果缓冲区没有要返回的数据，err 就是 io.EOF（除非 len(p) 为零）；否则就是 nil。
func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	rdbuf := make([]byte, 1)
	n, err := b.Read(rdbuf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n) // 1
	fmt.Println(b.String()) // bcde
	fmt.Println(string(rdbuf)) // a
}