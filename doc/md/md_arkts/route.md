## 插件HMRouter

### 依赖导包

```typescript
/*
./entry/oh-package.json5
ohpm install @hadss/hmrouter
*/
{
  "name": "entry",
  "version": "1.0.0",
  "description": "Please describe the basic information.",
  "main": "",
  "author": "",
  "license": "",
  "dependencies": {
    "@hadss/hmrouter":"^1.0.0-rc.5"
  }
}
```

```typescript
/*
编译插件
./hvigor/hvigor-config.json5
*/
{
  "modelVersion": "5.0.0",
  "dependencies": {
    "@hadss/hmrouter-plugin": "^1.0.0-rc.10"
  },
    ...
}   
```

### 修改编译配置

```typescript
/*
./entry/hvigorfile.ts
*/
import { hapTasks } from '@ohos/hvigor-ohos-plugin';
import {hapPlugin} from '@hadss/hmrouter-plugin';

export default {
    system: hapTasks,  /* Built-in plugin of Hvigor. It cannot be modified. */
    plugins:[hapPlugin()]         /* Custom plugin to extend the functionality of Hvigor. */
}
```

### 初始化

```typescript
/*
在入口EntryAbility类的onCreate回调中添加初始化HMRouter
*/
export default class EntryAbility extends UIAbility {
  onCreate(want: Want, launchParam: AbilityConstant.LaunchParam): void {
    HMRouterMgr.init({
      context:this.context
    })
  }
  ...
}  
```

### 根节点入口

```typescript
/*
必需用Column或Stack包裹
*/
import { HMDefaultGlobalAnimator, HMNavigation, HMRouterMgr } from '@hadss/hmrouter';
import { AttributeUpdater } from '@kit.ArkUI';
import { Routes } from '../router/routes';
@Entry
@Component
@Preview
struct Index {
  modifier: NavModifier = new NavModifier()

  build() {
    Column() {
      HMNavigation({
        navigationId: 'mainNavigation',
        options: {
          standardAnimator: HMDefaultGlobalAnimator.STANDARD_ANIMATOR,
          dialogAnimator: HMDefaultGlobalAnimator.DIALOG_ANIMATOR,
          modifier: this.modifier
        }
      }) {
        Row() {
          Text('点击跳转')
            .fontSize(30)
            .onClick(() => {
              HMRouterMgr.push({ pageUrl: Routes.pageB })
            })
        }
        .width('100%')
        .height('100%')
        .backgroundColor(Color.Yellow)
        .justifyContent(FlexAlign.Center)
      }
    }
    .height('100%')
    .width('100%')
  }
}
//
class NavModifier extends AttributeUpdater<NavigationAttribute> {
  initializeModifier(instance: NavigationAttribute): void {
    instance.mode(NavigationMode.Stack)
    instance.navBarWidth('100%')
    // instance.hideNavBar(true)
    instance.hideTitleBar(false)
    instance.hideToolBar(false)
    instance.title('Navigator')
  }
}
```

### 普通页面

```typescript
import { HMRouter, HMRouterMgr } from '@hadss/hmrouter'
import { Routes } from '../router/routes'

@HMRouter({ pageUrl: Routes.pageB })
@Component
export struct PageB {
  build() {
    Row() {
      Text('点击返回')
        .fontSize(50)
        .onClick(() => {
          HMRouterMgr.pop()
        })
    }
    .width('100%')
    .height('100%')
    .backgroundColor(Color.Red)
    .justifyContent(FlexAlign.Center)
  }
}
```

### HMRouter 硬编码问题

#### 编译插件

````typescript
//在 hvigorfile.ts修改
import { appTasks } from '@ohos/hvigor-ohos-plugin';
import { FileUtil, hvigor, getNode, HvigorNode, HvigorTask, HvigorPlugin } from '@ohos/hvigor';

