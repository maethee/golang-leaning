package main

import "fmt"

func main() {
	SampleChannelAndRange()
	SampleChannelAndSelect()
}

// when we want loop use the data from channel
// we should close send channed  because loop channel don't know when the last channel of data will send to it.
func SampleChannelAndRange() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // should close to inform who using this channel send data to channel is done.
	}()

	// loop channel data
	// when we using range with channel we have to close when send data is done.
	// range should know the last send data to channel
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Done close")
}

//select receiver to execute different statement till done
func sampleChannelAndSelectReceiver(eve, odd, quit <-chan int) {
	for {
		select {
		case v := <-eve:
			fmt.Println("Even chan is ", v)
		case v := <-odd:
			fmt.Println("Odd chan is ", v)
		case v, ok := <-quit: // we can use ok for checking the channel is close or not  true = not close false = close channel and will return value is 0
			fmt.Println("Quit chan is ", v, ok)
			return
		}

	}
}

//func to select send channel to send the data
func sampleChannelAndSelectSender(eve, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			odd <- i
		} else {
			eve <- i
		}
	}
	close(quit) // will send 0 seem like send 0 to quit  quit <- 0
}

func SampleChannelAndSelect() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sampleChannelAndSelectSender(eve, odd, quit)

	sampleChannelAndSelectReceiver(eve, odd, quit)

}
