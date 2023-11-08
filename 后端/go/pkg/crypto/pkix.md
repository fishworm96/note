## package pkix

软件包 pkix 包含用于 ASN.1 解析和序列化 X.509 证书、CRL 和 OCSP 的共享低级结构。

## Index

### type AlgorithmIdentifier

```go
type AlgorithmIdentifier struct {
  Algorithm asn1.ObjectIdentifier
  Parameters asn1.RawValue `asn1: "optional"`
}
```

AlgorithmIdentifier 表示同名的 ASN.1 结构。参见 RFC 5280 第 4.1.1.2 节。

### AttributeTypeAndValue

```go
type AttributeTypeAndValue struct {
  Type asn1.ObjectIdentifier
  Value any
}
```

AttributeTypeAndValue 反映了 RFC 5280 第 4.1.2.4 节中的同名 ASN.1 结构。

### type AttributeTypeAndValueSET 添加于1.3

```go
type AttributeTypeAndValueSET struct {
  Type asn1.ObjectIdentifier
  Value [][]AttributeTypeAndValue `asn1:"set"`
}
```

AttributeTypeAndValueSET 表示 RFC 2986（PKCS #10）中 AttributeTypeAndValue 序列的 ASN.1 序列集。

### type CertificateList 废除

#### func(certList *CertificateList) HasExpired(now time.Time) bool

### type Extension

```go
type Extension struct {
  Id asn1.ObjectIdentifier
  Critical bool `asn1:"optional"`
  Value []byte
}
```

扩展表示同名的 ASN.1 结构。参见 RFC 5280 第 4.2 节。

### type Name

```go
type Name struct {
  Country, Organization, OrganizationalUnit []string
  Locality, Province []string
  StreetAddress, PostalCode []string
  SerialNumber, CommonName string

  // // 名称包含所有已解析的属性。在解析区分名称时，可使用它来提取本软件包未解析的非标准属性。当 marshaling 到 RDNSequences 时，Names 字段将被忽略，请参阅 ExtraNames。
  Names []AttributeTypeAndValue

  // ExtraNames 包含的属性将被原始复制到任何已 marshaled 的区分名称中。其值覆盖具有相同 OID 的任何属性。解析时不会填充 ExtraNames 字段，请参阅 "名称"。
  ExtraNames []AttributeTypeAndValue
}
```

名称表示 X.509 识别名称。这只包括 DN 的常用元素。请注意，Name 只是 X.509 结构的近似值。如果需要准确的表示，可使用 asn1.Unmarshal 将原始主体或签发人作为 RDNSequence。

#### func (n *Name) FillFromRDNSequence(rdns *RDNSequence)

FillFromRDNSequence 从提供的 RDNSequence 中填充 n。多条目 RDN 会被扁平化，所有条目都会被添加到相关的 n 字段中，但不会保留分组。

#### func (n Name) String() string 添加于1.10

String 返回 n 的字符串形式，大致遵循 RFC 2253 区分名称的语法。

#### func (n Name) ToRDNSequence() (ret RDNSequence)

ToRDNSequence 将 n 转换为单个 RDNSequence。以下属性编码为多值 RDN：

- Country
- Organization
- OrganizationalUnit
- Locality
- Province
- StreetAddress
- PostalCode

每个 ExtraNames 条目都编码为一个单独的 RDN。

### type RDNSequence

```go
type RDNSequence []RelativeDistinguishedNameSET
```

#### func (r RDNSequence) String() string 添加于1.10

String 返回序列 r 的字符串表示，大致遵循 RFC 2253 区分名称语法。

### type RelativeDistinguishedNameSET

```go
type RelativeDistinguishedNameSET []AttributeTypeAndValue
```

### type RevokedCertificate

```go
type RevokedCertificate struct {
  SerialNumber *big.Int
  RevocationTime time.Time
  Extensions []Extension `asn1:"optional"`
}
```

RevokedCertificate 表示同名的 ASN.1 结构。参见 RFC 5280 第 5.1 节。

### type TBSCertificateList 废除