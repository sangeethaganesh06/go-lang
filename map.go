package main
 
import "fmt"
 
func main() {
    var map1 map[int]int
    if map1 == nil {			 // Checking if the map is nil or not
     
        fmt.Println("True")
    } else {
     
        fmt.Println("False")
    }
    map2 := map[int]string{90: "Dog",91: "Cat",92: "Cow",93: "Bird",94: "Rabbit"} //map[int]{
	
	fmt.Println("before update:",map2)
	map2[100]="ganesh"
	map2[101]="gomathi"			// update 
	fmt.Println("after update:",map2)
	delete(map2,100)			// delete
	fmt.Println("after delete:",map2)
    fmt.Println(map2[93])		// can access by particular key
	
	k,ok:=map2[94]
 fmt.Println(k,ok)				// v,ok helps to check the value and print value as well bol expressions
	}