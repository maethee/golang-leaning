package main

import (
	"fmt"
)

func main() {
	DeclareVariable()
}

var initZ = 42

//or
// if you don't assige value the program will automaticly set a default value of each TYPE
// int will set zero value to initX
// string will be "" empty string
// floats will be 0.0
var initX int

//multi declaration and cannot change value
const (
	oneDeclare = 1
	twoDeclare = "hello"
)

func DeclareVariable() {
	z := "maethee"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// we can't assigh z = 42 (int) go is static programing languate
	// the Valuable can store a certain TYPE.
}
