// Program to print the first 5 natural numbers

package main

import "fmt"

func main() {
	var n int // static & public var
	Sum := 0  //  private var
	fmt.Println("enter the no:")
	fmt.Scanf("%d", &n)
	for n > 0 { // for loop like while
		digit := n % 10
		// if Sum:=0 declar here cant acess println
		Sum = Sum*10 + digit
		n = n / 10
	}
	fmt.Println("after reverse:", Sum)
	for i := 1; i <= 10; i++ { // for loop
		fmt.Println(i)
	}
}
