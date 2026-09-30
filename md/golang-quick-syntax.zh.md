--
title:Golang 快速语法基础
create:2026-09-30
update:2026-09-30
category:Golang
tags:[golang,syntax,基础]
summary:由 01basic/02base/base.go 生成的 Go 语法速查：变量、常量、内置类型、复合类型、零值、控制流、运算符、类型转换、strconv 与 fmt。
top:1
copyright:true
--

# Golang 快速语法基础

> 一份简洁、以示例驱动的参考手册，由 [`01basic/02base/base.go`](../01basic/02base/base.go) 的代码生成。
> 下文每段示例都对应源文件中实际用到的语法。

## 1. 包与入口

每个可运行的 Go 程序都位于 `package main` 中，并暴露一个 `main()` 函数，它是程序的入口点。

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    variableExamples()
    loopExamples()
    operatorExamples()
    convertExamples()
}
```

其他函数只是普通的顶层声明，必须显式调用才会执行。

---

## 2. 变量

### 2.1 显式类型 + 初始值

```go
var name string = "Alice"
var a int = 10
var b float64 = 20.5
```

### 2.2 类型推断

当提供了初始值时，类型可以省略，Go 会自动推断：

```go
var age = 30          // 推断为 int
gender := "Female"   // 短声明，推断为 string
from, career := "USA", "Engineer"  // 多重短声明
```

> **注意：** `:=` 是*短声明*运算符，要求左侧**至少有一个新变量**。在同一作用域内用 `:=` 重新声明已声明过的变量会编译报错：
>
> ```go
> c := "hello"        // 此处 c 是 string
> // …… 同一作用域后面：
> c, err := strconv.ParseComplex(s, 128)  // ❌ no new variables on left side of :=
> ```
>
> 修正方法是换一个新名字：
>
> ```go
> comp, err := strconv.ParseComplex(s, 128) // ✅ comp 是新变量，err 被复用
> ```

### 2.3 多重赋值

```go
from, career := "USA", "Engineer"
x, y, z int = 1, 2, 3
```

---

## 3. 常量

常量用 `const` 声明。`const` 块用于成组声明相关常量。

```go
const pi = 3.14

const (
    e = 2.71
    g = 9.81
)
```

---

## 4. 内置类型

| 类别 | 类型 |
|------|------|
| 整型 | `int8 int16 int32 int64 int`、`uint8 uint16 uint32 uint64 uint uintptr` |
| 浮点 | `float32 float64` |
| 复数 | `complex64 complex128` |
| 布尔 | `bool` |
| 字符串 | `string` |

```go
var a int = 10
var b float64 = 20.5
c := 30
```

---

## 5. 复合类型

### 5.1 数组（定长）

```go
var arr [3]int = [3]int{1, 2, 3}
```

### 5.2 切片（动态）

```go
var slice []int = []int{4, 5, 6}
var emptySlice []int = []int{}
```

### 5.3 Map

```go
var m map[string]int = map[string]int{"one": 1, "two": 2}
var emptyMap map[string]int = map[string]int{}
```

### 5.4 结构体

```go
type Person struct {
    Name string
    Age  int
}
var p Person = Person{Name: "Bob", Age: 25}
```

### 5.5 指针

```go
var ptr *int = &a
fmt.Println(*ptr)   // 解引用
```

### 5.6 接口

```go
var i interface{} = "Hello"
```

### 5.7 函数作为值

```go
var f func(int, int) int = func(x, y int) int {
    return x + y
}
fmt.Println(f(2, 3))
```

### 5.8 通道（channel）

```go
ch := make(chan int)
go func() {
    ch <- 42
}()
fmt.Println(<-ch)
```

---

## 6. 零值与 `nil`

声明时未显式赋初值的变量，会取其**零值**：

| 类型 | 零值 |
|------|------|
| int | `0` |
| float | `0.0` |
| bool | `false` |
| string | `""` |
| 指针 / 切片 / map / chan / 函数 / 接口 | `nil` |

```go
var zeroInt int
var zeroFloat float64
var zeroBool bool
var zeroString string

var n *int = nil   // nil 指针
```

---

## 7. 控制流

### 7.1 `if` / `else if` / `else`

```go
score := 85
if score >= 80 {
    fmt.Println("Grade: A")
} else if score >= 60 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: C")
}
```

### 7.2 带 `fallthrough` 的 `switch`

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of the work week")
    fallthrough   // 继续执行下一个 case
case "Friday":
    fmt.Println("End of the work week")
default:
    fmt.Println("Midweek")
}
```

### 7.3 `for` 循环

Go 只有一个循环关键字 `for`，它有多种用法。

**经典三段式循环**

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**while 风格（仅条件）**

```go
j := 0
for j < 5 {
    fmt.Println(j)
    j++
}
```

**`for range`** —— 遍历数组、切片、map 和字符串：

```go
arr := []int{1, 2, 3, 4, 5}
for index, value := range arr {
    fmt.Println(index, value)
}

person := map[string]string{"name": "Alice", "city": "New York"}
for key, value := range person {
    fmt.Println(key, value)
}

text := "你好Go"
for index, char := range text {   // 遍历字符串得到的是 rune（字符）
    fmt.Printf("下标 %d,字符 %c\n", index, char)
}
```

**嵌套循环**

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}
```

### 7.4 `break` / `continue` / 带标签的 `break` / `goto`

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}

// 带标签的 break 用于跳出外层循环
Outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break Outer
            }
            fmt.Println(i, j)
        }
    }

// goto 跳转到同一作用域内的标签
i := 0
here:
    fmt.Println(i)
    i++
    if i < 5 {
        goto here
    }
```

