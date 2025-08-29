package pattern

import "fmt"

//实现IImage接口
type RealImage struct {
	Filename string
}
func (i *RealImage) Display() {
	fmt.Printf("Displaying: %v\n", i.Filename)
}
func NewRealImage(filename string) *RealImage {
	//加载图片的耗时操作
	fmt.Printf("Loading: %v\n", filename)
	return &RealImage{Filename: filename}
}

//实现IImage接口
type ProxyImage struct {
	Real    *RealImage
	Filename string
}
func (p *ProxyImage) Display() {
	if p.Real == nil {
		p.Real = NewRealImage(p.Filename)
	}
	p.Real.Display()
}
func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{Filename: fileName}
}