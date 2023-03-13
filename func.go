package main

import "fmt"

func add(n1 int, n2 int) int {
	total := n1 + n2
	return total

}
func main() {
	sum := add(12, 13)
	fmt.Println(sum)
	sum = add(12, 13)
	fmt.Println(sum)
	sum = add(13, 13)
	fmt.Println(sum)
	sum = add(11, 13)
	fmt.Println(sum)
	sum = add(10, 13)
	fmt.Println(sum)
}
