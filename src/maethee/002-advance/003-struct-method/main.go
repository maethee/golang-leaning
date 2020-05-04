package main

import "fmt"

func main() {
	SampleMethod()
}

type ClientUser struct {
	ID    string
	title string
}

// This will attach getTitle to ClientUser Type
func (c ClientUser) getTitle() {
	fmt.Println("I'am ", c.title)
}

func SampleMethod() {
	client := ClientUser{"007", "Maethee"}
	client.getTitle()
}

func Anonymous() {
	func(x int) { // can be using with goroutine for doing something in block scope
		//statement
	}(42)
}
