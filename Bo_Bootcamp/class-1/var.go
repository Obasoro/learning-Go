package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	var Name string = "Kunle"
	var v1, v2 = 1.3, 5.7
	var v3 = "Hello"
	v1 = 4.5

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.142))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf(''))

	cv1 := pl(a...any)
	cv2 := int(cv1)
	pl(cv2)

	cv3 := "500000"
	cv4, err := strconv.Atoi((cv3))
	pl.(cv4, err, reflect.TypeOf(cv4))

	cv5 := 500000
	cv6 := strconv.Itoa(cv5)
	pl.(cv6)

	cV7 := 31.42
	if cV8, err := strconv.ParserFloat(cV7, 64); 
	err == nil {
		pl(cV8)
	}
	cV9 := fmt.Sprint("%f", 31.42)
	pl(cV9s\n)
}