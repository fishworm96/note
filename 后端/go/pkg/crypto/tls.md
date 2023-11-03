## package tls

软件包 tls 部分实现了 RFC 5246 规定的 TLS 1.2 和 RFC 8446 规定的 TLS 1.3。

## Index

### Constants

```go
const (
  // TLS 1.0 - 1.2 密码套件。
	TLS_RSA_WITH_RC4_128_SHA                      uint16 = 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA                 uint16 = 0x000a
	TLS_RSA_WITH_AES_128_CBC_SHA                  uint16 = 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA                  uint16 = 0x0035
	TLS_RSA_WITH_AES_128_CBC_SHA256               uint16 = 0x003c
	TLS_RSA_WITH_AES_128_GCM_SHA256               uint16 = 0x009c
	TLS_RSA_WITH_AES_256_GCM_SHA384               uint16 = 0x009d
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA              uint16 = 0xc007
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          uint16 = 0xc009
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          uint16 = 0xc00a
	TLS_ECDHE_RSA_WITH_RC4_128_SHA                uint16 = 0xc011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA           uint16 = 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA            uint16 = 0xc013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA            uint16 = 0xc014
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       uint16 = 0xc023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256         uint16 = 0xc027
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256         uint16 = 0xc02f
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       uint16 = 0xc02b
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384         uint16 = 0xc030
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       uint16 = 0xc02c
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   uint16 = 0xcca8
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 uint16 = 0xcca9

  // TLS 1.3 密码套件。
	TLS_AES_128_GCM_SHA256       uint16 = 0x1301
	TLS_AES_256_GCM_SHA384       uint16 = 0x1302
	TLS_CHACHA20_POLY1305_SHA256 uint16 = 0x1303

  // TLS_FALLBACK_SCSV 不是标准密码套件，而是客户端正在执行版本回退的指示器。请参阅 RFC 7507。
  LS_FALLBACK_SCSV uint16 = 0x5600

  // 具有正确 _SHA256 后缀的相应密码套件的旧名称，保留以保持向后兼容性。
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305   = TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305 = TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
)
```

该软件包正在执行或已经执行的密码套件 ID 列表。
参见 https://www.iana.org/assignments/tls-parameters/tls-parameters.xml

```go
const (
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
	VersionTLS13 = 0x0304

  // 已弃用：SSLv3 在加密上已损坏，此包不再支持。请参见 golang.org/issue/32716。
  VersionSSL30 = 0x0300
)
```

```go
const (
	QUICEncryptionLevelInitial = QUICEncryptionLevel(iota)
	QUICEncryptionLevelEarly
	QUICEncryptionLevelHandshake
	QUICEncryptionLevelApplication
)
```

### func CipherSuiteName(id uint16) string 添加于1.14

CipherSuiteName 返回传入的密码套件 ID 的标准名称（例如 "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"），如果本软件包未实现该密码套件，则返回 ID 值的后备表示。

### func Listen(network, laddr string, config *Config) (net.Listener, error)

使用 net.Listen 创建一个 TLS 监听器，接受指定网络地址上的连接。配置必须为非零，且必须包含至少一个证书，否则将设置 GetCertificate。

### func NewListener(inner net.Listener, config *Config) net.Listener

NewListener 创建一个 Listener，它接受来自内部 Listener 的连接，并用 Server 封装每个连接。配置文件必须为非零，且必须包含至少一个证书，否则将设置 GetCertificate。

### VersionName(version uint16) string 添加于1.21.0

VersionName 返回所提供 TLS 版本号的名称（例如 "TLS 1.3"），如果该版本未在此软件包中实现，则返回该值的后备表示。

### type AlterError 添加于1.21.0

AlertError 是 TLS 警报。

使用 QUIC 传输时，QUICConn 方法将返回一个封装了 AlertError 的错误，而不是发送 TLS 警报。

#### func (e AlertError) Error() string 添加于1.21.0

### type Certificate

```go
type Certificate struct {
  Certificate [][]byte
  // PrivateKey 包含与 Leaf 中的公钥对应的私钥。这必须实现加密。具有 RSA、ECDSA 或 Ed25519 公钥的签名者。
  // 对于TLS 1.2及以下的服务器，它还可以实现加密。带有 RSA PublicKey 的解密器。

  PrivateKey crypto.PrivateKey
  // SupportedSignatureAlgorithms 是一个可选列表，用于限制 PrivateKey 可用于哪些签名算法。

  SupportedSignatureAlgorithms []SignatureScheme
  // OCSPStaple 包含一个可选的 OCSP 响应，该响应将提供给请求它的客户端。

  OCSPStaple []byte
  // SignedCertificateTimestamps 包含签名证书时间戳的可选列表，该列表将提供给请求它的客户端。

  SignedCertificateTimestamps [][]byte
  // Leaf 是叶证书的解析形式，可以使用 x509 进行初始化。ParseCertificate 用于减少每次握手处理。如果为 nil，则将根据需要解析叶证书。

  Leaf *x509.Certificate
}
```

证书是由一个或多个证书组成的链，叶片为先。

#### func LoadX509KeyPair(certFile, KeyFile string) (Certificate, error)

LoadX509KeyPair 可从一对文件中读取并解析一对公钥/私钥。文件必须包含 PEM 编码数据。证书文件可能包含叶子证书之后的中间证书，以形成证书链。成功返回时，Certificate.Leaf 将为零，因为证书的解析形式不会被保留。

#### func X509KeyPair(certPEMBlock, KeyPEMBlock []byte) (Certificate, error)

X509KeyPair 从一对 PEM 编码数据中解析一对公钥/私钥。成功返回时，Certificate.Leaf 将为零，因为解析后的证书形式不会被保留。

### type CertificateRequestInfo 添加于1.8

