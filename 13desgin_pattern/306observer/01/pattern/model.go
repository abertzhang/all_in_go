package pattern

import "fmt"
type Subject struct {
	Observers []IObserver
	Content   string
}
func (s *Subject) AddObserver(o IObserver) {
	s.Observers = append(s.Observers, o)
}
func (s *Subject) UpdateContext(content string) {
	s.Content = content
	s.notify()
}
func (s *Subject) notify() {
	for _, observer := range s.Observers {
		observer.Do(s)
	}
}

func NewSubject() *Subject {
	return &Subject{
		Observers: make([]IObserver, 0),
	}
}

//订阅者
type Reader struct {
	Name string
}

func (r *Reader) Do(s *Subject) {
	fmt.Printf("%s 收到更新: %s\n", r.Name, s.Content)
}
func NewReader(name string) *Reader {
	return &Reader{Name: name}
}