// 实现自定义插件，sync的时候自动生成路由类
export function RouterNamesGeneratePlugin(): HvigorPlugin {
    // 路由表JSON文件地址
    let namesFilePath = "config/router_names.json5"
    // 路由生成规则配置文件地址
    let configFilePath = "config/router_config.json5"
    // 路由类文件名称
    let fileName = "RouteUrl"

    return {
        pluginId: 'RouterNamesGenerate',
        apply(node: HvigorNode) {
            // 路由数据源
            let routerNames = FileUtil.readJson5(namesFilePath)
            let routerList = routerNames['rent'] // 路由对象列表
            if (!routerList) {
                console.log(`debug_hvigorfile: routerList is null`)
                return
            }

            // 路由配置表
            let routerConfig = FileUtil.readJson5(configFilePath)
            let moduleNames = routerConfig['moduleNames'] // 路由对象列表
            if (!moduleNames) {
                console.log(`debug_hvigorfile: moduleNames 配置的module为空，不生成路由类 `)
                return
            }
            console.log(`debug_hvigorfile: 路由配置表 ↓↓↓↓↓↓↓↓↓↓↓↓`)
            //moduleNames.forEach(element => {
            moduleNames.forEach((element: any) => {
                console.log(`debug_hvigorfile: moduleName=${element}`)
            })
            console.log(`debug_hvigorfile: 路由配置表 ↑↑↑↑↑↑↑↑↑↑↑↑`)

            console.log(`debug_hvigorfile: 路由表原始数据 ↓↓↓↓↓↓↓↓↓↓↓↓`)
            routerList.forEach((element: { [x: string]: any; }) => {
                console.log(`debug_hvigorfile: key=${element["key"]}, value=${element["value"]}`)
            })
            console.log(`debug_hvigorfile: 路由表原始数据 ↑↑↑↑↑↑↑↑↑↑↑↑`)

            // 获取所有节点
            const allNodes = hvigor.getAllNodes();
            for (let i = 0; i < allNodes.length; i++) {
                let nodeName = allNodes[i].getNodeName()
                let nodeDir = allNodes[i].getNodeDir()
                let getPath = nodeDir.getPath()

                // console.log(`debug_hvigorfile: module: nodeName=${nodeName}, getPath=${getPath}`)
                if (moduleNames.includes(nodeName)) {
                    let newFilePath = getPath + `/src/main/ets/routes/${fileName}.ets`
                    // 判断是否存在
                    FileUtil.ensureFileSync(newFilePath)
                    let exist2 = FileUtil.exist(newFilePath)

                    let date = new Date()
                    let formatDate = date.toLocaleString()
                    let tips = `// 自动生成,请勿修改 \n// 生成时间：${formatDate}`

                    let a = `${tips} \nexport class ${fileName} {\n\t`
                    let length = routerList.length
                    for (let j = 0; j < length; j++) {
                        let element = routerList[j]
                        // console.log(`debug_hvigorfile: 开始写入文件>>>>> key=${element["key"]}, value=${element["value"]}`)
                        let key = element["key"]
                        let value = element["value"]
                        let prefix = (j == length - 1) ? "\n" : "\n\t"
                        a += `public static readonly ${key}: string = "${value}"${prefix}`
                    }
                    a += `}`
                    // 备注：这个 写入文件，不要使用异步写入，异步写入的话会存在先执行 HMRouter 的插件的case，就容易出问题了，需要保证在执行 HMRouter 插件之前，生成完所有的路由映射类
                    FileUtil.writeFileSync(newFilePath, a)
                    console.log(`debug_hvigorfile: moduleName[${nodeName}]: 写入完成`)

                }
            }

            // FileUtil.readFile('entry/test.txt').then(function(data) {
            //   console.log('readFile: ' + data);
            // });
        }
    }
}

export default {
    system: appTasks,  /* Built-in plugin of Hvigor. It cannot be modified. */
    plugins:[RouterNamesGeneratePlugin()]         /* Custom plugin to extend the functionality of Hvigor. */
}

````

#### 编译插件配置

```typescript
//在工程级（根目录）在新建文件夹 config
//在 config 下添加 router_config.json5
{
  // 配置：需要生成路由类的module名字，需要在这配置一下
  "moduleNames": [
    "entry",
    "hary"
  ]
}
```

```typescript
//在工程级（根目录）在新建文件夹 config
//在 config 下添加 router_names.json5
{
  "rent": [
    {
      "key": "onePage",
      "value": "entry/onePage"
    },
    {
      "key": "twoPage",
      "value": "entry/twoPage"
    },
    //hary
    {
      "key": "mainPage",
      "value": "hary/mainPage"
    }
  ]
}

```



### 路由配置

