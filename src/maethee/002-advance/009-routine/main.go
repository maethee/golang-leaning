package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	SampleRuntime()
}

//runtime
//Concurrency and paralleliem
func oneRoutine() {
	fmt.Println("oneRoutine ")
	for i := 0; i <= 10; i++ {
		fmt.Println("oneRoutine ", i)
	}

	wg.Done()
}

func twoRoutine() {
	fmt.Println("twoRoutine ")
	for i := 0; i <= 10; i++ {
		fmt.Println("twoRoutine ", i)
	}
}

//any func can access wg
var wg sync.WaitGroup

func SampleRuntime() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("CPUS ", runtime.NumCPU())
	fmt.Println("Gorutines ", runtime.NumGoroutine())

	wg.Add(1) // amunt of goroutine we have to wait
	//seperate thread to execute
	go oneRoutine()

	//two will execute first because in this scope
	twoRoutine()

	fmt.Println("CPUS ", runtime.NumCPU())
	fmt.Println("Gorutines ", runtime.NumGoroutine())

	//wainting signal from go routine
	wg.Wait()
}
