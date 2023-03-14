package main

import (
	"fmt"
	"math/rand"
)

func add(n1 int, n2 int) int {
	total := n1 + n2
	return total

}
func mul(n1 int, n2 int) int {
	mul1 := n1 * n2
	return mul1
}
func div(n1 int, n2 int) int {
	div1 := n1 / n2
	return div1
}
func remin(n1 int, n2 int) int {
	remin1 := n1 % n2
	return remin1
}
func random() int { //random no
	random1 := rand.Intn(10)
	return random1
}

func Sum(nums ...int) int { //varidic function def

	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
func main() {
	sum := add(100, 10) // arith,atic operation
	fmt.Println(sum)
	sum = mul(100, 10)
	fmt.Println(sum)
	sum = div(100, 10)
	fmt.Println(sum)
	sum = remin(100, 10)
	fmt.Println(sum)
	sum = random()
	fmt.Println(sum)

	s1 := Sum(1, 2, 3) //varidic func with no of parameters
	s2 := Sum(1, 2, 3, 4)
	s3 := Sum(1, 2, 3, 4, 5)

	fmt.Println(s1, s2, s3)
}
