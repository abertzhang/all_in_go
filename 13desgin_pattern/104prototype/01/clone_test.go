package main

import (
	"testing"
)



type Type1 struct{
	name string
}
func (t *Type1) Clone() *Type1{
	tc:=*t
	return &tc
}

func TestClone(t *testing.T){
	t1:=&Type1{
		name: "type1",
	}
	t2:=t1.Clone()
	if t1==t2 {
		t.Fatal("error! get clone not working")
	} 
}