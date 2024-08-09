package main

import "fmt"

func ChageValueAtZeroIndex(array [2]int) { // putting size before type
	array[0] = 3 // 3 is assigned to index 0
	fmt.Println("Inside: ", array[0]) // will print 3
}

// Arrays, unlike slices, are not pointer wrapper types. 
//Passing an array as a function argument passes a copy:
func main() {
	x := [2] int {
		ChangeValueAtZeroIndex(x)
		fmt.Println(x) // it will print ??
	}
	x := [2] int{}
	ChageValueAtZeroIndex(x)
	fmt.Println(x) // it will print 0
}

// Slice

/*
- A slice is not statically sized
- A slice can grow to accommodate new values
/*

var x := [] int or x := [] int {}

x := []int{1,3,5,7} 

len(x) == 4
// index is a method in Golang

// x := []int{1, 3, 7, 7} the code below would achieve

index.x[2] == 7


// append is a method


x = append(x, 2)

y := x[1:3]


func DoAppend(s1 []int) {
	s1 = append(s1, 100)
	fmt.Println("Inside: ", s1)

}

func main(){
	x := []int{1, 2, 3}
	DoAppend(x)
	fmt.Println("outside: ", x)
}

func DoAppend(s1 []int) []int {
	return append(s1, 100)
}

func main() {
	x:= []int{1, 2, 3}
	x = DoAppend(x)
	fmt.Println("outside: ", x)
}

// extracting values C style aand Go

for i := 0; i < len(SomeSlice); i++ {
    fmt.Printf("Slice entry %d: %s\n", i, SomeSlice[i])
}

for index, val := range SomeSlice {
	fmt.Printf("slice entry %d, %s\n", index, val)
}

for _, val := range someSlice {
	fmt.Printf("Slice entry: %s\n", val )
}

func ChangeValue(word *string) {
	// Add "World" to the string pointer
	*word += "world"
}

func main() {
	say := "hello"
	ChangeValue(&say) //pointer added
	fmt.Println(say)
}

// Pointers

var intPtr = &

intPtr


The above function a pointer was added to it


Maps are a collections of key-value pair. 
In Python it is called dictionaries

var. counters = make(map[string]int, 10)


for _, val := range someSlice {
	fmt.Printf("slice entry: %s\n", val)
}

// Understanding Map

// Maps are a collection of key-value pairs that a user use to store data

var counters = make(map[string]int, 10)

modelToMake := map[string]string {
	"prius": "toyota",
	"chevelle": "chevy"
}

// Accessing value
carMake := modelToMake['chevelle"]
fmt.Println(carMake)

carMake := modelToMake["outback"]
fmt.Println(carMake)

if carmake, ok := modelToMake["outbuack"]; ok {
	fmt.Printf("car model \"outback\" has make %q", carMake)
}else{
	fmt.Printf("car model \"outback"\ has )
}