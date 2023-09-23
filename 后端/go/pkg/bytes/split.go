package main

import (
	"bytes"
	"fmt"
)

// Split 会将 s 分割成以 sep 分隔的所有子片段，并返回这些分隔符之间的子片段。如果 sep 为空，Split 会在每个 UTF-8 序列之后进行分割。它等同于计数为-1的 SplitN。

// 要在分隔符的第一个实例周围进行分割，请参阅 Cut。
func main() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(","))) // ["a" "b" "c"]
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a "))) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte(""))) // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins"))) // [""]
}