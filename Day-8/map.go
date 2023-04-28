package main

import "fmt"


func main () {
	fmt.Println("Maps is golang!!!")
	
	language := make(map[string]string)
	
	language["js"]= "JavaScript"
	language["py"] = "Python"
	language["rb"] = "Ruby"
	language["lua"] = "Lua"
	language["Go"] = "Golang"

	// fmt.Println("List of languages: ", language)

	// fmt.Println("js shorts for: ", language["js"])

	// delete(language, "rb")
	// fmt.Println(language)


	// Loops are interesting learning in Golang code

	for key, value := range language {
		fmt.Printf("for key %v, value %v\n", key, value)

	// for _, value := range language {
	// 	fmt.Printf("for key v, value %v\n", value)
	}

}