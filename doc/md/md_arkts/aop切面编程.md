## 装饰器分类

```
@ClassDecorator() // 类装饰器
class A {
  @PropertyDecorator() // 属性装饰器
  name: string;
  @MethodDecorator() // 方法装饰器
  fly(
    @ParameterDecorator() // 参数装饰器
    meters: number
  ) {
    // code
  }
  @AccessorDecorator() // 存取器装饰器
  get egg() {
    // code
  }
  set egg(e) {
    // code
  }
}
```



## 各种装饰器

```typescript
declare type ClassDecorator = <TFunction extends Function>(target: TFunction) => TFunction | void;
```

```typescript
declare type PropertyDecorator = (target: Object, propertyKey: string | symbol) => void;
```

```typescript
declare type ParameterDecorator = (target: Object, propertyKey: string | symbol, parameterIndex: number) => void;
```

```typescript
declare type MethodDecorator = <T>(target: Object, propertyKey: string | symbol, descriptor: TypedPropertyDescriptor<T>) => TypedPropertyDescriptor<T> | void;
```

## 接口TypedPropertyDescriptor

```typescript
interface TypedPropertyDescriptor<T> {
    enumerable?: boolean;
    configurable?: boolean;
    writable?: boolean;
    value?: T;
    get?: () => T;
    set?: (value: T) => void;
}
```



## 日志输出方法装饰器

```
function logMethodCall<T>(target: Object, propertyKey: string,
  descriptor: TypedPropertyDescriptor<T>): TypedPropertyDescriptor<T> | void {
  const originalMethod = descriptor.value
  console.log('target',JSON.stringify(target))
  console.log('propertyKey',JSON.stringify(propertyKey))
  console.log('descriptor',JSON.stringify(descriptor))
  // descriptor.value = function (){
  //   const  result =originalMethod.app
  // }
  return descriptor;
}

class MyClass {
  @logMethodCall
  myMethod() {
    console.log("Inside myMethod");
    return 42;
  }
}
```



## 节流

节流是忽略操作，在触发事件时，立即执行目标操作，如果在指定的时间区间内再次触发了事件，则会丢弃该事件不执行，只有超过了指定的时间之后，才会再次触发事件。

## 防抖

*防抖是延时操作*，在触发事件时，不立即执行目标操作，而是给出一个延迟的时间，如果在指定的时间区间内再次触发了事件，则重置延时时间，只有当延时时间走完了才会真正执行。

## 参考文档

[HarmonyOS ：扩展修饰器，实现节流、防抖、权限申请](https://juejin.cn/post/7373194499530244136)

[鸿蒙HarmonyOS应用开发-窥探：State装饰器](https://blog.csdn.net/shudaoshanQAQ/article/details/135351218)

[HarmonyOS开发实战：自定义装饰器实现Lifecycle组件](https://blog.csdn.net/m0_64422261/article/details/140326411?ops_request_misc=%257B%2522request%255Fid%2522%253A%252241288767-404B-4921-BA31-3ECE340E3CEC%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=41288767-404B-4921-BA31-3ECE340E3CEC&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-18-140326411-null-null.142^v100^pc_search_result_base5&utm_term=%E9%B8%BF%E8%92%99%E6%96%B9%E6%B3%95%E8%A3%85%E9%A5%B0%E5%99%A8&spm=1018.2226.3001.4187)