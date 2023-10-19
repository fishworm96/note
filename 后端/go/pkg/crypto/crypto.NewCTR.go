package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// 从安全的地方加载密匙，并在多次调用 NewCipher 时重复使用。(如果你想将口令转换成密钥，请使用合适的软件包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV 需要是唯一的，但并不安全。因此，通常把它放在密码文本的开头。
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 重要的是要记住，密文必须经过验证（即使用 crypto/hmac）和加密才能确保安全。

	// CTR 模式对于加密和解密都是一样的，因此我们也可以使用 NewCTR 对密文进行解密。

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2) // some plaintext
}
