### MQTT版本

MQTT 主流版本有 MQTT3.1.1 和 MQTT5。MQTT5 完全兼容 MQTT3.1.1，是在 MQTT3.1.1 的基础上进行完善补充。目前 MQTT3.1.1 的使用人数还是更多

### 概念

MQTT（Message Queuing Telemetry Transport，消息队列遥测传输）协议是基于 TCP/IP 协议栈构建的异步通信消息协议，是一种**轻量级**的、**客户端服务端**架构的、**发布/订阅模式**的消息传输协议。

MQTT协议最初版本是在1999年建立的，发明人是 Andy Stanford-Clark 和 Arlen Nipper。MQTT 是他们为了利用卫星通讯监控输油管道所开发的协议，由此可见，MQTT 就是专为低带宽、高延迟或不可靠的网络而设计的。

### MQTT 协议特点

- 简单易用，方便集成
- 安全可靠，支持TLS/SSL加密和认证机制
- 轻量级，占用带宽小，支持多种消息传输模式
- 灵活性，可知设备连接状态，可控数据传输质量

### 服务端

MQTT 服务端通常是一台服务器，它充当着 MQTT 信息传输的中心节点。其主要功能是接收来自 MQTT 客户端的信息并将其传递给其他 MQTT 客户端。此外，MQTT 服务端还负责管理客户端，确保客户端之间的通信畅通无阻，并确保 MQTT 消息被正确接收和准确投递。

服务端一般就是云平台，OneNET、阿里云、腾讯云等；也可以用 EMQ 或 Mosquitto 自己搭建服务端。

### 客户端

MQTT 客户端可以向服务端通过「发布」发送信息，也可以从服务端「订阅」来收取信息。

客户端一般就是我们的单片机，STM32、C51、树莓派等。

### 主题

在 MQTT 通讯中，客户端订阅的是一个个「主题」。MQTT 服务端在管理信息通讯时，使用「主题」来控制。

### 发布与订阅的特点

- 相互独立：客户端相互独立，彼此没有直接联系，不用知道对方的任何状态、情况。

- 空间分离：客户端只要连接同一个 MQTT 通讯网络，无论是互联网或者局域网都可以通讯。

- 时间异步：客户端的发布与订阅无需同步。若有客户端断连，服务端保存信息，待客户端上线后推送。

### MQTT报文

MQTT 协议通过交换预定义的 MQTT 控制报文来通信。MQTT 控制报文简称 MQTT 报文，接下来我将详细介绍 MQTT 报文。

### 报文结构

一个 MQTT 报文由固定报头、可变报头、有效载荷三部分组成：

- 固定报头（Fixed header），所有 MQTT 报文有，表示报文类型及报文的分组类标识。
- 可变报头（Variable header），部分 MQTT 报文有，报文类型决定了可变头是否存在及其具体内容。
- 有效载荷/消息体（Payload），部分 MQTT 报文有，存放报文的具体内容。

