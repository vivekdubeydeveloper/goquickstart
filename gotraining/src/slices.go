package main

import "fmt"

func main(){
	//initialize array
	a:=[10] int{1,2,3,4,5,6,7,8,9,10}

	//create slice from array print length and capacity
	b:=a[2:5]

	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("len=",len(b))
	fmt.Println("capacity=",cap(b))


	
	//prove reference type update any element of slices
	b[2]=125
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	
	//use append method
	b=append(b,123,124,125,126,127,128,129,130)

	fmt.Println("a=",a)
	fmt.Println("b=",b)
	

	//use make to create slices
	c:=make([] int, 5)

	for i:=1;i<6;i++{
		c[i-1]=i
	}

	fmt.Println(c)
	
}
