package main

import "studyGo/13desgin_pattern/304mediator/01/pattern"



func main() {
	ExampleMediator()

}

func ExampleMediator(){
	message:=pattern.Message{}
	p1:=&pattern.P1{}
	p2:=&pattern.P2{}
	p3:=&pattern.P3{}

	message.SendMessage(p1, "Hello from P1")
	message.SendMessage(p2, "Hello from P2")
	message.SendMessage(p3, "Hello from P3")
}
