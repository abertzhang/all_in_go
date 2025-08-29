package pattern

import "fmt"
//实现IShape接口
type Circle struct {
	Radius float64
}
func (c *Circle) Accept(visitor IVisitor) {
	visitor.VisitCircle(c)
}
func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}
//实现IShape接口
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor IVisitor) {
	visitor.VisitRectangle(r)
}
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

//周长--实现IVisitor接口
type SideCalculator struct {}

func (s *SideCalculator) VisitCircle(c *Circle) {
	fmt.Printf("Calculating sides for Circle with Radius: %f\n", c.Radius)
}
func (s *SideCalculator) VisitRectangle(r *Rectangle) {
	fmt.Printf("Calculating sides for Rectangle with Width: %f and Height: %f\n", r.Width, r.Height)
}
//面积--实现IVisitor接口
type AreaCalculator struct {}
func (a *AreaCalculator) VisitCircle(c *Circle) {
	fmt.Printf("Calculating area for Circle : %f\n", c.Radius*c.Radius*3.14)
}
func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	fmt.Printf("Calculating area for Rectangle : %f\n", r.Width*r.Height)
}