```go
type CertificateRequestInfo struct {
  // AcceptableCA 包含零个或多个 DER 编码的 X.501 可分辨名称。这些是服务器希望通过其对返回的证书进行签名的根 CA 或中间 CA 的名称。空切片表示服务器没有首选项。
  AcceptableCAs [][]byte

  // SignatureSchemes 列出服务器愿意验证的签名方案。
  SignatureSchemes []SignatureScheme

  // 版本是为此连接协商的 TLS 版本。
  Version uint16
  // 包含已筛选或未导出字段
}
```

CertificateRequestInfo 包含服务器 CertificateRequest 信息中的信息，用于要求客户端提供证书和控制证明。

#### func (c *CertificateRequestInfo) Context() context.Context 添加于1.17

Context 返回正在进行的握手的上下文。该上下文是传递给 HandshakeContext 的上下文的子上下文（如果有），并在握手结束时取消。

#### func (cri *CertificateRequestInfo) SupportsCertificate(c *Certificate) error 添加于1.14

如果发送 CertificateRequest 的服务器支持所提供的证书，则 SupportsCertificate 返回 nil。否则，它会返回一个错误，说明不兼容的原因。

### type CertificateVerificationError 添加于1.20

```go
type CertificateVerificationError struct {
  // 不应修改 UnverifiedCertificates 及其内容。
  UnverifiedCertificates []*x509.Certificate
  Error error
}
```

当证书验证在握手过程中失败时，将返回 CertificateVerificationError。

#### func (e *CertificateVerificationError) Error() string 添加于1.20

#### func (e *CertificateVerificationError) Unwrap() error 添加于1.20

### type CipherSuite 添加于1.14

```go
type CipherSuite struct {
  ID uint16
  Name string

  // 支持的版本是可以协商此密码套件的 TLS 协议版本列表。
  SupportedVersions []uint16

  // 如果密码套件由于其基元、设计或实现而存在已知的安全问题，则为不安全。
  Insecure bool
}
```

CipherSuite 是 TLS 密码套件。请注意，本软件包中的大多数函数都接受并公开密码套件 ID，而不是这种类型。

#### func CipherSuites() []*CipherSuite 添加于1.14

CipherSuites 返回此软件包当前实现的密码套件列表，不包括那些存在安全问题的套件，后者由 InsecureCipherSuites 返回。

该列表按 ID 排序。请注意，此软件包选择的默认密码套件可能取决于静态列表无法捕获的逻辑，因此可能与此函数返回的密码套件不一致。

#### func InsecureCipherSuites() []*CipherSuite 添加于1.14

InsecureCipherSuites 返回该软件包当前执行的存在安全问题的密码套件列表。

大多数应用程序不应使用此列表中的密码套件，而应只使用 CipherSuites 返回的密码套件。

### type ClientAuthType

```go
type ClientAuthType int
```

ClientAuthType 声明服务器将遵循的 TLS 客户端身份验证策略。

```go
const (
  // NoClientCert 指示在握手期间不应请求任何客户端证书，如果发送了任何证书，则不会对其进行验证
  NoClientCert ClientAuthType = iota
  // RequestClientCert 指示在握手期间应请求客户端证书，但不要求客户端发送任何证书
  RequestClientCert
  // RequireAnyClientCert 指示在握手期间应请求客户端证书，并且客户端至少需要发送一个证书，但该证书不需要有效。
  RequireAnyClientCert
  // VerifyClientCertIfGiven 指示在握手期间应请求客户端证书，但不要求客户端发送证书。如果客户端确实发送了证书，则该证书必须有效
  VerifyClientCertIfGiven
  // RequireAndVerifyClientCert 指示应请求客户端证书在握手期间，并且至少需要一份有效证书由客户端发送。
  RequireAndVerifyClientCert
)
```

#### func (i ClientAuthType) String() string 添加于1.15

### type ClientHelloInfo 添加于1.4

```go
type ClientHelloInfo struct {
  // CipherSuites 列出了客户端支持的 CipherSuite（例如 TLS_AES_128_GCM_SHA256、TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256）。
  CipherSuites []uint16

  // ServerName 表示客户端请求的服务器的名称为了支持虚拟主机。仅当客户端正在使用 SNI（请参阅 RFC 4366 的第 3.1 节）。
  ServerName string

  // SupportedCurves 列出客户端支持的椭圆曲线。仅当“支持的椭圆曲线”（Supported Elliptic Curves）正在使用扩展（请参阅 RFC 4492 的第 5.1.1 节）。
  SupportedCurves []CurveID

  // SupportedPoints 列出了客户端支持的点格式。仅当“支持的点格式扩展”时，才设置 SupportedPoints正在使用（请参阅 RFC 4492 的第 5.1.2 节）。
  SupportedPoints []uint8

  // SignatureSchemes 列出客户端的签名和哈希方案愿意验证。仅当 Signature正在使用算法扩展（请参阅 RFC 5246，第 7.4.1.4.1 节）。
  SignatureSchemes []SignatureScheme

  // SupportedProtos 列出了客户端支持的应用程序协议。仅当应用层协议正在使用协商扩展（请参阅 RFC 7301 的第 3.1 节）。

  // 服务器可以通过在GetConfigForClient 返回值。
  SupportedProtos []string

  // SupportedVersions 列出了客户端支持的 TLS 版本。对于低于 1.3 的 TLS 版本，这是从最大值推断的客户端公布的版本，因此除最大值外的值如果使用，可能会被拒绝。
  SupportedVersions []uint16

  // Conn 是底层网络。连接的 Conn。不要阅读从或写入此连接;这将导致 TLS连接失败。
  Conn net.Conn
  // 包含已筛选或未导出字段
}
```

ClientHelloInfo 包含来自 ClientHello 消息的信息，用于指导 GetCertificate 和 GetConfigForClient 回调中的应用程序逻辑。

#### func (c *ClientHelloInfo) Context() context.Context 添加于1.17

Context 返回正在进行的握手的上下文。该上下文是传递给 HandshakeContext 的上下文的子上下文（如果有），并在握手结束时取消。

