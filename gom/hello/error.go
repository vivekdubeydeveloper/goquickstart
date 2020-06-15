package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		We are going pass wrong file path command line
		we will try to open that file and will use panic

	*/
	fmt.Println("main started")
	fmt.Println(os.Args)
	test(os.Args[1])
	fmt.Println("main completed")
}

func test(filePath string) {
	defer test1()
	fmt.Println("test")
	_, err := os.Open(filePath)
	if err != nil {
		/*fmt.Println("Error in file opening")
		return*/
		panic("Error in file opening:" + err.Error())

	}

}

func test1() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from", r)
	}
	fmt.Println("test1")

}
