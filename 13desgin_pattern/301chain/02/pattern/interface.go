package pattern
type IHandler interface {
	//自身业务
	Do(*Context) error
	//设置下一个对象
	SetNext(IHandler) IHandler
	//执行
	Run(*Context) error
}