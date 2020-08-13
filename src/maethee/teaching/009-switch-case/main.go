package main

import "fmt"

func main() {
	switchStatements()
}

func switchStatements() {
	x := 4
	switch { // = true
	default:
		fmt.Println("Defaut")
	case x < 5:
		fmt.Println("Case one")
		fallthrough
	case x < 9:
		fmt.Println("Case two")
		fallthrough
	case true:
		fmt.Println("Case three")
	}
}