```typescript
/*
简单使用可以不用配置
hmrouter_config.json
*/
{
  "scanDir": [],
  "builderDir": "src/main/ets/generated",
  "saveGeneratedFile": true
}
```

### 路由拦截

#### 定义拦截器

```typescript
import { HMInterceptor, HMInterceptorAction, HMInterceptorInfo, HMRouterMgr, IHMInterceptor } from "@hadss/hmrouter";
import { Routes } from "./routes";

@HMInterceptor({ interceptorName: 'LoginCheckInterceptor' })
export class LoginCheckInterceptor implements IHMInterceptor {
  handle(info: HMInterceptorInfo): HMInterceptorAction {
    if ((info.srcName == Routes.pageB) && false) {
      return HMInterceptorAction.DO_NEXT
    } else {
      info.context.getPromptAction().showToast({ message: '请先登录' })
      HMRouterMgr.push({ pageUrl: Routes.pageB, skipAllInterceptor: true })
      return HMInterceptorAction.DO_REJECT
    }
  }
}
```

#### 拦截器使用-1

```typescript
import { HMRouter, HMRouterMgr } from '@hadss/hmrouter'
import { Routes } from '../router/routes'
@HMRouter({ pageUrl: Routes.pageB, interceptors: ['LoginCheckInterceptor'] })
@Component
export struct PageB {
  build() {
    Row() {
      Text('点击返回--shou')
        .fontSize(50)
        .onClick(() => {
          HMRouterMgr.pop()
        })
    }
    .width('100%')
    .height('100%')
    .backgroundColor(Color.Red)
    .justifyContent(FlexAlign.Center)
  }
}
```

#### 拦截器使用二

```typescript
HMRouterMgr.registerGlobalInterceptor({
  interceptor: new JumpInfoInterceptor(),
  interceptorName: 'JumpInfo',
  priority: 5
});
```

#### 拦截器使用-3

```typescript
@HMInterceptor({interceptorName: 'LoginStatusInterceptor', global: true})
export class LoginStatusInterceptor implements IHMInterceptor {
  handle(info: HMInterceptorInfo): HMInterceptorAction {
    console.log(`Login status is ${!!AppStorage.get('isLogin') ? 'Y' : 'N'}`);
    return HMInterceptorAction.DO_NEXT;
  }
}
```

#### 拦截页面跳转类型

```typescript
export class JumpInfoInterceptor implements IHMInterceptor {
  handle(info: HMInterceptorInfo): HMInterceptorAction {
    let connectionInfo: string = '';
    if(info.type === HMActionType.PUSH) {
      connectionInfo = 'jump to';
    } else {
      connectionInfo = 'back to';
    }

    console.log(`${info.srcName} ${connectionInfo} ${info.targetName}`);
    return HMInterceptorAction.DO_NEXT;
  }
}

```

```typescript
export declare class HMActionType {
    static POP: string;
    static PUSH: string;
    static REPLACE: string;
}
```

### 生命周期

#### 声明HMLifecycle

```typescript
@HMLifecycle({lifecycleName: 'ExitAppLifecycle'})
export class ExitAppLifecycle implements IHMLifecycle {
  private lastTime: number = 0;
  onBackPressed(ctx: HMLifecycleContext): boolean {
    let time = new Date().getTime();
    if(time - this.lastTime > 1000) {
      this.lastTime = time;
      ctx.uiContext.getPromptAction().showToast({
        message: CommonConstants.EXIT_TOAST,
        duration: 1000
      });
      return true;
    } else {
      return false;
    }
  }
}

```

#### 使用Lifecycle

```typescript
import { HMLifecycle, HMLifecycleContext, IHMLifecycle } from "@hadss/hmrouter";

@HMLifecycle({lifecycleName: 'ExitAppLifecycle'})
export class ExitAppLifecycle implements IHMLifecycle {
  private lastTime: number = 0;
  onBackPressed(ctx: HMLifecycleContext): boolean {
    let time = new Date().getTime();
    if(time - this.lastTime > 1000) {
      this.lastTime = time;
      ctx.uiContext.getPromptAction().showToast({
        message: '是否退出',
        duration: 1000
      });
      return false;
    } else {
      return true;
    }
  }
}
```

#### 页面停留时长

