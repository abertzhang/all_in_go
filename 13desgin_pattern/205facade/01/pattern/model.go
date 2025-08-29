package pattern

import "fmt"

//实现IApiA接口
type ApiRunA struct{}
func (a *ApiRunA) TestA() string {
	return "TestA"
}

func NewApiA() IApiA {
	return &ApiRunA{}
}
//实现IApiB接口
type ApiRunB struct{}
func (b *ApiRunB) TestB() string {
	return "TestB"
}

func NewApiB() IApiB {
	return &ApiRunB{}
}

type ApiRun struct{
	A IApiA
	B IApiB
}
func (r *ApiRun) Test() string {
	aRet := r.A.TestA()
	bRet := r.B.TestB()
	return fmt.Sprintf("%s---%s", aRet, bRet)
}

func NewApiRun() IApi {
	return &ApiRun{
		A: NewApiA(),
		B: NewApiB(),
	}
}