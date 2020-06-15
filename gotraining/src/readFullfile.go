package main
import (
	"fmt"
 a "io/ioutil"
)
//Read file at once
//Write file at once
func main(){
 data,err:=a.ReadFile("D:\\vivek\\work\\go\\goquickstart\\abc.txt")
 if err!=nil{
	 fmt.Println("Error in file reading ",err)
	 return 
 }
 fmt.Println("data is ",string(data))

 //owner read and write |Group read |Other read
 a.WriteFile("D:\\vivek\\work\\go\\goquickstart\\xyz.txt",data,0644)

 
}