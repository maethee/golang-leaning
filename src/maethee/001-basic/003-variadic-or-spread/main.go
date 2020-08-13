package main

import (
	"fmt"
)

func main(){
	UnlimitArgs("Unlimit", 42, true)

	//variadic with parameter 
	arr := []int{1, 2, 3}
	arrAppend := []int{4, 5, 6}

	//code below got same result
	//arr = append(arr, 4, 5, 6)
	arr = append(arr, arrAppend...)

	fmt.Printf("%T : %v", arr, arr)
}

// ... is unlimit args and interface is suport any type of the data, or empty interface
func UnlimitArgs(a ...interface{}) {
	fmt.Println(a...)
}
