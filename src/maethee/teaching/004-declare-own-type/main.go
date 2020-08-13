package main

import (
	"fmt"
)

type MyOwn int

func main() {

	var x MyOwn = 50
	var y int = 40

	y = int(x)

	fmt.Printf("%T %v \n", x, x)
	fmt.Printf("%T %v \n", y, y)

	Conversion(int32)
}

func Conversion(a int32){
	fmt.Printf("%T %v \n", a, a)
}