package pattern

import "fmt"
type Travel struct{
	Name string
	Strategy IStrategy
}
func (t *Travel) Traffic()  {
	t.Strategy.Traffic(t)
}
func (t *Travel) SetStrategy(strategy IStrategy)  {
	t.Strategy = strategy
}

func NewTravel(name string, strategy IStrategy) *Travel {
	return &Travel{
		Name:    name,
		Strategy: strategy,
	}
}


//
type Walk struct {}
func (w *Walk) Traffic(t *Travel)  {
	fmt.Println( t.Name+ "walk")
}
//
type Bus struct {}
func (b *Bus) Traffic(t *Travel)  {
	fmt.Println( t.Name+ "bus")
}