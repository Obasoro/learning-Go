package main

import "fmt"

const LoginToken string = "toyinkunle" // public



func main() {
	//fmt.Println("variable") for variable
	var username string = "kunle"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// Declaring Boolean

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.679799
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Aliases

	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)

	var website = "gmic.com.ng"
	fmt.Println(website)

	// Number of user

	numberofuser := 300000
	fmt.Println(numberofuser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T \n", LoginToken)


}

// package main 

// import "fmt"

// func main() {

// }