package pattern
type IObserver interface {
	Do(*Subject)
}