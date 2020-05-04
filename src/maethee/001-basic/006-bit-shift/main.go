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
}
