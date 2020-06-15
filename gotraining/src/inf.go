package main

import "fmt"

//declare shape interface with area method

type shape interface{
	area() int
}

//declare rect struct with length and width
 
type rect struct{
	l int
	w int
}



//implement interface
func (r rect) area() int{
 return r.l*r.w
}

//declare square struct with edge
type square struct{
	e int
	
}

//implement interface
func (s square) area() int{
	return s.e*s.e
   }

func main(){
 //delcare shape type variable
 var s shape
 //initialize with rect 
 s=rect{20,10}
 
 //call area function and print
 fmt.Println(s.area())
 s=square{30}

 fmt.Println(s.area())
 
}





