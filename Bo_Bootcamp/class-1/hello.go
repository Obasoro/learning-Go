package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"os"
	"log"
)

var pl = f.Println
func main() {
	pl("Hello Go")
	fmt.Print("Hello, what is your name? ")
	pl("What is your learning path")
	reader := bufio.NewReader(os.Stain)
	name, err := reader.ReadString("\n")
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}


}