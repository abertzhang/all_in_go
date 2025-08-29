package main

import "studyGo/13desgin_pattern/302command/01/pattern"



func main() {
	SonyTV()
}


func SonyTV(){
	tv:=pattern.TV{}
	onCommond:=pattern.OnCommand{Device:&tv}
	offCommond:=pattern.OffCommand{Device:&tv}
	button:=pattern.Button{}
	button.Command = &onCommond
	button.Press()
	button.Command = &offCommond
	button.Press()
}
