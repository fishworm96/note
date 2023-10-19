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
	// 从安全的地方加载密匙，并在多次 Seal/Open 调用中重复使用。(如果要将口令转换为密钥，请使用 bcrypt 或 scrypt 等合适的软件包。
	// 解码时，密钥应为 16 字节（AES-128）或 32 字节（AES-256）。
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("exampleplaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// 对于给定密钥，切勿使用超过 2^32 个随机非ces，因为存在重复的风险。
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext) // af18bfdf54be0e67db46c54f52b18360af9b969959f488a9998796e337e97504
}
