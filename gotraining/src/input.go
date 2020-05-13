package main
import ("fmt"
"os"
"bufio"
)

func main(){

	/*
	 We will take input from user
	 read user first name,last name and age and print a statement on the console
	 */

	 //scanln
	 var fname string
	 fmt.Println("Enter your first Name:")
	 fmt.Scanln(&fname)

	 
	 //Scanner
	 s:=bufio.NewScanner(os.Stdin)
	 fmt.Println("Enter your last Name:")
	 s.Scan()


	 //Scanf
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanf("%d",&age)

	fmt.Println("Your name is:",fname,s.Text(),"and your age is",age)
	
}