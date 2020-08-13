package main

import (
	"fmt"
)

func main() {
	ShifOperator()
}

func ShifOperator() {
	// << left shift
	// >> right shift

	x := 4 
	fmt.Printf("%T\t%v\t%b\n", x, x, x)

	//0 = 0
	//1 = 1
	//2 = 10
	//3 = 11
	//4 = 100
	
	y := x << 2
	fmt.Printf("%T\t%v\t%b\n", y, y, y)


	z := x >> 1
	fmt.Printf("%T\t%v\t%b\n", z, z, z)
}
