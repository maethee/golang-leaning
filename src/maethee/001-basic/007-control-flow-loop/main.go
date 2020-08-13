package main

import (
	"fmt"
)

func main() {
	ControlFlowLoop()
}

func ControlFlowLoop() {
	// for init statement; condition; post
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// print unicode
	// reference: https://golang.org/pkg/fmt/
	for j := 1; j <= 100; {
		fmt.Printf("%#U\n", j)
		j++;
	}


	//while loop
	for {
		fmt.Println("Loop")
		break;
	}

	for n := 0; n <= 4; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)

		//break in switch
		switch n {
		case 5: // skip space
			fmt.Printf("break in loop %d\n", n)
			break
		default:
			fmt.Printf("%c\n", n)
		}
	}
}
