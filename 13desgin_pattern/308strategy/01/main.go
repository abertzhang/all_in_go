package main

import "studyGo/13desgin_pattern/308strategy/01/pattern"



func main() {
	ExampleStrategy()
}
func ExampleStrategy() {
	walk := &pattern.Walk{}
	travel1 := pattern.NewTravel("小明", walk)
	travel1.Traffic()
	//
	bus := &pattern.Bus{}
	travel2 := pattern.NewTravel("小红", bus)
	travel2.Traffic()
	travel2.SetStrategy(walk)
	travel2.Traffic()
}