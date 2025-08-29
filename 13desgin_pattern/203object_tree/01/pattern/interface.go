package pattern
type IComponent interface {
	Parent() IComponent
	SetParent(IComponent)
	Name() string
	SetName(string)
	AddChild(IComponent)
	Search(string)
}