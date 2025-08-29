package main

import (
	"fmt"
	"studyGo/13desgin_pattern/206flyweight/01/pattern"
)

func main() {
	viewer1 := pattern.NewColorViewer("blue")
	viewer2 := pattern.NewColorViewer("green")
	viewer3 := pattern.NewColorViewer("blue")
	fmt.Println(viewer1.ColorFlyweight.Data)
	fmt.Println(viewer2.ColorFlyweight.Data)
	fmt.Println(viewer3.ColorFlyweight.Data)
	fmt.Println(viewer3.ColorFlyweight == viewer1.ColorFlyweight)

}