--
title:Golang 函数
create:2026-09-30
update:2026-09-30
category:Golang
tags:[golang,function,函数]
summary:函数定义与多返回值、错误处理（无 try/catch）、匿名函数与闭包、泛型函数、defer 与 LIFO、用接口实现多态。
top:1
copyright:true
--

# Go 语言学习笔记：函数

> 整理自掘金文章《[Go语言第八章(函数)](https://juejin.cn/post/7672244196402954303)》（作者：小满zs）。
> 函数是一段可复用的代码块：接收参数、执行逻辑，再通过 `return` 把结果交还给调用方。Go 的入口也是函数——程序从 `main` 开始跑。

---

## 一、函数定义

基本语法：

```go
func 函数名(参数名 参数类型) 返回值类型 {
	// 函数体
	return 返回值
}
```

拆开看每一部分：

| 部分 | 含义 |
|------|------|
| `func` | 声明函数的关键字 |
| 函数名 | 包内可调用的名字；首字母大写表示对外导出 |
| 参数列表 | `名字 类型`，多个参数用逗号分隔；同类型可简写为 `a, b int` |
| 返回值类型 | 写在参数列表后面；无返回值时可省略 |
| `return` | 结束函数并把结果返回给调用方；返回值类型必须与声明一致 |

最小例子：

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2)) // 3
}
```

参数同类型时可以合并写法：`func add(a, b int) int`。

### 1.1 多返回值

Go 支持一次返回多个值，常见于「结果 + 错误」（详见下一节「错误处理」）：

```go
package main

import "fmt"

func calc(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, sub := calc(1, 2)
	fmt.Println(sum, sub) // 3 -1
}
```

也可以给返回值命名（**命名返回值**）。此时 `return` 可以不带表达式，会返回当前命名变量的值：

```go
func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return // 等价于 return sum, sub
}
```

不需要某个返回值时，用 `_` 丢弃：

```go
sum, _ := calc(1, 2)
```

### 1.2 注意事项

1. **不能在函数内部声明具名函数**。这不只是 `main`，任意函数里都不允许再写 `func foo() {}` 这种带名字的声明。需要嵌套逻辑时，用**匿名函数**（见下文）。
2. 函数必须先声明再调用的说法不成立：**同一包内函数可以任意顺序定义**。
3. Go 是**按值传递**：传入的是副本。想改调用方的变量，要传指针，或使用切片 / map 这类引用语义类型。
4. 可变参数用 `...类型`，例如 `func sum(nums ...int) int`，调用时写成 `sum(1, 2, 3)`。

错误示例（无法编译）：

```go
func main() {
	// 不允许：函数里再声明具名函数
	func add(a, b int) int {
		return a + b
	}
}
```

正确写法：把内部逻辑写成匿名函数，赋给变量：

```go
func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
}
```

---

## 二、错误处理

Go **没有** `try / catch`。函数出错时，惯例是多返回值：**最后一个是 `error`**。调用方必须自己检查。

`error` 是内置接口，只有一个方法 `Error() string`。没出错时返回 `nil`，有错时返回非 `nil`。

```go
package main

import (
	"errors"
	"fmt"
)

// 除数为 0 时返回错误
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错了:", err)
		return
	}
	fmt.Println("结果:", result) // 5

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错了:", err) // 出错了: 除数不能为 0
	}
}
```

常用写法说明：

| 写法 | 作用 |
|------|------|
| `errors.New("说明")` | 创建一个简单错误 |
| `fmt.Errorf("a=%d 非法", a)` | 带格式化信息的错误 |
| `if err != nil { ... }` | 检查错误的标准姿势 |
| `return 0, err` | 把错误继续往上抛给调用方 |

> **注意**
> 1. **不要忽略 `error`**。写成 `result, _ := divide(...)` 等于把错误吞掉，出了问题很难排查。
> 2. 出错时，第一个返回值通常给零值（如 `0`、`""`、`nil`），真正该看的是 `err`。
> 3. 很多标准库函数也是这个模式，例如 `os.Open`、`strconv.Atoi`：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	fmt.Println(n) // 123

	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("转换失败:", err)
	}
}
```

