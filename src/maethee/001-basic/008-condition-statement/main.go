package main

import "fmt"

func main() {
	ConditionStatement()
}

func f() int {
	return 43
}

func conSwitch(n int) int {
	fmt.Println(n)
	return n
}

func ConditionStatement() {
	// condition
	if true {
		fmt.Println("first")
		fmt.Println("two")
	}

	// initate; condition
	if x := 42; x == 42 {
		fmt.Println("done")
	}

	//you can call func in initial state
	if y := f(); y == 42 {
		fmt.Println("done")
	} else if y == 43 {
		fmt.Println("value is 43")
	} else {
		fmt.Println("other value")
	}

	switch {
	case false:
		fmt.Println("case false")
	case (2 == 3):
		fmt.Println("case 2 == 3")
	case (2 == 2):
		fmt.Println("case 2 == 2")
		//tranfer control to next case

		fallthrough
	case false:
		fmt.Println("case 2 == 3")
	case (4 == 4):
		fmt.Println("case 4 == 4")
	default:
		fmt.Println("default")
	}

	// you can push func to switch condition
	switch conSwitch(4) {
	case conSwitch(1), conSwitch(2), conSwitch(3):
		fmt.Println("conSwitch")
	default:
		fmt.Println("default")
	}
}
