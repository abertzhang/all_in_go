package main

import (
	"studyGo/13desgin_pattern/07bridge/01/pattern"
)

func main() {
	ExampleBridge()
}

func ExampleBridge() {
	sms := pattern.NewSms()
	email := pattern.NewEmail()

	systemA := pattern.NewSystemA(sms)
	systemB := pattern.NewSystemB(email)

	systemA.SendMessage("Hello", "1234567890")
	systemB.SendMessage("Hello", "example@example.com")
}
