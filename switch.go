package main
import "fmt"
func main(){
fmt.Println("welcome to mathi's pizza")
fmt.Println("please rate out pizza between 1 to 5")
var r int           								//static variable declation
fmt.Scanf("%d",&r)
fmt.Printf("thanks for giving %d star rating  \n",r)
if r<=5{
switch r {
case 1:
fmt.Println("sorry to disappointing you.we are planned to send a free pizza for you,hope it be better that last one")
case 2:
fmt.Println("sorry to disappointing you.i'll make changes")
case 3:
fmt.Println("we are planned to make some chance in out pizza taste. hope it will be better one")
case 4:
fmt.Println("next time you may feel very good taste")
case 5:
fmt.Println("thank you smuch for your appreciation. ")
}
fmt.Println("congrats,you got 20% off on your next pizza")
}else if r>=5{
fmt.Println("please give a rating between 1-5")
}}