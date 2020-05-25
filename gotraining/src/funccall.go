package main
import "fmt"

func main(){
//swap two no 
a,b:=10,50
fmt.Println("Call By value")
callbyval(a,b)

fmt.Println("Within main a=",a,"b=",b)


fmt.Println("Call By Reference")
callbyref(&a,&b)
fmt.Println("Within main a=",a,"b=",b)
}

func callbyval(a,b int){
t:=a
a=b
b=t
fmt.Println("a=",a,"b=",b)
}

func callbyref(x,y *int){
t:=*x
*x=*y
*y=t
fmt.Println("x=",*x,"y=",*y)
}




