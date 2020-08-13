package main

import (
	"fmt"
)

func main() {
	StringType()
}

func StringType() {
	var s string = "Begin"
	fmt.Printf("%T %v\n", s, s)

	fmt.Printf("%T %v\n", len(s), len(s))

	//0 - 255
	fmt.Printf("%T %v\n", s[0], s[0])
	fmt.Printf("%T %v\n", s[len(s)-1], s[len(s)-1])

	fmt.Printf("%v\n", &s[0])
}
