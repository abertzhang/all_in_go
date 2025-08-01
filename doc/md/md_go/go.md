## 配置代理

```
export GOPROXY=https://goproxy.cn,direct
//win
set GOPROXY=https://goproxy.cn
//在项目文件夹中
go env -w GOPROXY=https://goproxy.cn,direct

```

```
//检查是否生效
go env | grep GOPROXY
```

## Linux执行不中断

```
nohup ./server_go &
nohup ./server_go > server.log 2>&1 &
tail -f server.log
tail -f nohup.out  # 实时跟踪日志输出


//
tail -f server.log
//查看进程
ps aux | grep server_go
//
kill <PID>
```

## go get

| 参数        | 说明                                                    |
| ----------- | ------------------------------------------------------- |
| `-d`        | 仅下载代码，不编译/安装（常用于仅获取代码）。           |
| `-u`        | 更新包及其依赖到最新版本（`-u=patch` 仅更新补丁版本）。 |
| `-t`        | 同时下载测试所需的依赖。                                |
| `-v`        | 显示详细日志。                                          |
| `-insecure` | 允许通过 HTTP 下载（默认强制 HTTPS）。                  |

## 参考文档

https://www.topgoer.com/

[Golang 中文学习文档](https://golang.halfiisland.com/)

http://www.golang.ltd/

https://gin-gonic.com/docs/

