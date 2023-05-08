package main

import "fmt"

func main() {
	fmt.Println("struct in Golang")
	// No inheritance in golanf: No classes

	kunle := User{"Kunle", "kunle@gmail.com", true, 16}
	fmt.Println(kunle)
	fmt.Printf("kunle's details are: %+v\n", kunle)
	fmt.Printf("Name is %v and email is %v\n", kunle.Name, kunle.Email)
	kunle.GetStatus()
	kunle.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is this user active: ", u.Status)

}

func (u User) NewEmail() {
	u.Email = "klawluv@gmail.com"
	fmt.Println("Email of user is update: ", u.Email)
}