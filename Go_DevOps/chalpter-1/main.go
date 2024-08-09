package main

import (
	"crypto/x509"
	"errors"
	"fmt"
	"strings"
	_ "sync" // Just an example of slide effect

	"github.com/devopsforgo/mypackage"
	jpackage "github.con/johnsiilver/mypackage"
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


// Using conditional if/else block
switch blocks

if [expression that evalaute to boolean] {

}

if [int statement]; [staement that evaluate to boolean]{

}

if err := SomeFunct(); err != nil {
	fmt.Println(err)
}

kunle, err := SomeFunct()
if err != nil {
	return err
} 
fmt.Println(kunle)
return nil


if x > 10 {
	fmt.Println("x is greater than 0")
	
} else if x = 11 {
	fmt.Println

	
result, remainder := DivNum(4, 6)
fmt.Printf("Result: %d, Remaider: %d", result, remainder)


// Variadic Arguments

func sum(num []int) int {
	sum := 0
	for _, n := range number {
		dum += n
	}
	return sum
}

func sum(number ...int) int {
	sum := 0
	for _, n := range number {
		sum += n

	}
	return sum
}
fmt.Println(sum(1,3,5,7))


// Go accept anonympus function

func main() {
	result := func(kunle, obasoro, adeyemi string) {

		return kunle + " " + adeyemi
	} {"Hello", "world"}
	fmt.Println(result)
}


// private, public and internally expoerted

package say

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}

func printWorld() {
	fmt.Println("world")
}

func PrintHelloWorld() {
	PrintHello()
	printWorld()
}

package main

import "github.com/repo/example/say"

func main() {
	say.PrintHello()
	say.printHelloWorld()
}
// Arrays

func chnageValueAtZero(array [2]int) {
	array[0] = 2
	fmt.Println("Inside: ", array[0] // will print 3)
}

func main() {
	x := [2] int {

	}

	chnageValueAtZero(x)
	fmt.Println(x) // will print zero


}

// slice

var x = [] int

x := []int{}

x := []int{2,4,6,8}

len(x) = 4

var i int
for ; i < 10; i++ {
	fmt.Println(i)
}
fmt.Println("i's the final value:"


// Accessing Values

modelToMake := map[string]string{
	"priouse": "toyota",
	"chevelle": "chevey"

}

carMake := ModelToMake["chevelle"]
fmt.Println(carMake) // 

if carMake, ok := modelToMake["outback"]; ok {
	fmt.Printf("car model \"outback\" has made %q,\n", carMake)
}else {
	fmt.Printf("car model \"outback\" has an unknown make")
}

// Adding new valuess

modelToMake["outback"] = "subaru"
counter["pageHit"] = 10

for key, val := range modelToMake {
	fmt.Printf("car model %q has %q,\n", key, val)
}

// Pointer in Go

var x int = 23
fmt.Println(&x)


func ChangeValue(word string) {
	word += "world"
}

func main() {
	say := "hello"
	changeValue(say)
	fmt.Println(say)
}

var intPtr *int

intPtr = &x

fmt.Println(x) // print 23
fmt.Println(*intPtr) // print 23
*intPtr = 80 // replace the value of x to 80
fmt. println(x) print 80


func changeValue(word *string) {
	*word += "world"
}

func main() {
	say := "hello"
	changeValue(&say) // pass a pointer
	fmt.Println(say) // prints Helloworls
}

// Struct

var record = struct {
	name string
	number int
}{
	name: "Obasoro Olakunle",
	Age: 38,
}

fmt.Println("%s is %d years old\n", record.name, record.Age)

type CarModel string

var mycar 

/*. Declaring MAps. */

var counters = make(map[string]int, 10)

// Custom Struct types

type Record struct {
	Name string
	Age int
}

func main() {
	kunle := Record{Name: "Obasoro Olakunle", Age: 28}
	Adeyemi := Record{Name: "Adeyemi Obasoro", Age: 34}
	fmt.Printf("%+v\n", kunle)
	fmt.Printf("%+v\n", Adeyemi)
}

// Removing init statement
var i int;
for ;i < 10;i++ {
	fmt.Println(i)
}
fmt.Println("i's final value: ", i)

// Remove the poststatement too and you have a while loop

var i int
for i < 10 {
	i++
}
b := true
for b{ // this is a while loop
        fmtPrintln("Hello")
	}
// Create an infinite loop
for {

	fmt.Println("Hello World!")
}

for {
	if err:= doSomothing(); err != nil {
		break
	}
	fmt.Println("Keep going")
}

for i := 0;i < 10; i++ {
	if i % 2 == 0 {
		continue
	}
	fmt.Println("Odd number: "; i)
}

// Loop Braces
***Correct loop braces***

for {
    fmt.Println("Hello Wor")

}

***Incorrect loop braces***
if x > 2 {
	fmt.Println("x is greater than 2")
}

if err := SomeFunction(); err != nil {
	fmt.Println(err)
}

/// Getting to use Struct

var record = struct{
	Name  string
	Age int
}{
	Name: "Obasoro Olakunle",
	Age: 120, // Yeah, not publsihing the real one
}
fmt.Printf("%s is %d years old\n", record.name, record.Age)

//type carModel string

var mayCar CarModel = "Chevelle"
myCar = CarModel("Chevelle")
myCarAsString := string(MyCar)

// Custom Struct type
type Record struct {
	Name string
	Age int
}

func main() {
	David := Record{Name: "David Obasoro", Age: 23}
	Sarah := Record{Name: "Sarah Obasoro", Age: 45}
	fmt.Printf("%+v\n", David)
	Fmt.Printf("v%+v\n", Sarah)
}

// Adding ethod to type

type Record struct {
	Name string
	Age int
	// string returns a csv representing our record
}
func (r Record) String() string {
	return fmt.Sprintf("%s, %d", r.Name, r.Age)
}

// Example

type dotnet_api struct {

}

kunle := Record{Name: "Obasoro Olakunle", Age: 38}
fmt.Println(kunle.String())

myRecord.Name = "Obasoro Olakunle"
fmt.Println(myRecord.Name)

func changeName (r Record) {
	r.Name = "Adeyemi"
	fmt.Println("Inside changeName:", r.Name)
}

func main() {
	rec := Record{Name: "Oluwaseyi"}
	changeName(rec)
	fmt.Println("main:", rec.Name)
}

func changeName(r *Record) {
	r.Name = "Adeyemi"
	fmt.Println("Inside changeName:", r.Name)
}

func main() {
	// create a pointer to a record
	rec := &Record{Name: "oluwaseyi"}
	changeName(rec)
	fmt.Println("main:", rec.Name)
}

func (r Record) IncAge() {
	r.Age++
}

// Constructor are declared language

func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be <= 0")
	}
	return &Record{Name: name, Age: age}, nil
}

rec, err := NewRecord("Adegbemiga", 100) 
if err != nil {
	return err
}


// Interface in Go

type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}
func (p Person) String() string {
	return fmt.Sprintf("%s%,%s", p.First, p.Last)
}

