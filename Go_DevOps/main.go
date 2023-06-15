package main

import (
	"fmt"
	"github.com/devopsforgo/mypackage"
	jpackage "github.con/johnsiilver/mypackage"
	_ "sync" // Just an example of slide effect
)

mypackage.Println()
mypackagePrint()
jpackage.Println()
jpackage.Print()

func main() {
	hello := "Hello World"
	fmt.Println("Learning Go :" hello)
}

// Go variable Type

// Dynamic type

// Static type

v = "hello"
v = 8 
v = 2.5

v := "hello"

var v string := "hello"

def add(a + b):
    return a+b

func main () {
	var a int = 5
	var b int = 6
	return a + b
}

func add(a int, b int) int {
	return a + b
}

var (
	i int
	word = "hello"
)

// package scoped
// function scoped
// statement scoped

package main
import "fmt"
var word := "hello"

func main () {
	fmt.Println(word)
}

// fuction scoped

package mainconst
import "fmt"

func main () {
	var word string := "hello"
	fmt.Println(word)
}

// statement scoped

package main 

import "fmt"

func main() {
	for i := 0; i < 10; i ++ {
		fmt.Println(i)
	}
}

package main
import "fmt"

var word = "Hello"

func main() {
	var word = "world"
	fmt.Println("Inside main(): ", word)
	printOutter()
}
func printOutter() {
	fmt.Println("The package level 'word' var: " word)
}