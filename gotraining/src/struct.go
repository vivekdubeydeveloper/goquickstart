package main
import "fmt"

//Define a struct

type Student struct{
	rollNo int
	name string
}

func main(){
//initialize struct print the value print single value
//s:=Student{1,"Ravi"}
s:=Student{rollNo:1,name:"Ravi"}

fmt.Println(s)
fmt.Println(s.name)

//declare pointer of struct assign value
var s1 *Student
s1=&s
//print single value 

fmt.Println(s1)
fmt.Println((*s1).name)
fmt.Println(s1.name)

}

