## package elliptic

import "crypto/elliptic"

椭圆软件包实现了质域上的标准 NIST P-224、P-256、P-384 和 P-521 椭圆曲线。

除了使用 crypto/ecdsa 所需的 P224、P256、P384 和 P521 值外，已不再直接使用该软件包。大多数其他用途应转用更高效、更安全的 crypto/ecdh，或转用第三方模块来实现较低级别的功能。

## Index

### func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error) 废除

### func Marshal(curve Curve, x, y *big.Int) []byte 废除

### func MarshalCompressed(curve Curve x, y *big.Int) []byte 添加于1.15

MarshalCompressed 将曲线上的点转换为 SEC 1（2.0 版）第 2.3.3 节规定的压缩形式。如果点不在曲线上（或为无穷远处的常规点），则行为未定义。

### func Unmarshal(curve Curve, data []byte) (x, y *big.Int) 废除

### func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int) 添加于1.15

UnmarshalCompressed 将 MarshalCompressed 序列化的点转换为 x、y 对。如果点不是压缩形式、不在曲线上或为无穷远点，则会出错。出错时，x = nil。

### type Curve

```go
type Curve interface {
  // Params 返回曲线的参数。
  Params()

  // IsOnCurve 报告给定的 (x,y) 是否位于曲线上。

	// 过时：这是一个低级的不安全 API。对于 ECDH，请使用 crypto/ecdh 软件包。crypto/ecdh 中 NIST 曲线的 NewPublicKey 方法接受与 Unmarshal 函数的编码相同，并执行曲线上检查。
  IsOnCurve(x, y *big.Int) bool

  // Add 返回 (x1,y1) 和 (x2,y2) 的和。

  // 已废弃：这是一个低级的不安全 API。
  Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

  // Double 返回 2 * (x, y)。

  // 已废弃：这是一个低级的不安全 API。
  Double(x1, y1 *big.Int) (x, y *big.Int)

  // ScalarMult 返回 k*(x,y)，其中 k 是大二进制整数。

  // 已废弃：这是一个不安全的低级 API。对于 ECDH，请使用 crypto/ecdh 软件包。大多数 ScalarMult 的使用都可以通过调用 crypto/ecdh 中 NIST 曲线的 ECDH crypto/ecdh 中的 NIST 曲线方法。
  ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

  // ScalarBaseMult 返回 k*G，其中 G 是群的基点 k 是大二进制整数。

  // 已弃用：这是一个不安全的低级应用程序接口。对于 ECDH，请使用 crypto/ecdh 软件包。ScalarBaseMult 的大多数用法可通过调用 crypto/ecdh 中的 crypto/ecdh 中的 PrivateKey.PublicKey 方法。
  ScalarBaseMult(k []byte) (x, y *big.Int)
}
```

A 曲线表示 a=-3 的简式韦尔斯特拉斯曲线。

当输入不是曲线上的点时，Add、Double 和 ScalarMult 的行为未定义。

需要注意的是，虽然 Add、Double、ScalarMult 或 ScalarBaseMult（但不包括 Unmarshal 或 UnmarshalCompressed 函数）可以返回位于无穷远处（0，0）的常规点，但该点不在曲线上。

除了 P224()、P256()、P384() 和 P521() 返回的曲线实现外，其他曲线实现已不再使用。

#### func P224() Curve

P224 返回一个实现 NIST P-224 的曲线（FIPS 186-3，D.2.2 节），也称为 secp224r1。该曲线的 CurveParams.Name 为 "P-224"。

多次调用该函数将返回相同的值，因此可用于相等检查和切换语句。

加密操作使用恒定时间算法实现。

#### func P256() Curve

P256 返回实现 NIST P-256 的曲线（FIPS 186-3，D.2.3 节），也称为 secp256r1 或 prime256v1。 该曲线的 CurveParams.Name 为 "P-256"。

多次调用此函数将返回相同的值，因此可用于相等检查和切换语句。

加密操作使用恒定时间算法实现。

#### func P384() Curve

P384 返回一个实现 NIST P-384（FIPS 186-3，D.2.4 节）的曲线，也称为 secp384r1。该曲线的 CurveParams.Name 为 "P-384"。

多次调用该函数将返回相同的值，因此可用于相等检查和切换语句。

加密操作使用恒定时间算法实现。

#### func P521() Curve

P521 返回实现 NIST P-521（FIPS 186-3，D.2.5 节）的曲线，也称为 secp521r1。该曲线的 CurveParams.Name 为 "P-521"。

多次调用该函数将返回相同的值，因此可用于相等检查和切换语句。

加密操作使用恒定时间算法实现。

### type CurveParams

```go
type CurveParams struct {
  P *big.Int // 下层字段的顺序
  N *big.Int // 基点顺序
  B *big.Int // 曲线方程的常数
  Gx, Gy *big.Int // (x,y) 的基点
  BitSize int // 下层字段的大小
  Name string // 曲线的统一名称
}
```

CurveParams 包含椭圆曲线的参数，还提供了一个通用的、非恒定时间的 Curve 实现。

通用曲线实现已被弃用，使用自定义曲线（非 P224()、P256()、P384() 和 P521() 返回的曲线）不能保证提供任何安全属性。

#### func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) 废除

#### func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int) 废除

#### func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool 废除

#### func (curve *CurveParams) Params() *CurveParams

#### func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int) 废除

#### func (curve *CurveParams) ScalarMult(Bx, By *big.Int) (*big.Int, *big.Int) 废除

