## 💡 核心特性

1. 极简设计
2. 轻量级实现
3. 跨平台支持
4. 优秀的性能
5. 丰富的生态系统

## 🛠️ 核心功能

1. 异步执行器
2. 任务调度
3. IO操作
4. 定时器
5. 通道通信
6. 网络编程

## 📊 应用场景

1. 小型Web服务
2. CLI工具
3. 嵌入式系统
4. 轻量级网络应用
5. 资源受限环境

## 安装smol

```rust
//方式一
cargo add smol
//方式二:toml添加
smol = "2.0.2"
```



## 简例-01

```rust
use std::future::join;
use std::time::Duration;
use smol::lock::futures;
use smol::Timer;
fn main() {
    //简例
    smol::block_on(run());
}    
async fn run() {
    print!("开始执行");
    //创建异步任务
    let task = async {
        Timer::after(Duration::from_secs(1)).await;
        println!("异步任务完成");
    };
    //执行任务
    task.await;
    println!("程序结束");
}

```



## 参考文档

[重磅揭秘：Smol异步运行时，比Tokio还要轻量的Rust并发利器！](https://mp.weixin.qq.com/s/1cjlKBr4EZaNSv0g6z8juQ)

[rust官方三方库](https://crates.io/)

[smol库](https://crates.io/crates/smol)

[smol文档](https://docs.rs/smol/2.0.2/smol/)