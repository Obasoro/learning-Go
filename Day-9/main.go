package main

import "fmt"

func main() {
	fmt.Println("struct in Golang")
	// No inheritance in golanf: No classes

	kunle := User{"Kunle", "kunle@gmail.com", true, 16}
	fmt.Println(kunle)
	fmt.Printf("kunle's details are: %+v\n", kunle)
	fmt.Printf("Name is %v and email is %v\n", kunle.Name, kunle.Emailtouch)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