**和 `panic` 的区别**（先建立直觉即可）：日常可预期的失败（文件不存在、输入不合法）用 `error`；程序无法继续的严重问题才用 `panic`。入门阶段优先掌握 `if err != nil`。

---

## 三、匿名函数

没有名字的函数叫匿名函数。可以立刻执行，也可以赋给变量、当作参数传递。

### 3.1 立刻执行（IIFE）

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("匿名函数")
	}() // 末尾 () 表示立刻调用
}
```

### 3.2 赋给变量后再调用

```go
package main

import "fmt"

func main() {
	greet := func(name string) {
		fmt.Println("你好,", name)
	}
	greet("Go")
}
```

### 3.3 注意事项

1. **闭包会捕获外层变量**。匿名函数引用外层变量时，拿到的是变量本身（引用语义），不是当时的拷贝。循环里起 goroutine 时尤其容易踩坑，要额外注意。
2. 匿名函数可以访问外层作用域的变量，适合做回调、延迟逻辑（配合 `defer`）。
3. 只有「函数类型」的值才能赋值、传递；具名函数本身不是变量，但可以写成 `f := add` 把函数当值用。
4. 需要反复复用、导出给其他包用的逻辑，优先写成包级具名函数，可读性和调试都更好。

### 3.4 闭包示例

```go
package main

import "fmt"

func makeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	next := makeCounter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
}
```

---

## 四、泛型函数

没有泛型时，想对 `int`、`float64` 做同样的加法，往往要写两套几乎一样的函数，很繁琐：

```go
func addInt(a, b int) int { return a + b }
func addFloat(a, b float64) float64 { return a + b }
```

Go 1.18 起支持**类型参数**。把「类型」也参数化，一套函数就能服务多种类型——这就是泛型的中心思想：

```go
package main

import "fmt"

// T 被约束为 int、uint 或 float64，这些类型都能做 +
func add[T int | uint | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add[int](1, 2))       // 显式指定类型参数
	fmt.Println(add[uint](1, 2))
	fmt.Println(add(1.5, 2.5))        // 编译器从实参推断 T 为 float64
}
```

- `[T int | uint | float64]`：类型参数 `T`，以及它允许的类型集合（**类型约束**）。
- 调用时可以写 `add[int](1, 2)`，多数情况也能靠实参**类型推断**省略。
- 约束里的类型必须都支持函数体里用到的运算；否则编译报错。

也可以用 `comparable`、自定义 interface 约束等更复杂的写法；入门先记住：**泛型用来消灭「同逻辑、多类型」的重复函数**。

---

## 五、defer 关键字

`defer` 会把一次函数调用推迟到**当前函数即将返回之前**执行，常用于关闭文件、解锁、收尾打印等「成对出现」的清理工作。

### 5.1 基本例子

```go
package main

import "fmt"

func main() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
}
```

输出：

```text
main 1
main 2
defer 1
```

### 5.2 多个 defer：后进先出（LIFO）

多个 `defer` 按栈顺序执行，**最后注册的最先跑**：

```go
package main

import "fmt"

func main() {
	defer fmt.Println("第一")
	defer fmt.Println("第二")
	defer fmt.Println("第三")
	fmt.Println("函数体")
}
```

输出：

```text
函数体
第三
第二
第一
```

### 5.3 实用案例：保证资源释放

打开文件后立刻 `defer Close()`，后面无论正常返回还是提前 `return`，都会关闭：

```go
package main

import (
	"fmt"
	"os"
)

func readFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close() // 函数返回前关闭，避免泄漏

	buf := make([]byte, 64)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}
	fmt.Println(string(buf[:n]))
	return nil
}
```

### 5.4 再记两点

1. **参数在 defer 语句执行时就已经算好**，不是等到真正调用时才算：

```go
package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("defer:", x) // 这里已经把 x=1 传进去了
	x = 2
	fmt.Println("main:", x)
}
```

输出是 `main: 2`，然后 `defer: 1`。

2. `defer` 绑定的是**函数**，不是代码块。写在 `for` 里的 `defer` 会等到外层函数返回才执行；循环里反复 `defer` 可能堆积很多调用，一般要包一层函数或换写法。

---

## 六、用接口模拟「面向对象」能力

先把话说清楚：

- Go **没有** class、没有基于类的继承，也**不是**传统 Java / C++ 那种面向对象语言。
- Go 提供的是：**结构体存数据**、**方法绑行为**、**接口描述能力**。多态主要靠 interface 实现。
- 所以更准确的说法是：Go 用「组合 + 方法 + 接口」达到类似面向对象里「多态、面向行为编程」的效果，而不是复制一套 class 体系。

接口只声明「要有哪些方法」，不关心具体类型：

```go
package main

import "fmt"

// Speaker 描述「能说话」这种能力
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": 汪汪"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": 喵喵"
}

// 只依赖接口，不依赖具体类型
func say(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	say(Dog{Name: "旺财"})
	say(Cat{Name: "咪咪"})
}
```

**要点**：

1. 类型**不用显式写 implements**。只要方法集满足接口，就自动实现了该接口。
2. 函数参数写成接口类型后，传入任何实现了这些方法的值都可以——这就是接口带来的多态。
3. 需要「复用字段 / 方法」时，优先用结构体**嵌入（组合）**，而不是继承。

> 常见误区：看到方法、接口就以为 Go 是 class OOP——其实它刻意避开了继承树，把重点放在接口和组合上。

---

## 七、练习题

1. **基本函数**：写一个 `max(a, b int) int`，返回较大的那个数；在 `main` 里调用并打印结果。
2. **多返回值**：写一个 `swap(a, b int) (int, int)`，交换两个整数并返回；用 `x, y := swap(1, 2)` 验证。
3. **错误处理**：写一个 `sqrt(n float64) (float64, error)`：`n < 0` 时返回错误（可用 `errors.New`），否则返回平方根（可用 `math.Sqrt`）。在 `main` 里分别用正数和负数调用，正确处理 `err`。
4. **匿名函数**：在 `main` 里用匿名函数实现「打印 1 到 5」，立刻执行；再把同一个逻辑赋给变量后调用一次。
5. **综合练习**：写一个 `parseAndDouble(s string) (int, error)`，用 `strconv.Atoi` 把字符串转成整数后乘以 2；转换失败则把错误返回。在 `main` 里测试 `"21"` 和 `"abc"` 两种输入。

---

## 八、核心要点速记

- **定义**：`func 名(参数) 返回值 { return ... }`；同类型参数可简写 `a, b int`；首字母大写即导出。
- **多返回值**：`(int, int)`；可**命名返回值**，此时裸 `return` 返回命名变量；不需要的值用 `_` 丢弃。
- **限制**：函数内**不能声明具名函数**（改用匿名函数）；同包内函数可任意顺序定义；**按值传递**（改外部请传指针或用引用语义类型）；可变参数 `...T`。
- **错误处理**：无 try/catch；惯例是最后返回 `error`，`nil` 表示无错；`if err != nil` 检查；**不要忽略 error**；可预期失败用 `error`，严重问题才 `panic`。
- **匿名函数**：可 IIFE 立刻执行 `func(){...}()`，也可赋给变量；**闭包捕获外层变量本身**（引用语义）。
- **泛型函数**：Go 1.18+ 支持 `func f[T 约束](a, b T) T`；可显式指定或靠实参推断；用于消灭「同逻辑多类型」的重复代码。
- **defer**：推迟到函数返回前执行；多个 defer **LIFO**；参数在 defer 语句处即求值；常用于 `defer f.Close()`。
- **接口多态**：Go 无 class/继承；「结构体 + 方法 + 接口」实现多态；**隐式实现**（无需 implements）；复用优先用嵌入组合。

---

*参考资料：[Go语言第八章(函数)](https://juejin.cn/post/7672244196402954303) · 作者：小满zs*
