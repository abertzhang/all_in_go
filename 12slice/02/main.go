package main

import (
	"fmt"
	"sort"
)	
func main() {
	intList:=[]int{5,3,45,523,12,34,45}
	sort.Ints(intList)
	fmt.Println(intList)
}