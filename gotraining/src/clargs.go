package main

import (
"fmt"
"os"

)

func main(){
	args:=os.Args
	fmt.Println(args)
	fmt.Println(args[1])
}
