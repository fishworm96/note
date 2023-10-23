## package ecdh

import "crypto/ecdh"

软件包 ecdh 通过 NIST 曲线和 Curve25519 实现了椭圆曲线 Diffie-Hellman 功能。

## Index

### type Curve

```go
type Curve interface {
  // GenerateKey 生成随机私钥。

  // 大多数应用程序应使用 [crypto/rand.Reader] 作为 rand。请注意返回的密钥并不确定地取决于从 rand 读取的字节在不同调用和/或不同版本之间可能会发生变化。
  GenerateKey(rand io.Reader) (*PrivateKey, error)

  // NewPrivateKey 会检查密钥是否有效，并返回一个私钥。

  // 对于 NIST 曲线，这遵循 SEC 1 2.0 版第 2.3.6 节，即相当于将字节解码为固定长度的大端整数，并检查结果是否低于曲线阶次。检查结果是否小于曲线的阶数。零私人密钥为零也会被拒绝，因为相应公钥的编码会不规则。密钥的编码将是不规则的。

  // 对于 X25519，这只检查标量长度。
  NewPrivateKey(key []byte) (*PrivateKey, error)

  // NewPublicKey 会检查密钥是否有效，并返回一个 PublicKey。

  // 对于 NIST 曲线，这将根据 SEC 1 对未压缩点进行解码2.0 版第 2.3.4 节。压缩编码和位于无穷大处的点都会被拒绝。

  // 对于 X25519，只检查 u 坐标长度。反向选择的公钥会导致 ECDH 返回错误。
  NewPrivateKey(key []byte) (*PublicKey, error)
  // 包含已过滤或未导出的方法
}
```

#### func P256() Curve

P256 返回一个实现 NIST P-256 的曲线（FIPS 186-3，D.2.3 节），也称为 secp256r1 或 prime256v1。

多次调用此函数将返回相同的值，可用于等价检查和切换语句。

#### func P384() Curve

P384 返回一个实现 NIST P-384 的曲线（FIPS 186-3，D.2.4 节），也称为 secp384r1。

多次调用该函数将返回相同的值，可用于等价检查和切换语句。

#### func P521() Curve

P521 返回实现 NIST P-521 的曲线（FIPS 186-3，D.2.5 节），也称为 secp521r1。

多次调用该函数将返回相同的值，可用于等价检查和切换语句。

#### func X25519 Curve

X25519 返回一个通过 Curve25519 实现 X25519 功能的 Curve（RFC 7748，第 5 节）。

多次调用该函数将返回相同的值，因此可用于相等检查和切换语句。

### PrivateKey

```go
type PrivateKey struct {
  // 包含已筛选或未导出字段
}
```

PrivateKey 是 ECDH 私钥，通常保密。

这些密钥可使用 crypto/x509.ParsePKCS8PrivateKey 进行解析，并使用 crypto/x509.MarshalPKCS8PrivateKey 进行编码。对于 NIST 曲线，解析后需要使用 crypto/ecdsa.PrivateKey.ECDH 进行转换。

#### func (k *PrivateKey) Bytes() []byte

Bytes 返回私人密钥编码的副本。

#### func (k *PrivateKey) Curve() Curve

#### func (k *PrivateKey) ECDH(remote *PublicKey) ([]byte, error)

ECDH 执行 ECDH 交换并返回共享密钥。私钥和公钥必须使用相同的曲线。

对于 NIST 曲线，按照 SEC 1 2.0 版第 3.3.1 节的规定执行 ECDH，并返回按照 SEC 1 2.0 版第 2.3.5 节编码的 x 坐标。结果永远不会是无穷远处的点。

对于 X25519，将执行 RFC 7748 第 6.1 节规定的 ECDH。如果结果为全零值，ECDH 将返回错误。

#### func (k *PrivateKey) Equal(x crypto.PrivateKey) bool

Equal 返回 x 是否代表与 k 相同的私人密钥。

请注意，可能存在编码不同的等价私人密钥，它们在此检查中返回 false，但作为 ECDH 的输入却表现相同。

只要密钥类型和它们的曲线匹配，这种检查就能在恒定时间内完成。

#### func (k *PrivateKey) Public() crypto.PublicKey

Public 实现了所有标准库私钥的隐式接口。请参阅 crypto.PrivateKey.Public 的文档。

#### func (k *PrivateKey) PublicKey() *PrivateKey

### type PublicKey

```go
type PublicKey struct {
  // 包含已筛选或未导出字段
}
```

PublicKey 是 ECDH 公钥，通常是通过网络发送的对等方的 ECDH 共享。

这些密钥可使用 crypto/x509.ParsePKIXPublicKey 进行解析，并使用 crypto/x509.MarshalPKIXPublicKey 进行编码。对于 NIST 曲线，解析后需要使用 crypto/ecdsa.PublicKey.ECDH 进行转换。

#### func (k *PrivateKey) Bytes() []byte

Bytes 返回公钥编码的副本。

#### func (k *PrivateKey) Curve() Curve

#### func (k *PrivateKey) Equal(x crypto.PublicKey) bool

Equal 返回 x 是否代表与 k 相同的公钥。

请注意，可能存在编码不同的等价公钥，这些等价公钥在此检查中返回 false，但作为 ECDH 的输入却表现相同。

只要密钥类型和它们的曲线匹配，这种检查就能在恒定时间内完成。