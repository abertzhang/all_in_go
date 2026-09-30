--
title:Golang 结构体
create:2026-09-30
update:2026-09-30
category:Golang
tags:[golang,struct,结构体]
summary:结构体定义与初始化、字段读写、嵌套与嵌入、方法与值/指针接收者、匿名结构体、泛型结构体。
top:1
copyright:true
--

# Go 语言学习笔记：结构体

> 整理自掘金文章《[Go语言第六章(结构体)](https://juejin.cn/post/7670398642598019118)》（作者：小满zs）。
> 结构体是 Go 中的一种**复合数据类型**，用于将多个不同类型的数据组合在一起，形成一个新的数据类型。

日常开发里，用结构体描述「一个人」「一辆车」「一条订单」这类对象非常常见。

---

## 一、定义结构体

用 `type` + `struct` 定义一种新类型：

```go
type 结构体名 struct {
	字段名 类型
	字段名 类型
	字段名 类型
}
```

实操：定义一个 `Person`，包含姓名、年龄和爱好：

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string
}

func main() {
	person := Person{
		Name:  "John",
		Age:   20,
		Hobby: []string{"reading", "swimming"},
	}
	fmt.Println(person)
}
```

`Person{...}` 是**字面量初始化**：按字段名赋值，顺序可以随便写，没写的字段会是该类型的零值。

### 1.1 字段读写

通过 `.` 访问和修改字段：

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "张三", Age: 18}
	fmt.Println(p.Name) // 张三

	p.Age = 20
	fmt.Println(p.Age) // 20
}
```

### 1.2 几种初始化方式

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. 按字段名初始化（推荐，可读性好）
	p1 := Person{Name: "张三", Age: 18}

	// 2. 按字段顺序初始化（字段多时容易写错，不推荐）
	p2 := Person{"李四", 20}

	// 3. 只给部分字段赋值，其余为零值
	p3 := Person{Name: "王五"} // Age 为 0

	// 4. 先声明再赋值（全部为零值）
	var p4 Person
	p4.Name = "赵六"
	p4.Age = 22

	// 5. 取地址，得到 *Person
	p5 := &Person{Name: "钱七", Age: 25}

	fmt.Println(p1, p2, p3, p4, p5)
}
```

指针形式 `p5` 访问字段时，Go 会**自动解引用**，写 `p5.Name` 即可，不必写 `(*p5).Name`。

---

## 二、结构体嵌套

一个结构体里可以包含另一个结构体类型的字段，用来组合更复杂的数据：

```go
package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

type Person struct {
	Name  string
	Age   int
	Hobby []string
	Car   Car
}

func main() {
	person := Person{
		Name:  "John",
		Age:   20,
		Hobby: []string{"reading", "swimming"},
		Car: Car{
			Brand: "Toyota",
			Model: "Camry",
			Year:  2020,
		},
	}
	fmt.Println(person)
	fmt.Println(person.Car.Brand) // Toyota
}
```

这里 `Car` 是**命名字段**，访问时要写 `person.Car.Brand`。

---

## 三、结构体嵌入

如果字段类型前**不写字段名**，就是嵌入（匿名字段）。嵌入后，被嵌入类型的字段和方法会「提升」到外层，可以直接用：

```go
package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

type Person struct {
	Name string
	Age  int
	Car        // 嵌入：没有字段名，只有类型
}

func main() {
	person := Person{
		Name: "John",
		Age:  20,
		Car: Car{
			Brand: "Toyota",
			Model: "Camry",
			Year:  2020,
		},
	}

	// 提升后可以直接访问
	fmt.Println(person.Brand) // Toyota
	// 也可以走完整路径
	fmt.Println(person.Car.Model) // Camry
}
```

> **嵌套 vs 嵌入**：嵌套是「有名字的组合」；嵌入是「把另一个类型的能力提升过来」。两者都能组合数据，嵌入更适合表达「is-a / 带有某种能力」的关系。

---

## 四、方法

方法是「绑在某个类型上的函数」。接收者写在 `func` 和函数名之间：

```go
func (接收者 类型) 方法名(参数列表) 返回值 {
	// ...
}
```

### 4.1 值接收者

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("我叫 %s，今年 %d 岁\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "张三", Age: 18}
	p.Introduce() // 我叫 张三，今年 18 岁
}
```

