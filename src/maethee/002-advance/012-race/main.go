package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	SampleRace()
}

// access share value at the sametime
// the value will not consistency
// buy try to run the command go run -race playground.go will handle share value and will return 100 that we expect

// Mutex can help to lock the share value when goroutime pull out and nobody can access till first goroutine return to it
func SampleRace() {
	fmt.Println("SampleRace")
	fmt.Println("CPUS ", runtime.NumCPU())
	fmt.Println("Gorutines ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup

	//Mutex
	var mu sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			// Mutex lock just let one go routine access the data
			mu.Lock()
			v := counter
			time.Sleep(time.Second)
			// swtich context when we have less thread. In real life not sure we have to used it.
			runtime.Gosched()
			v++
			counter = v
			//after done will allow other routine pull the data and update
			mu.Unlock()
			wg.Done()
		}()
		//show amount go routine runing at sametime
		fmt.Println("Gorutines in loop", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("CPUS ", runtime.NumCPU())
	fmt.Println("Gorutines ", runtime.NumGoroutine())
	fmt.Println("counter ", counter)
}
