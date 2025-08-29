package pattern

import "sync"
var (
	Instance *Singleton
	once sync.Once
)
type Singleton struct{
	Value int64
}
func (s *Singleton) GetValue()int64{
	return s.Value
}

type ISingleton interface{
	GetValue() int64
}
func GetInstance(v int64) Singleton{
	once.Do(func ()  {
		Instance=&Singleton{Value: v}
	})
	return *Instance
}