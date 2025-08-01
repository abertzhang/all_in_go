





## 基本数据类型

### 常用类型

```typescript
//
let isDone:boolean=false;
const isShow:boolean=false;
//所有数字都是浮点数，这些浮点数的类型是number
let decLiteral: number = 6;
let hexLiteral: number = 0xf00d;
const binaryLiteral: number = 0b1010;
const octalLiteral: number = 0o744;
//string
let name: string = "DHL";
const publicName: string = 'ByteCode';
const content = `My name is ${name}`
//数组,类型后面接上 []或Array<元素类型>
let list: number[] = [1, 2, 3];
let list: Array<number> = [1, 2, 3];
/*
如果可以从传递给泛型函数的参数中推断出具体类型，ArkTS 允许省略泛型类型实参。否则，需要指定泛型类型实参，不然会报错。
*/
//枚举
export enum ResponseCode{
 SUCCESS =  1000,
 FAILED = 2000
}
```



## ArkTs和TypeScript差异

### 禁止var

### 不支持any或unknown

### const和let区别

### 不支持解构赋值

```typescript
//解构:第一个元素赋给 head，并将剩余的元素以数组的形式赋给 tail
let [head,...tail]=[1,2,3,4]
//ArkTs写法
const data:Number[]=[1,2,3,4]
const len =data.length;
let head = data[0];
let tail:Number[]=[];
for(let i=1;i<len;++i){
  tail.push(data[i]);
}
```

### 不支持参数解构的函数声明

```typescript
request({location:[lat,lon],name:String}){
  console.log()
}
testRequest(){
  this.request({location:[10,10],name:"name"})
}
//ArkTs
request(location:number[],name:String){}
testRequest(){
  this.request([10,10],"name")
}
```

### 不支持catch语句标注类型

```typescript
try{

}catch(error:unknown){
//处理异常
}
//ArkTs
try{

}catch(error){
//处理异常
}
```

### 类型别名和变量不能同名

```
let value:string
type value=number[]
//ArkTs上面写法不允许
```

### 类中仅支持一个静态块

```
class Person{
static{
let name ="tony"
}
static{
let age =10
}
}
//ArkTs不能多个静态块
```

### 不支持构造中声明类字段

```
class Person{
	constructor(private name:string){
		this.name = name;
	}
	getFullName():string{
		return this.name;
	}
}
//ArkTs
class Person{
    private name:string;
    constructor( name:string){
        this.name = name;
    }
    getFullName(){
        return this.name;
    }
}
```

### 数组仅包含可推断类型的元素

```
export const person:Array<Person>=[
	{"name":"Join"},
	{"name":"byteCode"},
]
//ArkTs
export const person:Array<Person>=[
	new Person("Join");
	new Person("byteCode");
];
```

### 禁止for in

```typescript
/*
向对象中添加新的属性或方法
从对象中删除已有的属性或方法
将任意类型的值赋值给对象属性
*/
```

## 状态管理

### 概要

