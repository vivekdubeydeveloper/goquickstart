package main

import "fmt"

func main(){
/*
In first program we will pass int as parameter for switch
if int is 1 print case 1
if int is 2 print case 2
if int is 3 or 4 print case34
else print No case
*/
i:=12

switch i{
case 1:
	fmt.Println("Case 1")
case 2:
	fmt.Println("Case 2")
	
case 3,4:
	fmt.Println("Case 34")
	
default:
	fmt.Println("No Case")	

}

}