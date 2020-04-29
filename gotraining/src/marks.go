package main

import "fmt"

func main(){
m:=55
if m>75{
fmt.Println("D")
}else if m<75 && m>=60{
	fmt.Println("F")
}else if(m<60 && m>=45){
	fmt.Println("S")
}else if(m<45 && m>=33){
	fmt.Println("T")
}else{
	fmt.Println("FA")
}
	
}