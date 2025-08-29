package main

import "studyGo/13desgin_pattern/301chain/01/pattern"
func main(){
	ExampleChain()
}
func ExampleChain(){
	startPart := &pattern.EndPart{}
	aPart := &pattern.APart{}
	bPart := &pattern.BPart{}

	aPart.SetNext(startPart)
	bPart.SetNext(aPart)
	do := &pattern.Do{}
	// startPart.Execute(do)
	aPart.Execute(do)
}

func ExampleChain2() {
	startPart := &pattern.EndPart{}
	bPart := &pattern.BPart{}
	bPart.SetNext(startPart)

	aPart:=&pattern.APart{}
	aPart.SetNext(bPart)

	do := &pattern.Do{}
	aPart.Execute(do)
}