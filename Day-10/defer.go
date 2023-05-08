package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("world")
	fmt.Println("hello")
	MyDefer()
}

func MyDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}