值接收者拿到的是**副本**，方法里改字段**不会**影响原变量。

### 4.2 指针接收者

需要修改结构体本身时，用指针接收者：

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) GrowUp() {
	p.Age++
}

func (p Person) Introduce() {
	fmt.Printf("我叫 %s，今年 %d 岁\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "张三", Age: 18}
	p.GrowUp()      // 等价于 (&p).GrowUp()
	p.Introduce()   // 我叫 张三，今年 19 岁
}
```

调用时写 `p.GrowUp()` 即可，Go 会**自动取地址**。

> **习惯**：要改状态用指针接收者；只读、小结构体可以用值接收者。同一类型的方法最好**统一风格**，不要有的用值、有的用指针混着来。

---

## 五、匿名结构体

不提前 `type` 定义，直接在用的地方写 `struct { ... }`，适合只临时用一次的场景：

```go
package main

import "fmt"

func main() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  20,
	}
	fmt.Println(person.Name, person.Age)
}
```

配置、单次返回、测试数据等「用完就丢」的结构，用匿名结构体很方便。若多处复用，还是起个名字更清晰。

---

## 六、泛型结构体

Go 1.18 起，结构体也可以带类型参数：

```go
package main

import "fmt"

type Person[T string | int] struct {
	Name  string
	Age   int
	Phone T
}

func main() {
	person := Person[string]{
		Name:  "John",
		Age:   20,
		Phone: "010-10086",
	}
	person2 := Person[int]{
		Name:  "John",
		Age:   20,
		Phone: 10086,
	}
	fmt.Println(person)
	fmt.Println(person2)
}
```

`Person[T string | int]` 表示 `Phone` 只能是 `string` 或 `int`。创建时要写明具体类型，如 `Person[string]`、`Person[int]`。

---

## 七、练习题

1. **定义与初始化**：定义 `Book` 结构体（书名、作者、页数），用至少两种方式创建实例并打印。
2. **嵌套**：定义 `Address`（城市、街道）和 `User`（姓名、地址），创建用户并打印「姓名住在城市的街道」。
3. **嵌入**：把上一题的 `Address` 改成嵌入到 `User`，用提升后的字段直接打印城市。
4. **方法**：给 `Book` 写值接收者方法 `Info()` 打印信息；再写指针接收者方法 `AddPages(n int)` 增加页数，调用后验证页数有变。
5. **综合**：定义 `Rectangle`（宽、高），写方法 `Area()` 和 `Perimeter()` 返回面积和周长，并在 `main` 里验证。

---

## 八、核心要点速记

- **定义**：`type Name struct { ... }`，用 `type` + `struct` 声明新类型。
- **初始化**：推荐**按字段名**初始化（可读性好、可不写全）；按顺序初始化易错。
- **访问**：用 `.` 读写字段；指针结构体访问字段会自动解引用（`p.Name` 无需 `(*p).Name`）。
- **嵌套**：命名字段，走完整路径 `person.Car.Brand`。
- **嵌入**：匿名字段，字段/方法**提升**到外层，可直接 `person.Brand`。
- **方法**：接收者写在 `func` 与方法名之间；**值接收者**拿到副本（不改原值），**指针接收者**可修改原值；同类型方法风格要统一。
- **匿名结构体**：一次性的数据结构，就地定义。
- **泛型结构体**：Go 1.18+ 支持 `type Name[T Constraint] struct { ... }`。

---

*参考资料：[Go语言第六章(结构体)](https://juejin.cn/post/7670398642598019118) · 作者：小满zs*