### 7.5 `defer`

`defer` 调用会在所在函数返回之后执行（后进先出 LIFO）。

```go
defer fmt.Println("This will be printed last")
fmt.Println("This will be printed first")
```

---

## 8. 运算符

### 8.1 算术与比较

```go
a := 10
b := 3
fmt.Println(a + b)  // 13
fmt.Println(a - b)  // 7
fmt.Println(a * b)  // 30
fmt.Println(a / b)  // 3
fmt.Println(a % b)  // 1
fmt.Println(a == b) // false
fmt.Println(a > b)  // true
```

### 8.2 位运算

```go
fmt.Println(a & b)   // 2   与 AND
fmt.Println(a | b)   // 11  或 OR
fmt.Println(a ^ b)   // 9   异或 XOR
fmt.Println(a &^ b)  // 8   与非 AND NOT
fmt.Println(a << 1)  // 20  左移
fmt.Println(a >> 1)  // 5   右移
```

### 8.3 逻辑运算

```go
flag, isInit := true, false
fmt.Println(flag && isInit) // false
fmt.Println(flag || isInit) // true
fmt.Println(!flag)          // false
```

### 8.4 赋值与复合赋值

```go
c := a + b
c += 2   // 15
c -= 3   // 12
c *= 2   // 24
c /= 4   // 6
c %= 5   // 1
c <<= 1  // 2
c >>= 1  // 1
c &= 3   // 1
c |= 2   // 3
c ^= 1   // 2
c &^= 1  // 2
```

---

## 9. 类型转换

Go **不做**隐式的数值转换，必须使用显式的转换函数/强制转换。

```go
a := 123
b := 45.67
c := "hello"
d := []byte("world")

fmt.Println(int32(a))     // int -> int32
fmt.Println(float64(b))   // float -> float64
fmt.Println(c)            // string
fmt.Println(d)            // []byte

// string <-> []byte
s5 := "hello"
b2 := []byte(s5)
b3 := []byte{119, 111, 114, 108, 100}
s6 := string(b3)

// string <-> rune
s7 := "hello"
r := []rune(s7)
s8 := string(r)

// interface{} <-> 具体类型
s9 := "hello"
var in interface{} = s9
s10, ok := in.(string)   // 类型断言
if ok {
    fmt.Println(s10)
}
```

### 布尔与数值互转（手动）

```go
// bool -> int
b6 := true
i3 := 0
if b6 {
    i3 = 1
}
// int -> bool
i4 := 1
b7 := i4 != 0
// float -> bool
f5 := 0.0
b8 := f5 != 0
// bool -> float
b9 := true
f6 := 0.0
if b9 {
    f6 = 1.0
}
```

---

## 10. `strconv` —— 字符串转换

`strconv` 包用于解析和格式化基本类型。

```go
// string -> int
s := "123"
i, err := strconv.Atoi(s)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(i) // 123
}

// int -> string
i2 := 456
s2 := strconv.Itoa(i2)

// string -> float64
s3 := "45.67"
f, err := strconv.ParseFloat(s3, 64)
if err == nil {
    fmt.Println(f) // 45.67
}

// float64 -> string
f2 := 78.9
s4 := strconv.FormatFloat(f2, 'f', 2, 64) // "78.90"

// string -> bool
s11 := "true"
b4, err := strconv.ParseBool(s11)

// bool -> string
b5 := true
s12 := strconv.FormatBool(b5)

// string -> complex128
s13 := "1+2i"
comp, err := strconv.ParseComplex(s13, 128)
if err == nil {
    fmt.Println(comp) // (1+2i)
}
```

汇总表：

| 转换 | 函数 |
|------|------|
| string → int | `strconv.Atoi` |
| int → string | `strconv.Itoa` |
| string → float | `strconv.ParseFloat` |
| float → string | `strconv.FormatFloat` |
| string → bool | `strconv.ParseBool` |
| bool → string | `strconv.FormatBool` |
| string → 复数 | `strconv.ParseComplex`（`bitSize` 为 64 或 128） |

---

## 11. 使用 `fmt` 格式化

```go
// 打印结构体 / 值
fmt.Println(p)

// 格式化成字符串而不打印
c2 := complex(3, 4)
s14 := fmt.Sprintf("%v", c2) // "3+4i"
fmt.Println(s14)
```

`fmt.Println` 会以空格分隔各参数并换行输出。`fmt.Sprintf` 格式化方式相同，但返回字符串而不打印。

---

## 12. 速查表

- `var x T = v` —— 显式声明变量。
- `x := v` —— 短声明（类型自动推断；左侧需有新变量）。
- `const` —— 不可变的编译期常量。
- `for` —— 唯一的循环关键字，支持三段式、仅条件式、`range` 三种形式。
- `if` / `switch` / `defer` / `goto` / 带标签 `break` —— 控制流。
- `make(chan T)` —— 创建通道；`<-ch` 接收，`ch <- v` 发送。
- `&x` / `*p` —— 取地址与解引用。
- `T(v)` —— 显式转换（Go 无隐式数值转换）。
- `strconv.*` —— 在字符串与基本类型之间解析/格式化。
- `interface{}` + 类型断言 `v, ok := x.(T)` —— 处理任意类型。

---

*由 `01basic/02base/base.go` 生成。所有示例均直接取自源文件，因此可如所示编译运行。*
