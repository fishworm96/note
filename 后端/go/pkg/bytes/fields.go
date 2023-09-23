package main

import (
	"bytes"
	"fmt"
)

// Fields 将 s 解释为 UTF-8 编码的码位序列。它会按照 unicode.IsSpace 的定义，围绕一个或多个连续空白字符的每个实例分割片段 s，返回 s 的子片段，如果 s 只包含空白字符，则返回空片段。
func main() {
	fmt.Printf("Fields are:%q", bytes.Fields([]byte("  foo bar  baz   "))) // Fields are: ["foo" "bar" "baz"]

}