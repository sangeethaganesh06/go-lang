package main

import "fmt"

func main() {

	numch := make(chan int)

	go counter(numch)
	go squer(numch)
	for a := range numch {
		fmt.Println(a)
	}
}
func counter(out chan int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
}
func squer(in chan int) {

	out <- a * a
}
