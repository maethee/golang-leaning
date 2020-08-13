package main

import "fmt"

func main() {
	ConditionStatement()
}

func ConditionStatement() {

	// x := 3
	// max := 2

	// if x > max { // true
	// 	fmt.Printf("%T %v\n", x, x)
	// 	fmt.Printf("%T %v\n", max, max)
	// }

	// if false {
	// 	fmt.Println("True")
	// }

	if x := f(); x < 20 {
		fmt.Println("x < 50")
	} else if x == 44 {
		fmt.Println("x == 44")
	} else {
		fmt.Println("nothing match")
	}
}

func f() int {
	return 44
}
