package main

import "studyGo/13desgin_pattern/207proxy/01/pattern"
func main() {
ExampleProxy()
}

func ExampleProxy() {
	var sub pattern.ISubject= &pattern.Proxy{}
	sub.Proxy()
}
