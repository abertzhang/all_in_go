package main

import (
	"studyGo/13desgin_pattern/03builder/01/pattern"
)

func main() {
	ExampleBuilder()
}
func ExampleBuilder() {
	build := &pattern.Builder{}
	director := pattern.NewDirector(build)
	director.Construct()
}