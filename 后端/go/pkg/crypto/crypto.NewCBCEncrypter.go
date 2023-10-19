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
	// 从安全的位置加载您的密钥，并在多个 NewCipher 调用中重复使用它。（显然不要将此示例键用于任何实际操作。如果要将密码转换为密钥，请使用合适的包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("exampleplaintext")

	// CBC 模式适用于块，因此可能需要将明文填充到下一个整个块。有关此类填充的示例，请参阅 https://tools.ietf.org/html/rfc5246#section-6.2.3.2。在这里，我们假设明文已经具有正确的长度。
	if len(plaintext) % aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV需要是唯一的，但不是安全的。因此，通常将其包含在密文的开头。
	ciphertext := make([]byte, aes.BlockSize + len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 重要的是要记住，密文必须经过身份验证（即使用crypto/hmac）以及加密才能安全。

	fmt.Printf("%x\n", ciphertext) // 37464f0360489f7dc6bdcaba96b45b66c68fbe8b5ca953b7ac437b8be948d901
}