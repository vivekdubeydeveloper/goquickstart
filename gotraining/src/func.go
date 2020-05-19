package main

import (
	
	"fmt"
)

func main(){
//define a function for sum of 2 nos
fmt.Println(sum(12,10))

//how function help in designing
//calculate multiplication of 2 nos
//calculate bigger from 2 no
//calculate factorial of a no
a,b:=5,6


fmt.Println("m=",mul(a,b))

fmt.Println("greater =",greater(a,b))

fmt.Println("fact=",fact(a))

}


func sum(x int,y int) int{
	return x+y
}

func mul(x ,y int) int{
	return x*y
}

func greater(x ,y int) int{
	if x>y{
		return x
	}else{
		return y
	}
	 
}

func fact(a int) int{
	fact:=1

	for a>1{
		fact=fact*a
		a--
	}
	return fact
}