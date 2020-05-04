package main

import "fmt"

func main() {
	b := bar("Maethee")
	fmt.Println("bar return", b)

	m, ok := mouse("maethee", "cha", 1, 2, 3, 4)
	fmt.Println(m, ok)

		xi := []int{1, 2, 34, 5, 6, 7}
	r := sum(xi...)
	fmt.Println(r)
}

func bar(s string) string {
	fmt.Println("value", s)
	return fmt.Sprint("value\t", s)
}

// ...TYPE call variadic parameters
func mouse(f string, l string, i ...interface{}) (string, bool) {
	t := fmt.Sprint(f, l, i)
	return t, true
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}