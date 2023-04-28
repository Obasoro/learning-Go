package main

import "fmt"

func main() {
	fmt.Println("Learning Pointer ins Golang library")

	// var ptr *int
	// fmt.Println("Value pointer", ptr)

	myValue := 45

	var ptr = &myValue

	fmt.Println("Print my pointer memory allocation", ptr)
	fmt.Println("print pointer", *ptr)

	*ptr = *ptr + 3
	fmt.Println("My new value", myValue)

}