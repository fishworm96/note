## package crypto

import "crypto"

crypto 软件包收集了常用的加密常量。

## Index

### func RegisterHash(h Hash, f func() hash.Hash)

RegisterHash 注册一个函数，返回给定哈希函数的新实例。该函数用于在实现哈希函数的软件包中通过 init 函数调用。

### type Decrypter 添加于1.5

```go
type Decrypter interface {
  // Public 返回对应于不透明私钥的公钥。
  Public() PublicKey

  // Decrypt 解密消息。opts 参数应该适合所使用的基元。有关详细信息，请参阅每个实现中的文档。
  Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```

解密器是不透明私钥的接口，可用于非对称解密操作。硬件模块中保存的 RSA 密钥就是一个例子。

### type DecrypterOpts 添加于1.5

```go
type DecrypterOpts any
```

### type Hash

```go
type Hash uint
```

哈希标识另一个软件包中实现的加密哈希函数。

```go
const (
  MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // 未实现; MD5+SHA1 用于 TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b
)
```

#### func (h Hash) Available() bool

可用报告给定的哈希函数是否已链接到二进制文件中。

#### func (h Hash) HashFunc() Hash

HashFunc 只需返回 h 的值，这样 Hash 就实现了 SignerOpts。

#### func (h Hash) New() hash.Hash

New 返回一个计算给定哈希函数的新 hash.Hash。如果散列函数没有链接到二进制文件中，New 就会崩溃。

#### func (h Hash) Size() int

Size 返回给定散列函数产生的摘要的长度（以字节为单位）。它不要求将相关的哈希函数链接到程序中。

#### func (h Hahs) String() string 添加于1.5

### type PrivateKey

PrivateKey 表示使用未指定算法的私人密钥。

尽管出于向后兼容的原因，该类型是一个空接口，但标准库中的所有私钥类型都实现了以下接口

```go
interface {
  Public() crypto.PublicKey
  Equal(x crypto.PrivateKey) bool
}
```

以及 Signer 和 Decrypter 等特定用途接口，可用于提高应用程序的类型安全性。

### type PublicKey 添加于1.2

PublicKey 表示使用未指定算法的公钥。

尽管出于向后兼容的原因，该类型是一个空接口，但标准库中的所有公钥类型都实现了以下接口

```go
interface {
  Equal(x crypto.PublicKey) bool
}
```

可用于提高应用中的类型安全性。

### type Signer 添加于1.4

```go
type Signer interface {
  // Public 返回对应于不透明私钥的公钥。
  PUblic() PublicKey

  // 使用私钥对符号进行签名摘要，可能使用来自兰德的熵。对于 RSA 密钥，生成的签名应为 PKCS #1 v1.5 或 PSS 签名（由选项指示）。对于 （EC）DSA 密钥，它应该是 DER 序列化的 ASN.1 签名结构。

  // Hash 实现了 SignerOpts 接口，在大多数情况下，可以简单地传入用作 opts 的哈希函数。符号还可以尝试将断言类型选择类型转换为其他类型，以便获得算法特定的值。有关详细信息，请参阅每个包中的文档。

  // 请注意，当需要较大消息的哈希签名时，调用方负责对较大的消息进行哈希处理，并将哈希（作为摘要）和哈希函数（作为选择）传递给 Sign。
  Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

Signer 是不透明私钥的接口，可用于签名操作。例如，保存在硬件模块中的 RSA 密钥。

### type SignerOpts 添加于1.2

```go
type SignerOpts interface {
  // HashFunc 返回用于生成传递给 Signer.Sign 的消息的哈希函数的标识符，否则返回零表示未进行哈希处理。
  HashFunc() Hash
}
```