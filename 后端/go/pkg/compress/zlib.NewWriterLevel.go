package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func main() {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes()) // [120 156 202 72 205 201 201 215 81 40 207 47 202 73 225 2 4 0 0 255 255 33 231 4 147]
}