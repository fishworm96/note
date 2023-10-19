## package cipher

import "crypto/cipher"

软件包密码实现了标准的区块密码模式，可封装在低级区块密码实现中。参见 https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html 和 NIST 特别出版物 800-38A。

## Index

### type AEAD 添加于1.2

```go
type AEAD interface {
  // NonceSize 返回必须传递给 Seal and Open 的随机数的大小。
  NonceSize() int

  // 开销返回明文与其密文长度之间的最大差值。
  Overhead() int

  // Seal 对明文进行加密和身份验证，对附加数据进行身份验证并将结果附加到 DST，返回更新的切片。对于给定键，nonceSize（） 字节长度必须是 NonceSize（） 字节并且始终是唯一的。

  // 若要将纯文本的存储重用于加密输出，请使用明文 [：0] 作为 dst。否则，DST 的剩余容量不得与明文重叠。
  Seal(dst, nonce, plaintext, additionalData []byte) []byte

  // Open 解密并验证密文，验证其他数据，如果成功，则将生成的明文追加到 dst，返回更新的切片。nonce的长度必须是 NonceSize（）字节，并且它和其他数据都必须与传递给 Seal 的值匹配。

  // 要为解密输出重用密文的存储，请使用密文 [：0] 作为 dst。否则，DST 的剩余容量不得与明文重叠。

  // 即使函数失败，DST 的内容（直至其容量）也可能被覆盖。
  Open(dst, nonce, ciphertext, aditionalData []byte) ([]byte, error)
}
```

AEAD 是一种密码模式，提供带相关数据的验证加密。有关方法的说明，请参阅 https://en.wikipedia.org/wiki/Authenticated_encryption。

#### func NewGCM(cipher Block) (AEAD, error) 添加于1.2

NewGCM 返回以伽罗瓦计数器模式封装的 128 位块密码，并带有标准非密钥长度。

一般来说，GCM 实现执行的 GHASH 操作不是恒定时间操作。如果在硬件支持 AES 的系统上，底层块是由 aes.NewCipher 创建的，则属于例外。详情请参见 crypto/aes 软件包文档。

#### func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error) 添加于1.5

NewGCMWithNonceSize 返回以伽罗瓦计数器模式封装的 128 位块密码，该密码接受给定长度的非ces。长度不得为零。

只有在需要与使用非标准非ce长度的现有密码系统兼容时，才能使用此函数。所有其他用户都应使用 NewGCM，它的速度更快，抗误用能力更强。

#### func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error) 添加于1.11

NewGCMWithTagSize 返回以伽罗瓦计数器模式封装的 128 位块密码，它能生成具有给定长度的标记。

允许标签长度在 12 到 16 字节之间。

只有在需要与使用非标准标记长度的现有密码系统兼容时，才能使用此函数。所有其他用户都应使用 NewGCM，它更能防止误用。

### Block

```go
type Block interface {
  // BlockSize 返回密码的块大小。
  BlockSize() int

  // Encrypt 将 src 中的第一个数据块加密到 dst 中。
  // Dst 和 src 必须完全重叠或完全不重叠。
  Encrypt(dst, src []byte)

  // 解密将 src 中的第一个数据块解密到 dst 中。
  // Dst 和 src 必须完全重叠或完全不重叠。
  Decrypt(dst, src []byte)
}
```

区块表示使用给定密钥实现的区块密码。它具有加密或解密单个区块的功能。模式实现将这一功能扩展到区块流。

### BlockMode

```go
type BlockMode interface {
  // BlockSize 返回模式的数据块大小。
  BlockSize() int

  // CryptBlocks加密或解密许多块。src 的长度必须是块大小的倍数。Dst 和 src 必须完全重叠或根本不重叠。

  // 如果 len（dst） < len（src），CryptBlocks 应该会崩溃。传递比 src 大的 dst 是可以接受的，在这种情况下，CryptBlocks 将仅更新 DST[：LEN（src）]，不会触及 DST 的其余部分。

  // 对 CryptBlocks 的多次调用就像在一次运行中传递了 src 缓冲区的串联一样。也就是说，BlockMode 保持状态，并且不会在每次 CryptBlocks 调用时重置。
  CryptBlocks(dst, src []byte)
}
```

BlockMode 表示以基于区块的模式（CBC、ECB 等）运行的区块密码。

#### func NewCBCDecrypter(b Block, iv []byte) BlockMode

NewCBCDecrypter 返回一个 BlockMode，它使用给定的 Block 以密码块链模式进行解密。iv 的长度必须与 Block 的块大小相同，并且必须与用于加密数据的 iv 相匹配。

#### func NewCBCEncrypter(b Block, iv []byte) BlockMode

NewCBCEncrypter 返回一个 BlockMode，使用给定的 Block 以密码块链模式进行加密。iv 的长度必须与 Block 的块大小相同。

### type Stream

```go
type Stream interface {
  // XORKeyStream XOR 使用密码密钥流中的一个字节对给定切片中的每个字节进行 XOR。Dst 和 src 必须完全重叠或根本不重叠。

  // 如果 len（dst） < len（src），XORKeyStream 应该会恐慌。传递大于 src 的 dst 是可以接受的，在这种情况下，XORKeyStream 只会更新 dst[：len（src）]，而不会触及 dst 的其余部分。

  // 对 XORKeyStream 的多次调用的行为就像在一次运行中传递了 src 缓冲区的串联一样。也就是说，Stream 保持状态，并且不会在每个 CryptBlocks 调用时重置。
  XORKeyStream(dst, src []byte)
}
```

流表示流密码。

#### func NewCFBDecrypter(block Block, iv []byte) Stream

NewCFBDecrypter 返回使用给定区块以密码反馈模式解密的数据流。iv 的长度必须与区块的区块大小相同。

#### func NewCFBEncrypter(block Block, iv []byte) Stream

NewCFBEncrypter 返回使用给定区块以密码反馈模式加密的流。iv 的长度必须与块的块大小相同。

#### func NewCTR(block Block, iv []byte) Stream

NewCTR 返回一个在计数器模式下使用给定区块进行加密/解密的数据流。iv 的长度必须与块的块大小相同。

#### func NewOFB(b Block, iv []byte) Stream

NewOFB 返回一个在输出反馈模式下使用块密码 b 进行加密或解密的流。初始化向量 iv 的长度必须等于 b 的块大小。

### type StreamReader

```go
type StreamReader struct {
  S Stream
  R io.Reader
}
```

StreamReader 将 Stream 封装为 io.Reader。它调用 XORKeyStream 来处理通过的每个数据片段。

#### func (r StreamReader) Read(dst []byte) (n int, err error)

### type StreamWriter

```go
type StreamWriter struct [
  S Stream
  W io.Writer
  Err error // 闲置
]
```

StreamWriter 将 Stream 封装为 io.Writer。它调用 XORKeyStream 来处理通过的每个数据片段。如果任何 "写"（Write）调用返回短数据，则说明 StreamWriter 不同步，必须丢弃。StreamWriter 没有内部缓冲；无需调用 Close 来刷新写入数据。

#### func (w StreamWriter) Close() error

Close 关闭底层 Writer 并返回其 Close 返回值（如果 Writer 也是 io.Closer）。否则返回 nil。

#### func (w StreamWriter) Write(src []byte) (n int, err error)