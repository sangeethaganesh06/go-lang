package main

import "fmt"

func main() {
	numch := make(chan int)
	go counter(numch)
	for	s := range numch{
		fmt.Println(s)
	}
}
func counter(out chan int) {
	for i := 0; i < 5; i++ {
		out  <- i
	}
	close(out)
}
