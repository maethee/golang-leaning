package main

import (
	"fmt"
)

func main() {
	ShortOperator()
}

func ShortOperator() {
	name := "Being code"
	fmt.Printf("%T : %v \n", name, name)

	age := 50
	fmt.Printf("%T : %v \n", age, age)

	age = 45
	fmt.Printf("%T : %v \n", age, age)

	ageTwo := 50
	fmt.Printf("%T : %v \n", ageTwo, ageTwo)

	x, y := "Begin", 30
	fmt.Printf("%T : %v \n", x, x)
	fmt.Printf("%T : %v \n", y, y)
}