#### func (chi *ClientHelloInfo) SupportsCertificate(c *Certificate) error 添加于1.14

如果发送 ClientHello 的客户端支持所提供的证书，则 SupportsCertificate 返回 nil。否则会返回错误信息，说明不兼容的原因。

如果此 ClientHelloInfo 传递给了 GetConfigForClient 或 GetCertificate 回调，此方法将考虑相关的配置。请注意，如果 GetConfigForClient 返回了不同的 Config，本方法将不会考虑这一变化。

除非设置了 c.Leaf，否则此函数将调用 x509.ParseCertificate，这可能会产生很大的性能代价。

### type ClientSessionCache 添加于1.3

```go
type ClientSessionCache interface {
  // 获取与给定键关联的 ClientSessionState 的搜索。返回时，如果找到一个，则为 ok。
  Get(sessionKey string) (session *ClientSessionState, ok bool)

  // Put 使用给定键将 ClientSessionState 添加到缓存中。它可能如果 TLS 1.3 服务器提供多张会话票证。如果使用 nil *ClientSessionState 调用，它应该删除缓存条目。
  Put(sessionKey string, cs *ClientSessionState)
}
```

ClientSessionCache 是 ClientSessionState 对象的缓存，客户端可使用它来恢复与指定服务器的 TLS 会话。ClientSessionCache 的实现需要在不同的程序中同时调用。截至 TLS 1.2，只支持基于 ticket 的恢复，不支持基于 SessionID 的恢复。在 TLS 1.3 中，它们被合并为 PSK 模式，并通过此接口提供支持。

#### func NewLRUClientSessionCache(capacity int) ClientSessionCache 添加于1.3

NewLRUClientSessionCache 返回一个使用 LRU 策略、具有给定容量的 ClientSessionCache。如果容量小于 1，则使用默认容量。

### type ClientSessionState 添加于1.3

```go
type ClientSessionState struct {
  // 包含已筛选或未导出字段
}
```

ClientSessionState 包含客户端恢复之前 TLS 会话所需的状态。

#### func NewResumptionState(ticket []byte, state *SessionState) (*ClientSessionState, error) 添加于1.21.0

NewResumptionState 返回一个状态值，[ClientSessionCache.Get] 可以用它来恢复之前的会话。

NewResumptionState需要由ParseSessionState返回，并且票据和会话状态必须已由ClientSessionState.ResumptionState返回。

#### func (cs *ClientSessionState) ResumptionState() (ticket []byte, state *SessionState, err error) 添加于1.21.0

ResumptionState 返回服务器发送的会话票据（也称为会话标识）以及恢复该会话所需的状态。

它可被 [ClientSessionCache.Put] 调用，以序列化（使用 SessionState.Bytes）和存储会话。

### type Config

