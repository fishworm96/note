package main

import (
	"bytes"
	"fmt"
)

// Replace 会返回一个片段 s 的副本，并用 new 替换掉前 n 个不重叠的 old 实例。如果 old 为空，则会在片段开头和每个 UTF-8 序列后进行匹配，一个 k 符片段最多会有 k+1 个替换实例。如果 n < 0，则替换次数不受限制。
func main() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2)) // oinky oinky oink
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1)) // moo moo moo
}