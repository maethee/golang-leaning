package main

import (
	"fmt"
)

func main() {
	UntypeType()
}

func UntypeType() {

	// x variable , 4 untype int 
	x := 4.0

	fmt.Printf("%T\t%v\n", x, x)

	y := 4.0 //untype float64

	fmt.Printf("%T\t%v\n", y, y)

	sum := 4 + 4 // 8 untype int , 8.0 float64

	fmt.Printf("%T\t%v\n", sum, sum)

	var a int = 4 // convert to 4
	//var b float64 = 4.0

	sumTwo := 4.0 == a  // untype constant + type = 

	fmt.Printf("%T\t%v\n", sumTwo, sumTwo)
	
}
