package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// 从安全的地方加载密匙，并在多次调用 NewCipher 时重复使用。(如果你想把口令转换成密钥，请使用合适的软件包，如 bcrypt 或 scrypt。
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	bReader := bytes.NewReader([]byte("some secret text"))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 如果密钥对每个密码文本都是唯一的，那么使用零 IV 也是可以的。
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	// 将输入复制到输出缓冲区，边复制边加密。
	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	// 请注意，这个示例非常简单，因为它省略了对加密数据的任何验证。如果真的以这种方式使用 StreamReader，攻击者就可以在解密结果中翻转任意位。

	fmt.Printf("%x\n", out.Bytes()) // cf0495cc6f75dafc23948538e79904a9
}
