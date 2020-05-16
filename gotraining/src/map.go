package main
import "fmt"

func main(){
var mymap= map[string]int{
	"a":10,
	"b":20,
	"c":30,
}

mymap1:= map[string]int{
	"a":10,
	"b":20,
	"c":30,
}

c:=make(map[int]string)

c[1]="a"
c[2]="b"
c[3]="c"

/*mymap["a"]=10
mymap["b"]=20
mymap["c"]=30
mymap["d"]=40*/

fmt.Println(mymap)
fmt.Println(mymap1)
fmt.Println(c)
y,z:=c[20]
fmt.Println(y,"",z)
}