package main

import ("fmt"
"strings"
)
func main(){
	//intialize a string
	//print string and string length
	//print string bytes
	//print string in upper case
	//will check if a string contains a substring or not

	v:="vivek"

	fmt.Println(v)
	fmt.Println(len(v))

	for i:=0;i<len(v);i++{
		fmt.Println(v[i])
	}

	for i:=0;i<len(v);i++{
		fmt.Printf("%c",v[i])
	}
	
	fmt.Println("")
	fmt.Println(strings.ToUpper(v))
	fmt.Println(strings.Contains(v,"viv"))
	fmt.Println(strings.Contains(v,"abc"))
}