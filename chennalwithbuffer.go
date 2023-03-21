package main

import (
	"fmt"
	"time"
)

func main() {
											//chennals with buffer
	numCh := make(chan int)
	oddToMerger := make(chan int)
	evenToMerger := make(chan int, 10) // 10 is Buffer value
	squareToMerger := make(chan int)
	mergerToPrinter := make(chan int)
	done := make(chan struct{})

										//GoRoutines
	go counter(numCh)
	go oddEvenSplitter(numCh, oddToMerger, evenToMerger)
	go squarer(evenToMerger, squareToMerger)
	go merger(oddToMerger, squareToMerger, mergerToPrinter)
	go printer(mergerToPrinter, done)
	<-done			// done receiver
}

func counter(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(in chan int, out chan int) {
	for a := range in {
		time.Sleep(2 * time.Second)		//delay function
		out <- a * a
	}
	close(out)
}

func oddEvenSplitter(in chan int, odd chan int, even chan int) {
	for a := range in {
		if a%2 == 0 {
			even <- a
		} else {
			odd <- a
		}
	}
	close(odd)
	close(even)
}

func merger(oddIn chan int, evenIn chan int, out chan int) {
	for {
		select {
		case a, ok := <-oddIn:
			if ok {
				out <- a
			} else {
				oddIn = nil
			}

		case b, ok := <-evenIn:
			if ok {
				out <- b
			} else {
				evenIn = nil
			}
		}
		if oddIn == nil && evenIn == nil {
			break
		}
	}
	close(out)
}

func printer(in chan int, done chan struct{}) {
	for a := range in {
		fmt.Println(a)
	}
	done <- struct{}{}				//ends the input
}
