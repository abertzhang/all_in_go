package main

import "studyGo/13desgin_pattern/310visitor/01/pattern"

func main() {
	ExampleShape()
}

func ExampleShape() {
	circle := pattern.NewCircle(5)
	rectangle := pattern.NewRectangle(4, 6)
	side := pattern.SideCalculator{}
	area := pattern.AreaCalculator{}

	circle.Accept(&side)
	rectangle.Accept(&side)
	//
	circle.Accept(&area)
	rectangle.Accept(&area)

}