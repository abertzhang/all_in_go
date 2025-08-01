## 无证书get

```
curl https://myflutter.cn:5432
//查看握手
curl -v https://myflutter.cn:5432
```

## 发送POST请求

```
curl -X POST -d "data" http://example.com
```



```
curl -H "Header: value" http://example.com
```



```
curl -O http://example.com/file.txt
```



```
curl -F "file=@filename" http://example.com/upload
```



```
curl -u user http://example.com
```



```
curl -L http://example.com
```



```
curl -v http://example.com
```



```
curl -c cookies.txt http://example.com
```



```
curl -H "Content-Type: application/json" -d '{"key":"value"}' http://example.com
```

​	

```
//强制使用 TLS 1.2
curl --tlsv1.2 https://example.com
//信任自签名证书
//若服务端使用自签名证书，需通过 --cacert 指定 CA 证书
curl --cacert /path/to/ca-cert.pem https://example.com
//客户端证书认证
//若服务端要求客户端证书，需提供证书和私钥
curl --cert /path/to/client-cert.pem --key /path/to/client-key.pem https://example.com
//查看 TLS 握手详情
curl -v https://example.com
//忽略证书验证（仅测试环境）
//使用 -k 或 --insecure 跳过证书校验
curl -k https://example.com
```

```
CA 根证书‌：由可信机构颁发或自签的根证书（.pem 格式）‌58。
‌客户端证书‌：由 CA 签发的客户端证书（.pem 或 .crt 格式）‌13。
‌客户端私钥‌：与客户端证书配对的私钥（.pem 或 .key 格式）‌35。
‌注‌：若客户端证书为 PKCS#12 格式（.pfx 或 .p12），需转换为 .pem：
curl --cacert ca.pem \          # 验证服务端证书的根证书 ‌:ml-citation{ref="3,7" data="citationList"}
     --cert client.pem \        # 客户端证书文件 ‌:ml-citation{ref="2,3" data="citationList"}
     --key client-key.pem \     # 客户端私钥文件 ‌:ml-citation{ref="3,4" data="citationList"}
     https://api.example.com

```

```
# 准备文件
curl --cacert /etc/ssl/ca.pem \      # CA 根证书
     --cert /etc/ssl/client.pem \     # 客户端证书
     --key /etc/ssl/client-key.pem \  # 客户端私钥
     -H "Content-Type: application/json" \
     -X POST -d '{"data":"test"}' \
     https://internal-api.example.com/v1/endpoint

```



```
curl --cert client.pem --key client-key.pem --pass your_password

```

```
openssl pkcs12 -in client.pfx -out client.pem -nodes  # 转换并保留私钥

```

