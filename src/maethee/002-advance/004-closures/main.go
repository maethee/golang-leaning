package main

import "fmt"

func main() {
	SampleFunc()
}

//function return func() int
//useful like javascript when you want to preparing initial state to function that you want to execute
func FuncReturnFunc() func(int) int {
	// you can create initial state on this
	// or when you create share function each usecase has difference initial state
	return func(x int) int {
		return x + 52
	}
}

func FuncPassFunctpArgs(f func() func(int) int, x int) int {
	return f()(x)
}

var xGlobal int = 500

func SampleFunc() {
	fmt.Println("SampleFunc")
	x := FuncReturnFunc()
	y := x(2)
	fmt.Println("Value from func ", y)

	z := FuncPassFunctpArgs(FuncReturnFunc, 10)
	fmt.Println("Value from func that pass func in to it ", z)

	//code block we can define block of scope data like below
	//y in block scope can't share out site
	{
		y := 55
		fmt.Printf("Scope of y in code block is %d  global %d parent scope %d \n", y, xGlobal, z)
	}
}
