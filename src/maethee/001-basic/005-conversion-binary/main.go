package main

import (
	"fmt"
)

func main() {
	ConversionBinary()
}

func ConversionBinary() {
	s := `"Hello 
	
	world"`
	// we can use back ticks to wrap string that buddle new line
	fmt.Printf("%v", s)

	//conversion
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//each array is ASCII code

	//conversion to binary, hex , Unicode 8
	c := "H"
	cByte := []byte(c)
	fmt.Printf("%#x\n", c)
	fmt.Printf("%#U", cByte)
}
