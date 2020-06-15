package main
import ("fmt"
"time"
)

func main(){

	go display("First")
	go display("Second")
	go display("Third")
	time.Sleep(5*time.Second)
	fmt.Println("End Method")
}

func display(x string){
for i:=0;i<1000;i++{
	fmt.Println(x)
}	
 
}
