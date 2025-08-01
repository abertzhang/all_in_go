## 插件flutter_rust_bridge

### 环境配置

#### 依赖导包

```

```

#### 安装桥接工具

```dart
cargo install flutter_rust_bridge_codegen
```

#### 生成项目

```dart
//注意dart版本
flutter_rust_bridge_codegen create project_name
```

#### 构建产物

```
flutter build apk --target-platform android-arm64
//
flutter build windows
```

#### 重生桥接代码

```
flutter_rust_bridge_codegen generate
```

#### 报错信息

```
//NDK:23.1.7779620
ProcessException: No such file or directory
SEVERE:   Command: /Users/zhangchunhua/sdk/android_sdk/cmdline-tools/latest/bin/sdkmanager --install ndk;23.1.7779620
```

### 已有项目添加插件

#### 最快捷的方式

```
//项目根目录下执行命令
flutter_rust_bridge_codegen integrate
```



#### 添加rust_builder

```
//拷贝到项目根目录下,内含cargokit
```

#### 配置yaml

```
# rust_lib_项目名,查看生成名称
rust_lib_demo_rust:
  path: rust_builder
flutter_rust_bridge: 2.5.1
```

#### 添加rust项目文件夹

```
/*
在根目录下将rust项目的拷贝过来,文件名与配置一致
flutter_rust_bridge.yaml:
rust_root: rust/
*/
```

#### 生成dart桥接代码

```
flutter_rust_bridge_codegen generate
```



#### 新建生成器配置

```
/*
在项目根目录下新建flutter_rust_bridge.yaml
注意 dart_output 文件夹一定要存在
*/
rust_input: crate::api
rust_root: rust/
dart_output: lib/src/rust
```

## 插件项目使用Rust

### 脚本生成plugin项目

```
//安装脚本工具
cargo install frb_plugin_tool
//生成plugin,执行命令,如果已存在不影响
frb_plugin_tool
```

### 增加其他平台

```
flutter create -t plugin_ffi --platforms <platforms>
```

### 更新dart桥接代码

```
flutter_rust_bridge_codegen generate
```

### Cargokit

````
//subtreee形式添加
git subtree add --prefix cargokit https://github.com/irondash/cargokit.git main --squash
````



## 参考文档

[Flutter&Rust#01 | 突破能力瓶颈](https://juejin.cn/post/7411812014047166475?searchId=20241030103727EF8EF641B195DFCFA5FD)

[插件flutter_rust_bridge](https://pub-web.flutter-io.cn/packages/flutter_rust_bridge)

[Rust-crate:**flutter_rust_bridge**](https://crates.io/crates/flutter_rust_bridge)

[使用flutter_rust_bridge为dart构建第三方包](https://juejin.cn/post/7349182232441274380)

[用rust写个flutter插件并上传 pub.dev](https://juejin.cn/post/7412486655862734874)

[Flutter & Rust的一些探索与实践](https://juejin.cn/post/7168735276812992543?searchId=20241030103727EF8EF641B195DFCFA5FD)