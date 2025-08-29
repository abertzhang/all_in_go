package pattern

import "fmt"

type Do struct{
	APartDone bool
	BPartDone bool
	EndPartDone bool
}
type APart struct {
	next IDepartment
}

func (a *APart) Execute(d *Do) {
	if d.APartDone {
		fmt.Println("A 部分已完成")
		a.next.Execute(d)
		return
	}
	fmt.Println("A 部分开始处理")
	d.APartDone = true
	fmt.Println("A 部分已完成")
	a.next.Execute(d)
}

func (a *APart) SetNext(next IDepartment) {
	a.next = next
}
type BPart struct{
	next IDepartment
}
func (b *BPart) Execute(d *Do) {
	if d.BPartDone {
		fmt.Println("B 部分已完成")
		b.next.Execute(d)
		return
	}
	fmt.Println("B 部分开始处理")
	d.BPartDone = true
	fmt.Println("B 部分已完成")
	b.next.Execute(d)
}

func (b *BPart) SetNext(next IDepartment) {
	b.next = next
}

type EndPart struct{
	next IDepartment
}

func (e *EndPart) Execute(d *Do) {
	if d.EndPartDone {
		fmt.Println("结束部分已完成")
		return
	}
	fmt.Println("结束部分开始处理")
	d.EndPartDone = true
	fmt.Println("结束部分已完成")
}

func (e *EndPart) SetNext(next IDepartment) {
	e.next = next
}
