package pattern
//
type Base struct{}
func (b *Base) GetPrice() float64 {
	return 18.0
}
//西红柿馅料
type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() float64 {
	return t.Pizza.GetPrice() + 6.28
}

//芝士馅料
type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() float64 {
	return c.Pizza.GetPrice() + 8.98
}