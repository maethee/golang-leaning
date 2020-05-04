package main

import "fmt"

//try to declare each value (int,bool,string) and let see value and type
func main() {
	a := 42
	b := "Jame bond"
	c := true

	fmt.Println(a, b, c)

	s := fmt.Sprintf("%v\t%v\t\t%v", a, b, c)
	fmt.Println(s)
}
