package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst)) // JBSWY3DPFQQHO33SNRSCC===
}