```go
type Config struct {
  // Rand 为非 ces 和 RSA 屏蔽提供了熵源。如果 Rand 为零，TLS 将使用包 crypto/rand 中的加密随机读取器。
  // 读取器必须能被多个程序安全使用。
  Rand io.Reader

  // Time 返回当前时间，即从纪元开始的秒数。
  // 如果 Time 为空，TLS 将使用 time.Now。
  Time func() time.Time

  // 证书包含一个或多个证书链，以呈现给连接的另一端。将自动选择与对等方要求兼容的第一个证书。
  // 
  // 服务器配置必须设置 Certificates、GetCertificate 或 GetConfigForClient 之一。执行客户端身份验证的客户端可以设置 Certificates 或 GetClientCertificate。
  // 
  // 注意：如果有多个证书，并且它们没有可选字段 Leaf set，则证书选择将产生显著的每次握手性能成本。
  Certificates []Certificate

  // NameToCertificate 从证书名称映射到 Certificates 的元素。请注意，证书名称的格式可以是“*.example.com”，因此不必是域名。
  // 
  // 已弃用：NameToCertificate 仅允许将单个证书与给定名称相关联。将此字段保留为 nil，让库从 Certificates 中选择第一个兼容链。
  NameToCertificate map[string]*Certificate

  // GetCertificate 根据给定的 ClientHelloInfo 返回证书。仅当客户端提供 SNI 信息或证书为空时，才会调用它。
  // 
  // 如果 GetCertificate 为 nil 或返回 nil，则从 NameToCertificate 检索证书。如果 NameToCertificate 为 nil，则将使用 Certificates 的最佳元素。
  // 
  // 证书一经退回，不得修改。
  GetCertificate func(*ClientHelloInfo) (*Certificate, error)

  // GetClientCertificate（如果不是 nil）在服务器从客户端请求证书时调用。如果设置，则将忽略证书的内容。
  // 
  // 如果 GetClientCertificate 返回错误，则握手将中止，并返回该错误。否则，GetClientCertificate 必须返回非 nil 证书。如果 Certificate.Certificate 为空，则不会向服务器发送任何证书。如果服务器无法接受此操作，则可能会中止握手。
  // 
  // 如果发生重新协商或使用 TLS 1.3，则可以为同一连接多次调用 GetClientCertificate。
  // 
  // 证书一经退回，不得修改。
  GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error)

  // GetConfigForClient（如果不是 nil）是在从客户端收到 ClientHello 后调用的。它可能会返回一个非 nil Config，以便更改将用于处理此连接的 Config。如果返回的 Config 为 nil，则使用原始 Config。该回调返回的 Config 后续不得修改。
  // 
  // 如果 GetConfigForClient 为 nil，则传递给 Server（） 的 Config 将用于所有连接。
  // 
  // 如果在返回的 Config 上显式设置了 SessionTicketKey，或者在返回的 Config 上调用了 SetSessionTicketKeys，则将使用这些键。否则，将使用原始的 Config 密钥（如果它们被自动管理，则可能会轮换）。
  GetConfigForClient func(*ClientHelloInfo) (*Config, error)

  // VerifyPeerCertificate（如果不是 nil）是在 TLS 客户端或服务器进行正常证书验证后调用的。它接收对等方提供的原始 ASN.1 证书，以及正常处理发现的任何经过验证的链。如果它返回非 nil 错误，则握手将中止并导致该错误。
  // 
  // 如果正常验证失败，则握手将在考虑此回调之前中止。如果禁用了正常验证（在客户端上，当设置了 InsecureSkipVerify，或者在服务器上，当 ClientAuth 为 RequestClientCert 或 RequireAnyClientCert 时），则将考虑此回调，但 verifiedChains 参数将始终为 nil。当 ClientAuth 为 NoClientCert 时，不会在服务器上调用此回调。如果 ClientAuth 为 RequestClientCert 或 VerifyClientCertIfGiven，则服务器上的 rawCerts 可能为空。
  // 
  // 恢复连接时不会调用此回调，因为证书在恢复时不会重新验证。
  // 
  // 不得修改 verifiedChains 及其内容。
  VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

  // VerifyConnection（如果不是 nil）是在正常证书验证之后和 TLS 客户端或服务器在 VerifyPeerCertificate 之后调用的。如果它返回非 nil 错误，则握手将中止并导致该错误。
  // 
  // 如果正常验证失败，则握手将在考虑此回调之前中止。此回调将对所有连接运行，包括恢复，而不考虑 InsecureSkipVerify 或 ClientAuth 设置
  VerifyConnection func(ConnectionState) error

  // RootCA 定义客户端在验证服务器证书时使用的根证书颁发机构集。
  // 如果 RootCAs 为 nil，则 TLS 使用主机的根 CA 集。
  RootCAs *x509.CertPool

  // NextProtos 是受支持的应用程序级协议列表，按优先顺序排列。如果两个对等体都支持 ALPN，则所选协议将是此列表中的一个，如果没有相互支持的协议，则连接将失败。如果 NextProtos 为空或对等方不支持 ALPN，则连接将成功，并且 ConnectionState.NegotiatedProtocol 将为空。
  NextProtos []string

  // ServerName 用于验证返回的证书上的主机名，除非给出了 InsecureSkipVerify。它也包含在客户端的握手中以支持虚拟主机，除非它是 IP 地址。
  ServerName string

  // ClientAuth 确定服务器的 TLS 客户端身份验证策略。默认值为 NoClientCert。
  ClientAuth ClientAuthType

  // ClientCA 定义一组根证书颁发机构，如果需要通过 ClientAuth 中的策略验证客户端证书，服务器可以使用这些证书颁发机构。
  ClientCAs *x509.CertPool

  // InsecureSkipVerify 控制客户端是否验证服务器的证书链和主机名。如果 InsecureSkipVerify 为 true，则 crypto/tls 接受服务器提供的任何证书以及该证书中的任何主机名。在此模式下，除非使用自定义验证，否则 TLS 容易受到中间计算机攻击。这应仅用于测试或与 VerifyConnection 或 VerifyPeerCertificate 结合使用。
  InsecureSKipVerify bool

  // CipherSuites 是已启用的 TLS 1.0–1.2 密码套件的列表。列表的顺序将被忽略。请注意，TLS 1.3 密码套件不可配置。
  // 
  // 如果 CipherSuites 为 nil，则使用安全的默认列表。默认密码套件可能会随时间而更改。
  CipherSuites []uint16

  // PreferServerCipherSuites 是一个传统字段，没有任何作用。
  // 
  // 它用于控制服务器是遵循客户端的偏好还是服务器的首选项。现在，服务器根据考虑推断的客户端硬件、服务器硬件和安全性的逻辑选择相互支持的最佳密码套件。
  // 
  // 已废弃：PreferServerCipherSuites 已被忽略。
  PreferServerCipherSuites bool

  // 可以将 SessionTicketsDisabled 设置为 true 以禁用会话票证和 PSK（恢复）支持。请注意，在客户端上，如果 ClientSessionCache 为 nil，则会话票证支持也会被禁用。
  SessionTicketsDisabled bool

  // TLS 服务器使用 SessionTicketKey 来提供会话恢复。
  // 请参阅 RFC 5077 和 RFC 8446 的 PSK 模式。如果为零，则在第一次服务器握手之前，它将填充随机数据。
  // 
  // 已弃用：如果此字段保留为零，则会话票证密钥将每天自动轮换，并在 7 天后删除。若要自定义轮换计划或同步终止同一主机连接的服务器，请使用 SetSessionTicketKeys。
  SessionTicketKey [32]byte

  // ClientSessionCache 是用于 TLS 会话恢复的 ClientSessionState 条目的缓存。它仅供客户端使用。
  ClientSessionCache ClientSessionCache

  // 在服务器上调用 UnwrapSession，以将以前由 [WrapSession] 生成的票证/标识转换为可用的会话。
  // 
  // UnwrapSession 通常会解密票证中的会话状态（例如，使用 [Config.EncryptTicket]），或者使用票证作为句柄来恢复以前存储的状态。它必须使用 [ParseSessionState] 来反序列化会话状态。
  // 
  // 如果 UnwrapSession 返回错误，则连接将终止。如果它返回 （nil， nil），则忽略会话。crypto/tls 可能仍会选择不恢复返回的会话。
  UnwrapSession func(identity []byte, cs ConnectionState) (*SessionState, error)

  // 在服务器上调用 WrapSession 以生成会话票证/标识。
  // 
  // WrapSession 必须使用 [SessionState.Bytes] 序列化会话状态。然后，它可以对序列化状态进行加密（例如[Config.DecryptTicket]） 并将其用作票证，或存储状态并返回其句柄。
  // 
  // 如果 WrapSession 返回错误,连接将被终止。
  // 
  // 警告：返回值将以明文形式公开在网络上并提供给客户端。该应用程序负责对其进行加密和身份验证（以及轮换密钥）或返回高熵标识符。如果不能正确执行此操作，可能会危及当前、以前和将来的连接，具体取决于协议版本。
  WrapSession func(ConnectionState, *SessionState) ([]byte, error)

  // MinVersion 包含可接受的最小 TLS 版本。
  // 
  // 默认情况下，TLS 1.2 当前在充当客户端时用作最低要求，而在充当服务器时使用 TLS 1.0。TLS 1.0 是此软件包支持的最低版本，无论是作为客户端还是作为服务器。
  // 
  // 通过在 GODEBUG 环境变量中包含值“x509sha1=1”，可以暂时将客户端默认值恢复为 TLS 1.0。请注意，此选项将在 Go 1.19 中删除（但仍然可以将此字段显式设置为 VersionTLS10）。
  MinVersion uint16

  // MaxVersion 包含可接受的最大 TLS 版本。
  // 
  // 默认使用此包支持的最高版本，目前为 TLS 1.3。
  MaxVersion uint16

  // CurvePreferences 包含 ECDHE 握手中将使用的椭圆曲线,按首选项顺序排列。如果为空,则使用默认值。客户端将使用第一个首选项作为其在 TLS 1.3 中的密钥共享类型。这可能会在未来发生变化。
  CurvePreferences []CurveID

  // DynamicRecordSizingDisabled 禁用 TLS 记录的自适应调整大小。当为 true 时,始终使用最大可能的 TLS 记录大小。当为 false 时,TLS 记录的大小可能会进行调整,以尝试改善延迟。
  DynamicRecordSizingDisabled bool

  // 重新谈判控制支持哪些类型的重新谈判。默认值为 none,对于绝大多数应用程序都是正确的。
  Renegotiation RenegotiationSupport

  // KeyLogWriter 可以选择以 NSS 密钥日志格式指定 TLS 主密钥的目标，该目标可用于允许外部程序（如 Wireshark）解密 TLS 连接。请参见 https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format。使用 KeyLogWriter 会降低安全性，并且只能用于调试。
  KeyLogWriter io.Writer
  // 包含已筛选或未导出字段
}
```

