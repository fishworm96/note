package main

import (
	"bytes"
	"fmt"
	"os"
)

// 缓冲区是一个大小可变的字节缓冲区，具有读取和写入方法。Buffer 的零值表示缓冲区为空，可以随时使用。
func main() {
	var b bytes.Buffer // 缓冲区不需要初始化。
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!") // Hello world!
	b.WriteTo(os.Stdout)
}