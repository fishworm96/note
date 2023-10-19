package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io"
	"os"
)

// StreamReader 将 Stream 封装为 io.Reader。它调用 XORKeyStream 来处理通过的每个数据片段。
func main() {
	// 从安全的地方加载密匙，并在多次调用 NewCipher 时重复使用。(如果你想把口令转换成密钥，请使用合适的软件包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	encrypted, _ := hex.DecodeString("cf0495cc6f75dafc23948538e79904a9")
	bReader := bytes.NewReader(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 如果密钥对每个密码文本都是唯一的，那么使用零 IV 也是可以的。
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	reader := &cipher.StreamReader{S: stream, R: bReader}
	// 将输入复制到输出流，边复制边解密。
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}

	// 请注意，这个示例非常简单，因为它省略了对加密数据的任何验证。如果真的以这种方式使用 StreamReader，攻击者就可以在输出中翻转任意位。

}
