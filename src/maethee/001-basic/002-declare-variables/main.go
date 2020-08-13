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

var stock int64 = 10000
const stockName = "Begin store"

//multi declaration and cannot change value
const (
	oneDeclare = 1
	twoDeclare = "hello"
	threeDeclare
)

var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func DeclareVariable() {
	//predeclar variable

	inStock := "23"
	fmt.Printf("%T: %v %v\n", inStock, inStock, &inStock)
	z := "maethee"
	fmt.Printf("%T: %v \n", z, z)
	// we can't assigh z = 42 (int) go is static programing languate
	// the Valuable can store a certain TYPE.

	var a float32
	fmt.Printf("%T: %v \n", a, a)

	var b int16
	fmt.Printf("%T: %v \n", b, b)

	var c string
	fmt.Printf("%T: %v \n", c, c)

	var d bool
	fmt.Printf("%T: %v \n", d, d)
	d = true

	//scope value
	{
		fmt.Printf("%T: %v \n", d, d)
		d := "scope value"
		fmt.Printf("%T: %v \n", d, d)
	}

	if(true){
		fmt.Printf("%T: %v \n", d, d)
		d := "scope value"
		fmt.Printf("%T: %v \n", d, d)
	}

	fmt.Printf(" %T: %v \n", threeDeclare, threeDeclare)

	const Θ float64 = 3/2  
	fmt.Printf(" %T: %v \n", Θ, Θ)

	const Π float64 = 3/2.
	fmt.Printf(" %T: %v \n", Π, Π)
}
