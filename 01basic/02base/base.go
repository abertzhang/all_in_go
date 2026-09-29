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
func variableExamples() {
    // Declare and initialize variables
    var name string = "Alice"
    var age  = 30
    gender := "Female"
    from,career := "USA", "Engineer"    
    fmt.Println(name,age,gender,from,career)
    // Declare and initialize constants
    const pi = 3.14
    const (
        e = 2.71
        g = 9.81
    )
    fmt.Println(pi,e,g)
    // Declare and initialize numeric variables
    // int8, int16, int32, int64,int, uint8, uint16, uint32, uint64,uint,uintptr 
    // float32, float64, complex64, complex128
    var a int = 10
    var b float64 = 20.5
    c := 30
    fmt.Println(a,b,c)
    // Declare and initialize boolean variables
    var flag bool = true
    isInit := false
    fmt.Println(flag,isInit) 
    // Declare and initialize array variables
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)
    // Declare and initialize slice variables
    var slice []int = []int{4, 5, 6}
    fmt.Println(slice)
    // Declare and initialize map variables
    var m map[string]int = map[string]int{"one": 1, "two": 2}
    fmt.Println(m)
    // Declare and initialize struct variables
    type Person struct {
        Name string
        Age  int
    }
    var p Person = Person{Name: "Bob", Age: 25}
    fmt.Println(p)
    // Declare and initialize pointer variables
    var ptr *int = &a
    fmt.Println(*ptr)
    // Declare and initialize interface variables
    var i interface{} = "Hello"
    fmt.Println(i)
    // Declare and initialize function variables
    var f func(int, int) int = func(x, y int) int {
        return x + y
    }
    fmt.Println(f(2, 3))
    // Declare and initialize channel variables
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    fmt.Println(<-ch)
    // Declare and initialize nil variables
    // int-0, float-0.0, bool-false, string-"",array-nil, slice-nil, 
    // map-nil, pointer-nil, interface-nil, function-nil, channel-nil
    var n *int = nil
    fmt.Println(n)
    // Declare and initialize empty variables
    var emptySlice []int = []int{}
    var emptyMap map[string]int = map[string]int{}
    fmt.Println(emptySlice, emptyMap)
    // Declare and initialize zero value variables
    var zeroInt int
    var zeroFloat float64
    var zeroBool bool
    var zeroString string
    fmt.Println(zeroInt, zeroFloat, zeroBool, zeroString)
    // Declare and initialize multiple variables in a single line
    var x, y, z int = 1, 2, 3
    fmt.Println(x, y, z)
}

func loopExamples() {
    // if-else ,if-else if-else
   score := 85
   if score >= 80 {
       fmt.Println("Grade: A")
   }  else if score >= 60 {
       fmt.Println("Grade: B")
   } else {
       fmt.Println("Grade: C")
   }
   // switch-case-fallthrough
   day := "Monday"
   switch day {
   case "Monday":
       fmt.Println("Start of the work week")
       fallthrough
   case "Friday":
       fmt.Println("End of the work week")
   default:
       fmt.Println("Midweek")
   }
   // for loop
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }
       j := 0
    for j < 5 {
        fmt.Println(j)
        j++
    }
   // for range loop
   // for range loop over an array or slice or map or string
   arr := []int{1, 2, 3, 4, 5}
   for index, value := range arr {
       fmt.Println(index, value)
   }
       for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
    person := map[string]string{
		"name": "Alice",
		"city": "New York",
	}
	for key, value := range person {
		fmt.Println(key, value)
	}
    text := "你好Go"
	for index, char := range text {
		fmt.Printf("下标 %d,字符 %c\n", index, char)
	}
   // break and continue
   // break with tag
   for i := 0; i < 10; i++ {
       if i == 5 {
           break
       }
       if i%2 == 0 {
           continue
       }
       fmt.Println(i)
   }
   Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break Outer
			}
			fmt.Println(i, j)
		}
	}
   // goto statement
   i := 0
   here:
       fmt.Println(i)
       i++
       if i < 5 {
           goto here
       }
    // defer statement
    defer fmt.Println("This will be printed last")
    fmt.Println("This will be printed first")
    // while loop using for


}