Config 结构用于配置 TLS 客户端或服务器。在将 Config 传递给 TLS 函数后，不得对其进行修改。Config 可以重复使用；tls 软件包也不会修改它。

#### func (c *Config) BuildNameToCertificate() 废除

#### func (c *Config) Clone() *Config 添加于1.8

克隆会返回 c 的浅克隆值，如果 c 为 nil，则返回 nil。克隆被 TLS 客户端或服务器同时使用的 Config 是安全的。

#### func (c *Config) DecryptTicket(identity []byte, cs ConnectionState) (*SessionState, error) 添加于1.21.0

DecryptTicket 可解密由 Config.EncryptTicket 加密的票据。它可用作 [Config.UnwrapSession] 实现。

如果票据无法解密或解析，DecryptTicket 将返回（nil，nil）。

#### func (C *Config) EncryptTicket(cs ConnectionState, ss *SessionState) ([]byte, error) 添加于1.21.0

EncryptTicket 使用 Config 配置（或默认）的会话票据密钥加密票据。它可用作 [Config.WrapSession] 实现。

#### func (c *Config) SetSessionTicketKeys(keys [][32]byte) 添加于1.5

SetSessionTicketKeys 更新服务器的会话票据密钥。

第一个密钥将在创建新票据时使用，而所有密钥都可用于解密票据。在服务器运行时调用该函数以轮换会话票据密钥是安全的。如果 keys 为空，该函数将崩溃。

调用该函数将关闭会话票据密钥自动轮换功能。

如果多台服务器正在终止同一主机的连接，那么它们都应该拥有相同的会话 ticket 密钥。如果会话记录密钥泄漏，之前记录的和未来使用这些密钥的 TLS 连接都可能受到影响。

### type Conn

```go
type Conn struct {
  // 包含已筛选或未导出字段
}
```

Conn 表示安全连接。它实现了 net.Conn 接口。

#### func Client(conn net.Conn, config *Config) *Conn

客户端使用 conn 作为底层传输，返回一个新的 TLS 客户端连接。配置不能为零：用户必须在配置中设置 ServerName 或 InsecureSkipVerify。

#### func Dial(network, addr string, config *Config) (*conn, error)

Dial 使用 net.Dial 连接到给定的网络地址，然后启动 TLS 握手，并返回结果 TLS 连接。Dial 将 nil 配置等同于零配置；默认值请参阅 Config 文档。

#### func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error) 添加于1.3

DialWithDialer 使用 dialer.Dial 连接到给定的网络地址，然后启动 TLS 握手，并返回由此产生的 TLS 连接。拨号器中给出的任何超时或截止时间都适用于整个连接和 TLS 握手过程。

DialWithDialer 将 "nil "配置等同于 "0 "配置；默认值请参见 Config 文档。

DialWithDialer 内部使用 context.Background；要指定上下文，请使用 Dialer.DialContext，并将 NetDialer 设置为所需的拨号器。

#### func Server(conn net.Conn, config *Config) *Conn

服务器使用 conn 作为底层传输，返回一个新的 TLS 服务器端连接。配置 config 必须为非零，且必须包含至少一个证书，否则将设置 GetCertificate。

#### func (c *Conn) Close() error

Close 关闭连接。

#### func (c *Conn) CloseWrite() error 添加于1.8

CloseWrite 关闭连接的写入端。它只能在握手完成后调用，而不会在底层连接上调用 CloseWrite。大多数调用者应直接使用 CloseWrite。

#### func (c *Conn) ConnectionState() ConnectionState

ConnectionState 返回连接的基本 TLS 详情。

#### func (c *Conn) Handshake() error

如果尚未运行客户端或服务器握手协议，则运行握手协议。

本软件包的大多数用途都无需明确调用握手协议：第一次读取或写入时会自动调用。

要控制取消或设置握手超时，请使用 HandshakeContext 或 Dialer 的 DialContext 方法。

