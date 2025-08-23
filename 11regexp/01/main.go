package main

import (
	"fmt"
	"regexp"
)
func main() {
	match,_:=regexp.MatchString(`H.*!`, "Hello world")
	if match {
		println("Matched!")
	} else {
		println("Not matched!")
	}
	reg:=regexp.MustCompile(`foo.?`)

	fmt.Printf("%q\n", reg.FindString("seafood fool"))
}
 