## package rsa

import "crypto/rsa"

软件包 rsa 实现了 PKCS #1 和 RFC 8017 中规定的 RSA 加密。

RSA 是本软件包用于实现公钥加密或公钥签名的单一基本操作。

使用 RSA 进行加密和签名的原始规范是 PKCS #1，"RSA 加密 "和 "RSA 签名 "默认是指 PKCS #1 1.5 版。不过，该规范存在缺陷，新设计应尽可能使用第 2 版，通常仅使用 OAEP 和 PSS。

本软件包包含两套接口。当不需要更抽象的接口时，可使用 v1.5/OAEP 进行加密/解密，使用 v1.5/PSS 进行签名/验证。如果需要对公钥基元进行抽象，私钥类型将实现 crypto 软件包中的解密器和签名器接口。

除 GenerateKey、PrivateKey.Precompute 和 PrivateKey.Validate 外，该软件包中的其他操作均使用恒定时间算法实现。其他所有操作都只泄露相关值的比特大小，而这些值都取决于所选的密钥大小。

## Index

### Constants

```go
const (
  // PSSSaltLengthAuto 使 PSS 签名中的盐在签名时尽可能大，并在验证时自动检测。
  PSSSaltLengthAuto = 0
  // PSSSaltLengthEqualsHash 使盐长度等于签名中使用的哈希长度。
  PSSSaltLengthEqualsHash = -1
)
```

### Variables

```go
var ErrDecryption = error.New("crypto/rsa: decryption error")
```

ErrDecryption 表示信息解密失败。为了避免自适应攻击，它故意含糊其辞。

```go
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA key size")
```

ErrMessageTooLong 会在尝试加密或签名密钥过大的报文时返回。使用 SignPSS 时，如果盐的大小过大，也会返回 ErrMessageTooLong。

```go
var ErrVerification = errors.New("crypto/rsa: verification error")
```

ErrVerification 表示验证签名失败。为了避免自适应攻击，它故意含糊其辞。

### func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, ...) ([]byte, error)

DecryptOAEP 使用 RSA-OAEP 解密密文。

OAEP 由散列函数参数化，散列函数用作随机神谕。给定信息的加密和解密必须使用相同的哈希函数，sha256.New() 是一个合理的选择。

随机参数是传统参数，可以忽略，也可以为零。

标签参数必须与加密时给出的值相匹配。详情请参阅 EncryptOAEP。

### func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)

DecryptPKCS1v15 使用 RSA 和 PKCS #1.5 中的填充方案对明文进行解密。随机参数是传统参数，可以忽略，也可以为零。

请注意，该函数是否返回错误会泄露秘密信息。如果攻击者能使该函数重复运行，并知道每个实例是否返回错误，那么他们就能像拥有私钥一样解密和伪造签名。请参阅 DecryptPKCS1v15SessionKey 了解解决此问题的方法。

### func DecryptPKCS1v15SessionKey(random io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error

DecryptPKCS1v15SessionKey 使用 RSA 和 PKCS #1.5 中的填充方案对会话密钥进行解密。随机参数是传统参数，可以忽略，也可以为零。

如果密文长度错误或密文大于公有模数，DecryptPKCS1v15SessionKey 将返回错误。否则，不会返回错误信息。如果填充有效，生成的明文信息将被复制到密钥中。否则，密钥保持不变。这些替代过程在恒定时间内完成。该函数的用户最好事先生成一个随机会话密钥，然后使用生成的密钥值继续执行协议。

需要注意的是，如果会话密钥太小，攻击者有可能对其进行暴力破解。如果攻击者能做到这一点，他们就能知道是否使用了随机值（因为对于相同的密文，随机值是不同的），从而知道填充是否正确。这也违背了该功能的初衷。使用至少 16 字节的密钥可以防止这种攻击。

此方法实现了 RFC 3218 第 2.3.2 节 [1] 中描述的针对 Bleichenbacher 选择密文攻击 [0] 的保护措施。虽然这些保护措施大大增加了 Bleichenbacher 攻击的难度，但只有在使用 DecryptPKCS1v15SessionKey 的协议的其他部分在设计时考虑到这些因素，这些保护措施才会有效。特别是，如果使用解密会话密钥的任何后续操作泄露了密钥的任何信息（例如，它是静态密钥还是随机密钥），那么缓解措施就会失效。这种方法的使用必须非常谨慎，通常只有在与现有协议（如 TLS）兼容性绝对必要时才能使用，因为现有协议在设计时就考虑到了这些特性。

