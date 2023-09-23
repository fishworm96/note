package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// ToUpperSpecial 将 s 作为 UTF-8 编码字节处理，并返回一个将所有 Unicode 字母映射为大写字母的副本，同时优先考虑特殊大小写规则。
func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToUpperSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str)) // Original : ahoj vývojári golang
	fmt.Println("ToUpper : " + string(totitle)) // ToUpper : AHOJ VÝVOJÁRİ GOLANG
}