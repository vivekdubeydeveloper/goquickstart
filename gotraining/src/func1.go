package main

import "fmt"

func main(){
//for multiple return area and perimeter of a rectangle
a,p:=cal(10,20)
fmt.Println("a=",a,"p=",p)

//named return substraction of 2 nos
fmt.Println("m=",minus(15,10))
}

func cal(l,w int) (int,int){
	a:=l*w
	p:=(l+w)*2
	return a,p
}

func minus(x,y int) (m int){
	m=x-y
	return 
}
