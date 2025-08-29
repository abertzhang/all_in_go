package main

import (
	"fmt"
	"studyGo/13desgin_pattern/205facade/01/pattern"
)

func main() {
	api := pattern.NewApiRun()
	fmt.Println(api.Test())
}


