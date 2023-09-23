package main

import (
	"bytes"
	"fmt"
	"unicode"
)

// ToTitleSpecial 将 s 作为 UTF-8 编码字节处理，并返回一份将所有 Unicode 字母映射为标题大小写的副本，其中优先考虑特殊大小写规则。
func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToTitleSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str)) // Original : ahoj vývojári golang
	fmt.Println("ToTitle : " + string(totitle)) // ToTitle : AHOJ VÝVOJÁRİ GOLANG
}