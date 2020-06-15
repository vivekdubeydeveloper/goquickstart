package main
import ("fmt"
"os"
"bufio"
)
//read file in chunk
//Write file in chunk
func main(){
 f,err:=os.Open("D:\\vivek\\work\\go\\goquickstart\\abc.txt")
 f1,err1:=os.Create("D:\\vivek\\work\\go\\goquickstart\\xyz11.txt")
 if err!=nil{
	fmt.Println("Error in file reading ",err)
	return 
}

if err1!=nil{
	fmt.Println("Error in file creating ",err1)
	return 
}
defer f1.Close()
defer f.Close()
defer sayHello()
r:=bufio.NewReader(f)
b:=make([]byte,10)

for{

	n,err:=r.Read(b)
	
	if err!=nil{
		fmt.Println("Error in file reading ",err)
		return
	}
	fmt.Println("n=",n)
	fmt.Println(string(b[0:n]))
    f1.Write(b)
}

fmt.Println("++++END FOR LOOP+++++")

}


func sayHello(){
	fmt.Println("++++SAY HELLO+++++")
}
