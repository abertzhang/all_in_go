package pattern

import "fmt"

//管理类
type Director struct {
	builder IBuilder
}
func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()	
}

//
type Builder struct {}

func (b *Builder) Part1() {
	fmt.Println("建造 Part 1-完成")
}
func (b *Builder) Part2() {
	fmt.Println("建造 Part 2-完成")
}
func (b *Builder) Part3() {
	fmt.Println("建造 Part 3-完成")
}