![img](https://lxlinux.superbed.verylink.top/item/655d9f94c458853aefd4b7c7.jpg)

### 消息类型（message type）

### QoS-服务质量

QoS（Quality of Service，服务质量）。在数据通信的过程中，有的消息很重要，不可以丢失；有的消息不重要，丢了也没关系。所以在 MQTT 中可以配置 QoS，给不同重要的消息不同的服务质量。

MQTT 协议有三种服务质量级别：

- QoS = 0：最多发一次
- QoS = 1：最少发一次
- QoS = 2：保证收一次

对于不同重要级的消息选择不同的 QoS，较为重要消息的使用 QoS = 1 和 QoS = 2

### MQTT共有15种协议

1. CONNECT

2. CONNACK

3. PUBLISH

4. PUBACK

5. PUBREC

6. PUBREL

7. PUBCOMP

8. DISCONNECT

9. PINGREQ

10. PINGRES

11. SUBSCRIBE

12. SUBACK

13. UNSUBSCRIBE

14. UNSUBACK

15. AUTH

### MQTT共有7种数据格式

1. one Byte
2. Two Byte
3. Four Byte
4. UTF-8 String
5. UTF-8 String Pair
6. Variable Byte Integer
7. Binary Data

### **Host类型**

```
协议一般有 4 种：
基于普通 TCP 的 MQTT、
基于 SSL/TLS 的 MQTT、
基于 WebSocket 的 MQTT，
基于加密 WebSocket 的 MQTT
```



### 参数类MqttQos

```typescript
export type MqttQos = 0 | 1 | 2;
```

### 参数类MqttPersistenceType

```
export type MqttPersistenceType = 0 | 1 | 2;
```

### 参数类MqttClientOptions

```
export interface MqttClientOptions {
  url: string;
  clientId: string;  
  persistenceType?: MqttPersistenceType;
}
```

### 参数类MqttConnectOptions

```
export interface MqttConnectOptions {
  cleanSession?: boolean; // default is true.
  connectTimeout?: number; // default is 30s.
  keepAliveInterval?: number; // default is 60s.
  serverURIs?: Array<string>;
  retryInterval?: number; // default is 0.  
  sslOptions?: {
      enableServerCertAuth?: boolean;
      verify?: boolean;
      caPath?: string;
      trustStore?: string;
      keyStore?: string;
      privateKey?: string;
      privateKeyPassword?: string;
      enabledCipherSuites?: string;
      sslVersion?: MQTT_SSL_VERSION;
  }
  willOptions?: {
    topicName: string;
    message: string;
    retained?: boolean;
    qos?: MqttQos;
  };  
  MQTTVersion?: number;
  automaticReconnect?: boolean; // default is false.
  minRetryInterval?: number; // default is 1s.
  maxRetryInterval?: number; // default is 60s.  
  userName: string;  
  password: string;  
}
```

### 参数MqttSubscribeOptions

```
export interface MqttSubscribeOptions {
  topic: string;
  qos: MqttQos;
}
```

### 参数MqttPublishOptions

```
export interface MqttPublishOptions {
  topic: string;
  payload: string;
  qos: MqttQos;
  retained?: boolean;  
  dup?: boolean;  
  msgid?: number;  
}  
```

### MqttResponse

```
export interface MqttResponse {
  message: string;
  code: number;
}
```

### MqttMessage

```
export interface MqttMessage {
  topic: string;
  payload: string;
  qos: MqttQos;
  retained: number;
  dup: number;
  msgid: number;
}
```

### MqttClient

```
export interface MqttClient {
  connect(options: MqttConnectOptions, callback: AsyncCallback<MqttResponse>): void;
  connect(options: MqttConnectOptions): Promise<MqttResponse>;
  destroy(): Promise<boolean>;  
  disconnect(callback: AsyncCallback<MqttResponse>): void;
  disconnect(): Promise<MqttResponse>;  
  messageArrived(callback: AsyncCallback<MqttMessage>): void;  
  connectLost(callback: AsyncCallback<MqttResponse>): void;  
  publish(options: MqttPublishOptions, callback: AsyncCallback<MqttResponse>): void;
  publish(options: MqttPublishOptions): Promise<MqttResponse>;
  subscribe(options: MqttSubscribeOptions, callback: AsyncCallback<MqttResponse>): void;
  subscribe(options: MqttSubscribeOptions): Promise<MqttResponse>;  
  unsubscribe(options: MqttSubscribeOptions, callback: AsyncCallback<MqttResponse>): void;
  unsubscribe(options: MqttSubscribeOptions): Promise<MqttResponse>;  
  isConnected(): Promise<boolean>;  
  reconnect(): Promise<boolean>;  
  setMqttTrace(level: MQTTASYNC_TRACE_LEVELS): void;  
}
```

### 枚举类MQTTASYNC_TRACE_LEVELS

```
export enum MQTTASYNC_TRACE_LEVELS {
  MQTTASYNC_TRACE_MAXIMUM = 1,
  MQTTASYNC_TRACE_MEDIUM,
  MQTTASYNC_TRACE_MINIMUM,
  MQTTASYNC_TRACE_PROTOCOL,
  MQTTASYNC_TRACE_ERROR,
  MQTTASYNC_TRACE_SEVERE,
  MQTTASYNC_TRACE_FATAL
}
```

### 枚举类MQTT_SSL_VERSION

```
export enum MQTT_SSL_VERSION {
  MQTT_SSL_VERSION_DEFAULT = 0,
  MQTT_SSL_VERSION_TLS_1_0,
  MQTT_SSL_VERSION_TLS_1_1,
  MQTT_SSL_VERSION_TLS_1_2,
}
```

### 类MqttAsync

```typescript
import mqttAsync from 'libmqttasync.so';
import { MqttClientOptions, MqttClient } from 'libmqttasync.so'
class MqttAsync {
  public static createMqtt(options: MqttClientOptions): MqttClient {
    console.log('AsyncMqtt createMqtt_napi')
    return mqttAsync.createMqttSync(options)
  };
}
export default MqttAsync
```

### 服务器EMQ

#### 连地址

```
jba41a51.ala.cn-hangzhou.emqxsl.cn
MQTT over TLS/SSL 端口：8883
WebSocket over TLS/SSL 端口：8084

```

#### Api地址

```
https://jba41a51.ala.cn-hangzhou.emqxsl.cn:8443/api/v5
```

### MQTTX服务器

```
mqtt://broker.emqx.io:1883
mqttx_454ec92a
```

### 创建mqtt

```typescript
const clientOptions: MqttClientOptions = {  url: '192.168.xxx.xxx/:1883',  clientId: 'client_id_' + new Date().getTime(),  persistenceType: 1, }
const connectOptions: MqttConnectOptions = {  userName: '',  password: '',  connectTimeout: 30, }  
this.mqClient = MqttAsync.createMqtt(this.clientOptions);
this.mqClient.connect(this.connectOptions, (data: MqttResponse) => {      console.log(TAG+" data: "+JSON.stringify(data)); 
 if (data.code == 0) {        
     this.messageArrived();        
     this.subscribe('主设备号/#');     
 	}   
});
```

### 监听mqtt

```typescript
//接收消息，使用此接口后，当订阅的主题有消息发布时，会自动接收到消息。
public messageArrived(): void {  this.mqClient.messageArrived((err, data) => {    console.log(TAG+"messageArrived!!!!!!!!!!!");    console.log(TAG+"messageArrived data:"+JSON.stringify(data)); }); }
```

### 订阅消息

```typescript
public subscribe(topic: string, qos: QoS = 1): void {  const subscribeOption: MqttSubscribeOptions = { topic, qos };  this.mqClient.subscribe(subscribeOption, (err, data)=>{      this.handleMessage(data) }); }
```

### 发布消息

```typescript
public publish<T>(topic: string, payload: string | Record<string, any>,  qos: QoS = 0): void {    if (typeof payload !== 'string') {      payload = JSON.stringify(payload)   }
                                                                                              const payloadLen = payload.length;    
                                                                                              const publishOption: MqttPublishOptions = { topic, payload, qos, payloadLen };    console.log(TAG, 'publishOption data: ' + JSON.stringify(publishOption));    this.mqClient.publish(publishOption, (err, data)=>{       console.log(TAG+"publish!!!!!!!!!!!");    console.log(TAG+"publish data:"+JSON.stringify(data)); })); }
```



### 参考文档

[HarmonyOS ArkTS 实现MQTT协议(2)](https://developer.huawei.com/consumer/cn/blog/topic/03140280038355014)

[万字猛文：MQTT原理及案例](https://juejin.cn/post/7332331306212835365?searchId=20240817160014D6C0C251B7841E08E5EB)

[HarmonyOS ArkTS 实现MQTT协议](https://juejin.cn/post/7318619321420202024)

[MQTT 协议入门：基础知识和快速教程](https://juejin.cn/post/7249205118532190269)

[OpenHarmony中使用Mqtt](https://developer.huawei.com/consumer/cn/blog/topic/03145131922336026)  @ohos/mqtt

[EMQ服务官网](https://cloud.emqx.com/console/deployments/new)

[EMQX官网](https://www.emqx.com/zh)