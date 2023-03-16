package main

import "fmt"

type date struct {
	day   int
	month int
	year  int
}

func (d *date) setdate() {
	fmt.Println("enter the dob:")
	fmt.Scan(&d.day)
	fmt.Scan(&d.month)
	fmt.Scan(&d.year)
}
func (d date) getdate() {
	fmt.Print("dob:", d.day, "/", d.month, "/", d.year)

}

type employee struct {
	name   string
	age    int
	phno   int
	dept   string
	salary int
	id     int
	dob    date //composition happens here
}

func (e *employee) setinfo() {
	fmt.Println("enter the details:")
	fmt.Scan(&e.name, &e.age, &e.phno, &e.dept, &e.salary, &e.id)
	e.dob.setdate()

}
func (e employee) getinfo() {
	fmt.Println("name:", e.name)
	fmt.Println("age:", e.age)
	fmt.Println("ph.no:", e.phno)
	fmt.Println("dept:", e.dept)
	fmt.Println("salary:", e.salary)
	fmt.Println("id:", e.id)
	e.dob.getdate()
}
func main() {
	emp := employee{}
	emp.setinfo()
	emp.getinfo()
}
