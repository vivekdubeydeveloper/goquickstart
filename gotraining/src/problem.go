package main

import ("fmt"
"time"
"math"
"math/rand"
)


func main(){
//print current date
fmt.Println("Current Date is :",time.Now())

/*
want to print square root of any no
want to print random no
want to calculate area of circle
*/
fmt.Println("Square root of no is :",math.Sqrt(4))
fmt.Println("Random no is :",rand.Intn(100))
fmt.Println("Area of circle is: ",math.Pi*10*10)
}
