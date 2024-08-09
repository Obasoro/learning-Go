package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to golang time module study")

	presentTime := time.Now()
	fmt.Println("present time: ", presentTime)
	fmt.Println(presentTime.Format(01-02-2006 15:04:05 Monday))

	CreatedDate := time.Date(2023, time.January, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(CreatedDate)
	fmt.Println(CreatedDate.Format("2006-01-02 Monday))
}