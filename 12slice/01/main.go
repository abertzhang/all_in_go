package main

import (
	"fmt"
	"slices"
)

func main() {
	a:=make([]int,5,6)
	var sliceA []int
	sliceA = append(sliceA,9,10 )
	fmt.Println(sliceA)
	fmt.Printf("%v---%T",a, a)
	arr:=[5]int{1,2,6,7}
	fmt.Println(arr)
	sliceB := arr[:]
	fmt.Println(sliceB)
	sliceC := append(sliceB, sliceA...)
	fmt.Println(sliceC)
}
