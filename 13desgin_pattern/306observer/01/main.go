package main

import "studyGo/13desgin_pattern/306observer/01/pattern"

func main() {
	ExampleObserver()
}

func ExampleObserver(){
	subject := pattern.NewSubject()
	boy := pattern.NewReader("小男孩")
	girl := pattern.NewReader("小女孩")
	subject.AddObserver(boy)
	subject.AddObserver(girl)
	subject.UpdateContext("今天的天气真好")

}