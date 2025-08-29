package pattern

import "fmt"

//实现ISubject接口
type Proxy struct {
	Real RealSubject
}

func (p *Proxy) Proxy() string {
	var res string
	//在调用真实对象之前，检查缓存，判断权限，等等
	p.Real.Real()
	// 调用真实对象
	p.Real.Pre()
	// 调用之后的操作，如缓存结果，对结果进行处理，等等
	p.Real.After()
	return res
}

type RealSubject struct{}
func (r *RealSubject) Real()  {
	fmt.Println("real")
}
func (r *RealSubject) Pre()  {
	fmt.Println("pre")
}
func (r *RealSubject) After()  {
	fmt.Println("after")
}