![image](http://qiniu-article.myflutter.cn/img/17303920926f49d2a7cea65053f0105e~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp)

![image](http://qiniu-article.myflutter.cn/img/7f11de8ed2d14f12b27c458df0ac7217~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp)

### State

![image](http://qiniu-article.myflutter.cn/img/b838b8721f044df4823e767f434a52f9~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp)

```
@State变量装饰器只支持Object、class、string、number、boolean、enum类型，以及这些类型的数组。
不支持复杂类型（比如Date类型）
ComponentV2不支持State装饰器
```

#### 父子组件初始化和传递装饰图

![image](http://qiniu-article.myflutter.cn/img/f3a2b0459afe48448c25e7d6c2a4cb54~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp)

### Prop

### Link

### Provide

### Consume

### Observed

### ObjectLink

### LocalStorge

#### 参考文档

[官方 LocalStorge](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/arkts-localstorage-V5)

### AppStorge

#### 参考文档

[官方 AppStorge](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/arkts-appstorage-V5)

## 设计模式

### 模式汇总

```
创建型模式，共五种：
工厂方法模式、抽象工厂模式、单例模式、建造者模式、原型模式。
//
结构型模式，共七种：
适配器模式、装饰器模式、代理模式、外观模式、桥接模式、组合模式、享元模式。
//
行为型模式，共十一种：
策略模式、模板方法模式、观察者模式、迭代子模式、责任链模式、命令模式、备忘录模式、状态模式、访问者模式、中介者模式、解释器模式。
```

### 单例模式

```typescript
//Singleton
class Person {
  static instance: Person | null = null
  static getInstance() {
    if (!Person.instance) {
      Person.instance = new Person()
    }
    return Person.instance
  }
  //防止外面new Person()
  private constructor() {}
}
```



## 配置 app.json5

```typescript
{
  "app": {
    //必选，
    "bundleName": "cn.myFlutter.studyArkTs",//app包名
    "versionCode": 1000000,//版本 code
    "versionName": "1.0.0",//版本号
    "icon": "$media:app_icon",//app图标
    "label": "$string:app_name",//app名称
    //非必未全
    "vendor": "example",
    "accessible": false,
  }
}
/*
vendor：
*/
```

### 文档链接

[module.json5配置文件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/module-configuration-file-V5)

## Want

### 显式Want

在启动目标应用组件时，调用方传入的[want](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-app-ability-want-V5)参数中指定了abilityName和bundleName，称为显式Want。

```typescript
import { Want } from '@kit.AbilityKit';
let wantInfo: Want = {
  deviceId: '', // deviceId为空表示本设备
  bundleName: 'com.example.myapplication',
  abilityName: 'FuncAbility',
}
```
### 隐式Want

在启动目标应用组件时，调用方传入的[want](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-app-ability-want-V5)参数中未指定abilityName，称为隐式Want。

```typescript
import { Want } from '@kit.AbilityKit';
let wantInfo: Want = {
  // uncomment line below if wish to implicitly query only in the specific bundle.
  // bundleName: 'com.example.myapplication',
  action: 'ohos.want.action.search',
  // entities can be omitted
  entities: [ 'entity.system.browsable' ],
  uri: 'https://www.test.com:8080/query/student',
  type: 'text/plain',
};
```

### skill 简例

```typescript
{
  "abilities": [
    {
      "skills": [
        {
          "actions": [
            "ohos.want.action.home"
          ],
          "entities": [
            "entity.system.home"
          ],
          "uris": [
            {
              "scheme":"http",
              "host":"example.com",
              "port":"80",
              "path":"path",
              "type": "text/*",
              "linkFeature": "Login"
            }
          ],
          "permissions": [],
          "domainVerify": false
        }
      ]
    }
  ]
}
```



### skills标签

### 文档链接

[skills标签](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/module-configuration-file-V5#skills%E6%A0%87%E7%AD%BE)

## 异步

### emitter

#### 依赖导包

```
import { emitter } from '@kit.BasicServicesKit';
```

#### 常用方法

```
//api11+
on(event: InnerEvent, callback: Callback<EventData>): void
//api12+
on<T>(eventId: string, callback: Callback<GenericEventData<T>>): void
```

```
//单次订阅
once(event: InnerEvent, callback: Callback<EventData>): void
//API-11+
once(eventId: string, callback: Callback<EventData>): void
//api12+
once<T>(eventId: string, callback: Callback<GenericEventData<T>>): void
```

```
//API-10+
off(eventId: number, callback: Callback<EventData>): void
//API-11+
off(eventId: string, callback: Callback<EventData>): void
//API-12+
off<T>(eventId: string, callback: Callback<GenericEventData<T>>): void
```

```
//API-11+
emit(eventId: string, data?: EventData): void
//API-12+
emit<T>(eventId: string, data?: GenericEventData<T>): void
//API-11+
emit(eventId: string, options: Options, data?: EventData): void
//API-12+
emit<T>(eventId: string, options: Options, data?: GenericEventData<T>): void
```

```
```



### 类 InnerEvent

```
getListenerCount(eventId: number | string): number
```



### 类 EventData

```typescript
export interface EventData {
	data?: {
            [key: string]: any;
        };
}
```



### 类 EventPriority

| 名称      | 值   | 说明                                                |
| :-------- | :--- | :-------------------------------------------------- |
| IMMEDIATE | 0    | 表示事件被立即投递。                                |
| HIGH      | 1    | 表示事件先于LOW优先级投递。                         |
| LOW       | 2    | 表示事件优于IDLE优先级投递，事件的默认优先级是LOW。 |
| IDLE      | 3    | 表示在没有其他事件的情况下，才投递该事件。          |

### 类GenericEventData

```
export interface GenericEventData<T> {
        data?: T;
    }
```

### TaskPool

#### 依赖导包

```
import { taskpool } from '@kit.ArkTS';
```



```typescript
//在TaskPool中执行的函数
@Concurrent
function add(a: number, b: number): number {
  return a + b
}
//调用这个函数即可
async function test() {
  let task: taskPool.Task = new taskPool.Task(add, 1, 3)
  let result = await taskPool.execute(task)
  console.log('taskPool result:', result.toString())
}
```

### 公共事件CommonEvent

#### 依赖导包

```typescript
import { commonEventManager } from '@kit.BasicServicesKit'
```

```typescript
//示例
import { commonEventManager } from '@kit.BasicServicesKit'; 

let subscriber:commonEventManager.CommonEventSubscriber; 
let subscribeInfo: commonEventManager.CommonEventSubscribeInfo = { 
  events: ['usual.event.SCREEN_OFF'], // 订阅灭屏公共事件 
  priority:80 
} 
commonEventManager.createSubscriber(subscribeInfo, (err, data) => { 
  if (err) { 
    console.error(`Failed to create subscriber. Code is ${err.code}, message is ${err.message}`); 
    return; 
  } 
  console.info('Succeeded in creating subscriber1.'); 
  subscriber = data; 
  // 订阅公共事件回调 
  commonEventManager.subscribe(subscriber, (err, data) => { 
    if (err) { 
      console.error(`Failed to subscribe common event. Code is ${err.code}, message is ${err.message}`); 
      return; 
    } else { 
      console.info(`Succeeded in subscribe common event Succeeded1 `); 
    } 
  }) 
})
```

#### 公共事件

```typescript
'usual.event.SCREEN_OFF'//订阅灭屏公共事件
'COMMON_EVENT_BOOT_COMPLETED'//开机事件
```



### 文档链接

[HarmonyOS Next 浅谈 发布-订阅模式](https://juejin.cn/post/7439334253827702834)

[状态管理最佳实践](https://developer.huawei.com/consumer/cn/doc/best-practices-V5/bpta-status-management-V5)

[@ohos.events.emitter (Emitter)](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-emitter-V5#emitteron)

[如何监听系统公共事件，如熄屏、亮屏、开机等](https://developer.huawei.com/consumer/cn/doc/harmonyos-faqs-V5/faqs-notification-kit-5-V5)

## Background Tasks Kit

### 短时任务

依赖导包

```typescript
import { backgroundTaskManager } from '@kit.BackgroundTasksKit';
import { BusinessError } from '@kit.BasicServicesKit';
```

常用方法

```typescript
//申请短时任务。
requestSuspendDelay(reason: string, callback: Callback<void>): DelaySuspendInfo
//获取对应短时任务的剩余时间
getRemainingDelayTime(requestId: number): Promise<number>
//取消短时任务
cancelSuspendDelay(requestId: number): void
```

### 长时任务

依赖导包

```typescript
 import { backgroundTaskManager } from '@kit.BackgroundTasksKit';
 import { AbilityConstant, UIAbility, Want } from '@kit.AbilityKit';
 import { window } from '@kit.ArkUI';
 import { rpc } from '@kit.IPCKit'
 import { BusinessError } from '@kit.BasicServicesKit';
 import { wantAgent, WantAgent } from '@kit.AbilityKit';
```

长时任务类型

| 参数名                  | 描述                     | 配置项                | 场景举例                                             |
| :---------------------- | :----------------------- | :-------------------- | :--------------------------------------------------- |
| DATA_TRANSFER           | 数据传输                 | dataTransfer          | 后台下载大文件，如浏览器后台下载等。                 |
| AUDIO_PLAYBACK          | 音视频播放               | audioPlayback         | 音乐类应用在后台播放音乐，投播。支持在元服务中使用。 |
| AUDIO_RECORDING         | 录制                     | audioRecording        | 录音机在后台录音。                                   |
| LOCATION                | 定位导航                 | location              | 导航类应用后台导航。                                 |
| BLUETOOTH_INTERACTION   | 蓝牙相关                 | bluetoothInteraction  | 通过蓝牙传输分享的文件。                             |
| MULTI_DEVICE_CONNECTION | 多设备互联               | multiDeviceConnection | 分布式业务连接。支持在元服务中使用。                 |
| TASK_KEEPING            | 计算任务（仅对2in1开放） | taskKeeping           | 杀毒软件。                                           |

常用方法

```
//申请长时任务
startBackgroundRunning(context: Context, bgMode: BackgroundMode, wantAgent: WantAgent): Promise<void>
//取消长时任务
stopBackgroundRunning(context: Context): Promise<void>
```

权限申请

```
ohos.permission.KEEP_BACKGROUND_RUNNING
```

配置 module.json5

```typescript
 "module": {
     "abilities": [
         {
             "backgroundModes": [
              // 长时任务类型的配置项
             "audioRecording"
             ], 
             "skills": [
                 // 必填项：申请长时任务时entities和actions值
                 {
                     "entities": [
                         "entity.system.home"
                     ],
                     "actions": [
                         "action.system.home"
                     ]    
                 },
                 // 可选项：添加deeplink、applink等跳转功能
                 {
                     "entities": [
                         "test"
                     ],
                     "actions": [
                         "test"
                     ],
                     "uris": [
                         {
                             "scheme": "test"
                         }
                     ]
                 }
             ]
         }
     ],
     ...
 }
```



### 延迟任务

常用方法

| 接口名                                                       | 接口描述                                             |
| :----------------------------------------------------------- | :--------------------------------------------------- |
| startWork(work: WorkInfo): void;                             | 申请延迟任务                                         |
| stopWork(work: WorkInfo, needCancel?: boolean): void;        | 取消延迟任务                                         |
| getWorkStatus(workId: number, callback: AsyncCallback<WorkInfo>): void; | 获取延迟任务状态（Callback形式）                     |
| getWorkStatus(workId: number): Promise<WorkInfo>;            | 获取延迟任务状态（Promise形式）                      |
| obtainAllWorks(callback: AsyncCallback<Array<WorkInfo>>): void; | 获取所有延迟任务（Callback形式）                     |
| obtainAllWorks(): Promise<Array<WorkInfo>>;                  | 获取所有延迟任务（Promise形式）                      |
| stopAndClearWorks(): void;                                   | 停止并清除任务                                       |
| isLastWorkTimeOut(workId: number, callback: AsyncCallback<boolean>): void; | 获取上次任务是否超时（针对RepeatWork，Callback形式） |
| isLastWorkTimeOut(workId: number): Promise<boolean>;         | 获取上次任务是否超时（针对RepeatWork，Promise形式）  |

| 接口名                                          | 接口描述               |
| :---------------------------------------------- | :--------------------- |
| onWorkStart(work: workScheduler.WorkInfo): void | 延迟调度任务开始的回调 |
| onWorkStop(work: workScheduler.WorkInfo): void  | 延迟调度任务结束的回调 |

WorkInfo参数

| 名称            | 类型                                                         | 必填 | 说明                                                         |
| :-------------- | :----------------------------------------------------------- | :--- | :----------------------------------------------------------- |
| workId          | number                                                       | 是   | 延迟任务ID。                                                 |
| bundleName      | string                                                       | 是   | 延迟任务所在应用的包名。                                     |
| abilityName     | string                                                       | 是   | 包内ability名称。                                            |
| networkType     | [NetworkType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-resourceschedule-workscheduler-V5#networktype) | 否   | 网络类型。                                                   |
| isCharging      | boolean                                                      | 否   | 是否充电。- true表示充电触发延迟回调，false表示不充电触发延迟回调。 |
| chargerType     | [ChargingType](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-resourceschedule-workscheduler-V5#chargingtype) | 否   | 充电类型。                                                   |
| batteryLevel    | number                                                       | 否   | 电量。                                                       |
| batteryStatus   | [BatteryStatus](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-resourceschedule-workscheduler-V5#batterystatus) | 否   | 电池状态。                                                   |
| storageRequest  | [StorageRequest](https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/js-apis-resourceschedule-workscheduler-V5#storagerequest) | 否   | 存储状态。                                                   |
| isRepeat        | boolean                                                      | 否   | 是否循环任务。- true表示循环任务，false表示非循环任务。      |
| repeatCycleTime | number                                                       | 否   | 循环间隔，单位为毫秒。                                       |
| repeatCount     | number                                                       | 否   | 循环次数。                                                   |
| isPersisted     | boolean                                                      | 否   | 注册的延迟任务是否可保存在系统中。- true表示可保存，即系统重启后，任务可恢复。false表示不可保存。 |
| isDeepIdle      | boolean                                                      | 否   | 是否要求设备进入空闲状态。- true表示需要，false表示不需要。  |
| idleWaitTime    | number                                                       | 否   | 空闲等待时间，单位为毫秒。                                   |
| parameters      | [key: string]: number \| string \| boolean                   | 否   | 携带参数信息。                                               |

### 代理提醒

依赖导包

```typescript
import { reminderAgentManager } from '@kit.BackgroundTasksKit';
import { notificationManager } from '@kit.NotificationKit';
import { BusinessError } from '@kit.BasicServicesKit';
```

常用方法

| 接口名                                                       | 描述                                                 |
| :----------------------------------------------------------- | :--------------------------------------------------- |
| publishReminder(reminderReq: ReminderRequest): Promise<number> | 发布一个定时提醒类通知                               |
| cancelReminder(reminderId: number): Promise<void>            | 取消一个指定的提醒类通知                             |
| getValidReminders(): Promise<Array<ReminderRequest>>         | 获取当前应用设置的所有有效的提醒                     |
| cancelAllReminders(): Promise<void>                          | 取消当前应用设置的所有提醒                           |
| addNotificationSlot(slot: NotificationSlot): Promise<void>   | 注册一个提醒类需要使用的通知通道（NotificationSlot） |
| removeNotificationSlot(slotType: notification.SlotType): Promise<void> | 删除指定的通知通道（NotificationSlot）               |

### 文档资料

[Background Tasks Kit（后台任务开发服务）](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/background-task-kit-V5)

## Performance Analysis Kit

### HiLog

### HiAppEvent

HiAppEvent是在系统层面为应用开发者提供的一种事件打点机制，帮助应用记录在运行过程中发生的故障信息、统计信息、安全信息、用户行为信息，支撑开发者分析应用的运行情况。以便进一步统计分析访问数、日常用户活跃数量、用户操作习惯以及其他影响用户使用产品的关键因素。

打点： 记录由用户操作引起的变化，提供业务数据信息，供开发、产品、运维分析。

#### 事件类型

- 行为事件：记录用户日常操作行为的事件，例如按钮点击、界面跳转等行为。
- 故障事件：定位和分析应用故障的事件，例如界面卡顿、掉网掉话等异常。
- 统计事件：统计和度量应用关键行为的事件，例如对使用时长、访问数等的统计。
- 安全事件：记录涉及应用安全行为的事件，例如密码修改、用户授权等行为。

#### 订阅应用事件（ArkTS）

依赖导包

```typescript
import { BusinessError } from '@kit.BasicServicesKit';
import { hiAppEvent, hilog } from '@kit.PerformanceAnalysisKit';
```

常用方法

| 接口名                                                       | 描述                                                 |
| :----------------------------------------------------------- | :--------------------------------------------------- |
| write(info: AppEventInfo, callback: AsyncCallback<void>): void | 应用事件异步打点方法，使用callback方式作为异步回调。 |
| write(info: AppEventInfo): Promise<void>                     | 应用事件异步打点方法，使用Promise方式作为异步回调。  |

| 接口名                                              | 描述                                         |
| :-------------------------------------------------- | :------------------------------------------- |
| addWatcher(watcher: Watcher): AppEventPackageHolder | 添加应用事件观察者，以添加对应用事件的订阅。 |
| removeWatcher(watcher: Watcher): void               | 移除应用事件观察者，以移除对应用事件的订阅。 |

示例代码

```typescript
//在onCreate函数中添加对用户点击按钮事件的订阅 
hiAppEvent.addWatcher({
   // 开发者可以自定义观察者名称，系统会使用名称来标识不同的观察者
   name: "watcher1",
   // 开发者可以订阅感兴趣的应用事件，此处是订阅了按钮事件
   appEventFilters: [{ domain: "button" }],
   // 开发者可以设置订阅回调触发的条件，此处是设置为事件打点数量满足1个
   triggerCondition: { row: 1 },
   // 开发者可以自行实现订阅回调函数，以便对订阅获取到的事件打点数据进行自定义处理
   onTrigger: (curRow: number, curSize: number, holder: hiAppEvent.AppEventPackageHolder) => {
     // 返回的holder对象为null，表示订阅过程发生异常，因此在记录错误日志后直接返回
     if (holder == null) {
       hilog.error(0x0000, 'testTag', "HiAppEvent holder is null");
       return;
     }
     hilog.info(0x0000, 'testTag', `HiAppEvent onTrigger: curRow=%{public}d, curSize=%{public}d`, curRow, curSize);
     let eventPkg: hiAppEvent.AppEventPackage | null = null;
     // 根据设置阈值大小（默认为512KB）去获取订阅事件包，直到将订阅数据全部取出
     // 返回的事件包对象为null，表示当前订阅数据已被全部取出，此次订阅回调触发结束
     while ((eventPkg = holder.takeNext()) != null) {
       // 开发者可以对事件包中的事件打点数据进行自定义处理，此处是将事件打点数据打印在日志中
       hilog.info(0x0000, 'testTag', `HiAppEvent eventPkg.packageId=%{public}d`, eventPkg.packageId);
       hilog.info(0x0000, 'testTag', `HiAppEvent eventPkg.row=%{public}d`, eventPkg.row);
       hilog.info(0x0000, 'testTag', `HiAppEvent eventPkg.size=%{public}d`, eventPkg.size);
       for (const eventInfo of eventPkg.data) {
         hilog.info(0x0000, 'testTag', `HiAppEvent eventPkg.info=%{public}s`, eventInfo);
       }
     }
   }
 });
```

```typescript
//添加一个按钮并在其onClick函数中进行事件打点，以记录按钮点击事件  
Button("writeTest").onClick(()=>{
    // 在按钮点击函数中进行事件打点，以记录按钮点击事件
    let eventParams: Record<string, number> = { 'click_time': 100 };
    let eventInfo: hiAppEvent.AppEventInfo = {
      // 事件领域定义
      domain: "button",
      // 事件名称定义
      name: "click",
      // 事件类型定义
      eventType: hiAppEvent.EventType.BEHAVIOR,
      // 事件参数定义
      params: eventParams,
    };
    hiAppEvent.write(eventInfo).then(() => {
      hilog.info(0x0000, 'testTag', `HiAppEvent success to write event`)
    }).catch((err: BusinessError) => {
      hilog.error(0x0000, 'testTag', `HiAppEvent err.code: ${err.code}, err.message: ${err.message}`)
    });
  })
```

#### 订阅应用事件（C/C++）

打点接口功能

| 接口名                                                       | 描述                                 |
| :----------------------------------------------------------- | :----------------------------------- |
| int OH_HiAppEvent_Write(const char *domain, const char *name, enum EventType type, const ParamList list) | 实现对参数为列表类型的应用事件打点。 |

订阅接口功能

| 接口名                                                       | 描述                                         |
| :----------------------------------------------------------- | :------------------------------------------- |
| int OH_HiAppEvent_AddWatcher (HiAppEvent_Watcher *watcher)   | 添加应用事件观察者，以添加对应用事件的订阅。 |
| int OH_HiAppEvent_RemoveWatcher (HiAppEvent_Watcher *watcher) | 移除应用事件观察者，以移除对应用事件的订阅。 |

#### 系统事件分类

- **[崩溃事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/crash-events-V13)**
- **[卡死事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/freeze-events-V13)**
- **[资源泄漏事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/resource-leak-events-V13)**
- **[踩内存事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/address-sanitizer-events-V13)**
- **[主线程超时事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/main-thread-jank-events-V13)**
- **[启动耗时事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/u542f_u52a8_u8017_u65f6_u4e8b_u4ef6-V13)**
- **[滑动丢帧事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/u6ed1_u52a8_u4e22_u5e27_u4e8b_u4ef6-V13)**
- **[CPU高负载事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/u9ad8_u8d1f_u8f7d_u4e8b_u4ef6-V13)**
- **[24h功耗器件分解统计事件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/24h_u529f_u8017_u5668_u4ef6_u5206-V13)**

#### 崩溃事件

##### 订阅崩溃事件（ArkTS）

事件自定义参数设置接口功能

| 接口名                                                       | 描述                     |
| :----------------------------------------------------------- | :----------------------- |
| setEventParam(params: Record<string, ParamType>, domain: string, name?: string): Promise<void> | 事件自定义参数设置方法。 |

订阅接口功能

| 接口名                                              | 描述                                         |
| :-------------------------------------------------- | :------------------------------------------- |
| addWatcher(watcher: Watcher): AppEventPackageHolder | 添加应用事件观察者，以添加对应用事件的订阅。 |
| removeWatcher(watcher: Watcher): void               | 移除应用事件观察者，以移除对应用事件的订阅。 |

##### 订阅崩溃事件（C/C++）

#### 卡死事件

### HiTraceMeter

### HiTraceChain

### HiChecker

### HiDebug

### HiCollie

### 文档资料

[订阅应用事件（ArkTS）](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V13/hiappevent-watcher-app-events-arkts-V13)

## Network Kit

### HTTP

### WebSocket

权限

```
ohos.permission.INTERNET
```

接口

| 接口名               | 描述                                      |
| :------------------- | :---------------------------------------- |
| createWebSocket()    | 创建一个WebSocket连接。                   |
| connect()            | 根据URL地址，建立一个WebSocket连接。      |
| send()               | 通过WebSocket连接发送数据。               |
| close()              | 关闭WebSocket连接。                       |
| on(type: 'open')     | 订阅WebSocket的打开事件。                 |
| off(type: 'open')    | 取消订阅WebSocket的打开事件。             |
| on(type: 'message')  | 订阅WebSocket的接收到服务器消息事件。     |
| off(type: 'message') | 取消订阅WebSocket的接收到服务器消息事件。 |
| on(type: 'close')    | 订阅WebSocket的关闭事件。                 |
| off(type: 'close')   | 取消订阅WebSocket的关闭事件               |
| on(type: 'error')    | 订阅WebSocket的Error事件。                |
| off(type: 'error')   | 取消订阅WebSocket的Error事件。            |

示例

```typescript
import { webSocket } from '@kit.NetworkKit';
import { BusinessError } from '@kit.BasicServicesKit';

let defaultIpAddress = "ws://";
let ws = webSocket.createWebSocket();
ws.on('open', (err: BusinessError, value: Object) => {
  console.log("on open, status:" + JSON.stringify(value));
  // 当收到on('open')事件时，可以通过send()方法与服务器进行通信
  ws.send("Hello, server!", (err: BusinessError, value: boolean) => {
    if (!err) {
      console.log("Message sent successfully");
    } else {
      console.log("Failed to send the message. Err:" + JSON.stringify(err));
    }
  });
});
ws.on('message', (err: BusinessError, value: string | ArrayBuffer) => {
  console.log("on message, message:" + value);
  // 当收到服务器的`bye`消息时（此消息字段仅为示意，具体字段需要与服务器协商），主动断开连接
  if (value === 'bye') {
    ws.close((err: BusinessError, value: boolean) => {
      if (!err) {
        console.log("Connection closed successfully");
      } else {
        console.log("Failed to close the connection. Err: " + JSON.stringify(err));
      }
    });
  }
});
ws.on('close', (err: BusinessError, value: webSocket.CloseResult) => {
  console.log("on close, code is " + value.code + ", reason is " + value.reason);
});
ws.on('error', (err: BusinessError) => {
  console.log("on error, error:" + JSON.stringify(err));
});
ws.connect(defaultIpAddress, (err: BusinessError, value: boolean) => {
  if (!err) {
    console.log("Connected successfully");
  } else {
    console.log("Connection failed. Err:" + JSON.stringify(err));
  }
});
```



### Socket

### MDNS管理

### 文档资料

## 鸿蒙高证编程题

### 连续整数之和

```typescript
/*
一个整数有可能可以被表示成 m（m>1）个连续正整数之和
15=1+2+3+4+5
15=4+5+6
15=7+8
*/
process.stdin.resume();
process.stdin.setEncoding('utf-8');
let input = '';
process.stdin.on('data', (data) => {
    input += data;
});
process.stdin.on('end', () => {
    let inputArray = input.split('\n');
    let n = parseInt(inputArray[0].trim(), 10);
 
    function doFunc() {
        for (let m = 2; m * (m - 1) / 2 < n; m++) {
            if ((n - m * (m - 1) / 2) % m === 0) {
                console.log("YES");
                return;
            }
        }
        console.log("NO");
    }
 
    doFunc();
    process.exit();
});
```

### 重复字母连续次数

```typescript
/*
一个字符串，只包含大写字母，求同一个字母连续出现的最大次数。
AAAABBCDHHH,同一个字母连续出现的最大次数是 4，因为一开始 A 连续出现 4 次
*/
process.stdin.resume();
process.stdin.setEncoding('utf-8');
let input = '';
process.stdin.on('data', (data) => {
    input += data;
});
process.stdin.on('end', () => {
    let inputArray = input.split('\n');
    let str = inputArray[0].trim();
 
    function doFunc() {
        let maxCount = 0;
        let currentCount = 1;
 
        for (let i = 1; i < str.length; i++) {
            if (str[i] === str[i - 1]) {
                currentCount++;
            } else {
                if (currentCount > maxCount) {
                    maxCount = currentCount;
                }
                currentCount = 1;
            }
        }
        
        // 最后一段连续字符的处理
        if (currentCount > maxCount) {
            maxCount = currentCount;
        }
 
        console.log(maxCount);
    }
 
    doFunc();
    process.exit();
});
```

### 两数求和

```typescript
/*
输入 1 1
		2 3
输出 2
		5
*/
process.stdin.resume();
process.stdin.setEncoding('utf-8');
let input = '';
process.stdin.on('data', (data) => {
  input += data;
});
process.stdin.on('end', () => {
  let inputArray = input.trim().split('\n');
 
  function doFunc() {
    for (let i = 0; i < inputArray.length; i++) {
      let nums = inputArray[i].split(' ').map(Number);
      if (nums.length === 2) {
        let a = nums[0];
        let b = nums[1];
        let sum = a + b;
        console.log(a);
        console.log(b);
        console.log(sum);
      }
    }
  }
 
  doFunc();
  process.exit();
});
```

### 单词重量

```typescript
/*
每个句子由多个单词组成，句子中的每个单词的长度都可能不一样，
我们假设每个单词的长度Ni为该单词的重量，你需要做的就是给出整个句子的平均重量V
*/
process.stdin.resume();
process.stdin.setEncoding('utf-8');
let input = '';
process.stdin.on('data', (data) => {
    input += data;
});
process.stdin.on('end', () => {
    let inputArray = input.split('\n');
 
    function doFunc() {
        let sentence = inputArray[0].trim();
        let words = sentence.split(' ');
        let totalLength = 0;
        
        for (let word of words) {
            totalLength += word.length;
        }
        
        let averageLength = totalLength / words.length;
        console.log(averageLength.toFixed(2));
    }
 
    doFunc();
    process.exit();
});
```



## 原子服务

### 卡片服务

#### 创建卡片服务

```
在Entry(某个模块)右键选择"service widget",然后点击'static widget'或'dynamic widget'
```

#### 卡片页面

```typescript
./entry/src/ets/widget/pages/WidgetCard.ets
```



#### form_config.json

```typescript
{
  "forms": [
    {
      "name": "widget",
      "displayName": "$string:widget_display_name",
      "description": "$string:widget_desc",
      "src": "./ets/widget/pages/WidgetCard.ets",
      "uiSyntax": "arkts",
      "window": {
        "designWidth": 720,
        "autoDesignWidth": true
      },
      "colorMode": "auto",
      "isDynamic": true,
      "isDefault": true,
      "updateEnabled": false,
      "scheduledUpdateTime": "10:30",
      "updateDuration": 1,
      "defaultDimension": "2*2",
      "supportDimensions": [
        "2*2",
        "2*4",
        "4*4"
      ]
    }
  ]
}
```



### 文档资料

[鸿蒙HarmonyOS元服务项目初探](https://juejin.cn/post/7319541661149773862)

[什么是HarmonyOS元服务](https://juejin.cn/post/7427653609583296547)

[HarmonyOS元服务开发环境搭建与项目创建初始化](https://juejin.cn/post/7429236134373572659)

## 画布Canvas

![img](http://qiniu-article.myflutter.cn/img/11887a5580a3421d8d88b4f2437d5eec~tplv-k3u1fbpfcp-jj-mark:3024:0:0:0:q75.awebp)

### 文档资料

[鸿蒙纪·梦始卷#11 | 画板绘制 - 认识绘制](https://juejin.cn/post/7437332158529978420?searchId=20241204165229F8B92788D715E99778FF)

[鸿蒙--canvas实现球面运动动画](https://juejin.cn/post/7316357622847881235?searchId=20241204165229F8B92788D715E99778FF)

[利用鸿蒙NAPI实现高效绘制技术](https://juejin.cn/post/7403527810642706495)

## 类WrappedBuilder

### 源码

```typescript
declare class WrappedBuilder<Args extends Object[]> {
  builder: (...args: Args) => void;
  constructor(builder: (...args: Args) => void);
}
```

### 简例01

```typescript
@Component
export struct HomePage {
  @State text: string = 'wrappedBuilder'
  build() {
    Column({ space: 20 }) {
      //wrappedBuilder
      globalBuilder.builder(this.text, 20)
    }
  }
}
@Builder
function MyBuilder(value: string, size: number) {
  Text(value)
    .fontSize(size)
}

let globalBuilder: WrappedBuilder<[string, number]> = wrapBuilder(MyBuilder)
```



## 函数wrappedBuilder

```typescript
declare function wrapBuilder<Args extends Object[]>(
builder: (...args: Args) => void): WrappedBuilder<Args>;
```

## 实用类型

### Partial

```typescript
/**
 * Make all properties in T optional
 */
type Partial<T> = {
    [P in keyof T]?: T[P];
};
interface User {
  id: number;
  name: string;
  age: number;
}

type UserWithoutId = Partial<User>
//变成可选参数
let user: UserWithoutId = {
  id: 12
}
```

### Required

```typescript
/**
 * Make all properties in T required
 */
type Required<T> = {
    [P in keyof T]-?: T[P];
};
//转为必选参数
interface Address {
  room?: string
  street?: string
}

type AddressRequired = Required<Address>
let add: AddressRequired = {
  room: '',
  street: ''
}
```

### Readonly

```typescript
/**
 * Make all properties in T readonly
 改为只读属性
 */
type Readonly<T> = {
    readonly [P in keyof T]: T[P];
};
```

### Pick

```typescript
/**
 * From T, pick a set of properties whose keys are in the union K
 */
type Pick<T, K extends keyof T> = {
    [P in K]: T[P];
};
```

### Omit

```typescript
/**
 * Construct a type with the properties of T except for those in type K.
 */
type Omit<T, K extends keyof any> = Pick<T, Exclude<keyof T, K>>;
```



## 参考文档

[揭晓 ArkTS，重塑语法，打造更健壮和可靠的代码](https://juejin.cn/post/7329404406492102690?searchId=20240720182244E5C60CF1489D6495255C)

[5分钟秒懂ArkTs,不能错过的知识点解析](https://mp.weixin.qq.com/s/ss7sUTdkHLDlDegYjfs_Zg)

[鸿蒙开发之ArkTS基础知识](https://juejin.cn/post/7257922419330334780?from=search-suggest)

[初识ArkTS语言](https://juejin.cn/post/7304844128734412838?from=search-suggest)

[浅析HarmonyOS开发语言ArkTS](https://juejin.cn/post/7310109331504922636?from=search-suggest)

[鸿蒙应用开发ArkTS基础组件的使用](https://www.yuque.com/maxiaobai-svip/harmonyos/uc89poyssocv12b5)

[初识ArkTS语言-鸿蒙开发者官网](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/arkts-get-started-V5)

[HarmonyOS应用开发必备-ArkTS基础知识](https://juejin.cn/post/7365830295393828883?from=search-suggest)

[基础组件](https://developer.huawei.com/consumer/cn/doc/harmonyos-guides-V5/ui-js-basic-components-V5)

[鸿蒙开发 - 相对布局 RelativeContainer](https://developer.huawei.com/consumer/cn/blog/topic/03156767520624033)

[多图预览框架](https://developer.huawei.com/consumer/cn/blog/topic/03156612283719012)

[OpenHarmony HDC工具详解](https://developer.huawei.com/consumer/cn/blog/topic/03763247825380098)

[UI组件性能优化](https://developer.huawei.com/consumer/cn/doc/best-practices-V5/bpta-ui-component-performance-optimization-V5)

[鸿蒙图书手册教程](https://gitee.com/openharmony/docs/tree/master/zh-cn/application-dev/reference#/openharmony/docs/blob/master/zh-cn/application-dev/reference/syscap.md)

[HarmonyOS实战开发-ArkTS语言（状态管理）](https://juejin.cn/post/7353459702929702975)

[鸿蒙开发效率手册](https://juejin.cn/post/7372577541113413644)

[鸿蒙中的长列表「LazyForEach」：起猛了，竟然在鸿蒙系统上看到了「RecyclerView」？]()

[HarmonyOS 鸿蒙基础-实现页面滑动方法其一Scroll](https://juejin.cn/post/7410583247984803879)

[【代码案例】HarmonyOS NEXT图形锁屏案例](https://developer.huawei.com/consumer/cn/forum/topic/0201165508127488135?fid=0109140870620153026)

[鸿蒙(HarmonyOS)常见的三种弹窗方式](https://juejin.cn/post/7408812253829218323)
