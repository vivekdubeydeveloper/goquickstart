package main
import "fmt"

func main(){
//we are going to initialize a int and store its address n pointer
//we will print address of variable and value variable using pointer
//we will update value of variable using pointer and print variable
x:=10
var y *int

y=&x

fmt.Println("address=",y," value=",*y)

*y=20

fmt.Println("x=",x)

}

