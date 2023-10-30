package main

import (
	"crypto/sha256"
	"crypto/rsa"
	"crypto"
	"fmt"
	"os"
)

func main() {
	message := []byte("message to be signed")

	// // 只有较小的信息可以直接签名；因此签名的是信息的哈希值，而不是信息本身。这就要求散列函数具有抗碰撞性。在撰写本文时（2016 年），SHA-256 是用于此目的的强度最低的哈希函数。
hashed := sha256.Sum256(message)
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature)
}