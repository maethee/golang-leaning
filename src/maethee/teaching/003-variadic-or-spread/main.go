package main

import (
	"fmt"
)

func main() {
	s := []string{"Joe", "Anna", "Eileen"}
	prefix := "Prefix"
	Greeting(prefix, s...)
	fmt.Printf("%T : %v \n", prefix, prefix)	
}

func Greeting(prefix string, who ...string) {
	fmt.Printf("%T : %v \n", prefix, prefix)	
	prefix = "Bee"
}	
