package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello Gopher!")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst) // 48656c6c6f20476f7068657221
}