package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	ciphertext, _ := hex.DecodeString("4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460")
	label := []byte("orders")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating RSA key pair: %s\n", err)
		return
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))

	// 请记住，加密仅提供机密性。这在假定真实性之前，应先对密文进行签名，甚至然后，考虑消息可能会被重新排序。
}
