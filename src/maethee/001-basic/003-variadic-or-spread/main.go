package main

import (
	"fmt"
)

func main(){
	UnlimitArgs("Unlimit", 42, true)
}

// ... is unlimit args and interface is suport any type of the data, or empty interface
func UnlimitArgs(a ...interface{}) {
	fmt.Println(a...)
}
