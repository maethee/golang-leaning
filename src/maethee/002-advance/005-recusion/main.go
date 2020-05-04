package main

import "fmt"

func main() {
	SampleRecursion()
}

// Recusion func
//
// ** we prefer loop if youcan, because recursive function consume the memory for than event loop
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

func SampleRecursion() {
	fmt.Println("SampleRecursion")
	x := factorial(5)
	// 5 * 4 * 3 * 2 * 1
	fmt.Println(x)

	y := loopFact(5)
	fmt.Println(y)
}
