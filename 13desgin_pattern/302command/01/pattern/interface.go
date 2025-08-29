package pattern
type ICommand interface {
	Execute()
}
type IDevice interface {
	On()
	Off()
}