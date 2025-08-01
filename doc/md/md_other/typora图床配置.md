

分类

其他

### Typora命令式图床设置

#### Typora设置



<img src="http://qiniu-article.myflutter.cn/img/image-20230319173359284.png" alt="image-20230319173359284" style="zoom: 67% title=图片;" />

#### picgo工具安装

##### 1.安装Node.js

在Node官方网站https://nodejs.org/zh-cn下载安装到电脑,

##### 2.安装picgo

```
//安装node后自带npm
npm install picgo
//安装相关插件
picgo install rename-file
//打开配置文件
vim ~/.picgo/config.json
//七牛云的配置参数放入config.json
{
    "picBed": {
      "uploader": "qiniu",
      "qiniu": {
        "accessKey": "",//在七牛控制台点击用户头像弹出选择密钥管理
        "secretKey": "",
        "bucket": "", // 存储空间名
        "url": "", // 自定义域名
        "area":  "", // 存储区域编号
        "options": "", // 网址后缀，比如？s1g.cn
        "path": "img/" // 在存储空间名下的自定义存储路径，比如 img/
      }
    },
    "picgoPlugins": {}
}

```

##### .其他配置参数

```
//github
{
  "picBed": {
    "uploader": "githubPlus",
    "current": "githubPlus",
    "githubPlus": {
      "branch": "main",
      "customUrl": "https://cdn.jsdelivr.net/gh/your_github_id/your_repo@main", //比原声raw.github.xxxx要快
      "origin": "github",
      "repo": "your_github_id/your_repo", // 替换仓库
      "path": "",// 存放图片的仓库目录下的文件夹
      "token": "xxxx"// 访问github的仓库的token
    },
    "picgoPlugins": {
      "picgo-plugin-github-plus": true,
      "picgo-plugin-rename-file": false
    },
    "picgo-plugin-rename-file": {
      "format": "{y}/{m}/{d}/{hash}-{origin}-{rand:6}"
    }
}
```

