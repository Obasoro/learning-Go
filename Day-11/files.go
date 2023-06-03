package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learning how to load files in Golang")
	content := "This is how we learn programming online"

	file, err := os.Create("./learning.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)

	}

	fmt.Println("the length is: ", length)
	defer file.Close()
	readFile("./learning.txt")

	func readFile(filename, string) {

		databyte, err := ioutil.ReadFile(filename)
	    if err != nil {
		    panic(err)
	    }

		
	}

	
	fmt.Println("The text data inside the file \n", databyte)
}

func checkNilErr(err error) {
	if err := nil {
		panic(err)
	}
}