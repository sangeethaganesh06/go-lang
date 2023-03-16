package main

import "fmt"

type person struct {
	name string
	age  int
	phno int
}
type employee struct {
	person			//person embedded inside the employee
	dept   string
	salary int
	id     int
}

func (e *employee) setinfo() {
	fmt.Println("enter the details:")
	fmt.Scan(&e.name, &e.age, &e.phno, &e.dept, &e.salary, &e.id)
}
func (e employee) getinfo() {
	fmt.Println("name:", e.name)
	fmt.Println("age:", e.age)
	fmt.Println("ph.no:", e.phno)
	fmt.Println("dept:", e.dept)
	fmt.Println("salary:", e.salary)
	fmt.Println("id:", e.id)
}
func main() {
	emp := employee{}
	emp.setinfo()
	emp.getinfo()
}
