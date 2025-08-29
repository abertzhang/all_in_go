package main

import (
	"fmt"
	"studyGo/13desgin_pattern/102abstract/02/pattern"
)

func main() {
	factory := &pattern.ShopFactory{}
	pen := factory.CreatePen()
	book := factory.CreateBook()
	fmt.Println(pen.GetPen())
	fmt.Println(book.GetBook())
}
