package main

import "fmt"

func ChageValueAtZeroIndex(array [2]int) {
	array[0] = 3
	fmt.Println("Inside: ", array[0]) // will print 3
}

func main() {
	x := [2] int {
		ChangeValueAtZeroIndex(x)
		fmt.
	}
}

// Slice

var x := [] int or x := [] int {}

x := []int{1,3,5,7}

len(x) == 4

index.x[2] == 7

// append

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

for index, val := rangeSomeSlice {
	fmt.Printf("slice entry %d, %s\n", index, val)
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

var intPtr *int

intPtr = &X

// The above function a pointer was added to it


