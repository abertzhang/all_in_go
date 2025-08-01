## 子命令

```go
func main() {
    if len(os.Args) < 2 {
        fmt.Println("请指定子命令：start/stop")
        return
    }

    switch os.Args {
    case "start":
        cmd := flag.NewFlagSet("start", flag.ExitOnError)
        port := cmd.Int("port", 8080, "端口号")
        cmd.Parse(os.Args[2:])
        fmt.Println("启动服务，端口:", *port)
    case "stop":
        cmd := flag.NewFlagSet("stop", flag.ExitOnError)
        force := cmd.Bool("force", false, "强制停止")
        cmd.Parse(os.Args[2:])
        fmt.Println("停止服务，强制模式:", *force)
    default:
        fmt.Println("未知命令")
    }
}

```

```
./app start -port=9000    # 输出：启动服务，端口: 9000
./app stop -force         # 输出：停止服务，强制模式: true
```

