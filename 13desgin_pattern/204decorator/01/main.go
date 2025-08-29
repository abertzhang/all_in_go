package main

import (
	"fmt"
	"studyGo/13desgin_pattern/204decorator/01/pattern"
)			

func main() {
	var pizza pattern.IPizza
	pizza = &pattern.Base{}
	fmt.Println("Base pizza price:", pizza.GetPrice())

	pizza = &pattern.TomatoTopping{Pizza: pizza}
	fmt.Println("Pizza with tomato topping price:", pizza.GetPrice())

	pizza = &pattern.CheeseTopping{Pizza: pizza}
	fmt.Println("Pizza with cheese topping price:", pizza.GetPrice())
}
