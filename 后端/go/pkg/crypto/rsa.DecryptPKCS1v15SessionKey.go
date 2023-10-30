package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// 混合方案应至少使用 16 字节的对称密钥。这里我们读取了 RSA 解密未解密时将使用的随机密钥格式良好。
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic("RNG failure")
	}

	rsaCiphertext, _ := hex.DecodeString("aabbccddeeff")

	if err := rsa.DecryptPKCS1v15SessionKey(nil, rsaPrivateKey, rsaCiphertext, key); err != nil {
		// 由此产生的任何错误都将是“公开的”——这意味着它们可以在没有任何秘密信息的情况下确定。（对于实例，如果给定 RSA 的密钥长度是不可能的公钥。
		fmt.Fprintf(os.Stderr, "Error from RSA decryption: %s\n", err)
		return
	}

	// 给定生成的密钥，可以使用对称方案来解密更大的密文。
	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes.NewCipher failed: " + err.Error())
	}

	// 由于密钥是随机的，因此使用固定的随机数是可以接受的，因为（key， nonce） 对仍将是唯一的，根据需要。
	var zeroNonce [12]byte
	aead, err := cipher.NewGCM(block)
	if err != nil {
		panic("cipher.NewGCM failed: " + err.Error())
	}
	ciphertext, _ := hex.DecodeString("00112233445566")
	plaintext, err := aead.Open(nil, zeroNonce[:], ciphertext, nil)
	if err != nil {
		// RSA 密文格式不正确;解密将此处失败，因为 AES-GCM 密钥不正确。
		fmt.Fprintf(os.Stderr, "Error decrypting: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))
}