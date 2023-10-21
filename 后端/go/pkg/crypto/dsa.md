## package dsa

import "crypto/dsa"

软件包 dsa 实现了 FIPS 186-3 中定义的数字签名算法。

该软件包中的 DSA 操作不是通过恒定时间算法实现的。

已废弃：DSA 是一种传统算法，应使用 Ed25519（由 crypto/ed25519 软件包实现）等现代替代算法。使用 1024 位模数（L1024N160 参数）的密钥在密码学上比较弱，而更大的密钥则不被广泛支持。请注意，FIPS 186-5 不再允许使用 DSA 生成签名。

## Index

### Variables

```go
var ErrInvalidPublicKey = error.New("crypto/dsa: invalid public key")
```

ErrInvalidPublicKey 会在本代码无法使用公钥时产生。FIPS 对 DSA 密钥的格式要求相当严格，但其他代码可能不那么严格。因此，在使用其他代码生成的密钥时，必须处理此错误。

### func GenerateKey(priv *PrivateKey, rand io.Reader) error

GenerateKey 生成一对公钥和私钥。私钥的参数必须已经有效（请参阅 GenerateParameters）。

### func GenerateParameters(params *Parameters, rand io.Reader, sizes ParameterSizes) error

GenerateParameters 会将一组随机、有效的 DSA 参数放入 params 中。即使在速度很快的机器上，这个函数也会耗时数秒。

### func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

Sign 使用私人密钥对任意长度的哈希值（应为较大信息的哈希值）进行签名。它以一对整数的形式返回签名。私钥的安全性取决于 rand 的熵。

请注意，FIPS 186-3 第 4.6 节规定，哈希值应截断为子组的字节长度。此函数本身不执行截断。

请注意，使用攻击者控制的私钥调用 Sign 可能会占用大量 CPU。

### func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool

Verify 使用公开密钥 pub 验证哈希值 r、s 中的签名。它报告签名是否有效。

请注意，FIPS 186-3 第 4.6 节规定，哈希值应截断为子组的字节长度。此函数本身不执行截断。

### type ParameterSizes

```go
type ParameterSizes int
```

ParameterSizes 是一组 DSA 参数中可接受的素数位长的枚举。参见 FIPS 186-3，第 4.2 节。

```go
const (
  L1024N160 ParameterSizes = iota
	L2048N224
	L2048N256
	L3072N256
)
```

### type Parameters

```go
type Parameters struct {
  P, Q, G *big.Int
}
```

参数表示密钥的域参数。这些参数可以在多个密钥中共享。Q 的位长必须是 8 的倍数。

### type PrivateKey

```go
type PrivateKey struct {
  PublicKey
  X *big.Int
}
```

PrivateKey 表示 DSA 私钥。

### type PublicKey

```go
type PublicKey struct {
  Parameters
  Y *big.Int
}
```

PublicKey 表示 DSA 公钥。