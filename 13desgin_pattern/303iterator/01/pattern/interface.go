package pattern
type ICollection interface {
	CreateIterator() IIterator
}

type IIterator interface {
	HasNext() bool
	// Next() *Part
	Next() interface{}
}