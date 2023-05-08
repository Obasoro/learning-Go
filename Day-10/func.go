package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to function in Golang")
	greeter()
	greetersTwo()
}

func greetersTwo() {
    fmt.Println("Another method within a function")

	result := adder(5, 3)
	fmt.Println("Result is: ", result)

	proRes := ProAdder(4, 6, 8, 4)

	fmt.Println("This is result of: ", proRes)
	}

func adder(valeOne int, valueTwo int) int{
	return valeOne + valueTwo
}

func ProAdder(values ...int) int {
	total := 0

	for _, val := range values{
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Namste from Golang")
}