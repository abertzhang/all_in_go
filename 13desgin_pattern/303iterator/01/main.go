package main

import (
	"fmt"
	"studyGo/13desgin_pattern/303iterator/01/pattern"
)
func main() {
  ExampleIterator()
}

func ExampleIterator(){
	part1:=pattern.Part{Title:"part1",Number:10}	
	part2:=pattern.Part{Title:"part2",Number:20}
	part3:=pattern.Part{Title:"part3",Number:30}
	collection := pattern.PartCollection{Parts: []*pattern.Part{&part1, &part2, &part3}}
	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		part := iterator.Next().(*pattern.Part)
		fmt.Printf("Title: %s, Number: %d\n", part.Title, part.Number)
	}
}