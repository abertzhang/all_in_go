## vue 指令

### v-html

```typescript
<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      rawHtml: '<span style="color: red">这是一段红色的文字</span>'
    }
  }
}
</script>
<template>
<div class="hello">
  <p>
    {{rawHtml}}
  </p>
  <p v-html="rawHtml">
  </p>
</div>
</template>
```

### v-bind

```typescript
<div class="hello" v-bind:id="dynamicId">
  动态 id: {{ dynamicId }}
</div>
//

```

### 模板表达式

单个表达式，js 表达式

```
<p>{{num+1}}</p>

```

### v-if

v-if子组件销毁和重建

也是惰性的

有较高的切换开销

### v-else

### v-show

元素总是被渲染

有较高的初始渲染开销

### v-for

key 属性

 ```
    <ul>
       <li v-for="name in names" :key="name">
       //<li v-for="(item,index) in names" :key="item">
         {{ name }}
       </li>
     </ul>
 ```

### v-on

缩写@符号

```typescript
//声明
<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      count: 0,
    }
  }
}
</script>
//使用
<button @click="count++">
   Click me {{ count }} times
</button>

```

### v-model

表单双向绑定

修饰符 lazy  trim  number

## 组件

### prop

```typescript
//父组件
<script  lang="ts">
import HelloWorld from './components/HelloWorld.vue'
export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data() {
    return {
      title: '我是一个来自父级消息',
    }
  }
}
</script>

<template>
  <HelloWorld :title="title"/>
</template>
```

```typescript
//子组件
<script  lang="ts">
export default {
  name: 'HelloWorld',
  props: {
    title: {
      type: String,
      default: 'Hello World'
    },
      names:{
      type:Arrary<string>,
      default:()=>[]
    }
  }
}
</script>
<template>
  <h3>
    prop传递数据:{{title}}
  </h3>
</template>

```

### emit

```typescript
<script  lang="ts">
import HelloWorld from './components/HelloWorld.vue'
export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data() {
    return {
      msg: '',
    }
  },
  methods: {
    getDataHandle(msg:string) {
      this.msg = msg
    }
  }
}
</script>
<template>
  <HelloWorld @onEvent="getDataHandle"/>
  <h3>
    父组件自定义事件:{{msg}}</h3>
</template>


```

```typescript
<script  lang="ts">
export default {
  name: 'HelloWorld',
  data() {
    return {
      msg: '我是一个来自子级消息',
    }
  },
  methods: {
    sendHandleClick() {
      this.$emit('onEvent', this.msg)
    }
  },
}
</script>

<template>
  <h3>
    子组件自定义事件:{{msg}}
  </h3>
  <button @click="sendHandleClick">
    点击我传递到父组件
  </button>
</template>

```



### 组件生命周期

```typescript
<script  lang="ts">
export default {
  beforeCreate() {
    console.log('子组件创建前')
  },
  created() {
    console.log('子组件创建')
  },
  beforeMount() {
    console.log('子组件挂载钱')
  },
  mounted() {
   console.log('子组件挂载')
  },

  beforeUpdate() {
    console.log('子组件更新')
  },
  updated() {
    console.log('子组件更新')
  },
  beforeunmount() {
    console.log('子组件销毁')
  },
  unmounted() {
    console.log('子组件销毁')
  },
  activated() {
    console.log('子组件激活')
  },
  deactivated() {
    console.log('子组件失活')
  },
  render() {
    console.log('子组件渲染')
  },
  renderTracked() {
    console.log('子组件渲染')
  },
  renderTriggered() {
    console.log('子组件渲染')  
  },
  errorCaptured() {
    console.log('子组件错误')
  }
}
</script>
```

## pnpm

### 安装方式

#### npm

```
npm install -g pnpm
```

#### 命令

```
//mac
curl -fsSL https://get.pnpm.io/install.sh | sh -
//win
iwr -useb https://get.pnpm.io/install.ps1 | iex
```

#### node

```
corepack enable
```



### 常用命令

```
pnpm --version
pnpm add -g pnpm@latest
pnpm config set registry http://registry.npmmirror.com
pnpm install
pnpm add <package-name>
pnpm run <script-name>
pnpm init
pnpm remove <package-name>
pnpm add <package-name> --save-dev
pnpm list
pnpm list -g
pnpm update <package-name>
pnpm update
pnpm config set <key> <value>
pnpm config set registry https://registry.npmjs.org
//清理缓存
pnpm store prune
pnpm create vite my-vue-project -- --template vue
```

## 镜像



### 推荐资料

[pnpm 的安装 ](https://www.cnblogs.com/z5337/p/18700738)

## 三方库

### pinia

## 参考文档

