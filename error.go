package main

import (
	"errors"
	"fmt"
)

func div(a int, b int) (int, error) {
	if b == 0 {
		return 1, errors.New("arithmatic error")
	}
	return a / b, nil
}

func main() {
	var x int
	var y int
	fmt.Println("enter x value:")
	fmt.Scanf("%d", &x)
	fmt.Println("enter y value:")
	fmt.Scanf("%d", &y)
	fmt.Print("x:", x, "and", "y:", y)
	answer, err := div(x, y)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}
