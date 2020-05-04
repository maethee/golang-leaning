package main

import (
	"fmt"
)

func main() {
	SampleChannel()
}

func doSomeThing(n int) int {
	return n * 40
}

// channel is small piece you can send the data to it.
func funcSendChan(c chan<- int) {
	fmt.Println("funcSendChan")
	c <- 500
}

func funcReceiveChan(c <-chan int) {
	fmt.Println("funcReceiveChan", <-c)

}

func SampleChannel() {
	fmt.Println("SampleChannel")

	ch := make(chan int, 1)

	/*
		can not declare and send the data immediatly like below
		ch := make(chan int)
		ch <- 40 like send the data hand to hand can not happend at same time

		but if you allowcate size of chan you can send the data to this block in below
		make(chan int, 1)
	*/

	go func() {
		ch <- doSomeThing(4)
	}()

	//print receive data
	fmt.Println(<-ch)
	fmt.Println("End")

	cr := make(<-chan int) // receive channel
	cs := make(chan<- int) // send channel

	// both above you can not assign the value like cr <- 4 or cs <- 4 golang will inform you invalid
	// but when cr = ch ,ch is nomally chan  golang allow to do that
	// but ch = cr is not allow
	//cr = ch
	//ch = cs

	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// pass chan to other func and using goroutine
	go funcSendChan(ch)

	// try to using goroutine does not woking like below
	// wrong: go funcReceiveChan(ch)

	// it can receive in main thread
	funcReceiveChan(ch)
	fmt.Println("Done")
}
