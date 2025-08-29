package main

import "studyGo/13desgin_pattern/305memento/01/pattern"

func main() {
	ExampleText()
}

func ExampleText() {
	text := &pattern.Text{}
	text.Write("How are you?")
	text.Show()
	memento := text.Save()
	text.Write("Hello, Go!")
	text.Show()
	text.Load(memento)
	text.Show()
}