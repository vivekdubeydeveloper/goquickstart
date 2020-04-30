package main

import "fmt"

func main(){
	/*
	In second program we will pass the mark of student for switch
	if mark is more than or equal to 75, print D
	if mark is less than 75 and greater than or equal to 60, print F
	if mark is less than 60 and greater than or equal to 45, print S
	if mark is less than 45 and greater than or equal to 33, print T
    if mark is less than 33 print FA
	*/
   
	m:=80
	switch{
	case m>=75:
		fmt.Println("D")
	case m<75 && m>=60:
		fmt.Println("F")
	case m<60 && m>=45:
		fmt.Println("S")
    case m<45 && m>=33:
		fmt.Println("T")
	default:
		fmt.Println("FA")		
	}
}

