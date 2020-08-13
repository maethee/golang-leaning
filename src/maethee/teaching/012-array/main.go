package main

import "fmt"

func main() {
	SampleArray()
}

func SampleArray() {
	var x [5]int

	var y []int
	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", y, y)

	if y == nil {
		fmt.Println("Y is nil")
	}
}