```typescript
import { HMLifecycle, HMLifecycleContext, IHMLifecycle } from "@hadss/hmrouter";

@HMLifecycle({ lifecycleName: 'PageDurationLifecycle' })
export class PageDurationLifecycle implements IHMLifecycle {
  private timeMap: Map<string, number> = new Map()

  onShown(ctx: HMLifecycleContext): void {
    const pageName = ctx.navContext?.pathInfo.name
    if (!pageName) {
      return
    }
    this.timeMap.set(pageName, new Date().getTime());
  }

  onDisAppear(ctx: HMLifecycleContext): void {
    const pageName = ctx.navContext?.pathInfo.name
    if (pageName && this.timeMap.has(pageName)) {
      const duration = new Date().getTime() - (this.timeMap.get(pageName) as number)
      this.timeMap.delete(pageName)
      ctx.uiContext.getPromptAction().showToast({
        message: `Page ${pageName} 停留 ${duration} 毫秒`,
        duration: 1000
      });
    }
  }
}
```

### 路由动画

```typescript
import { HMAnimator, HMAnimatorHandle, IHMAnimator } from "@hadss/hmrouter";

@HMAnimator({ animatorName: 'LiveCommentsAnimator' })
export class LiveCommentsAnimator implements IHMAnimator {
  effect(enterHandle: HMAnimatorHandle, exitHandle: HMAnimatorHandle): void {
    //入场动画
    enterHandle.start((translateOption, scaleOption) => {
      translateOption.y = '100%'
      scaleOption.x = 0
    })
    enterHandle.finish((translateOption, scaleOption) => {
      translateOption.y = '0'
      scaleOption.x = 1
    })
    enterHandle.duration = 800
    //出场动画
    exitHandle.start((translateOption, scaleOption,
      opacityOption) => {
      translateOption.y = '0'
    })
    exitHandle.finish((translateOption, scaleOption,
      opacityOption) => {
      translateOption.y = '100%'
    })
    exitHandle.duration = 500
  }

  interactive?(handle: HMAnimatorHandle): void {
  }
}
```

```typescript
import { HMRouter, HMRouterMgr, HMService } from '@hadss/hmrouter'
import { Routes } from '../router/routes'

@HMRouter({
  pageUrl: Routes.pageB,
  interceptors: ['LoginCheckInterceptor'],
  lifecycle: 'PageDurationLifecycle',
  animator: 'LiveCommentsAnimator'
})
@Component
export struct PageB {
  @HMService({ serviceName: 'testConsole' })
  info() {
    console.log('调用HMService');
  }

  build() {
    Row() {
      Text('点击返回--shou')
        .fontSize(50)
        .onClick(() => {
          this.info()
          HMRouterMgr.pop()
        })
    }
    .width('100%')
    .height('100%')
    .backgroundColor(Color.Red)
    .justifyContent(FlexAlign.Center)
  }
}
```



### 服务路由

```typescript
import { HMService } from "@hadss/hmrouter";

export class CustomService {
  @HMService({ serviceName: 'testConsole' })
  testConsole(): void {
    console.log('调用HMService-方法装饰器');
  }

  @HMService({ serviceName: 'testFunWithReturn' })
  testFunWithReturn(param1: string, param2: string): string {
    return `调用服务 testFunWithReturn:${param1} ${param2}`
  }

  @HMService({ serviceName: 'testAsyncFun', singleton: true })
  async asyncFunction(): Promise<string> {
    return new Promise((resolve) => {
      resolve('调用异步服务 testAsyncFun')
    })
  }
}
```



### 类IHMLifecycle

```typescript
export interface IHMLifecycle {
    onPrepare?(ctx: HMLifecycleContext): void;
    onAppear?(ctx: HMLifecycleContext): void;
    onDisAppear?(ctx: HMLifecycleContext): void;
    onShown?(ctx: HMLifecycleContext): void;
    onHidden?(ctx: HMLifecycleContext): void;
    onWillAppear?(ctx: HMLifecycleContext): void;
    onWillDisappear?(ctx: HMLifecycleContext): void;
    onWillShow?(ctx: HMLifecycleContext): void;
    onWillHide?(ctx: HMLifecycleContext): void;
    onReady?(ctx: HMLifecycleContext): void;
    onBackPressed?(ctx: HMLifecycleContext): boolean;
}

```



### 类HMNavigation

