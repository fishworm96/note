package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// FieldsFunc 将 s 解释为 UTF-8 编码的码位序列。如果 s 中的所有代码点都满足 f(c)，或 len(s) == 0，则返回空片段。

// FieldsFunc 不保证调用 f(c) 的顺序，并假定 f 总是为给定的 c 返回相同的值。
func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,bar3..."), f)) // Fields are: ["foo1" "bar2" "bar3"]
}