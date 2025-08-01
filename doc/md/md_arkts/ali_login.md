## 管理类PhoneNumberAuthHelper

### 创建实例

```typescript
public static getInstance(context: Context, listener: TokenResultListener): PhoneNumberAuthHelper;
```

### 配置授权页面

```typescript
public  setAuthConfig(config: AuthUiConfig): void;
```

### 设置密钥

```
public  setAuthConfig(config: AuthUiConfig): void;
```

### 获取token

```
//totalTime 超时时间单位ms
public getLoginToken(totalTime: number): Promise<void>;
```

```
//
this.helper.getLoginToken(5000)
```

### 清除监听

```
public clearAuthListener():void
```

### 退出授权页面

```
public quitLoginPage(): Promise<void>;
```

### 设置监听

```
public setAuthListener(listener: TokenResultListener): void
```

### 检测环境

```typescript
//number--SDK功能类型。取值：1，一键登录；2，是本机号码校验
public checkEnvAvailable(type: number): void
```

### 一键登录加速接口

```typescript

public accelerateLoginPage(
  overdueTimeMills: number, //超时时间，单位：ms
  listener: PreLoginResultListener
): void
```

### 点击事件监听接口

```typescript
public setUIClickListener(
  listener: AuthUIControlClickListener
): void
```

## 接口AuthUIControlClickListener

## 接口PreLoginResultListener

## 接口TokenResultListener

````typescript
class TokenListener implements TokenResultListener {
    private page: Index
    constructor(page: Index) {
        this.page = page
    }
    onSuccess(msg: string): void {
        console.log("auth:onSuccess:" + msg)
        const result: object = JSON.parse(msg)
        const code = result['_code'] + ''
        const token = result['_token'] + ''
        if (code === "600000") {
            this.page.quitLoginPage()
        }
    }
    onFailure(ret: string): void {
        console.log("auth:onFailure:" + ret)
        this.page.updateAuth()
    }
}
````



## 类AuthUiConfig



## 获取AppInfo

### appId-签名-appIdentifier

```typescript
  let flag = bundleManager.BundleFlag.GET_BUNDLE_INFO_WITH_SIGNATURE_INFO;
  let bundleInfo = bundleManager.getBundleInfoForSelfSync(flag)
  let appId = bundleInfo.signatureInfo.appId;
  console.log("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");
  console.log('volvo--',appId);
  console.log("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");
  console.log('volvo--',bundleInfo.signatureInfo.appIdentifier);
 console.log('volvo--',bundleInfo.signatureInfo.fingerprint);
```

## 权限配置

```typescript
    "requestPermissions": [
      {
        "name": 'ohos.permission.INTERNET',
        "name": 'ohos.permission.GET_NETWORK_INFO',
        "name":"ohos.permission.SET_NETWORK_INFO"
      }
    ]
```

| **权限**         | **说明**                                       |
| ---------------- | ---------------------------------------------- |
| INTERNET         | 允许应用程序联网，用于访问网关和认证服务器     |
| GET_NETWORK_INFO | 获取网络状态，判断是否数据、wifi等             |
| SET_NETWORK_INFO | 允许应用配置数据网络，取号时需要切换到蜂窝网络 |

## 鸿蒙端错误码

| **返回码** | **返回码描述**           |
| ---------- | ------------------------ |
| 600000     | 成功                     |
| 600001     | 拉起授权页成功           |
| 600002     | 拉起授权页失败           |
| 600005     | 设备终端不安全           |
| 600007     | 未检测到sim卡            |
| 600008     | 蜂窝网络未开启           |
| 600009     | 无法判断运营商           |
| 600010     | 未知异常                 |
| 600011     | 获取token失败            |
| 600012     | 获取token成功            |
| 600015     | 请求超时                 |
| 600017     | secret解析失败           |
| 600024     | 设备终端支持认证         |
| 600026     | 授权页已存在不可再次拉起 |
| 700000     | 用户取消授权页           |
| 700002     | 用户点击授权页登录按钮   |
| 700003     | 用户点击授权页协议复选框 |
| 700008     | 用户点击二次弹窗确认按钮 |
| 700009     | 用户点击二次弹窗取消按钮 |

## 一键登录流程图

https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/4636531861/p629216.png

## 参考文档

[一键登录鸿蒙端](https://help.aliyun.com/zh/pnvs/developer-reference/harmony-client-access?spm=a2c4g.11186623.0.0.3217f1cbDRoWEJ)

