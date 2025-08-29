package main

import (
	"studyGo/13desgin_pattern/207proxy/02/pattern"
)
func main() {
	proxy := pattern.NewProxyImage("test.jpg")
	proxy.Display()
}