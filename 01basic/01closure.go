package main

import "fmt"

func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		fmt.Println(x)
		return x
	}
}

func main() {
	a := adder()
	a(2)
	a(4)
	a(100)
}
