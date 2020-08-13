package main

import (
	"fmt"
)

func main() {
	BitShift()
}

func BitShift() {
	fmt.Println("BitShift")
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 2
	fmt.Printf("%d\t\t%b\n", y, y)

	z := y >> 2
	fmt.Printf("%d\t\t%b\n", z, z)

	//var s uint = 1
	//var i = 1.0 << 2

	var o int = 2
	var p int = 3
	i := o == p
	fmt.Printf("%d\t\t%b\n", i, i)

	//Sample two 

	//untype 23, the binary is 10111  

	//left shift = 101110
	// 23 * 2 = 46
	y := 23 << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	//right shift = 10111
	// (23 - 1) / 2) = 11
	z := 23 >> 3
	fmt.Printf("%d\t\t%b\n", z, z)

	//var fl int = 4.0

	var r int = 3.1 + 3

	fmt.Printf("%T\t\t%v\n", r, r)

	// 1 = 1
	// 2 = 10
	// 3 = 11
	// 4 = 100
	// 5 = 101
	// 6 = 110
	
	x := 4
	var fl float64 = 4.0
	y := 4 == fl
	
	fmt.Printf("%T\t%v", y, y)

	sum := 1.0 + 1
	
	fmt.Printf("%T\t%v", sum, sum)

}