```typescript
export declare struct HMNavigation {
    @Require
    navigationId: string;
    @State
    navigationEnable: boolean;
    homePageUrl?: string;
    options?: HMNavigationOption;
    hmRouterStore: IHMRouterStore;
    @BuilderParam
    closer: () => void;
    private hmRouterMgrService;
    private pageStack?;
    private customTransition?;
    private hideNavBar?;
    private navigationKey?;
    private navbarLifecycleMgr;
    private pageBuilder?;
    @Builder
    closerBuilder(): void;
    getPageBuilder(name: string): boolean;
    @Builder
    dynamicPageBuilder(name: string, arg: ESObject): void;
    aboutToAppear(): void;
    onDidBuild(): void;
    aboutToDisappear(): void;
    build(): void;
}
```



### 类HMRouterMgr

#### 类定义

```typescript
import { HMInterceptorInstance, HMLifecycleInstance, HMPageInstance } from '../store/ComponentInstance';
import { HMRouterPathCallback } from './HMRouterPathCallback';
import { HMRouterPathInfo } from './HMRouterPathInfo';
import { IHMAnimator } from './IHMAnimator';
import { HMRouterConfig } from './HMRouterConfig'
import { HMPageLifecycle } from '../lifecycle/HMPageLifecycle';
import { IHMLifecycleOwner } from '../lifecycle/interface/IHMLifecycleOwner';
import { HMServiceResp } from './IHMService';
export declare class HMRouterMgr {
...
}
```

#### 静态属性

```typescript
static isInit: boolean;
```

#### 私有属性

```typescript
private static service;
private static routerStore;
private static pageLifecycleMgr;
```

#### 静态方法

```typescript
static init(config: HMRouterConfig): void;
static openLog(level: 'DEBUG' | 'INFO'): void;
static getPathStack(navigationId: string): NavPathStack | null;
static getCurrentParam(): Object | null;
static getCurrentLifecycleOwner(): IHMLifecycleOwner | null;

```

```typescript
static push(
  pathInfo: HMRouterPathInfo, 
  callback?: HMRouterPathCallback
): void;
```

```typescript
static replace(
  pathInfo: HMRouterPathInfo, 
  callback?: HMRouterPathCallback
): void;
```

```typescript
static pop(
  pathInfo?: HMRouterPathInfo, 
  skipedLayerNumber?: number
): void;
```

```typescript
static registerGlobalInterceptor(
  interceptorInstance: HMInterceptorInstance
): void;
static unRegisterGlobalInterceptor(
  interceptorName: string): boolean;
```

```typescript
static registerGlobalLifecycle(
	lifecycleInstance: HMLifecycleInstance): void;
static unRegisterGlobalLifecycle(
  lifecycleName: string): boolean;
static generatePageLifecycleId(): string;
static getPageLifecycleById(
  pageLifecycleId: string
): HMPageLifecycle | undefined;
```

```typescript
static registerGlobalAnimator(
  navigationId: string, 
  key: 'standard' | 'dialog', 
  animator: IHMAnimator
): void;
static unRegisterGlobalAnimator(
  navigationId: string, 
  key: 'standard' | 'dialog'
): boolean;
```

```typescript
static registerPageBuilder(
  pageInstance: HMPageInstance
): boolean;
```

```typescript
static request(
  serviceName: string, ...args: Object[]
): HMServiceResp;
```

### 接口HMRouterConfig

```typescript
import { common } from '@kit.AbilityKit';
export interface HMRouterConfig {
    context: common.Context;
    initWithTaskPool?: boolean;
}
```

### 接口HMRouterPathInfo

```typescript
import { IHMAnimator } from './IHMAnimator';
import { IHMInterceptor } from './IHMInterceptor';
/**
 * 路由跳转/返回参数
 */
export interface HMRouterPathInfo {
    navigationId?: string;
    pageUrl?: string;
    param?: ESObject;
    interceptors?: IHMInterceptor[];
    animator?: IHMAnimator | boolean;
    skipAllInterceptor?: boolean;
}
```

### 接口HMRouterPathCallback

```typescript
export interface HMRouterPathCallback {
    /**
     * 页面返回回调
     * @param popInfo
     */
    onResult?: (popInfo: HMPopInfo) => void;
    /**
     * 目标页面跳转完成回调
     */
    onArrival?: () => void;
    /**
     * 目标页面找不到回调
     */
    onLost?: () => void;
}
```

