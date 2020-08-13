package main

import "fmt"

func main() {
	typeSwitch()
}

func typeSwitch() {

	var x interface{} = 3.1
	switch i := x.(type) {
	case nil:
		fmt.Println("Nil", i)
	case int:
		//block
		fmt.Println("Int", i)
	case string:
		//block
		fmt.Println("String", i)
	default:
		//block
		fmt.Println("Default", i)
	}
}
