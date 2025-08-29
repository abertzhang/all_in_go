package pattern

import "fmt"

type Text struct {
	Content string
}

type TextMemento struct {
	Content string
}
func (t *Text) Write(content string)  {
	t.Content = content
}
func (t *Text) Save() IMemento {
	return &TextMemento{Content: t.Content}
}
func (t *Text) Load(m IMemento)  {
	tm:=m.(*TextMemento)
	t.Content = tm.Content
}
func (t *Text) Show()  {
	fmt.Println("Current Content:", t.Content)
}