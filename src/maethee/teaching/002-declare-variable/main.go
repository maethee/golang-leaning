package main

import "fmt"

var stock int = 100

func main() {
	addStock()
	DeclareVariable()
	removeStock()
	Access()
	fmt.Printf("Stock %T : %v \n", stock, stock)
}

func DeclareVariable() {

	fmt.Printf("Stock %T : %v \n", stock, stock)
}

func addStock(){
	stock++
}

func removeStock(){
	stock--
}