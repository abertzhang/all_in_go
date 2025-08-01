## 下载地址

```
https://www.openssl.org/
https://github.com/openssl/openssl
//mac自带openssl
```

## 支持功能模块

```
# 查看标准命令
openssl list-standard-commands  
//查看加密算法（如AES、DES）
openssl list-cipher-algorithms     
//查看哈希算法（如SHA256、MD5)
openssl list-message-digest-commands 

```



## 证书密钥

### 证书要求

服务端和客户端证书必须由同一CA签发，且证书的`Common Name`需与通信地址匹配‌

证书文件需使用PEM格式（Go原生支持），若使用`.crt`文件需确认编码正确性‌

Go的`tls`包支持`.pem`和`.crt`格式，但需确保CA证书与服务端/客户端证书的签发关系

服务端证书需包含`subjectAltName`字段（如`DNS:localhost`）‌

### 根CA私钥

```
//ok
openssl genrsa -out ca.key 4096
```

### 根CA证书conf

```
文本
```

### 根CA证书csr

```
//生成ca.csr  需要ca.key和ca.conf
openssl req -new -key ca.key -out ca.csr -config ca.cnf
```

```
//ok
openssl req -new -key ca.key -out ca.csr -subj "/CN=myflutter.cn"
```

### 根CA证书

```
//ok-生成CA根证书,需要ca.csr和ca.key
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.pem
```

```
//生成CA根证书,需要ca.conf和ca.key
openssl req -new -x509 -days 3650 -key ca.key -out ca.pem -config ca.conf
```

```
//
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt -extfile ca_csr.cnf -extensions v3_ca
```

```
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

### 服务端私钥

```
# ok-生成 2048 位 RSA 私钥
openssl genrsa -out server.key 2048
```

### 服务端csr

```
openssl req -new -key server.key -out server.csr -subj "/CN=myflutter.cn"
```

```
#ok- 生成CSR请求文件,需要ca.conf和server.key
openssl req -new -key server.key -out server.csr -config ca.conf
```

### 服务端证书

```
//ok----生成server.crt 需要server.csr和ca.pem和ca.key
openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server.pem -days 365
```

### 客户端私钥

```
# ok-生成客户端私钥和证书（步骤同服务端）
openssl genrsa -out client.key 2048
```

### 客户端csr

```
//
openssl req -new -key client.key -out client.csr -subj "/CN=myflutter.cn"
```

```
//ok-生成client.csr 需要ca.conf和client.key
openssl req -new -key client.key -out client.csr -config ca.conf
```

### 客户端证书

```
//ok---生成client.crt 需要client.csr和ca.pem和ca.key
openssl x509 -req -in client.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out client.pem -days 365
```

## ca.conf

```
# ========== CA 基础配置段 ==========
[ CA_default ]
dir             = ./ssl             # CA 工作根目
certs           = $dir/certs        # 已签发证书存储路径
crl_dir         = $dir/crl          # CRL 吊销列表目录
new_certs_dir   = $dir/newcerts     # 新签发证书存放路
database        = $dir/index.txt    # 证书签发记录数据库文件
serial          = $dir/serial       # 证书序列号文件（16进制
private_key     = $dir/private/ca.key  # CA 私钥路
certificate     = $dir/certs/ca.crt    # CA 证书路径

# ========== 策略控制 ==========
[ policy_match ]
countryName            = match     # 必须与 CA 证书国家代码一
stateOrProvinceName    = match     # 必须与 CA 省份一致
organizationName       = match     # 必须与 CA 组织名称一致
commonName             = supplied  # 通用名需用户提供

# ========== 证书请求默认参数 ==========
[ req ]
default_bits        = 4096         # 密钥长度（根 CA 建议 ≥4096
default_md          = sha256       # 摘要算法（禁止 MD5/SHA1
distinguished_name  = dn
x509_extensions     = v3_ca

# ========== X.509 扩展字段 ==========
[ v3_ca ]
basicConstraints        = critical, CA:TRUE  # 声明 CA 身
keyUsage                = critical, keyCertSign, cRLSign  # 限定用
subjectKeyIdentifier    = hash
authorityKeyIdentifier  = keyid:always
```

## server.conf

```
[ req ]
default_bits        = 2048
prompt              = no
default_md          = sha256
distinguished_name  = dn
req_extensions      = req_ext

[ dn ]
CN                  = myflutter.cn  # 服务端域名（必须与实际访问地址一致）
O                   = MyCompany
L                   = Beijing
C                   = CN

[ req_ext ]
subjectAltName      = @alt_names

[ alt_names ]
DNS.1               = myflutter.cn
IP.1                = 192.168.1.100  # 可选：若通过IP访问需添加

```

## 自签名证书

```
openssl genrsa -out ca.key 2048
openssl req -new -key ca.key -out ca.csr -subj "/CN=myflutter.cn"
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.pem
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=myflutter.cn"
openssl x509 -req -in server.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out server.pem -days 365
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=myflutter.cn"
openssl x509 -req -in client.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out client.pem -days 365
```

## 证书转换

```
//DER编码的CER文件转换
openssl x509 -inform der -in certificate.cer -outform pem -out certificate.pem
```

```
//BASE64编码的CER文件转换
openssl x509 -in certificate.cer -out certificate.pem
```

```
//处理含私钥的PFX文件
openssl pkcs12 -in certname.pfx -nocerts -out key.pem -nodes  # 提取私钥
openssl pkcs12 -in certname.pfx -nokeys -out cert.pem        # 提取证书
cat cert.pem key.pem > combined.pem                          # 合并为完整PEM文件
```

```
//查看PEM文件内容
openssl x509 -text -in certificate.pem -noout

```

