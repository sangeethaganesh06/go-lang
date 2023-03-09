package main
import ("fmt")

func main() {
  var x = 10
  x +=5
  fmt.Println(x)
   var x1 = 5
  var y = 3
  fmt.Println(x1>y) // returns 1 (true) because 5 is greater than 3
   var p = 5
  var q = 3
  fmt.Println(p<q && p>q) // logical operator
  fmt.Println(!(x < 5 && x < 10)) 
  x2:=12
  x3:=10
  fmt.Printf("x2 = %b\n",x2) // bitwise operator
  fmt.Printf("x3= %b\n",x3) 
    
  fmt.Printf("x2 | x3 is %d  and its bit value :%b\n",x2|x3,x2|x3)
  
}