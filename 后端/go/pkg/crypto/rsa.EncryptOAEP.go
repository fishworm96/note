package main

import (
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
	"os"
)

func main() {
	secretMessage := []byte("send reinforcements, we're going to advance")
	label := []byte("orders")

	// crypto/rand.Reader 是随机化加密函数的一个很好的熵源。
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256(), rng, &test2048Key.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// 由于加密是一个随机函数，所以每次加密的密文都会不同。
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}