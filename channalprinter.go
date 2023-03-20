package main

import "fmt"

func main() {

	numch := make(chan int)
	done := make(chan struct{})
	go counter(numch)
	go squrer(numch, done)

	<-done
}

func counter(out chan int) {
	for i := 0; i < 15; i++ {
		if i%2==0{
		out <- i
	}
}
	close(out)
}
func squrer(out chan int, done chan struct{}) {
	for a := range out {
		fmt.Println(a + a)
	}
	done <- struct{}{}
}
