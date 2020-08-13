package main

import (
	"fmt"
)

func main() {
	//typeSwitch()
	variousInterface()
}

func typeSwitch() {

	var x interface{} = 3
	y, isString := x.(int)
	fmt.Printf("%T %v \n", x, x)
	if isString {
		fmt.Printf("%T %v \n", y, y)
	}

	var name string = string(y)
	fmt.Printf("%T %v \n", name, name)
}

type Custom int

func variousInterface() {
	//var a Custom = 3
	var x interface{} = nil
	switch i := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case string:
		fmt.Println("string", i)
	default:
		fmt.Printf("default %T %v", i, i)
	}
}
