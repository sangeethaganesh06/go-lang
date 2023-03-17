package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return (x * factorial(x-1)) //recursion occurs
	}
}
func main() {
	var n int
	fmt.Println("enter the value of n:")
	fmt.Scanf("%d", &n)
	f := factorial(n)
	fmt.Printf("%d", f)
}
