package main

import "fmt"

type human interface {
	getinfo()
	setinfo()
}
type student struct {
	name  string
	class int
	rank  int
}
type employee struct {
	name string
	age  int
	id   int
}

func (s *student) getinfo() {
	fmt.Print("enter student details:")
	fmt.Scan(&s.name, &s.class, &s.rank)
}
func (s student) setinfo() {
	fmt.Println(s.name, s.class, s.rank)
}
func (e *employee) getinfo() {
	fmt.Print("enter employee details:")
	fmt.Scan(&e.name, &e.age, &e.id)
}
func (e employee) setinfo() {
	fmt.Println(e.name, e.age, e.id)
}
func main() {
	stu := student{}
	emp := employee{}
	hum := []human{&stu, &emp} //using slice
	//	hum[0].getinfo()
	//	hum[1].getinfo()
	//	hum[0].setinfo()
	//	hum[1].setinfo()
	for _, h := range hum {
		h.getinfo()
		h.setinfo()
	}
}
