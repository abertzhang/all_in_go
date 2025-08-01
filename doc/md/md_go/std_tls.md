## Config

### 源码构造器

```go
type Config struct {
    Rand                                io.Reader
    Time                        func() time.Time
    Certificates                        []Certificate
    NameToCertificate                   map[string]*Certificate
    GetCertificate              func(*ClientHelloInfo) (*Certificate, error)
    GetClientCertificate        func(*CertificateRequestInfo) (*Certificate, error)
    GetConfigForClient          func(*ClientHelloInfo) (*Config, error)
    VerifyPeerCertificate       func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error
    VerifyConnection            func(ConnectionState) error
    RootCAs                             *x509.CertPool
    NextProtos                          []string
    ServerName                          string
    ClientAuth                          ClientAuthType
    ClientCAs                           *x509.CertPool
    InsecureSkipVerify                  bool
    CipherSuites                        []uint16
    PreferServerCipherSuites            bool
    SessionTicketsDisabled              bool
    SessionTicketKey                    [32]byte
    ClientSessionCache                  ClientSessionCache
    UnwrapSession               func(identity []byte, cs ConnectionState) (*SessionState, error)
    WrapSession                 func(ConnectionState, *SessionState) ([]byte, error)
    MinVersion                          uint16
    MaxVersion                          uint16
    CurvePreferences                    []CurveID
    DynamicRecordSizingDisabled         bool
    Renegotiation                       RenegotiationSupport
    KeyLogWriter                        io.Writer
    EncryptedClientHelloConfigList      []byte
    EncryptedClientHelloRejectionVerify func(ConnectionState) error
    EncryptedClientHelloKeys            []EncryptedClientHelloKey
    mutex                               sync.RWMutex
    sessionTicketKeys                   []ticketKey
    autoSessionTicketKeys               []ticketKey
}
```
### 方法

```go
//方法对象: (*Config):
ticketKeyFromBytes(b [32]byte) (key tls.ticketKey)
Clone() *tls.Config
initLegacySessionTicketKeyRLocked()
ticketKeys(configForClient *tls.Config) []tls.ticketKey
SetSessionTicketKeys(keys [][32]byte)
rand() io.Reader
time() time.Time
cipherSuites() []uint16
supportedVersions(isClient bool) []uint16
maxSupportedVersion(isClient bool) uint16
curvePreferences(version uint16) []tls.CurveID
supportsCurve(version uint16, curve tls.CurveID) bool
mutualVersion(isClient bool, peerVersions []uint16) (uint16, bool)
getCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error)
BuildNameToCertificate()
writeKeyLog(label string, clientRandom []byte, secret []byte) error
EncryptTicket(cs tls.ConnectionState, ss *tls.SessionState) ([]byte, error)
encryptTicket(state []byte, ticketKeys []tls.ticketKey) ([]byte, error)
DecryptTicket(identity []byte, cs tls.ConnectionState) (*tls.SessionState, error)
decryptTicket(encrypted []byte, ticketKeys []tls.ticketKey) []byte
```



## ClientHelloInfo

```go
type ClientHelloInfo struct {
    CipherSuites      []uint16
    ServerName        string
    SupportedCurves   []CurveID
    SupportedPoints   []uint8
    SignatureSchemes  []SignatureScheme
    SupportedProtos   []string
    SupportedVersions []uint16
    Extensions        []uint16
    Conn              net.Conn
    config            *Config
    ctx               context.Context
}
```

## ConnectionState

### 定义

```go
type ConnectionState struct {
    Version                     uint16
    HandshakeComplete           bool
    DidResume                   bool
    CipherSuite                 uint16
    NegotiatedProtocol          string
    NegotiatedProtocolIsMutual  bool
    ServerName                  string
    PeerCertificates            []*x509.Certificate
    VerifiedChains              [][]*x509.Certificate
    SignedCertificateTimestamps [][]byte
    OCSPResponse                []byte
    TLSUnique                   []byte
    ECHAccepted                 bool
    ekm                         func(label string, context []byte, length int) ([]byte, error)
    testingOnlyDidHRR           bool
    testingOnlyCurveID          CurveID
}
```

### 方法

```
//方法对象: (*ConnectionState):
ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)
```



## ClientAuthType

### 源码

```go
type ClientAuthType int
const (
	NoClientCert ClientAuthType = iota
	RequestClientCert
	RequireAnyClientCert
	VerifyClientCertIfGiven
	RequireAndVerifyClientCert
)
```

| 值                                   | 行为描述                                                     | 典型场景                               |
| ------------------------------------ | ------------------------------------------------------------ | -------------------------------------- |
| ‌**`tls.NoClientCert`**‌               | 不要求客户端提供证书，服务端仅验证自身身份（单向认证）       | 公开API服务或匿名访问场景‌12            |
| ‌**`tls.RequestClientCert`**‌          | 请求客户端证书但不强制验证，客户端可选择性提供证书（服务端忽略验证结果） | 混合环境（部分客户端需认证）‌13         |
| ‌**`tls.RequireAnyClientCert`**‌       | 强制客户端必须提供证书，但仅校验证书存在性（不验证签名、有效期或信任链） | 低安全性内部网络环境‌23                 |
| ‌**`tls.VerifyClientCertIfGiven`**‌    | 若客户端提供证书则验证其合法性；若无证书仍允许连接           | 兼容匿名与认证客户端的场景‌23           |
| ‌**`tls.RequireAndVerifyClientCert`**‌ | 强制客户端提供证书，并严格验证其签名链、有效期及信任关系（需配置`ClientCAs`） | 高安全性场景如金融交易或微服务间通信‌13 |

### 验证严格性

`NoClientCert` < `RequestClientCert` < `RequireAnyClientCert` < `VerifyClientCertIfGiven` < `RequireAndVerifyClientCert`‌

### 依赖配置

- `RequireAndVerifyClientCert`必须配合`ClientCAs`指定可信CA池，否则验证失败‌14
- `VerifyClientCertIfGiven`需结合`ClientCAs`实现有条件验证‌

## 参考文档

[官方TLS文档](https://pkg.go.dev/crypto/tls)