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

var i int
for ;i < 10; 1++ {
	fmt.Println(i)
}
fmt.Println("i's final value: ". i)

var i int
for i < 10 {
	i++
}

b := true
for b { // This will loop forever
	fmt.Println("Hello")
}

for {

	if err := DoSomething() ; err !=nil {
		break
	}
	fmt.Println("keep going")
}

for i := 0; i < 10; i++ {
	if i % 2 == 0 {
		continue
	}
	fmt.Println("Odd number: ", i)
}

if x > 2 {
	fmt.Println("X is greater than 2")
}

if err := SomeFunction(); err := nil {
	fmt.Println(err)
}

// Understanding condition

// if condition {
	function1()

}else {
	function2()
}

// Example

if v, err := SomeFunct(); err:= nil {
	return err
} else {
	fmt.Println(v)
	return nil
}

v, err := SomeFunct()
if err != nil {
	return err
}
fmt.Println(v)
return nil

if x > 0 {
	fmt.Println("X is greater than Zero")

}else if x < 0 {
	fmt.Println("X is less than zero")

} else {
	fmt.Println("X is equal to Zero")
}

if x > 0 
{ // This must go up on the previous line
    fmt.Println("hello")
}

else { // This line must start on the previous line
	fmt.Println("World")
	
}

// Example for Switch

switch [value] {
case [match]:
	[statement]
case [matcj]:
	[satement]
default:
	[statement]
}

switch x {
case 3:
	fmt.Println("X is 3")
	case 4, 5 // execute if x is 4 or 5
	fmt.Println("X is 4 or 5")
default:
	fmt.Println(x is unknown)
}

switch x := SomeFunct() {
case 3:
	fmt.Println("x is 3")
}

switch {
case x > 0:
	fmt.Println("X is greater than 0")
case x < 0:
	fmt.Println("X is lesser than 0")
case x = 0:
	fmt.Println("X is equal 0")
default:
	fmt.Println("X must be 0")
}

// Function in Go are what you expect in modern programming

// Multiple return value are supported
// Variadic arguement
// Nmaed return values

func fucntionNmae ([Varname] [VarType], ...) ([return value] [return value], ...) {

}

func add(x int, y int) int {
	return x+y

}

return := add(2, 3)
fmt.Println(result)


func DivNum (num, div int) (res, rem int) {
	result = num / div
	remainder = num % div
	return res, rem
}

result, remainder := DivNum(3,2)
fmt.Printf("Result: %d, Remaider: %d", result, remainder)