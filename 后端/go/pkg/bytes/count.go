package main

import (
	"bytes"
	"fmt"
)

func main()  {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e"))) // 3
	// 每个 rune 之前或之后
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // 5
}