type StrList []string
func (s StrList) String() string {
	return strings.Join(s, ", ")
}

// PrintStringer prints the value of a Stringer to stdout

func PrintStringer(s String) {
	fmt.Println(s.String())
}

func main() {
	yemi := Person{First: "Adeyemi", Last: "Obasoro"}
	var nameList Stringer = StrList{"Obasoro", "Olakunle"}

	PrintStringer(yemi)
	PrintStringer(nameList)
}

// Interfaces

/* The first thing aboug interface is tha value must implement every method defined in the method
interface{} is go universal value container used to pass value
*/
var i interface{}
i = 10
i = "hello world"
i = 3.5
i = Person{First: "John"}

func Println(a...interface{}) (n int, err error)
func Printf(format string, a ....interface{}) (n int, err error)

/* Assertion*/
if v, ok := i.(string); ok {
	fmt.Println(v)
}

switch v := i.(type) {
	case int:
		fmt.Printf("i was %d\n", i)
	case string:
		fmt.Printf("i was %s\n", i)
	case float:
		fmt.Printf("i was %v\n", i)
	case Person, *Person:
		fmt.Printf("i was %v\n", i)
	default:
		// %T will print l's underlying the type
		fmt.Printf("i was an unsupporteed type %T\n")
case condition:
	
}
