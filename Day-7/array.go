package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang arrays!")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Orange"
	fruitlist[3] = "Tomato"

	fmt.Println("Fruilist: ", fruitlist)
	fmt.Println("Fruitlist: ", len(fruitlist))

	var vegList = [5]string{"Apple", "Orange", "Tomato"}
	fmt.Println("Veganlist: ", len(vegList))
}

