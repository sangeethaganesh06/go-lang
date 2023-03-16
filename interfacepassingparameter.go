package main

import "fmt"

type human interface {
	getinfo()
	setinfo()
}
type breakfast struct {
	dishname  string
	rate      int
	hotelname string
}
type lunch struct {
	breakfast //structure embedding
	juice     string
}

func (s *lunch) getinfo() {
	fmt.Print("enter details:")
	fmt.Scan(&s.hotelname, &s.dishname, &s.rate, &s.juice)
}
func (s lunch) setinfo() {
	fmt.Println(s.hotelname, s.dishname, s.rate, s.juice)
}
func process(h human) { //passes interface into function
	h.getinfo()
	h.setinfo()
	h.setinfo()
	h.setinfo()
	h.setinfo()

}
func main() {

	humans := []human{new(lunch)}
	for _, hum := range humans {
		process(hum)

	}
}
