package main

import (
	"fmt"
)

func main() {
	StringType()
}

func StringType() {
	s := `"Hello 
	
	world"`
	// we can use back ticks to wrap string that buddle new line
	fmt.Printf("%v", s)

	//conversion
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//each array is ASCII code

	fmt.Printf("%v\n", s[0])

	//conversion to binary, hex , Unicode 8
	uint8 b
}