func operatorExamples() {
   // && || !
   // + - * / %
   // & | ^ &^ << >>
   // = += -= *= /= %= <<= >>= &= ^= |=
   a := 10
   b := 3
   flag, isInit := true, false
   fmt.Println(a + b) //13
   fmt.Println(a - b) //7
   fmt.Println(a * b) //30
   fmt.Println(a / b) //3
   fmt.Println(a % b) //1
   fmt.Println(a & b) //2
   fmt.Println(a | b) //11
   fmt.Println(a ^ b) //9
   fmt.Println(a &^ b) //8
   fmt.Println(a << 1) //20
   fmt.Println(a >> 1) //5
   fmt.Println(a == b) //false
   fmt.Println(a != b) //true
   fmt.Println(a > b) //true
   fmt.Println(a < b) //false
   fmt.Println(a >= b) //true
   fmt.Println(a <= b) //false
   fmt.Println(flag && isInit) //false
   fmt.Println(flag || isInit) //true
   fmt.Println(!flag) //false
   c := a + b
   fmt.Println(c) //13
   c += 2
   fmt.Println(c) //15 
   c -= 3
   fmt.Println(c) //12
   c *= 2
   fmt.Println(c) //24
   c /= 4
   fmt.Println(c) //6
   c %= 5
   fmt.Println(c) //1
   c <<= 1
   fmt.Println(c) //2
   c >>= 1
   fmt.Println(c) //1
   c &= 3
   fmt.Println(c) //1
   c |= 2
   fmt.Println(c) //3
   c ^= 1
   fmt.Println(c) //2
   c &^= 1
   fmt.Println(c) //2
}

func convertExamples() {
    // int(),int8(),int16(),int32(),int64()
    //uint(),uint8(),uint16(),uint32(),uint64()
    // float32(), float64()
    // string(), []byte()
    a := 123
    b := 45.67
    c := "hello"
    d := []byte("world")
    fmt.Println(int32(a))
    fmt.Println(float64(b)) 
    fmt.Println(c)
    fmt.Println(d)
   // string to int
   s := "123"
   i, err := strconv.Atoi(s)
   if err != nil {
       fmt.Println(err)
   } else {
       fmt.Println(i) //123
   }
   // int to string
   i2 := 456
   s2 := strconv.Itoa(i2)
   fmt.Println(s2) //"456"
   // string to float
   s3 := "45.67"
   f, err := strconv.ParseFloat(s3, 64)
   if err != nil {
       fmt.Println(err)
   } else {
       fmt.Println(f) //45.67
   }
   // float to string
   f2 := 78.9
   s4 := strconv.FormatFloat(f2, 'f', 2, 64)
   fmt.Println(s4) //"78.90"
   // string to []byte
   s5 := "hello"
   b2 := []byte(s5)
   fmt.Println(b2) // [104 101 108 108 111] 
   // []byte to string
   b3 := []byte{119, 111, 114, 108, 100}
   s6 := string(b3)
   fmt.Println(s6) //"world"
   // string to rune
   s7 := "hello"
   r := []rune(s7)
   fmt.Println(r) // [104 101 108 108 111]
   // rune to string
   s8 := string(r)
   fmt.Println(s8) //"hello"
   // string to interface{}
   s9 := "hello"
   var in interface{} = s9
   fmt.Println(i) //"hello"
   // interface{} to string
   s10, ok := in.(string)
   if ok {
       fmt.Println(s10) //"hello"
   }
   // string to bool
   s11 := "true"
   b4, err := strconv.ParseBool(s11)
   if err != nil {
       fmt.Println(err)
   } else {
       fmt.Println(b4) //true
   }    
   // bool to string
   b5 := true
   s12 := strconv.FormatBool(b5)
   fmt.Println(s12) //"true"
   // bool to int
   b6 := true
   i3 := 0
   if b6 {
       i3 = 1
   }
   fmt.Println(i3) //1
   // int to bool
   i4 := 1
   b7 := false
   if i4 != 0 {
       b7 = true
   }
   fmt.Println(b7) //true
   // float to int
   f3 := 12.34
   i5 := int(f3)
   fmt.Println(i5) //12
   // int to float
   i6 := 12
   f4 := float64(i6)
   fmt.Println(f4) //12.0
   // float to bool
   f5 := 0.0
   b8 := false
   if f5 != 0 {
       b8 = true
   }
   fmt.Println(b8) //false
   // bool to float
   b9 := true
   f6 := 0.0
   if b9 {
       f6 = 1.0
   }
   fmt.Println(f6) //1.0
   // string to complex128
   s13 := "1+2i"
   // 注意：本作用域内 c（string）与 err 已声明，故此处用新变量名 comp，避免 := 无新变量报错
   comp, err := strconv.ParseComplex(s13, 128)
   if err != nil {
       fmt.Println(err)
   } else {
       fmt.Println(comp) //(1+2i)
   }
   // complex128 to string
   c2 := complex(3, 4)
   s14 := fmt.Sprintf("%v", c2)
   fmt.Println(s14) //"3+4i"        
}