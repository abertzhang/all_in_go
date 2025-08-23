package pattern

import "fmt"

type Sms struct {
}

func NewSms() *Sms {
	return &Sms{}
}
func (s *Sms) Send(text, to string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, text)
}
type Email struct {
}

func NewEmail() *Email {
	return &Email{}
}
func (e *Email) Send(text, to string) {
	fmt.Printf("Sending Email to %s: %s\n", to, text)
}

type SystemA struct {
	Method ISendMessage
}

func NewSystemA(method ISendMessage) *SystemA {
	return &SystemA{Method: method}
}
func (s *SystemA) SendMessage(text, to string) {
	s.Method.Send(fmt.Sprintf("[System A] %s", text), to)
}
type SystemB struct {
	Method ISendMessage
}

func NewSystemB(method ISendMessage) *SystemB {
	return &SystemB{Method: method}
}
func (s *SystemB) SendMessage(text, to string) {
	s.Method.Send(fmt.Sprintf("[System B] %s", text), to)
}