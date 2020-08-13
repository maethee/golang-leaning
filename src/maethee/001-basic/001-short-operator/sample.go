package main

import "fmt"

func main() {
	 name1 := "begin code"
	 fmt.Printf("%T : %v\n", name1, name1)

	 age := 50
	 fmt.Printf("%T : %v\n", age, age)

	 age = 40
	 fmt.Printf("%T : %v\n", age, age)

	 age2 := 30
	 fmt.Printf("%T : %v\n", age2, age2)

	 x, y := "string" , 66
	 fmt.Printf("%T : %v\n", x, x)
	 fmt.Printf("%T : %v\n", y, y)

}