package main
import "fmt"

func main(){
//declare array using approach1
var a [3]int
//initialize array
for i:=0;i<len(a);i++{
	a[i]=(i+1)*10
}
fmt.Println(a)
//print length of array
fmt.Println(len(a))
//declare array with sorthand operator and ellipsis
//b:=[3] int{1,2,3}
b:=[...] int{1,2,3}

//range form of for loop

for   i,v :=range b{
fmt.Println("i=",i,"v=",v)
}


//array are value type not reference type
c:=[3] int{1,2,3}
d:=c
d[1]=5
fmt.Println("c=",c)
fmt.Println("d=",d)
//compare two array==
e:=[3] int{1,2,3}
f:=[3] int{1,4,3}

fmt.Println("e==f",e==f)

}