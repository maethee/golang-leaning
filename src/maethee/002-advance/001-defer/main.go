package main

import "fmt"

func main() {
	SampleDefer()
}

//Defer
func foo() {
	fmt.Println("foo")
}
func fooTwo() {
	fmt.Println("fooTwo")
}

// defer gonna run when function exit or surrounding func return
func SampleDefer() {
	defer foo()
	fooTwo()
}
