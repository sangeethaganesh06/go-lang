package main
import(		
			"fmt"
			"sort"
			)
func main(){
   var car [4]string												// static  decl
		car[0]="innova"
		car[1]="urban cruser"
		car[2]="glanza"
		car[3]="fortuner"
	fmt.Println("toyoto cars are:",car)
	fmt.Println("my fav toyoto cars :",len(car))
   bike:=[3] string {"unicorn","duke","bmw"}						//dynamic decl
	fmt.Println("my fav bikes are:",bike)
	
// slice-discriptor of array
//slice with data type

s:=[]int{}		// syntax
	fmt.Println(s)	 												//print empty slice
s1:=[]string{"q","r","g","w"}
	fmt.Println(s1)													// print[q r g w]
	fmt.Println("length of slice:",len(s1))							// print length of slice
	fmt.Println("capacity of slice:",cap(s1))						// "capacity"-no of element in underlying array
// slice from array
slc:=car[1:]	
	fmt.Println(slc)			
slc= append(slc,"i10","i20")										//append function
	fmt.Println(slc)
	fmt.Println(slc[1:4])											//we can slice the slice 
	fmt.Println("length of slice:",len(slc))						// print length of slice
	fmt.Println("capacity of slice:",cap(slc))			

// slice using make () function	- array with abstraction
mark:=make([] int,5)
mark[0]=80
mark[1]=99
mark[2]=65
mark[3]=70
mark[4]=90
	fmt.Println(mark)
mark=append(mark,89,63,95)
	fmt.Println(mark)
sort.Ints(mark)														//slice sorting
	fmt.Println(mark)
	fmt.Println("the slice is sorted:",sort.IntsAreSorted(mark))	//shows boolean expressions
totalmark := make([]int, len(mark))
  				   							
  // copy content of numbers to numbers1 
  fmt.Println("orginal slices:",mark)
	copy(totalmark,mark)												// copy slice to new one
  fmt.Println("copied slice:",totalmark) 
   var index int=2		
mark=append(mark[:index],mark[index+1:]...)								// remve a element of second intex
fmt.Println(mark) 
}