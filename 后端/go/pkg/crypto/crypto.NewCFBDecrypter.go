package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// 从安全的地方加载密匙，并在多次调用 NewCipher 时重复使用。(如果你想把口令转换成密钥，请使用合适的软件包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("7dd015f06bec7f1b8f6559dad89f4131da62261786845100056b353194ad")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV 需要是唯一的，但并不安全。因此，通常把它放在密码文本的开头。
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// 如果两个参数相同，XORKeyStream 可以就地工作。
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext) // some plaintext
}