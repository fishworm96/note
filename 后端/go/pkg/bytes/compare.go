package main

import (
	"bytes"
)

// Compare 返回一个整数，将两个字节片段按词典顺序进行比较。如果 a == b，结果为 0；如果 a < b，结果为 -1 ；如果 a > b，结果为 +1 。
func main() {
	// 通过与零比较来解释比较的结果。
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		// a 小于 b
	}
	if bytes.Compare(a, b) <= 0 {
		// a 小于等于 b
	}
	if bytes.Compare(a, b) > 0 {
		// a 大于 b
	}
	if bytes.Compare(a, b) >= 0 {
		// a 大于等于 b
	}

	// 在进行相等比较时，首选 Equal（相等），而不是 Compare（比较）。
	if bytes.Equal(a, b) {
		// a 等于 b
	}
	if !bytes.Equal(a, b) {
		// a 不等于 b
	}
}