package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// ToLowerSpecial 将 s 作为 UTF-8 编码字节处理，并返回一个将所有 Unicode 字母映射为小写字母的副本，同时优先考虑特殊大小写规则。
func main() {
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original :" + string(str)) // Original :AHOJ VÝVOJÁRİ GOLANG
	fmt.Println("ToLower :" + string(totitle)) // ToLower :ahoj vývojári golang
}