package main

import (
	"bytes"
	"fmt"
)

// 映射会返回一个字节片 s 的副本，并根据映射函数修改其所有字符。如果映射返回的是负值，该字符将从字节片中删除，不会被替换。s 中的字符和输出将被解释为 UTF-8 编码的码位。
func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' +(r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s\n", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher..."))) // 'Gjnf oevyyvt naq gur fyvgul tbcure...
}