package main

import (
	"crypto/des"
)

func main() {
	// 当需要使用 EDE2 时，也可以通过复制 16 字节密钥的前 8 字节来使用 NewTripleDESCipher。
	ede2Key := []byte("example key 1234")

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, ede2Key[:16]...)
	tripleDESKey = append(tripleDESKey, ede2Key[:8]...)

	_, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	// 有关如何使用加密块进行加密和解密，请参阅 crypto/cipher。
}