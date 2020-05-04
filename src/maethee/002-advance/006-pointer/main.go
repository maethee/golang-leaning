package main

import "fmt"

func main() {
	SamplePointer()
}

//we get a address of value
func ChangeValueByAddress(y *int) {
	*y = 55
}

func SamplePointer() {
	x := 40
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\n", &x)

	var y *int = &x
	fmt.Println("address of value ", y)
	fmt.Println("value of address ", *y)

	// & giva you a address
	// * before variable give you a value
	// * before TYPE is a pointer

	//y is address and we will put value to this address  for x and y that using same address
	// the value x and y will be the same
	*y = 45
	fmt.Println("value of address x ", x)
	fmt.Println("value of address y ", *y)

	c := 0
	//we pass address or pointer to this
	ChangeValueByAddress(&c)
	fmt.Println("value of c ", c)
}