- [0] "针对基于 RSA 加密标准 PKCS #1 的协议的自选密文攻击"，Daniel Bleichenbacher，《密码学进展》（Crypto '98)
- [1] RFC 3218，《防止对 CMS 的百万信息攻击》，https://www.rfc-editor.org/rfc/rfc3218.html。

### func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)

EncryptOAEP 使用 RSA-OAEP 对给定信息进行加密。

OAEP 由散列函数参数化，散列函数用作随机oracle。给定信息的加密和解密必须使用相同的哈希函数，sha256.New() 是一个合理的选择。

随机参数用作熵源，以确保对同一信息加密两次不会得到相同的密文。大多数应用程序应使用 crypto/rand.Reader 作为随机参数。

标签参数可包含不会被加密的任意数据，但这些数据会给出信息的重要上下文。例如，如果给定的公钥用于加密两种类型的信息，那么可以使用不同的标签值来确保用于一种目的的密文不会被攻击者用于另一种目的。如果不需要，标签值可以为空。

信息长度不得超过公开模长度减去哈希长度的两倍，再减去 2。

### func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

EncryptPKCS1v15 使用 RSA 和 PKCS #1.5 中的填充方案对给定信息进行加密。信息长度不得超过公有模数长度减去 11 字节。

随机参数用作熵源，以确保对同一信息加密两次不会得到相同的密文。大多数应用程序应使用 crypto/rand.Reader 作为随机参数。请注意，返回的密文并不确定地取决于从 random 中读取的字节，在不同调用和/或不同版本之间可能会发生变化。

警告：使用此函数加密非会话密钥明文是危险的。请在新协议中使用 RSA OAEP。

### func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)

SignPKCS1v15 使用 RSA PKCS #1.5 的 RSASSA-PKCS1-V1_5-SIGN 计算散列的签名。请注意，散列必须是使用给定散列函数对输入信息散列后的结果。如果哈希值为零，则直接签署哈希值。除互操作性外，这种做法并不可取。

随机参数是传统参数，可以忽略，也可以为零。

该函数是确定性的。因此，如果可能的信息集很小，攻击者就有可能建立一个从信息到签名的映射，并识别签名信息。与以往一样，签名提供的是真实性，而不是保密性。

### func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, ...) ([]byte, error) 添加于1.2

SignPSS 使用 PSS 计算摘要的签名。

digest 必须是使用给定的散列函数对输入信息散列后的结果。opts 参数可以为 nil，在这种情况下将使用合理的默认值。如果设置了 opts.Hash，则会覆盖 hash。

签名将根据信息、密钥和盐的大小，使用 rand 中的字节进行随机化。大多数应用程序都应使用 crypto/rand.Reader 作为 rand。

### func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error

hashed 是使用给定散列函数对输入信息散列的结果，sig 是签名。返回零错误表示签名有效。如果哈希值为零，则直接使用哈希值。除互操作性外，这种做法并不可取。

### func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error 添加于1.2

VerifyPSS 验证 PSS 签名。

返回 nil 错误表示签名有效。摘要必须是使用给定的散列函数对输入信息散列后的结果。opts 参数可以为 nil，在这种情况下将使用合理的默认值。

### type CRTValue

```go
type CRTValue struct {
  Exp *big.Int // D mod (prime-1).
  Coeff *big.Int // R-Coeff ≡ 1 mod Prime。
  R *big.Int // 之前的素数乘积（包括 p 和 q）。
}
```

CRTValue 包含预计算的中文余数定理值。

### type OAEPOptions 添加于1.5

```go
type OAEPOptions struct {
  // Hash 是生成掩码时使用的哈希函数。
  Hash crypto.Hash

  // MGFHash 是 MGF1 使用的哈希函数。
  // 如果为零，则使用哈希值。
  MGFHash crypto.Hash

  // Label 是任意字节字符串，必须等于加密时使用的值。
  Label []byte
}
```

OAEPOptions 是使用 crypto.Decrypter 接口向 OAEP 解密传递选项的接口。

### type PKCS1v15DecryptOptions 添加于1.5

