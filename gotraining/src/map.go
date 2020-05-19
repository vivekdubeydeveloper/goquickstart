package main
import "fmt"

func main(){

//declaring map
//var a map[string]int

//initializing map
//a["x"]=1
//declaring and initializing map
var b=map[string]int{
	"a":1,
	"b":2,
	"c":3,
}

//declaring and initializing map using sorthand
c:=map[string]int{
	"a":1,
	"b":2,
	"c":3,
}
 
//using make function to declare the map and initialize 
d:=make(map[string]int)
d["a"]=1
d["b"]=2
d["c"]=3

fmt.Println("b=",b)
fmt.Println("c=",c)
fmt.Println("d=",d)

//range form of for loop
for _,v:=range d{
	fmt.Println("v=",v)
}

//search key
x,y:=d["abc"]
fmt.Println("x=",x,"y",y)

}