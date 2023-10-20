package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// 从安全的位置加载您的密钥，并在多个 NewCipher 调用中重复使用它。（显然不要将此示例键用于任何实际操作。）如果要将密码转换为密钥，请使用合适的包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("73c86d43a9d700a253a96c85b0f6b03ac9792e0e757f869cca306bd3cba1c62b")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// IV需要是唯一的，但不是安全的。因此，通常将其包含在密文的开头。
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC 模式始终以整块为单位运行。
	if len(ciphertext) % aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	// 如果两个参数相同，CryptBlocks 可以就地工作。
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	// 如果原始明文长度不是块大小的倍数，则在加密时必须添加填充，此时将删除填充。有关示例，请参阅 https://tools.ietf.org/html/rfc5246#section-6.2.3.2。但是，请务必注意，密文在解密之前必须进行身份验证（即使用crypto/hmac），以避免创建填充甲骨文。

	fmt.Printf("%s\n", ciphertext) // exampleplaintext
}