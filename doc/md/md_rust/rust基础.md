## Cargo

### workspace

#### 常用命令

```rust
//在空文件夹内-待验证
//cargo workspace init
//添加项目
//cargo worksapce add <项目名> <路径>
//例子
cargo workspace add main ./main
//运行某个项目
cargo run --package 项目名
//构建某个项目
cargo build --package 项目名
```

#### 配置toml

```rust
假设你有两个项目：frontend和backend。frontend项目需要使用backend项目提供的API
[workspace]
members = [ "frontend", "backend" ]
[dependencies]
backend = { path = "backend" }

```

### 文档资料

[Rust：最全cargo 命令（建议收藏）]()

## Vec

### 源码

```rust
pub struct Vec<T, A: Allocator = Global> {
    buf: RawVec<T, A>,
    len: usize,
}
impl<T> Vec<T> {
    #[inline]
    pub const fn new() -> Self {
        Vec { buf: RawVec::NEW, len: 0 }
    }
//...略...
}

```



### 生成Vec

```
```



## 参考资料

[官方地址](https://www.rust-lang.org/)

[Rust开发者必看！深入解析Vec数据结构，提升你的编程效率](https://mp.weixin.qq.com/s/ilOf97xZp_Ym8nsaxkF10A)

[vec中文文档](https://www.rustwiki.org.cn/zh-CN/std/vec/index.html)

[Rust 动态数组Vec基本概念及其用法](https://juejin.cn/post/7261245655129129021?searchId=2024102608152176F97ADB239B5DCA75ED)

[Rust-- Vec, Array and Slice](https://juejin.cn/post/7150905997811253279?searchId=2024102608152176F97ADB239B5DCA75ED)

[【Rust学习之旅】动态数组 Vector 、String、 HashMap（八）](https://juejin.cn/post/7219345788743352375?from=search-suggest)

[Rust（官方文档重点总结）](https://blog.csdn.net/shsjssnn/category_12572899.html)

[掌握Rust与Cargo Workspaces：高效管理多个项目](https://blog.csdn.net/silenceallat/article/details/138415765)

[学Rust不学Cargo，等于没学Rust：workspace详解](https://blog.csdn.net/weixin_37561180/article/details/135313939?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7Ebaidujs_baidulandingword%7ECtr-1-135313939-blog-122254087.235%5Ev43%5Epc_blog_bottom_relevance_base7&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7Ebaidujs_baidulandingword%7ECtr-1-135313939-blog-122254087.235%5Ev43%5Epc_blog_bottom_relevance_base7&utm_relevant_index=2)

[不想再用mod.rs了？试试Rust模块新的命名约定](https://juejin.cn/post/7407277812262731795?searchId=20241031103830EC8D9174B6E9C36A8C6B)

[Rust工作空间（workspace）实践](https://juejin.cn/post/7366899841454768138?searchId=20241031103830EC8D9174B6E9C36A8C6B)

[Rust中的智能指针:Box<T> Rc<T> Arc<T> Cell<T> RefCell<T> Weak<T>](https://juejin.cn/post/7223761453433225272?searchId=2024110319165191BEAFABD55EF82CDB0F)

[Rust：详解 Copy 和 Clone](https://article.juejin.cn/post/7232665460713373755)

[带你了解 Rust 中的move, copy, clone](https://article.juejin.cn/post/7046638487851761694?from=search-suggest)
