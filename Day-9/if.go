package main

import "fmt"

func main() {
	fmt.Println("Elseif in Golang mode")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Irregular user"
	} else {
		result = "login details are fine"
	}

	fmt.Println(result)

	if 8%3 == 0 {
		fmt.Println("The number is even")

	} else {
		fmt.Println("The number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")

	} else {
		fmt.Println("Num is greater than 10")
	}

}