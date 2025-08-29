package main

import "fmt"
func main() {
	f := PhoneDecorate(func(phone string) string {
		return phone
	})
	s := f("iPhone")
	fmt.Println(s)
}

type PlusPhone func(phone string) string
func PhoneDecorate(fn PlusPhone) PlusPhone {
	return func(phone string) string {
		phone=phone+"plus"
		return fn(phone)
	}
}

func CasePhoneDecorate(fn PlusPhone) PlusPhone {
	return func(phone string) string {
		phone=phone+"case"
		return fn(phone)
	}
}