## package ecdsa

import "crypto/ecdsa"

软件包 ecdsa 实现了 FIPS 186-4 和 SEC 1 2.0 版中定义的椭圆曲线数字签名算法。

该软件包生成的签名不是确定性的，而是将熵与私钥和信息混合在一起，从而在随机源失效的情况下达到相同的安全级别。

## Index

### func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

Sign 使用私人密钥对一个哈希值（应该是对较大信息的哈希值）进行签名。如果哈希值长度大于私钥曲线阶的比特长度，哈希值将被截断到该长度。它以一对整数的形式返回签名。大多数应用程序应使用 SignASN1 而不是直接处理 r、s。

### func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error) 添加于1.15

SignASN1 使用私人密钥对一个哈希值（应该是对较大信息的哈希值）进行签名。如果哈希值长度大于私人密钥曲线阶的比特长度，哈希值将被截断到该长度。它会返回 ASN.1 编码的签名。

签名是随机的。大多数应用程序应使用 crypto/rand.Reader 作为 rand。请注意，返回的签名并不确定地取决于从 rand 读取的字节，在不同调用和/或不同版本之间可能会发生变化。

### func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool

Verify 使用公钥 pub 验证哈希值 r, s 中的签名。其返回值记录了签名是否有效。大多数应用程序应使用 VerifyASN1 而不是直接处理 r、s。

### func VerifyASN1(pub *PrivateKey, hash, sig []byte) bool 添加于1.15

VerifyASN1 使用公钥 pub 验证散列的 ASN.1 编码签名 sig。其返回值记录了签名是否有效。

### type PrivateKey

```go
type PrivateKey struct {
  PublicKey
  D *big.Int
}
```

PrivateKey 表示 ECDSA 私钥。

#### func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)

GenerateKey 为指定曲线生成新的 ECDSA 私钥。

大多数应用程序应使用 crypto/rand.Reader 作为 rand。请注意，返回的密钥并不确定地取决于从 rand 读取的字节，可能会在不同调用和/或不同版本之间发生变化。

#### func (k *PrivateKey) ECDH() (*ecdh.PrivateKey, error) 添加于1.20

ECDH 将 k 作为 ecdh.PrivateKey 返回。如果根据 ecdh.Curve.NewPrivateKey 的定义，密钥无效，或者如果 crypto/ecdh.Curve.NewPrivateKey 不支持该曲线，则返回错误。

#### func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool 添加于1.15

Equal 报告 priv 和 x 的值是否相同。

有关如何比较曲线的详细信息，请参阅 PublicKey.Equal。

#### func (priv *PrivateKey) Public() crypto.PublicKey 添加于1.4

Public 返回与 priv 对应的公钥。

#### func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) 添加于1.4

Sign 使用 priv 对摘要进行签名，并从 rand 读取随机值。opts 参数目前没有使用，但为了与 crypto.Signer 接口保持一致，它应该是用于摘要信息的哈希函数。

该方法实现了 crypto.Signer，这是一个支持密钥的接口，其中私人部分保存在硬件模块等中。普通用途可直接使用此软件包中的 SignASN1 函数。

### type PublicKey

```go
type PublicKey struct {
  elliptic.Curve
  X, Y *big.Int
}
```

PublicKey 表示 ECDSA 公钥。

#### func (k *PublicKey) ECDH() (*ecdh.PublicKey, error) 添加于1.20

ECDH 将 k 作为 ecdh.PublicKey 返回。如果根据 ecdh.Curve.NewPublicKey 的定义，该密钥无效，或者如果 crypto/ecdh.Curve.NewPublicKey 不支持该曲线，则返回错误。

#### func (pub *PublicKey) Equal(x crypto.PublicKey) bool 添加于1.15

Equal 报告 pub 和 x 是否具有相同的值。

只有当两个密钥具有相同的曲线值时，才会被认为具有相同的值。请注意，elliptic.P256() 和 elliptic.P256().Params() 的值是不同的，因为后者是通用实现，而不是恒定时间实现。