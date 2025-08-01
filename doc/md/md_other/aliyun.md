

## 阿里 ECS

```
//
aliyun.com
Zhang&
//公网ip
106.15.73.200
```

## 安装Golang

```
//
https://go.dev/dl/
```



```
//解压
tar -C /usr/local -xzf go*.tar.gz
//删除

```

```
vim /etc/profile
//
export PATH=$PATH:/opt/software/go/bin
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
```

```
go version
//清理缓存并重新构建
rm -rf $GOPATH/pkg/* $GOPATH/bin/*
```

## 安装Nodejs

```
/*
这里，-x选项表示解压，-v表示详细模式（可选，用于显示解压过程），-f指定文件名。
*/
tar -xvf node-v22.14.0-linux-arm64.tar.xz
```