```go
type PKCS1v15DecryptOptions struct {
  // SessionKeyLen 是要解密的会话密钥的长度。如果不是零，则解密期间的填充错误将导致返回此长度的随机纯文本，而不是错误。这些替代方案在恒定的时间内发生。
  SessionKeyLen int
}
```

PKCS1v15DecryptOptions 用于使用 crypto.Decrypter 接口向 PKCS #1.5 解密传递选项。

### type PSSOptions 添加于1.2

```go
type PSSOptions struct {
  // SaltLength 控制 PSS 签名中使用的盐的长度。它可以是正数字节，也可以是特殊的 PSSSaltLength 常量之一。
  SaltLenth int

  // 哈希是用于生成消息摘要的哈希函数。如果不是零，它将覆盖传递给 SignPSS 的哈希函数。使用 PrivateKey.Sign 时是必需的。
  Hash crypto.Hash
}
```

PSSOptions 包含用于创建和验证 PSS 签名的选项。

#### func (opts *PSSOptions) HashFunc() crypto.Hash

HashFunc 返回 opts.Hash，这样 PSSOptions 就实现了 crypto.SignerOpts.HashFunc 的功能。

### type PrecomputedValues

```go
type PrecomputedValues struct {
  Dp, Dq *big.Int // D mod (P-1)（或 mod Q-1）
  Qinv *big.Int // Q^-1 mod P

  // CRTValues 用于第 3 个和后续素数。由于历史事故，前两个素数的 CRT 在 PKCS #1 中的处理方式不同，互操作性非常重要，因此我们对此进行了镜像。

  // 已弃用：为了向后兼容，Precompute 仍会填充这些值，但不会使用这些值。多素 RSA 非常罕见，并且由此软件包实现，没有进行 CRT 优化以限制复杂性。
  CRTValues []CRTValue
  // 包含过滤或未导出的字段
}
```

### type PrivateKey

```go
type PrivateKey struct {
  PublicKey // 公共部分。
  D *big.Int // 私有指数
  Primes []*big.Int // N 的质因数，有 >= 2 个元素。

  // Precomputed 包含可加快 RSA 操作速度的预计算值（如果可用）。它必须通过调用 PrivateKey.Precompute 生成，并且不得修改。
  Precomputed PrecomputedValues
}
```

私钥代表 RSA 密钥

#### func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)

GenerateKey 按给定的比特大小生成随机 RSA 私钥。

大多数应用程序应使用 crypto/rand.Reader 作为 rand。请注意，返回的密钥并不确定地取决于从 rand 读取的字节，在不同调用和/或不同版本之间可能会发生变化。

#### func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error) 废除

#### func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error) 添加于1.5

解密用 priv 对密文进行解密。如果 opts 为 nil 或 *PKCS1v15DecryptOptions 类型，则执行 PKCS #1.5 解密。否则，opts 必须是 *OAEPOptions 类型，否则将执行 OAEP 解密。

#### func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool 添加于1.5

Equal 报告 priv 和 x 的值是否相等。它忽略预计算值。

#### func (priv *PrivateKey) Precompute()

预计算会执行一些计算，以加快未来的私钥操作。

#### func (priv *PrivateKey) Public() crypto.PublicKey 添加于1.4

Public 返回与 priv 对应的公钥。

#### func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) 添加于1.4

用 priv 签名摘要，从 rand 读取随机性。如果 opts 是 *PSSOptions，则将使用 PSS 算法，否则将使用 PKCS #1.5 版。摘要必须是使用 opts.HashFunc() 对输入信息进行散列的结果。

该方法实现了 crypto.Signer，这是一个支持私钥部分保存在硬件模块等中的密钥的接口。普通用途应直接使用此软件包中的 Sign* 函数。

#### func (riv *PrivateKey) Validate() error

Validate 对密钥进行基本的正确性检查。如果密钥有效，则返回 nil，否则返回错误信息。

### type PublicKey

```go
type PublicKey struct {
  N *big.Int // 模量
  E int // 公共指数
}
```

公钥代表 RSA 密钥的公开部分。

#### func (pub *PublicKey) Equal(x crypto.PublicKey) bool 添加于1.15

Equal 报告 pub 和 x 是否具有相同的值。

#### func (pub *PublicKey) Size() int 添加于1.11

Size 返回以字节为单位的模大小。该公开密钥的原始签名和密文的大小相同。