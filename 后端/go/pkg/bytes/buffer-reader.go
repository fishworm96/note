package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

// 缓冲区是一个大小可变的字节缓冲区，具有读取和写入方法。Buffer 的零值表示缓冲区为空，可以随时使用。
func main() {
	// 缓冲区可将字符串或[]字节转换为 io.Reader.NX 格式。
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec) // Gophers rule!
}