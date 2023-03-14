package main
import "fmt"
func main(){
arr:=[5]int {11,22,33,44,55}
for i,n:=range arr{ 					// range using array
fmt.Printf("arr[%d]:%d\n",i,n)
}
slc:=[]int{111,222,33}
for i1,n1:=range slc{ 					// range using slic e
fmt.Printf("slice[%d]:%d\n",i1,n1)
}
str:="sangeethamathigomathiganesh"
for i,n:=range arr{ 					// range using array
fmt.Printf("arr[%d]:%d\n",i,n)
}