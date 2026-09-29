package main

import (
	"fmt"
	"studyGo/13desgin_pattern/201adapter/01/pattern"
)

func main() {
	into := 10.56
	length, err := pattern.NewLengthAdapter().GetLength("m", into)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(length)
	//
	lengthMeter := pattern.NewLengthMeter().GetLength(into)
	fmt.Println(lengthMeter)
}