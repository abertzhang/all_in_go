package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString("abc123def456", -1))
}
