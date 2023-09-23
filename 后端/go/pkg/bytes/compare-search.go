package main

import (
	"bytes"
	"sort"
)

// Compare 返回一个整数，将两个字节片段按词典顺序进行比较。如果 a == b，结果为 0；如果 a < b，结果为 -1 ；如果 a > b，结果为 +1 。
func main() {
	// 二进制搜索，找到匹配的字节片。
	var needle []byte
	var haystack [][]byte
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		// 找到了
	}
}