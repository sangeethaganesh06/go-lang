package main

import "fmt"

type variable interface {
	int | int64 | float32 | string
}

func max[T variable](x, y T) {
	if x > y {
		fmt.Printf("max number is %d", x)
	}
	fmt.Printf("max number is %d", y)
}
func main() {
	a := 12
	b := 15
	max(a, b)
}
