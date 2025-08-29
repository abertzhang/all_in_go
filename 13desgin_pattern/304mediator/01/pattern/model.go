package pattern

import "fmt"
type P1 struct {}
func (p *P1) GetMessage(data string)  {
	fmt.Println("P1 received:", data)
}

type P2 struct {}
func (p *P2) GetMessage(data string)  {
	fmt.Println("P2 received:", data)
}

type P3 struct {}
func (p *P3) GetMessage(data string)  {
	fmt.Println("P3 received:", data)
}
//
type Message struct {
	P1 *P1
	P2 *P2
	P3 *P3
}

func (m *Message) SendMessage(i interface{}, data string)  {
	switch i.(type) {
	case *P1:
		m.P2.GetMessage(data)
	case *P2:
		m.P1.GetMessage(data)
	case *P3:
		m.P1.GetMessage(data)
		m.P2.GetMessage(data)
	}
}
