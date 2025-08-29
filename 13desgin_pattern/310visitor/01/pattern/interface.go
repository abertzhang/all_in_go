package pattern
type IShape interface {
	Accept(visitor IVisitor)
}

type IVisitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}