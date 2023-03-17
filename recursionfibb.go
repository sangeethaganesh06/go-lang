package main

import "fmt"

func fibbonic(x int) int {

	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return (fibbonic(x-1) + fibbonic(x-2)) //recursion
	}
}
func main() {
	var n int
	fmt.Println("enter the valje of n:")
	fmt.Scanf("%d", &n)
	f := fibbonic(n)
	fmt.Printf("%d", f)
}
