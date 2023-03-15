// Program to pass pointer as a function argument

package main

import "fmt"

// function definition with a pointer argument
func update(num *int) {

	fmt.Println("The number is", num)
}
func main() {

	var number = 66

	// function call
	update(&number)

}
