### npm

#### 临时忽略SSL证书验证

```javascript
npm config set strict-ssl false
npm config set strict-ssl true
```

#### 更新

```
npm install -g npm@latest
//
```

#### 更换npm镜像源

```
npm config set registry https://registry.npmjs.org/
npm config set registry https://mirrors.huaweicloud.com/repository/npm/
```

#### 清除缓存

```
npm cache clean --force
```

#### 安装卸载

```javascript
//此命令用于安装 npm 包和特定包所依赖的其他包。它将安装在本地node_modules
//文件夹中。
npm install <packagename>
npm i <packagename>
//  
npm uninstall <packagename>
npm un <packagename>  
```

#### 更新

```
npm update <packagename> 
npm update
npm up <packagename>
```

#### 弃用

```
//此命令将通过向所有尝试安装它的人提供弃用警告或消息来更新包的 npm 注册表项
npm deprecate <pkg>[@<version range>] <message>
//==注意==：要取消弃用特定包，请为消息参数指定一个空字符串 ("")。请注意，您
//必须使用双引号，它们之间不能有空格。
npm deprecate <pkg>[@<version range>] ""
```

### 安装Express

```
npm i -g express-generator@4
```



### 参考文档

[作为开发人员你需要知道的 npm 命令](https://juejin.cn/post/6999492660314505246?searchId=202407012026384FBB3B64600A8B1350E0)

[Sequelize官方](https://www.sequelize.cn/)