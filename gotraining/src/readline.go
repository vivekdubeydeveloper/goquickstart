package main
import ("fmt"
"bufio"
"os"
)
//Read file line by line
//Write file line by line

func main(){
	f,err:=os.Open("D:\\vivek\\work\\go\\goquickstart\\abc.txt")
	f1,err1:=os.Create("D:\\vivek\\work\\go\\goquickstart\\xyz12.txt")
 if err!=nil{
	fmt.Println("Error in file reading ",err)
	return 
}

if err1!=nil{
	fmt.Println("Error in file crating ",err1)
	return 
}

defer f1.Close()
defer f.Close()
s:=bufio.NewScanner(f)

for s.Scan(){
	//fmt.Println(s.Text())
	f1.WriteString(s.Text())
	f1.WriteString("\n")
}

if s.Err!=nil{
	fmt.Println("Error in file reading ",s.Err)
}

}