### 接口HMPopInfo

```typescript
interface HMNavPathInfo {
    name: string;
    param?: Object;
}
/**
 * 路由返回回调参数类型
 */
export interface HMPopInfo extends PopInfo {
    srcPageInfo: HMNavPathInfo;
}
```

### 接口IHMInterceptor

```typescript
import { HMInterceptorAction } from '../router/HMInterceptorAction';
import { HMInterceptorInfo } from '../router/HMInterceptorInfo';
export interface IHMInterceptor {
    handle(info: HMInterceptorInfo): HMInterceptorAction;
}
```

### 接口IHMAnimator

```typescript
import { HMAnimatorHandle } from '../animator/HMAnimatorHandle';
export interface IHMAnimator {
    effect(enterHandle: HMAnimatorHandle, exitHandle: HMAnimatorHandle): void;
    interactive?(handle: HMAnimatorHandle): void;
}
export declare namespace IHMAnimator {
    interface EffectOptions {
        direction?: IHMAnimator.Direction;
        opacity?: IHMAnimator.OpacityOption;
        scale?: IHMAnimator.ScaleOption;
    }
    interface TranslateOption {
        x?: number | string;
        y?: number | string;
        z?: number | string;
    }
    interface ScaleOption {
        x?: number;
        y?: number;
        centerX?: number | string;
        centerY?: number | string;
    }
    interface OpacityOption {
        opacity?: number;
    }
    enum Direction {
        RIGHT_TO_LEFT = 0,
        LEFT_TO_RIGHT = 1,
        BOTTOM_TO_TOP = 2,
        TOP_TO_BOTTOM = 3
    }
    class Effect {
        effectOption: IHMAnimator.EffectOptions;
        constructor(effectOption: IHMAnimator.EffectOptions);
        toAnimator(): IHMAnimator;
    }
}
```

### 接口IHMLifecycleOwner

```typescript
import { HMLifecycleCallback, IHMLifecycle } from '../../api/IHMLifecycle';
import { HMLifecycleState } from '../../api/HMLifecycleState';
export interface IHMLifecycleOwner {
    addObserver(state: HMLifecycleState, callback: HMLifecycleCallback, priority?: number): void;
    getLifecycle(): IHMLifecycle | undefined;
}
```

### 枚举HMLifecycleState

```typescript
export declare enum HMLifecycleState {
    onDisAppear = "onDisAppear",
    onShown = "onShown",
    onHidden = "onHidden",
    onWillDisappear = "onWillDisappear",
    onWillShow = "onWillShow",
    onWillHide = "onWillHide",
    onBackPressed = "onBackPressed"
}
export declare enum InnerLifecycleState {
    onPrepare = "onPrepare",
    onAppear = "onAppear",
    onWillAppear = "onWillAppear",
    onReady = "onReady"
}
export type AllLifecycleState = HMLifecycleState | InnerLifecycleState;
```

## FWRouter

### 文档资料

[Navigation页面管理-鸿蒙@fw/router框架源码解析（二）](https://juejin.cn/post/7404005062177456165)



## 参考文档

[组件导航 (Navigation)(推荐)](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/arkts-navigation-navigation-V5)

[导航与切换](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/4_5_u5bfc_u822a_u4e0e_u5207_u6362-V5)

[Harmony HMRouter Interceptor 的基本使用](https://juejin.cn/post/7439275048965554187?from=search-suggest)

[Harmony HMRouter HMLifecycle 基本使用](https://juejin.cn/post/7439258070193782825)

[[HarmonyOS] 鸿蒙开发解决HMRouter路由地址无法抽取的问题](https://juejin.cn/post/7440469508991320075)

[喜大普奔！鸿蒙官方开源发布了路由管理组件HMRouter](https://juejin.cn/post/7411187555776315427?searchId=2024120515143428BBB6B048435918A50D)

[鸿蒙应用开发-初见：Hvigor](https://juejin.cn/post/7307723756411027483)

[《京东金融APP的鸿蒙之旅系列专题》鸿蒙工程化：Hvigor构建技术](https://juejin.cn/post/7425193631583354932)

[鸿蒙应用开发从入门到入魔：Navigation路由管理为什么这么麻烦？](https://juejin.cn/post/7402204074714267685?searchId=20241215190150665CFAA7F698BBD26C8F)