package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "SGVsbG8sIHdvcmxkIQ=="
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
	}
	dst =dst[:n]
	fmt.Printf("%q\n", dst) // "Hello, world!"
}