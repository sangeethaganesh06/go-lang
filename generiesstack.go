package main

import "fmt"

type stack[T any] struct {
	list []T
}

func (s *stack[T]) push(x T) {
	s.list = append(s.list, x)
}
func (s *stack[T]) pull() T {
	val := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return val
}
func (s *stack[T]) peek() T {
	return s.list[len(s.list)-1]

}
func main() {
	intstack := stack[int]{}
	intstack.push(1)
	intstack.push(10)
	intstack.push(100)
	intstack.push(1000)
	fmt.Println("the elements in the stack are", intstack)
	fmt.Println("1 st element in the stack is", intstack.peek())
	fmt.Println(" pull 1 st element", intstack.pull())
	fmt.Println("after pull,the elements in the stack are", intstack)
	fmt.Println("after pull,1 st element in the stack is", intstack.peek())
}
