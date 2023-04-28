// Slicine functionality in Go!!!
package main

import (
	"fmt"
	"sort"
)

func main () {
	// fmt.Println("Welcome to Day8")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruilist is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Orange")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	//highScore := make([]int, 4)

	// highScore[0] = 234
	// highScore[1] = 300
	// highScore[2] = 800
	// highScore[3] = 450
	//highScore[4] = 700

	// highScore = append(highScore, 555, 420, 600)

	// fmt.Println(highScore)

	// fmt.Println(sort.IntsAreSorted(highScore))
	// sort.Ints(highScore)
	// fmt.Println(highScore)

	// How to remove a value from slices

	var courses = []string{"react.js", "javascript.js", "swift", "python", "ruby"}
	fmt.Println("courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)






}
