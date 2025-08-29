package pattern
const (
	LeafNode = iota
	CompositeNode
)

type Component struct {}

func NewComponent(kind int, name string) IComponent {
	var c IComponent
	switch kind {
	case LeafNode:
		// c=NewLeaf(name)
	case CompositeNode:
		// c=NewComposite(name)
	}
	return c
}