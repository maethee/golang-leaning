package main

import "fmt"

func main() {
	ControlFlowLoop()
}

func ControlFlowLoop() {

	// x := 1

	// for x < 5 {
	// 	fmt.Printf("%T %v\n",x ,x)
	// 	x++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%T %v\n",i, i)
	// }

	for {
		fmt.Println("In Loop")
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Mod")
			continue
		}

		switch i {
			case 5:
				fmt.Println("Break")
				break
			default:
				fmt.Println("Default")
		}
		//below
		fmt.Println(i)
		break
	}
}