为避免拒绝服务攻击，TLS 服务器或客户端发送的证书中允许的最大 RSA 密钥大小限制为 8192 位。可以通过在 GODEBUG 环境变量中设置 tlsmaxrsasize 来覆盖此限制（例如 GODEBUG=tlsmaxrsasize=4096）。

#### func (c *Conn) HandshakeContext(ctx context.Context) error 添加于1.17

如果客户端或服务器握手协议尚未运行，HandshakeContext 会运行该协议。

提供的上下文必须为非空。如果在握手完成前取消上下文，握手将被中断并返回错误信息。一旦握手完成，取消上下文将不会影响连接。

此软件包的大多数用途都不需要明确调用 HandshakeContext：第一次读取或写入会自动调用它。

#### func (c *Conn) LocalAddr() net.Addr

LocalAddr 返回本地网络地址。

#### func (c *Conn) NetConn() net.Conn 添加于1.18

注意，直接写入或读取该连接会破坏 TLS 会话。

#### func (c *Conn) OCSPResponse() []byte

注意，直接写入或读取该连接会破坏 TLS 会话。

#### func (c *Conn) Read(b []byte) (int, error)

OCSPResponse 返回 TLS 服务器的装订 OCSP 响应（如果有）。(仅对客户端连接有效）。

#### func (c *Conn) RemoteAddr() net.Addr

读取从连接中读取数据。

由于 "读 "调用 "握手"，为了防止无限阻塞，必须在握手尚未完成时调用 "读 "之前为 "读 "和 "写 "设置一个截止日期。请参阅 SetDeadline、SetReadDeadline 和 SetWriteDeadline。

#### func (c *Conn) SetDeadline(t time.Time) error

RemoteAddr 返回远程网络地址。

#### func (c *Conn) SetReadReadDeadline(t time.Time) error

SetDeadline 设置与连接相关的读写截止时间。t 值为零表示读取和写入不会超时。写入超时后，TLS 状态将被破坏，所有未来的写入都将返回相同的错误。

#### func (c *Conn) SetReadDeadline(t time.Time) error)

SetReadDeadline 设置底层连接的读取截止时间。t 值为零表示读取不会超时。

#### func (c *Conn) SetWriteDeadline(t time.Time) error

SetWriteDeadline 设置底层连接的写入截止时间。t 值为零表示写入不会超时。写入超时后，TLS 状态将被破坏，所有未来的写入都将返回相同的错误。

#### func (c *Conn) VerifyHostname(host string) error

VerifyHostname 检查对等证书链是否可用于连接主机。如果有效，则返回 nil；如果无效，则返回错误信息，说明问题所在。

#### func (c *Conn) Write(b []byte) (int, error)

写入将数据写入连接。

由于 "写 "调用 "握手"，为了防止无限阻塞，必须在握手尚未完成时调用 "写 "之前为 "读 "和 "写 "设置一个截止日期。请参阅 SetDeadline、SetReadDeadline 和 SetWriteDeadline。

### type ConnectionState

```go
type ConnectionState struct {
  // 版本是连接使用的 TLS 版本（如 VersionTLS12）。
  Version uint16

  // 如果握手结束，则 HandshakeComplete 为 true。
  HandshakeComplete bool

  // 如果此连接已使用会话票证或类似机制从上一个会话成功恢复，则 DidResume 为 true。
  DidResume bool

  // CipherSuite 是为连接协商的密码套件（e.g.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256、TLS_AES_128_GCM_SHA256）。
  CipherSuite uint16

  // NegotiatedProtocol 是与 ALPN 协商的应用协议。
  NegotiatedProtocol string

  // NegotiatedProtocolIsMutual 用于表示相互 NPN 协商。
  // 
  // 已废弃：该值始终为 true。
  NegotiatedProtocolIsMutual bool

  // erverName 是客户端发送的服务器名称指示扩展的值。它在服务器端和客户端都可用。
  ServerName string

  // PeerCertificates 是 Peer 发送的已解析证书，按发送顺序排列。第一个元素是验证连接所依据的叶证书。
  // 
  // 在客户端，它不能为空。在服务器端，如果 Config.ClientAuth 不是 RequireAnyClientCert 或 RequireAndVerifyClientCert，它可以为空。
  // 
  // 不得修改 PeerCertificates 及其内容。
  PeerCertificates []*x509.Certificate

  // VerifiedChains 是一个或多个链的列表，其中第一个元素是 PeerCertificates[0]，最后一个元素来自 Config.RootCAs（在客户端）或 Config.ClientCAs（在服务器端）。
  // 
  // 在客户端，如果 Config.InsecureSkipVerify 为 false，则设置它。在服务器端，如果 Config.ClientAuth 为 VerifyClientCertIfGiven（并且对等方提供了证书）或 RequireAndVerifyClientCert，则设置它。
  // 
  // 不得修改 VerifiedChains 及其内容。
  VerifiedChains [][]*x509.Certificate

  // SignedCertificateTimestamps 是对等方通过叶证书的 TLS 握手提供的 SCT 列表（如果有）。
  SignedCertificateTimestamps [][]byte

  // OCSPResponse 是对等方为叶证书（如果有）提供的装订联机证书状态协议 （OCSP） 响应。
  OCSPResponse []byte

  // TLSUnique 包含“tls-unique”通道绑定值（请参阅 RFC 5929 的第 3 节）。对于 TLS 1.3 连接和不支持扩展主密钥 （RFC 7627） 的已恢复连接，此值将为 nil。
  TLSUnique []byte
  // 包含已筛选或未导出字段
}
```

ConnectionState 记录连接的基本 TLS 详情。

#### func (cs *ConnectionState) ExportKeyingMateria(label string, context []byte, length int) ([]byte, error) 添加于1.11

ExportKeyingMaterial 返回 RFC 5705 中定义的新片段中导出密钥材料的长度字节。如果上下文为空，则不作为种子的一部分使用。如果连接已通过 Config.Renegotiation 设置为允许重新协商，则此函数将返回错误信息。

在某些情况下，返回的值对连接来说可能不是唯一的。请参阅 RFC 5705 和 RFC 7627 的安全考虑部分，以及 https://mitls.org/pages/attacks/3SHAKE#channelbindings。

### type CurveID 添加于1.3

CurveID 是椭圆曲线的 TLS 标识符类型。请参见 https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8。

在 TLS 1.3 中，该类型被称为 NamedGroup，但目前该库只支持基于椭圆曲线的组。参见 RFC 8446，第 4.2.7 节。

```go
const (
  CurveP256 CurveID = 23
  CurveP384 CurveID = 24
  CurveP521 CurveID = 25
  x25519 CurveID = 29
)
```

#### func (i CurveID) String() string 添加于1.15

### type Dialer 添加于1.15

```go
type Dialer struct {
  NetDialer *net.Dialer

  Config *Config
}
```

Dialer 会根据底层连接的配置和拨号器拨号 TLS 连接。

#### func (d *Dialer) Dial(network, addr string) (net.Conn, error) 添加于1.15

Dial 连接到给定的网络地址，并启动 TLS 握手，返回由此产生的 TLS 连接。

返回的 Conn（如果有）将始终是 *Conn 类型。

Dial 内部使用 context.Background；要指定上下文，请使用 DialContext。

#### func (d *Dialer) DialContext(ctx context.Context, addr string) (net.Conn, error) 添加于1.15

DialContext 会连接到给定的网络地址并启动 TLS 握手，然后返回 TLS 连接。

提供的上下文必须非零。如果上下文在连接完成前过期，则会返回错误信息。一旦连接成功，任何上下文过期都不会影响连接。

返回的 Conn（如果有）将始终是 *Conn 类型。

### type QUICConfig 添加于1.21.0

```go
type QUIConfig struct {
  TLSConfig *Config
}
```

QUICConfig 配置 QUICConn。

### type QUICConn 添加于1.21.0

```go
type QUICConn struct {
  // 包含已筛选或未导出字段
}
```

QUICConn 表示使用 QUIC 实现作为底层传输的连接，如 RFC 9001 所述。

QUICConn 的方法不能安全并发使用。

#### func QUICClient(config *QUICConfig) *QUICConn 添加于1.21.0

QUICClient 返回一个使用 QUICTransport 作为底层传输的新 TLS 客户端连接。配置不能为空。

配置的 MinVersion 必须至少为 TLS 1.3。

#### func QUICServer(config *QUICConfig) *QUICConn 添加于1.21.0

QUICServer 返回一个使用 QUICTransport 作为底层传输的新 TLS 服务器端连接。配置不能为空。

配置的 MinVersion 必须至少为 TLS 1.3。

#### (q *QUICConn) Close() error 添加于1.21.0

关闭连接并停止任何正在进行的握手

#### (q *QUICConn) ConnectionState() ConnectionState 添加于1.21.0

ConnectionState 返回连接的基本 TLS 详情。

#### (q *QUICConn) HandleData(level QUICEncryptionLevel, data []byte) error 添加于1.21.0

HandleData 处理从对等程序接收到的握手字节。它可能会产生连接事件，这些事件可通过 NextEvent 读取。

#### (q *QUICConn) NextEvent() QUICEvent 添加于1.21.0

NextEvent 返回连接上发生的下一个事件。如果没有可用事件，它将返回一个 Kind 为 QUICNoEvent 的事件。

#### (q *QUICConn) SendSessionTicket(opts QUICSessionTicketOptions) error 添加于1.21.0

SendSessionTicket 会向客户端发送会话票据。它会产生连接事件，可使用 NextEvent 读取这些事件。目前，它只能被调用一次。

#### (q *QUICConn) SetTransportParameters(params []byte) 添加于1.21.0

SetTransportParameters 设置要发送给对等设备的传输参数。

服务器连接可能会延迟设置传输参数，直到收到客户端的传输参数。请参阅 QUICTransportParametersRequired。

#### (q *QUICConn) Start(ctx context.Context) error 添加于1.21.0

Start 启动客户端或服务器握手协议。它可能会产生连接事件，这些事件可以用 NextEvent 读取。

Start 最多只能调用一次。

### type QUICEncryptionLevel 添加于1.21.0

QUICEncryptionLevel 表示用于传输握手信息的 QUIC 加密级别。

#### func (I QUICEncryptionLevel) String() string 添加于1.21.0

### QUICEvent 添加于1.21.0

```go
type QUICEvent struct {
  Kind QUICEventKind

  // 为 QUICSetReadSecret、QUICSetWriteSecret 和 QUICWriteData 设置。
  Level QUICEncryptionLevel

  // 为 QUICTransportParameters、QUICSetReadSecret、QUICSetWriteSecret 和 QUICWriteData 设置。
  // 内容归 crypto/tls 所有，在调用下一次 NextEvent 之前一直有效。
  Data []byte

  // 为 QUICSetReadSecret 和 QUICSetWriteSecret 设置。
  Suite uint16
}
```

QUICE 事件是 QUIC 连接上发生的事件。

事件类型由 Kind 字段指定。其他字段的内容与事件类型有关。

### QUICEventKind 添加于1.21.0

QUICEventKind 是 QUIC 连接上的一种操作类型。

```go
const (
  // QUICNoEvent 表示没有可用的事件。
  QUICNoEvent QUICEventKind = iota

  // QUICSetReadSecret 和 QUICSetWriteSecret 提供给定加密级别的读取和写入密钥。
  // 设置了 QUICEvent.Level、QUICEvent.Data 和 QUICEvent.Suite。
  // 
  // 初始加密级别的密钥派生自初始目标连接 ID，而不是由 QUICConn 提供。
  QUICSetReadSecret
  QUICSetWriteSecret

  // QUICWriteData 提供在 CRYPTO 帧中发送给对等设备的数据。
  // QUICEvent.Data 已设置。
  QUICWriteData

  // QUICTransportParameters 提供对等方的 QUIC 传输参数。
  // QUICEvent.Data 已设置。
  QUICTransportParameters

  // QUICTransportParametersRequired 指示调用方必须提供 QUIC 传输参数才能发送到对等方。调用方应使用 QUICConn.SetTransportParameters 设置传输参数，并再次调用 QUICConn.NextEvent。
  // 
  // 如果在调用 QUICConn.Start 之前设置了传输参数，则连接永远不会生成 QUICTransportParametersRequired 事件。
  QUICTransportParametersRequired

  // QUICRejectedEarlyData 表示服务器拒绝了 0-RTT 数据，即使我们提供了它。在返回 QUICEncryptionLevelApplication 密钥之前返回它。
  QUICRejectedEarlyData

  // QUICHandshakeDone 表示 TLS 握手已完成。
  QUICHandshakeDone
)
```

### QUICSessionTicketOptions 添加于1.21.0

```go
type QUICSessionTicketOptions struct {
  // EarlyData 指定票证是否可用于 0-RTT。
  EarlyData bool
}
```

### RecordHeaderError 添加于1.6

```go
type RecordHeaderError struct {
  // Msg 包含一个描述错误的可读字符串。
  Msg string
  // RecordHeader 包含触发错误的 TLS 记录标头的 5 个字节。
  RecordHeader [5]byte
  // Conn 提供底层网络。Conn 在客户端发送的初始握手看起来不像 TLS 的情况下。
  // 如果已经握手或已将 TLS 警报写入连接，则为 n。
  Conn net.Conn
}
```

当 TLS 记录头无效时，将返回 RecordHeaderError。

#### func (e RecordHeaderError) Error() string 添加于1.6

### type RenegotiationSupport 添加于1.7

RenegotiationSupport 枚举了 TLS 重新协商的不同支持级别。TLS 重新协商是在第一次握手后对连接执行后续握手的行为。这大大增加了状态机的复杂性，并引发了许多微妙的安全问题。虽然不支持启动重新协商，但可以启用对接受重新协商请求的支持。

即使启用了重新协商功能，服务器也不能在两次握手之间改变其身份（即叶证书必须相同）。此外，不允许同时进行握手和应用数据流，因此重新协商只能用于与重新协商同步的协议，如 HTTPS。

TLS 1.3 中未定义重新协商。

```go
const (
  // RenegotiateNever 禁用重新协商。
  RenegotiateNever RenegotiationSupport = iota

  // RenegotiateOnceAsClient 允许远程服务器为每个连接请求一次重新协商。
  RenegotiateOnceAsClient

  // 重新协商FreelyAsClient 允许远程服务器重复请求重新协商。
  RenegotiateFreelyAsClient
)
```

### type SessionState 添加于1.21.0

```go
type SessionState struct {
  // Extra 被 crypto/tls 忽略，但由 [SessionState.Bytes] 编码并由 [ParseSessionState] 解析。
  // 
  // 这允许 [Config.UnwrapSession]/[Config.WrapSession] 和 [ClientSessionCache] 实现在此会话旁边存储和检索其他数据。
  // 
  // 若要允许协议堆栈中的不同层共享此字段，应用程序必须仅追加到该字段，而不能替换该字段，并且必须使用即使顺序不一也可以识别的条目（例如，以 id 和版本前缀开头）。
  Extra [][]byte

  // EarlyData 指示票证是否可以用于 QUIC 连接中的 0-RTT。如果拒绝提供 0-RTT 是真的，即使支持，应用程序也可以将其设置为 false。
  EarlyData bool
  // 包含已筛选或未导出字段
}
```

会话状态是一个可恢复的会话。

#### func ParseSessionState(data []byte) (*SessionState, error) 添加于1.21.0

ParseSessionState 会解析由 SessionState.Bytes 编码的会话状态。

#### func (s *SessionState) Bytes() ([]byte, error) 添加于1.21.0

字节对会话（包括任何私有字段）进行编码，以便 ParseSessionState 对其进行解析。该编码包含对未来和可能过去会话的安全性至关重要的秘密值。

具体的编码应视为不透明，在不同的 Go 版本之间可能会发生不兼容的变化。

### type SignatureScheme 添加于1.8

```go
type SignatureScheme uint16
```

SignatureScheme 标识 TLS 支持的签名算法。参见 RFC 8446，第 4.2.3 节。

```go
const (
  // RSASSA-PKCS1-v1_5 算法。
	PKCS1WithSHA256 SignatureScheme = 0x0401
	PKCS1WithSHA384 SignatureScheme = 0x0501
	PKCS1WithSHA512 SignatureScheme = 0x0601

  // 使用公钥 OID rsaEncryption 的 RSASSA-PSS 算法。
  PSSWithSHA256 SignatureScheme = 0x0804
	PSSWithSHA384 SignatureScheme = 0x0805
	PSSWithSHA512 SignatureScheme = 0x0806

	// ECDSA 算法。在 TLS 1.3 中仅受限于特定曲线。
	ECDSAWithP256AndSHA256 SignatureScheme = 0x0403
	ECDSAWithP384AndSHA384 SignatureScheme = 0x0503
	ECDSAWithP521AndSHA512 SignatureScheme = 0x0603

	// EdDSA 算法。
	Ed25519 SignatureScheme = 0x0807

	// TLS 1.2 的传统签名和散列算法。
	PKCS1WithSHA1 SignatureScheme = 0x0201
	ECDSAWithSHA1 SignatureScheme = 0x0203
)
```

#### func (i SignatureScheme) String() string 添加于1.15

### Bugs

crypto/tls 软件包仅针对 CBC 模式加密的 Lucky13 攻击和 SHA1 变体实施了一些应对措施。请参见 http://www.isg.rhul.ac.uk/tls/TLStiming.pdf 和 https://www.imperialviolet.org/2013/02/04/luckythirteen.html。