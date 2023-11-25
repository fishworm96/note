package main

import (
	"encoding/base32"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input) // MZXW6ADCMFZA====
	// 编码完成后必须关闭编码器，以清除任何部分数据块。如果注释掉下面一行，最后一个部分块 "r "将不会被编码。
	encoder.Close()
}