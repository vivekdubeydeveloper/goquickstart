package main

import (
"fmt"
"os"

)

func main(){
	/*
	 we will give 2 command line parameters and 
	 print the parameters on console
	*/
	args:=os.Args

	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
	
}